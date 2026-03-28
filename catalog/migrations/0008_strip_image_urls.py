from django.db import migrations

BASE_URL = 'https://static.farmec.ie/'


def strip_image_urls(apps, schema_editor):
    Supplier = apps.get_model('catalog', 'Supplier')
    Machine = apps.get_model('catalog', 'Machine')
    Product = apps.get_model('catalog', 'Product')
    Spareparts = apps.get_model('catalog', 'Spareparts')

    for obj in Supplier.objects.all():
        changed = False
        for field in ('logo_image', 'marketing_image'):
            val = getattr(obj, field).name
            if val and val.startswith(BASE_URL):
                setattr(obj, field, val[len(BASE_URL):])
                changed = True
        if changed:
            obj.save(update_fields=['logo_image', 'marketing_image'])

    for model, field in [(Machine, 'machine_image'), (Product, 'product_image'), (Spareparts, 'parts_image')]:
        for obj in model.objects.all():
            val = getattr(obj, field).name
            if val and val.startswith(BASE_URL):
                setattr(obj, field, val[len(BASE_URL):])
                obj.save(update_fields=[field])


def restore_image_urls(apps, schema_editor):
    Supplier = apps.get_model('catalog', 'Supplier')
    Machine = apps.get_model('catalog', 'Machine')
    Product = apps.get_model('catalog', 'Product')
    Spareparts = apps.get_model('catalog', 'Spareparts')

    for obj in Supplier.objects.all():
        changed = False
        for field in ('logo_image', 'marketing_image'):
            val = getattr(obj, field).name
            if val and not val.startswith('http'):
                setattr(obj, field, BASE_URL + val)
                changed = True
        if changed:
            obj.save(update_fields=['logo_image', 'marketing_image'])

    for model, field in [(Machine, 'machine_image'), (Product, 'product_image'), (Spareparts, 'parts_image')]:
        for obj in model.objects.all():
            val = getattr(obj, field).name
            if val and not val.startswith('http'):
                setattr(obj, field, BASE_URL + val)
                obj.save(update_fields=[field])


class Migration(migrations.Migration):

    dependencies = [
        ('catalog', '0007_imagefield_conversion'),
    ]

    operations = [
        migrations.RunPython(strip_image_urls, restore_image_urls),
    ]
