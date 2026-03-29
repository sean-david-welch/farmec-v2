import sqlite3
import uuid
from datetime import datetime
from decimal import Decimal, InvalidOperation
from pathlib import Path

from django.core.management.base import BaseCommand
from django.utils import timezone

from catalog.models import Machine, Product, Spareparts, Supplier, Video
from content.models import Blog, Carousel, Exhibition, Timeline
from legal.models import Privacy, Terms
from support.models import Machineregistration, Partsrequired, Warrantyclaim
from team.models import Employee


def parse_date(value):
    if not value:
        return None
    for fmt in ('%Y-%m-%d', '%Y-%m-%dT%H:%M:%SZ', '%Y-%m-%d %H:%M:%S'):
        try:
            return datetime.strptime(value[:10], '%Y-%m-%d').date()
        except (ValueError, TypeError):
            pass
    return None


def parse_datetime(value):
    if not value:
        return timezone.now()
    for fmt in ('%Y-%m-%dT%H:%M:%SZ', '%Y-%m-%d %H:%M:%S', '%Y-%m-%d'):
        try:
            dt = datetime.strptime(value, fmt)
            return timezone.make_aware(dt)
        except (ValueError, TypeError):
            continue
    return timezone.now()


def parse_decimal(value):
    if not value:
        return None
    try:
        return Decimal(str(value))
    except InvalidOperation:
        return None


def parse_int(value):
    if not value:
        return 0
    try:
        return int(value)
    except (ValueError, TypeError):
        return 0


S3_BASE = 'https://static.farmec.ie/'


def strip_url(value):
    if not value:
        return ''
    if value.startswith(S3_BASE):
        return value[len(S3_BASE):]
    return value


def base_fields(row):
    return {
        'order': 1,
        'publish': True,
        'uid': uuid.uuid4(),
        'created': parse_datetime(row['created'] if 'created' in row.keys() else None),
        'modified': timezone.now(),
    }


class Command(BaseCommand):
    help = 'Migrate data from Go SQLite database to Django'

    def add_arguments(self, parser):
        parser.add_argument(
            '--db',
            default='database/go-database.db',
            help='Path to the Go SQLite database file',
        )

    def handle(self, *args, **options):
        db_path = Path(options['db'])
        if not db_path.exists():
            self.stderr.write(f'Database not found: {db_path}')
            return

        conn = sqlite3.connect(db_path)
        conn.row_factory = sqlite3.Row
        cur = conn.cursor()

        self.migrate_suppliers(cur)
        self.migrate_machines(cur)
        self.migrate_products(cur)
        self.migrate_spareparts(cur)
        self.migrate_videos(cur)
        self.migrate_blogs(cur)
        self.migrate_carousels(cur)
        self.migrate_exhibitions(cur)
        self.migrate_timelines(cur)
        self.migrate_employees(cur)
        self.migrate_privacy(cur)
        self.migrate_terms(cur)
        self.migrate_warrantyclaims(cur)
        self.migrate_partsrequired(cur)
        self.migrate_machineregistrations(cur)

        conn.close()
        self.stdout.write(self.style.SUCCESS('Migration complete.'))

    def migrate_suppliers(self, cur):
        cur.execute('SELECT * FROM Supplier')
        objs = [
            Supplier(
                id=row['id'],
                name=row['name'],
                logo_image=strip_url(row['logo_image']),
                marketing_image=strip_url(row['marketing_image']),
                description=row['description'],
                social_facebook=row['social_facebook'],
                social_twitter=row['social_twitter'],
                social_instagram=row['social_instagram'],
                social_youtube=row['social_youtube'],
                social_linkedin=row['social_linkedin'],
                social_website=row['social_website'],
                slug=row['slug'],
                **base_fields(row),
            )
            for row in cur.fetchall()
        ]
        Supplier.objects.bulk_create(objs, ignore_conflicts=True)
        self.stdout.write(f'  Suppliers: {len(objs)}')

    def migrate_machines(self, cur):
        supplier_ids = set(Supplier.objects.values_list('id', flat=True))
        cur.execute('SELECT * FROM Machine')
        rows = [r for r in cur.fetchall() if r['supplier_id'] in supplier_ids]
        objs = [
            Machine(
                id=row['id'],
                supplier_id=row['supplier_id'],
                name=row['name'],
                machine_image=strip_url(row['machine_image']),
                description=row['description'],
                machine_link=row['machine_link'],
                slug=row['slug'],
                **base_fields(row),
            )
            for row in rows
        ]
        Machine.objects.bulk_create(objs, ignore_conflicts=True)
        self.stdout.write(f'  Machines: {len(objs)}')

    def migrate_products(self, cur):
        machine_ids = set(Machine.objects.values_list('id', flat=True))
        cur.execute('SELECT * FROM Product')
        rows = [r for r in cur.fetchall() if r['machine_id'] in machine_ids]
        objs = [
            Product(
                id=row['id'],
                machine_id=row['machine_id'],
                name=row['name'],
                product_image=strip_url(row['product_image']),
                description=row['description'],
                product_link=row['product_link'],
                slug=row['slug'],
                **base_fields(row),
            )
            for row in rows
        ]
        Product.objects.bulk_create(objs, ignore_conflicts=True)
        self.stdout.write(f'  Products: {len(objs)}')

    def migrate_spareparts(self, cur):
        supplier_ids = set(Supplier.objects.values_list('id', flat=True))
        cur.execute('SELECT * FROM SpareParts')
        rows = [r for r in cur.fetchall() if r['supplier_id'] in supplier_ids]
        objs = [
            Spareparts(
                id=row['id'],
                supplier_id=row['supplier_id'],
                name=row['name'],
                parts_image=strip_url(row['parts_image']),
                spare_parts_link=row['spare_parts_link'],
                slug=row['slug'],
                **base_fields(row),
            )
            for row in rows
        ]
        Spareparts.objects.bulk_create(objs, ignore_conflicts=True)
        self.stdout.write(f'  Spare parts: {len(objs)}')

    def migrate_videos(self, cur):
        supplier_ids = set(Supplier.objects.values_list('id', flat=True))
        cur.execute('SELECT * FROM Video')
        rows = [r for r in cur.fetchall() if r['supplier_id'] in supplier_ids]
        objs = [
            Video(
                id=row['id'],
                supplier_id=row['supplier_id'],
                web_url=row['web_url'],
                title=row['title'],
                description=row['description'],
                video_id=row['video_id'],
                thumbnail_url=row['thumbnail_url'],
                **base_fields(row),
            )
            for row in rows
        ]
        Video.objects.bulk_create(objs, ignore_conflicts=True)
        self.stdout.write(f'  Videos: {len(objs)}')

    def migrate_blogs(self, cur):
        cur.execute('SELECT * FROM Blog')
        objs = [
            Blog(
                id=row['id'],
                title=row['title'],
                date=parse_date(row['date']),
                main_image=strip_url(row['main_image']),
                subheading=row['subheading'],
                body=row['body'],
                slug=row['slug'],
                **base_fields(row),
            )
            for row in cur.fetchall()
        ]
        Blog.objects.bulk_create(objs, ignore_conflicts=True)
        self.stdout.write(f'  Blogs: {len(objs)}')

    def migrate_carousels(self, cur):
        cur.execute('SELECT * FROM Carousel')
        objs = [
            Carousel(
                id=row['id'],
                name=row['name'],
                image=strip_url(row['image']),
                **base_fields(row),
            )
            for row in cur.fetchall()
        ]
        Carousel.objects.bulk_create(objs, ignore_conflicts=True)
        self.stdout.write(f'  Carousels: {len(objs)}')

    def migrate_exhibitions(self, cur):
        cur.execute('SELECT * FROM Exhibition')
        objs = [
            Exhibition(
                id=row['id'],
                title=row['title'],
                date=parse_date(row['date']),
                location=row['location'],
                info=row['info'],
                **base_fields(row),
            )
            for row in cur.fetchall()
        ]
        Exhibition.objects.bulk_create(objs, ignore_conflicts=True)
        self.stdout.write(f'  Exhibitions: {len(objs)}')

    def migrate_timelines(self, cur):
        cur.execute('SELECT * FROM Timeline')
        objs = [
            Timeline(
                id=row['id'],
                title=row['title'],
                date=parse_date(row['date']),
                body=row['body'],
                **base_fields(row),
            )
            for row in cur.fetchall()
        ]
        Timeline.objects.bulk_create(objs, ignore_conflicts=True)
        self.stdout.write(f'  Timelines: {len(objs)}')

    def migrate_employees(self, cur):
        cur.execute('SELECT * FROM Employee')
        objs = [
            Employee(
                id=row['id'],
                name=row['name'],
                email=row['email'],
                role=row['role'],
                profile_image=strip_url(row['profile_image']),
                **base_fields(row),
            )
            for row in cur.fetchall()
        ]
        Employee.objects.bulk_create(objs, ignore_conflicts=True)
        self.stdout.write(f'  Employees: {len(objs)}')

    def migrate_privacy(self, cur):
        cur.execute('SELECT * FROM Privacy')
        objs = [
            Privacy(
                id=row['id'],
                title=row['title'],
                body=row['body'],
                **base_fields(row),
            )
            for row in cur.fetchall()
        ]
        Privacy.objects.bulk_create(objs, ignore_conflicts=True)
        self.stdout.write(f'  Privacy: {len(objs)}')

    def migrate_terms(self, cur):
        cur.execute('SELECT * FROM Terms')
        objs = [
            Terms(
                id=row['id'],
                title=row['title'],
                body=row['body'],
                **base_fields(row),
            )
            for row in cur.fetchall()
        ]
        Terms.objects.bulk_create(objs, ignore_conflicts=True)
        self.stdout.write(f'  Terms: {len(objs)}')

    def migrate_warrantyclaims(self, cur):
        cur.execute('SELECT * FROM WarrantyClaim')
        objs = [
            Warrantyclaim(
                id=row['id'],
                dealer=row['dealer'],
                dealer_contact=row['dealer_contact'],
                owner_name=row['owner_name'],
                machine_model=row['machine_model'],
                serial_number=row['serial_number'],
                install_date=parse_date(row['install_date']),
                failure_date=parse_date(row['failure_date']),
                repair_date=parse_date(row['repair_date']),
                failure_details=row['failure_details'],
                repair_details=row['repair_details'],
                labour_hours=parse_decimal(row['labour_hours']),
                completed_by=row['completed_by'],
                **base_fields(row),
            )
            for row in cur.fetchall()
        ]
        Warrantyclaim.objects.bulk_create(objs, ignore_conflicts=True)
        self.stdout.write(f'  Warranty claims: {len(objs)}')

    def migrate_partsrequired(self, cur):
        warranty_ids = set(Warrantyclaim.objects.values_list('id', flat=True))
        cur.execute('SELECT * FROM PartsRequired')
        rows = [r for r in cur.fetchall() if r['warranty_id'] in warranty_ids]
        objs = [
            Partsrequired(
                id=row['id'],
                warranty_id=row['warranty_id'],
                part_number=row['part_number'],
                quantity_needed=parse_int(row['quantity_needed']),
                invoice_number=row['invoice_number'],
                description=row['description'],
                **base_fields(row),
            )
            for row in rows
        ]
        Partsrequired.objects.bulk_create(objs, ignore_conflicts=True)
        self.stdout.write(f'  Parts required: {len(objs)}')

    def migrate_machineregistrations(self, cur):
        cur.execute('SELECT * FROM MachineRegistration')
        objs = [
            Machineregistration(
                id=row['id'],
                dealer_name=row['dealer_name'],
                dealer_address=row['dealer_address'],
                owner_name=row['owner_name'],
                owner_address=row['owner_address'],
                machine_model=row['machine_model'],
                serial_number=row['serial_number'],
                install_date=parse_date(row['install_date']),
                invoice_number=row['invoice_number'],
                complete_supply=bool(row['complete_supply']),
                pdi_complete=bool(row['pdi_complete']),
                pto_correct=bool(row['pto_correct']),
                machine_test_run=bool(row['machine_test_run']),
                safety_induction=bool(row['safety_induction']),
                operator_handbook=bool(row['operator_handbook']),
                date=parse_date(row['date']),
                completed_by=row['completed_by'],
                **base_fields(row),
            )
            for row in cur.fetchall()
        ]
        Machineregistration.objects.bulk_create(objs, ignore_conflicts=True)
        self.stdout.write(f'  Machine registrations: {len(objs)}')
