import logging
import sqlite3
import tempfile
from pathlib import Path

import boto3
from botocore.client import BaseClient
from django.conf import settings
from django.db import connection
from django.tasks import task
from django.utils import timezone

from farmec.logging import log_error
from farmec.utils import EmailClient

logger: logging.Logger = logging.getLogger(__name__)


def build_client() -> BaseClient:
    """
    Build the S3 client used for backups from the project's AWS settings.
    """
    return boto3.client(
        "s3",
        region_name=settings.AWS_S3_REGION_NAME,
        aws_access_key_id=settings.AWS_S3_ACCESS_KEY_ID,
        aws_secret_access_key=settings.AWS_S3_SECRET_ACCESS_KEY,
    )


def create_snapshot(destination: Path) -> None:
    """
    Write a consistent copy of the default SQLite database to ``destination``.

    Uses SQLite's online backup API, so it is safe while the app is writing.

    :param destination: File path to create.
    """
    target: sqlite3.Connection = sqlite3.connect(destination)
    try:
        connection.ensure_connection()
        connection.connection.backup(target)
    finally:
        target.close()


def upload_backup(client: BaseClient, snapshot: Path, key: str) -> int:
    """
    Upload a snapshot to the backup bucket and confirm it landed.

    :param client: S3 client to upload with.
    :param snapshot: Local file to upload.
    :param key: Destination object key.
    :returns: Size in bytes of the uploaded object, as reported by S3.
    :raises RuntimeError: When the stored object size differs from the local file.
    """
    client.upload_file(str(snapshot), settings.BACKUP_BUCKET, key)
    head: dict = client.head_object(Bucket=settings.BACKUP_BUCKET, Key=key)
    stored: int = head["ContentLength"]
    expected: int = snapshot.stat().st_size
    if stored != expected:
        raise RuntimeError(
            "Uploaded {key} is {stored} bytes, expected {expected}".format(
                key=key, stored=stored, expected=expected
            )
        )
    return stored


def notify_failure(error: Exception) -> None:
    """
    Report a failed backup to Sentry and by email, without raising.

    :param error: The exception that aborted the backup.
    """
    log_error(
        logger, "Database backup failed: %s", error,
        context={"bucket": settings.BACKUP_BUCKET}, tags={"feature": "backup"}, exc_info=True,
    )
    try:
        text: str = "The nightly database backup failed:\n\n{error!r}".format(error=error)
        EmailClient().send(
            subject="Farmec database backup FAILED", text=text, html="<pre>{text}</pre>".format(text=text),
        )
    except Exception:
        logger.exception("Could not send backup failure email")


@task
def backup_database() -> str:
    """
    Snapshot the SQLite database and upload it to S3 as ``database-YYYYMMDD.db``.

    Failures are logged, reported, then re-raised.

    :returns: The S3 object key that was written.
    """
    key: str = "database-{today:%Y%m%d}.db".format(today=timezone.now())
    try:
        with tempfile.TemporaryDirectory() as tmp:
            snapshot: Path = Path(tmp) / key
            create_snapshot(snapshot)
            size: int = upload_backup(build_client(), snapshot, key)
    except Exception as error:
        notify_failure(error)
        raise
    logger.info(
        "Database backup uploaded to s3://%s/%s (%d bytes)",
        settings.BACKUP_BUCKET,
        key,
        size,
    )
    return key
