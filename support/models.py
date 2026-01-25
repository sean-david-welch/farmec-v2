from django.db import models
from django.utils.translation import gettext_lazy as _
from base_model import BaseModel


class Warrantyclaim(BaseModel):
    id = models.TextField(primary_key=True, verbose_name=_('ID'))
    dealer = models.CharField(max_length=255, verbose_name=_('dealer'), help_text=_('Authorized dealer name'))
    dealer_contact = models.CharField(max_length=255, blank=True, null=True, verbose_name=_('dealer contact'), help_text=_('Dealer contact person or phone'))
    owner_name = models.CharField(max_length=255, verbose_name=_('owner name'), help_text=_('Machine owner/customer name'))
    machine_model = models.CharField(max_length=255, verbose_name=_('machine model'), help_text=_('Model name/number of the machine'))
    serial_number = models.CharField(max_length=100, db_index=True, verbose_name=_('serial number'), help_text=_('Machine serial number for identification'))
    install_date = models.DateField(blank=True, null=True, verbose_name=_('install date'), help_text=_('Date machine was installed'))
    failure_date = models.DateField(blank=True, null=True, verbose_name=_('failure date'), help_text=_('Date failure occurred'))
    repair_date = models.DateField(blank=True, null=True, verbose_name=_('repair date'), help_text=_('Date repair was completed'))
    failure_details = models.TextField(blank=True, null=True, verbose_name=_('failure details'), help_text=_('Description of the failure/damage'))
    repair_details = models.TextField(blank=True, null=True, verbose_name=_('repair details'), help_text=_('Description of work performed'))
    labour_hours = models.DecimalField(max_digits=8, decimal_places=2, blank=True, null=True, verbose_name=_('labour hours'), help_text=_('Hours of labour for repair'))
    completed_by = models.CharField(max_length=255, blank=True, null=True, verbose_name=_('completed by'), help_text=_('Technician or person who completed repair'))

    class Meta:
        managed = True
        db_table = 'WarrantyClaim'

    def __str__(self):
        return f"Claim {self.id} - {self.owner_name}"


class Partsrequired(BaseModel):
    id = models.TextField(primary_key=True, verbose_name=_('ID'))
    warranty = models.ForeignKey('Warrantyclaim', on_delete=models.CASCADE, blank=True, null=True, verbose_name=_('warranty claim'), help_text=_('Warranty claim this part is required for'))
    part_number = models.CharField(max_length=100, blank=True, null=True, db_index=True, verbose_name=_('part number'), help_text=_('Supplier part number or SKU'))
    quantity_needed = models.PositiveIntegerField(verbose_name=_('quantity needed'), help_text=_('Number of units required'))
    invoice_number = models.CharField(max_length=100, blank=True, null=True, verbose_name=_('invoice number'), help_text=_('Supplier invoice reference'))
    description = models.TextField(blank=True, null=True, verbose_name=_('description'), help_text=_('Part name and specifications'))

    class Meta:
        managed = True
        db_table = 'PartsRequired'

    def __str__(self):
        return f"{self.part_number} x{self.quantity_needed}"


class Machineregistration(BaseModel):
    id = models.TextField(primary_key=True, verbose_name=_('ID'))
    dealer_name = models.CharField(max_length=255, verbose_name=_('dealer name'), help_text=_('Authorized dealer business name'))
    dealer_address = models.CharField(max_length=500, blank=True, null=True, verbose_name=_('dealer address'), help_text=_('Dealer location address'))
    owner_name = models.CharField(max_length=255, verbose_name=_('owner name'), help_text=_('Machine owner/customer name'))
    owner_address = models.CharField(max_length=500, blank=True, null=True, verbose_name=_('owner address'), help_text=_('Owner location address'))
    machine_model = models.CharField(max_length=255, verbose_name=_('machine model'), help_text=_('Model name/number of machine'))
    serial_number = models.CharField(max_length=100, db_index=True, verbose_name=_('serial number'), help_text=_('Machine serial number'))
    install_date = models.DateField(blank=True, null=True, verbose_name=_('install date'), help_text=_('Date of installation'))
    invoice_number = models.CharField(max_length=100, blank=True, null=True, verbose_name=_('invoice number'), help_text=_('Sales invoice reference'))
    complete_supply = models.BooleanField(default=False, verbose_name=_('complete supply'), help_text=_('All items and documents included with delivery'))
    pdi_complete = models.BooleanField(default=False, verbose_name=_('PDI complete'), help_text=_('Pre-Delivery Inspection completed'))
    pto_correct = models.BooleanField(default=False, verbose_name=_('PTO correct'), help_text=_('Power Take-Off correctly configured'))
    machine_test_run = models.BooleanField(default=False, verbose_name=_('machine test run'), help_text=_('Machine tested and working'))
    safety_induction = models.BooleanField(default=False, verbose_name=_('safety induction'), help_text=_('Owner received safety training'))
    operator_handbook = models.BooleanField(default=False, verbose_name=_('operator handbook'), help_text=_('Owner manual provided'))
    date = models.DateField(blank=True, null=True, verbose_name=_('date'), help_text=_('Registration completion date'))
    completed_by = models.CharField(max_length=255, blank=True, null=True, verbose_name=_('completed by'), help_text=_('Person who completed registration'))

    class Meta:
        managed = True
        db_table = 'MachineRegistration'

    def __str__(self):
        return f"{self.owner_name} - {self.machine_model}"
