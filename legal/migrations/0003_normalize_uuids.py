from django.db import migrations


def normalize_uuids(apps, schema_editor):
    with schema_editor.connection.cursor() as cursor:
        cursor.execute("UPDATE Privacy SET id = REPLACE(id, '-', '') WHERE id LIKE '%-%'")
        cursor.execute("UPDATE Terms SET id = REPLACE(id, '-', '') WHERE id LIKE '%-%'")


class Migration(migrations.Migration):

    dependencies = [
        ('legal', '0002_alter_privacy_id_alter_terms_id'),
    ]

    operations = [
        migrations.RunPython(normalize_uuids, migrations.RunPython.noop),
    ]
