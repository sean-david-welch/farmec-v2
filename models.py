# This is an auto-generated Django model module.
# You'll have to do the following manually to clean this up:
#   * Rearrange models' order
#   * Make sure each model has one field with primary_key=True
#   * Make sure each ForeignKey and OneToOneField has `on_delete` set to the desired behavior
#   * Remove `managed = False` lines if you wish to allow Django to create, modify, and delete the table
# Feel free to rename the models, but don't rename db_table values or field names.
from django.db import models


class Blog(models.Model):
    id = models.TextField(primary_key=True)
    title = models.TextField()
    date = models.TextField(blank=True, null=True)
    main_image = models.TextField(blank=True, null=True)
    subheading = models.TextField(blank=True, null=True)
    body = models.TextField(blank=True, null=True)
    created = models.TextField(blank=True, null=True)
    slug = models.TextField(blank=True, null=True)

    class Meta:
        managed = False
        db_table = 'Blog'


class Carousel(models.Model):
    id = models.TextField(primary_key=True)
    name = models.TextField()
    image = models.TextField(blank=True, null=True)
    created = models.TextField(blank=True, null=True)

    class Meta:
        managed = False
        db_table = 'Carousel'


class Employee(models.Model):
    id = models.TextField(primary_key=True)
    name = models.TextField()
    email = models.TextField()
    role = models.TextField()
    profile_image = models.TextField(blank=True, null=True)
    created = models.TextField(blank=True, null=True)

    class Meta:
        managed = False
        db_table = 'Employee'


class Exhibition(models.Model):
    id = models.TextField(primary_key=True)
    title = models.TextField()
    date = models.TextField(blank=True, null=True)
    location = models.TextField(blank=True, null=True)
    info = models.TextField(blank=True, null=True)
    created = models.TextField(blank=True, null=True)

    class Meta:
        managed = False
        db_table = 'Exhibition'


class Lineitems(models.Model):
    id = models.TextField(primary_key=True)
    name = models.TextField()
    price = models.FloatField()
    image = models.TextField(blank=True, null=True)

    class Meta:
        managed = False
        db_table = 'LineItems'


class Machine(models.Model):
    id = models.TextField(primary_key=True)
    supplier = models.ForeignKey('Supplier', models.DO_NOTHING)
    name = models.TextField()
    machine_image = models.TextField(blank=True, null=True)
    description = models.TextField(blank=True, null=True)
    machine_link = models.TextField(blank=True, null=True)
    created = models.TextField(blank=True, null=True)
    slug = models.TextField(blank=True, null=True)

    class Meta:
        managed = False
        db_table = 'Machine'


class Machineregistration(models.Model):
    id = models.TextField(primary_key=True)
    dealer_name = models.TextField()
    dealer_address = models.TextField(blank=True, null=True)
    owner_name = models.TextField()
    owner_address = models.TextField(blank=True, null=True)
    machine_model = models.TextField()
    serial_number = models.TextField()
    install_date = models.TextField(blank=True, null=True)
    invoice_number = models.TextField(blank=True, null=True)
    complete_supply = models.IntegerField(blank=True, null=True)
    pdi_complete = models.IntegerField(blank=True, null=True)
    pto_correct = models.IntegerField(blank=True, null=True)
    machine_test_run = models.IntegerField(blank=True, null=True)
    safety_induction = models.IntegerField(blank=True, null=True)
    operator_handbook = models.IntegerField(blank=True, null=True)
    date = models.TextField(blank=True, null=True)
    completed_by = models.TextField(blank=True, null=True)
    created = models.TextField(blank=True, null=True)

    class Meta:
        managed = False
        db_table = 'MachineRegistration'


class Partsrequired(models.Model):
    id = models.TextField(primary_key=True)
    warranty = models.ForeignKey('Warrantyclaim', models.DO_NOTHING)
    part_number = models.TextField(blank=True, null=True)
    quantity_needed = models.TextField()
    invoice_number = models.TextField(blank=True, null=True)
    description = models.TextField(blank=True, null=True)

    class Meta:
        managed = False
        db_table = 'PartsRequired'


class Privacy(models.Model):
    id = models.TextField(primary_key=True)
    title = models.TextField()
    body = models.TextField(blank=True, null=True)
    created = models.TextField(blank=True, null=True)

    class Meta:
        managed = False
        db_table = 'Privacy'


class Product(models.Model):
    id = models.TextField(primary_key=True)
    machine = models.ForeignKey(Machine, models.DO_NOTHING)
    name = models.TextField()
    product_image = models.TextField(blank=True, null=True)
    description = models.TextField(blank=True, null=True)
    product_link = models.TextField(blank=True, null=True)
    slug = models.TextField(blank=True, null=True)

    class Meta:
        managed = False
        db_table = 'Product'


class Spareparts(models.Model):
    id = models.TextField(primary_key=True)
    supplier = models.ForeignKey('Supplier', models.DO_NOTHING)
    name = models.TextField()
    parts_image = models.TextField(blank=True, null=True)
    spare_parts_link = models.TextField(blank=True, null=True)
    slug = models.TextField(blank=True, null=True)

    class Meta:
        managed = False
        db_table = 'SpareParts'


class Supplier(models.Model):
    id = models.TextField(primary_key=True)
    name = models.TextField()
    logo_image = models.TextField(blank=True, null=True)
    marketing_image = models.TextField(blank=True, null=True)
    description = models.TextField(blank=True, null=True)
    social_facebook = models.TextField(blank=True, null=True)
    social_twitter = models.TextField(blank=True, null=True)
    social_instagram = models.TextField(blank=True, null=True)
    social_youtube = models.TextField(blank=True, null=True)
    social_linkedin = models.TextField(blank=True, null=True)
    social_website = models.TextField(blank=True, null=True)
    created = models.TextField(blank=True, null=True)
    slug = models.TextField(blank=True, null=True)

    class Meta:
        managed = False
        db_table = 'Supplier'


class Terms(models.Model):
    id = models.TextField(primary_key=True)
    title = models.TextField()
    body = models.TextField(blank=True, null=True)
    created = models.TextField(blank=True, null=True)

    class Meta:
        managed = False
        db_table = 'Terms'


class Timeline(models.Model):
    id = models.TextField(primary_key=True)
    title = models.TextField()
    date = models.TextField(blank=True, null=True)
    body = models.TextField(blank=True, null=True)
    created = models.TextField(blank=True, null=True)

    class Meta:
        managed = False
        db_table = 'Timeline'


class Video(models.Model):
    id = models.TextField(primary_key=True)
    supplier = models.ForeignKey(Supplier, models.DO_NOTHING)
    web_url = models.TextField(blank=True, null=True)
    title = models.TextField(blank=True, null=True)
    description = models.TextField(blank=True, null=True)
    video_id = models.TextField(blank=True, null=True)
    thumbnail_url = models.TextField(blank=True, null=True)
    created = models.TextField(blank=True, null=True)

    class Meta:
        managed = False
        db_table = 'Video'


class Warrantyclaim(models.Model):
    id = models.TextField(primary_key=True)
    dealer = models.TextField()
    dealer_contact = models.TextField(blank=True, null=True)
    owner_name = models.TextField()
    machine_model = models.TextField()
    serial_number = models.TextField()
    install_date = models.TextField(blank=True, null=True)
    failure_date = models.TextField(blank=True, null=True)
    repair_date = models.TextField(blank=True, null=True)
    failure_details = models.TextField(blank=True, null=True)
    repair_details = models.TextField(blank=True, null=True)
    labour_hours = models.TextField(blank=True, null=True)
    completed_by = models.TextField(blank=True, null=True)
    created = models.TextField(blank=True, null=True)

    class Meta:
        managed = False
        db_table = 'WarrantyClaim'


class GooseDbVersion(models.Model):
    version_id = models.IntegerField()
    is_applied = models.IntegerField()
    tstamp = models.TextField(blank=True, null=True)  # This field type is a guess.

    class Meta:
        managed = False
        db_table = 'goose_db_version'
