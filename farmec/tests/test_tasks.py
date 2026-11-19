import sqlite3
import tempfile
from pathlib import Path

from django.core import mail
from django.test import TestCase, override_settings

from farmec import tasks


class FakeS3:
    """Minimal in-memory stand-in for the two S3 calls the backup makes."""

    def __init__(self, size_offset: int = 0) -> None:
        self.size_offset: int = size_offset
        self.sizes: dict[str, int] = {}

    def upload_file(self, filename: str, bucket: str, key: str) -> None:
        self.sizes[key] = Path(filename).stat().st_size + self.size_offset

    def head_object(self, Bucket: str, Key: str) -> dict:
        return {"ContentLength": self.sizes[Key]}


class SnapshotTest(TestCase):
    def test_create_snapshot__is_valid_sqlite_file(self):
        with tempfile.TemporaryDirectory() as tmp:
            path: Path = Path(tmp) / "snap.db"
            tasks.create_snapshot(path)
            snapshot = sqlite3.connect(path)
            try:
                result = snapshot.execute("PRAGMA integrity_check").fetchone()
            finally:
                snapshot.close()
        self.assertEqual(result, ("ok",))


@override_settings(BACKUP_BUCKET="test-bucket")
class UploadBackupTest(TestCase):
    def test_upload__returns_stored_size(self):
        with tempfile.TemporaryDirectory() as tmp:
            snapshot: Path = Path(tmp) / "snap.db"
            snapshot.write_bytes(b"12345")
            size: int = tasks.upload_backup(FakeS3(), snapshot, "key.db")
        self.assertEqual(size, 5)

    def test_upload__size_mismatch_raises(self):
        with tempfile.TemporaryDirectory() as tmp:
            snapshot: Path = Path(tmp) / "snap.db"
            snapshot.write_bytes(b"12345")
            with self.assertRaises(RuntimeError):
                tasks.upload_backup(FakeS3(size_offset=-1), snapshot, "key.db")


@override_settings(EMAIL_USER="owner@example.com")
class NotifyFailureTest(TestCase):
    def test_notify_failure__sends_email(self):
        tasks.notify_failure(RuntimeError("boom"))
        self.assertEqual(len(mail.outbox), 1)
        self.assertIn("boom", mail.outbox[0].body)
