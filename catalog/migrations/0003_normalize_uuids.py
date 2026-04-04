from django.db import migrations


def normalize_uuids(apps, schema_editor):
    with schema_editor.connection.cursor() as cursor:
        # Normalize FK columns first (before changing PKs)
        cursor.execute("UPDATE Machine SET supplier_id = REPLACE(supplier_id, '-', '') WHERE supplier_id LIKE '%-%'")
        cursor.execute("UPDATE Product SET machine_id = REPLACE(machine_id, '-', '') WHERE machine_id LIKE '%-%'")
        cursor.execute("UPDATE SpareParts SET supplier_id = REPLACE(supplier_id, '-', '') WHERE supplier_id LIKE '%-%'")
        cursor.execute("UPDATE Video SET supplier_id = REPLACE(supplier_id, '-', '') WHERE supplier_id LIKE '%-%'")

        # Normalize PK columns
        cursor.execute("UPDATE Supplier SET id = REPLACE(id, '-', '') WHERE id LIKE '%-%'")
        cursor.execute("UPDATE Machine SET id = REPLACE(id, '-', '') WHERE id LIKE '%-%'")
        cursor.execute("UPDATE Product SET id = REPLACE(id, '-', '') WHERE id LIKE '%-%'")
        cursor.execute("UPDATE SpareParts SET id = REPLACE(id, '-', '') WHERE id LIKE '%-%'")
        cursor.execute("UPDATE Video SET id = REPLACE(id, '-', '') WHERE id LIKE '%-%'")


class Migration(migrations.Migration):

    dependencies = [
        ('catalog', '0002_alter_machine_id_alter_product_id_and_more'),
    ]

    operations = [
        migrations.RunPython(normalize_uuids, migrations.RunPython.noop),
    ]
