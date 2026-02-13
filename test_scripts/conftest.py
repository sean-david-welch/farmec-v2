import sqlite3
from pathlib import Path

import pytest


def get_old_db_schema() -> dict[str, str]:
    """
    Extract schema from old database.

    :return: Dictionary mapping table names to CREATE TABLE statements
    """
    old_db_path: Path = Path(__file__).parent.parent / "server" / "database" / "database.db"

    if not old_db_path.exists():
        return {}

    schema = {}
    try:
        old_conn: sqlite3.Connection = sqlite3.connect(str(old_db_path))
        old_cursor: sqlite3.Cursor = old_conn.cursor()
        old_cursor.execute(
            "SELECT name FROM sqlite_master WHERE type='table' AND name NOT LIKE 'sqlite_%'",
        )
        tables: list[str] = [row[0] for row in old_cursor.fetchall()]
        for table_name in tables:
            old_cursor.execute(
                "SELECT sql FROM sqlite_master WHERE type='table' AND name=?",
                (table_name,),
            )
            result = old_cursor.fetchone()
            if result and result[0]:
                schema[table_name] = result[0]
        old_conn.close()
    except sqlite3.OperationalError:
        pass
    return schema


@pytest.fixture(scope="session", autouse=True)
def django_db_setup(django_db_blocker):
    """
    Set up Django test database with old database schema.

    Since tests run without migrations, we manually create the old schema tables
    that the Django models expect.
    """
    with django_db_blocker.unblock():
        from django.conf import settings
        from django.db import DEFAULT_DB_ALIAS

        schema = get_old_db_schema()
        if not schema:
            yield
            return
        db_config: str = settings.DATABASES[DEFAULT_DB_ALIAS]
        if db_config["ENGINE"] == "django.db.backends.sqlite3":
            db_path = db_config["NAME"]
            try:
                test_conn = sqlite3.connect(str(db_path))
                test_cursor = test_conn.cursor()
                for table_name, create_sql in schema.items():
                    try:
                        test_cursor.execute(create_sql)
                    except sqlite3.OperationalError:
                        pass
                test_conn.commit()
                test_conn.close()
            except sqlite3.OperationalError:
                pass
    yield
