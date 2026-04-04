from django.db import migrations


def normalize_uuids(apps, schema_editor):
    with schema_editor.connection.cursor() as cursor:
        # Normalize FK columns first
        cursor.execute("UPDATE PartsRequired SET warranty_id = REPLACE(warranty_id, '-', '') WHERE warranty_id LIKE '%-%'")

        # Normalize PK columns
        cursor.execute("UPDATE WarrantyClaim SET id = REPLACE(id, '-', '') WHERE id LIKE '%-%'")
        cursor.execute("UPDATE PartsRequired SET id = REPLACE(id, '-', '') WHERE id LIKE '%-%'")
        cursor.execute("UPDATE MachineRegistration SET id = REPLACE(id, '-', '') WHERE id LIKE '%-%'")


class Migration(migrations.Migration):

    dependencies = [
        ('support', '0002_alter_machineregistration_id_alter_partsrequired_id_and_more'),
    ]

    operations = [
        migrations.RunPython(normalize_uuids, migrations.RunPython.noop),
    ]
