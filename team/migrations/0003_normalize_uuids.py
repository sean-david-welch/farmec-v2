from django.db import migrations


def normalize_uuids(apps, schema_editor):
    with schema_editor.connection.cursor() as cursor:
        cursor.execute("UPDATE Employee SET id = REPLACE(id, '-', '') WHERE id LIKE '%-%'")


class Migration(migrations.Migration):

    dependencies = [
        ('team', '0002_alter_employee_id'),
    ]

    operations = [
        migrations.RunPython(normalize_uuids, migrations.RunPython.noop),
    ]
