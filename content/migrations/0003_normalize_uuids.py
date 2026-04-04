from django.db import migrations


def normalize_uuids(apps, schema_editor):
    with schema_editor.connection.cursor() as cursor:
        cursor.execute("UPDATE Blog SET id = REPLACE(id, '-', '') WHERE id LIKE '%-%'")
        cursor.execute("UPDATE Carousel SET id = REPLACE(id, '-', '') WHERE id LIKE '%-%'")
        cursor.execute("UPDATE Exhibition SET id = REPLACE(id, '-', '') WHERE id LIKE '%-%'")
        cursor.execute("UPDATE Timeline SET id = REPLACE(id, '-', '') WHERE id LIKE '%-%'")


class Migration(migrations.Migration):

    dependencies = [
        ('content', '0002_alter_blog_id_alter_carousel_id_alter_exhibition_id_and_more'),
    ]

    operations = [
        migrations.RunPython(normalize_uuids, migrations.RunPython.noop),
    ]
