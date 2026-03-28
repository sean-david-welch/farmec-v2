from django.db import migrations

BASE_URL = 'https://static.farmec.ie/'


def strip_image_urls(apps, schema_editor):
    Employee = apps.get_model('team', 'Employee')

    for obj in Employee.objects.all():
        val = obj.profile_image.name
        if val and val.startswith(BASE_URL):
            obj.profile_image = val[len(BASE_URL):]
            obj.save(update_fields=['profile_image'])


def restore_image_urls(apps, schema_editor):
    Employee = apps.get_model('team', 'Employee')

    for obj in Employee.objects.all():
        val = obj.profile_image.name
        if val and not val.startswith('http'):
            obj.profile_image = BASE_URL + val
            obj.save(update_fields=['profile_image'])


class Migration(migrations.Migration):

    dependencies = [
        ('team', '0005_imagefield_conversion'),
    ]

    operations = [
        migrations.RunPython(strip_image_urls, restore_image_urls),
    ]
