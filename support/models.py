from django.db import models


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


class Partsrequired(models.Model):
    id = models.TextField(primary_key=True)
    warranty = models.TextField(blank=True, null=True)
    part_number = models.TextField(blank=True, null=True)
    quantity_needed = models.TextField()
    invoice_number = models.TextField(blank=True, null=True)
    description = models.TextField(blank=True, null=True)

    class Meta:
        managed = False
        db_table = 'PartsRequired'


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
