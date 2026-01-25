from django.db import models
from base_model import BaseModel


class Warrantyclaim(BaseModel):
    id = models.TextField(primary_key=True)
    dealer = models.CharField(max_length=255, verbose_name='dealer')
    dealer_contact = models.CharField(max_length=255, blank=True, null=True, verbose_name='dealer contact')
    owner_name = models.CharField(max_length=255, verbose_name='owner name')
    machine_model = models.CharField(max_length=255, verbose_name='machine model')
    serial_number = models.CharField(max_length=100, db_index=True, verbose_name='serial number')
    install_date = models.DateField(blank=True, null=True, verbose_name='install date')
    failure_date = models.DateField(blank=True, null=True, verbose_name='failure date')
    repair_date = models.DateField(blank=True, null=True, verbose_name='repair date')
    failure_details = models.TextField(blank=True, null=True, verbose_name='failure details')
    repair_details = models.TextField(blank=True, null=True, verbose_name='repair details')
    labour_hours = models.DecimalField(max_digits=8, decimal_places=2, blank=True, null=True, verbose_name='labour hours')
    completed_by = models.CharField(max_length=255, blank=True, null=True, verbose_name='completed by')

    class Meta:
        managed = True
        db_table = 'WarrantyClaim'

    def __str__(self):
        return f"Claim {self.id} - {self.owner_name}"


class Partsrequired(BaseModel):
    id = models.TextField(primary_key=True)
    warranty = models.TextField(blank=True, null=True, verbose_name='warranty id')
    part_number = models.CharField(max_length=100, blank=True, null=True, db_index=True, verbose_name='part number')
    quantity_needed = models.PositiveIntegerField(verbose_name='quantity needed')
    invoice_number = models.CharField(max_length=100, blank=True, null=True, verbose_name='invoice number')
    description = models.TextField(blank=True, null=True, verbose_name='description')

    class Meta:
        managed = True
        db_table = 'PartsRequired'

    def __str__(self):
        return f"{self.part_number} x{self.quantity_needed}"


class Machineregistration(BaseModel):
    id = models.TextField(primary_key=True)
    dealer_name = models.CharField(max_length=255, verbose_name='dealer name')
    dealer_address = models.CharField(max_length=500, blank=True, null=True, verbose_name='dealer address')
    owner_name = models.CharField(max_length=255, verbose_name='owner name')
    owner_address = models.CharField(max_length=500, blank=True, null=True, verbose_name='owner address')
    machine_model = models.CharField(max_length=255, verbose_name='machine model')
    serial_number = models.CharField(max_length=100, db_index=True, verbose_name='serial number')
    install_date = models.DateField(blank=True, null=True, verbose_name='install date')
    invoice_number = models.CharField(max_length=100, blank=True, null=True, verbose_name='invoice number')
    complete_supply = models.BooleanField(default=False, verbose_name='complete supply')
    pdi_complete = models.BooleanField(default=False, verbose_name='PDI complete')
    pto_correct = models.BooleanField(default=False, verbose_name='PTO correct')
    machine_test_run = models.BooleanField(default=False, verbose_name='machine test run')
    safety_induction = models.BooleanField(default=False, verbose_name='safety induction')
    operator_handbook = models.BooleanField(default=False, verbose_name='operator handbook')
    date = models.DateField(blank=True, null=True, verbose_name='date')
    completed_by = models.CharField(max_length=255, blank=True, null=True, verbose_name='completed by')

    class Meta:
        managed = True
        db_table = 'MachineRegistration'

    def __str__(self):
        return f"{self.owner_name} - {self.machine_model}"
