from django.db import migrations

BASE_URL = 'https://static.farmec.ie/'


def strip_image_urls(apps, schema_editor):
    Blog = apps.get_model('content', 'Blog')
    Carousel = apps.get_model('content', 'Carousel')

    for model, field in [(Blog, 'main_image'), (Carousel, 'image')]:
        for obj in model.objects.all():
            val = getattr(obj, field).name
            if val and val.startswith(BASE_URL):
                setattr(obj, field, val[len(BASE_URL):])
                obj.save(update_fields=[field])



def restore_image_urls(apps, schema_editor):
    Blog = apps.get_model('content', 'Blog')
    Carousel = apps.get_model('content', 'Carousel')

    for model, field in [(Blog, 'main_image'), (Carousel, 'image')]:
        for obj in model.objects.all():
            val = getattr(obj, field).name
            if val and not val.startswith('http'):
                setattr(obj, field, BASE_URL + val)
                obj.save(update_fields=[field])


class Migration(migrations.Migration):

    dependencies = [
        ('content', '0005_imagefield_conversion'),
    ]

    operations = [
        migrations.RunPython(strip_image_urls, restore_image_urls),
    ]
