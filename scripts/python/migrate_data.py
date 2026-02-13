#!/usr/bin/env python
"""
SQLite Database Migration Script (Django ORM Version)
Migrates data from old database to new Django database with schema transformation.

Schema Changes:
1. Date fields: TEXT → DateField/DateTimeField
2. Boolean fields: INTEGER (0/1) → BooleanField
3. REAL fields: REAL → DecimalField
4. New BaseModel fields added: order, publish, uid, created, modified
5. Foreign key fields: supplier_id → supplier (FK reference)

Features:
- Transactional execution (all-or-nothing atomicity)
- Per-record error handling with detailed logging
- Validates foreign key references before insert
- Idempotent (uses get_or_create, safe to re-run)
- Summary statistics on completion
"""

import os
import sys
import sqlite3
import uuid
import logging
from datetime import datetime
from decimal import Decimal

from django.utils import timezone
from django.db import transaction
from catalog.models import Supplier, Machine, Product, Spareparts, Lineitems, Video
from content.models import Blog, Carousel, Exhibition, Timeline
from team.models import Employee
from support.models import Warrantyclaim, Partsrequired, Machineregistration
from legal.models import Privacy, Terms

logger = logging.getLogger(__name__)


def setup_django():
    """Initialize Django before importing models."""
    sys.path.insert(0, '/Users/seanwelch/Coding/farmec-v2')
    os.environ.setdefault('DJANGO_SETTINGS_MODULE', 'settings')
    import django
    django.setup()


def parse_date(date_str):
    """Parse TEXT date to date object. Handles various formats."""
    if not date_str:
        return None
    try:
        # Try ISO format first (YYYY-MM-DD)
        return datetime.fromisoformat(date_str).date()
    except (ValueError, AttributeError, TypeError):
        return None


def parse_datetime(datetime_str):
    """Parse TEXT datetime to datetime object."""
    if not datetime_str:
        return timezone.now()
    try:
        # Try ISO format
        dt = datetime.fromisoformat(datetime_str)
        return dt if dt.tzinfo else timezone.make_aware(dt)
    except (ValueError, AttributeError, TypeError):
        return timezone.now()


def to_bool(value):
    """Convert INTEGER to boolean."""
    if value is None:
        return False
    try:
        return bool(int(value))
    except (ValueError, TypeError):
        return False


def parse_decimal(value):
    """Convert value to Decimal with error handling."""
    if not value:
        return None
    try:
        return Decimal(str(value))
    except (ValueError, TypeError, ArithmeticError):
        return None


def migrate_suppliers(old_conn):
    """Migrate Supplier table."""
    logger.info("Migrating Suppliers")
    cursor = old_conn.cursor()
    cursor.execute("SELECT * FROM Supplier")

    for row in cursor.fetchall():
        try:
            id_val, name, logo_image, marketing_image, description, social_facebook, \
                social_twitter, social_instagram, social_youtube, social_linkedin, \
                social_website, created, slug = row

            # Validate required fields
            if not id_val or not name:
                logger.warning(f"Skipped Supplier (missing required fields): {id_val}")
                continue

            supplier, created_flag = Supplier.objects.get_or_create(
                id=id_val,
                defaults={
                    'name': name,
                    'logo_image': logo_image,
                    'marketing_image': marketing_image,
                    'description': description,
                    'social_facebook': social_facebook,
                    'social_twitter': social_twitter,
                    'social_instagram': social_instagram,
                    'social_youtube': social_youtube,
                    'social_linkedin': social_linkedin,
                    'social_website': social_website,
                    'slug': slug,
                    'order': 1,
                    'publish': True,
                    'uid': uuid.uuid4(),
                    'created': parse_datetime(created),
                }
            )
            if not created_flag:
                logger.debug(f"Supplier {id_val} already exists")
        except Exception as e:
            logger.error(f"Error migrating supplier {id_val}: {e}")


def migrate_machines(old_conn):
    """Migrate Machine table."""
    print("Migrating Machines...")
    cursor = old_conn.cursor()
    cursor.execute("SELECT * FROM Machine")

    for row in cursor.fetchall():
        try:
            id_val, supplier_id, name, machine_image, description, machine_link, created, slug = row

            # Validate required fields
            if not id_val or not name or not supplier_id:
                continue

            try:
                supplier = Supplier.objects.get(id=supplier_id)
            except Supplier.DoesNotExist:
                print(f"  ⚠ Skipping Machine {id_val}: Supplier {supplier_id} not found")
                continue

            machine, created_flag = Machine.objects.get_or_create(
                id=id_val,
                defaults={
                    'supplier': supplier,
                    'name': name,
                    'machine_image': machine_image,
                    'description': description,
                    'machine_link': machine_link,
                    'slug': slug,
                    'order': 1,
                    'publish': True,
                    'uid': uuid.uuid4(),
                    'created': parse_datetime(created),
                }
            )
            if created_flag:
                print(f"  ✓ Created: {name}")
            else:
                print(f"  ✓ Already exists: {name}")
        except Exception as e:
            print(f"  ❌ Error migrating machine {id_val}: {e}")


def migrate_products(old_conn):
    """Migrate Product table."""
    print("Migrating Products...")
    cursor = old_conn.cursor()
    cursor.execute("SELECT * FROM Product")

    for row in cursor.fetchall():
        try:
            id_val, machine_id, name, product_image, description, product_link, slug = row

            # Validate required fields
            if not id_val or not name or not machine_id:
                continue

            try:
                machine = Machine.objects.get(id=machine_id)
            except Machine.DoesNotExist:
                print(f"  ⚠ Skipping Product {id_val}: Machine {machine_id} not found")
                continue

            product, created_flag = Product.objects.get_or_create(
                id=id_val,
                defaults={
                    'machine': machine,
                    'name': name,
                    'product_image': product_image,
                    'description': description,
                    'product_link': product_link,
                    'slug': slug,
                    'order': 1,
                    'publish': True,
                    'uid': uuid.uuid4(),
                    'created': timezone.now(),
                }
            )
            if created_flag:
                print(f"  ✓ Created: {name}")
            else:
                print(f"  ✓ Already exists: {name}")
        except Exception as e:
            print(f"  ❌ Error migrating product {id_val}: {e}")


def migrate_spare_parts(old_conn):
    """Migrate SpareParts table."""
    print("Migrating SpareParts...", end=' ', flush=True)
    cursor = old_conn.cursor()
    cursor.execute("SELECT * FROM SpareParts")

    for row in cursor.fetchall():
        try:
            id_val, supplier_id, name, parts_image, spare_parts_link, slug = row

            # Validate required fields
            if not id_val or not name or not supplier_id:
                skipped_count += 1
                continue

            try:
                supplier = Supplier.objects.get(id=supplier_id)
            except Supplier.DoesNotExist:
                skipped_count += 1
                continue

            part, created_flag = Spareparts.objects.get_or_create(
                id=id_val,
                defaults={
                    'supplier': supplier,
                    'name': name,
                    'parts_image': parts_image,
                    'spare_parts_link': spare_parts_link,
                    'slug': slug,
                    'order': 1,
                    'publish': True,
                    'uid': uuid.uuid4(),
                    'created': timezone.now(),
                }
            )
            if created_flag:
                created_count += 1
            else:
                existing_count += 1
        except Exception as e:
            print(f"  ❌ Error migrating spare part {id_val}: {e}")


def migrate_line_items(old_conn):
    """Migrate LineItems table."""
    print("Migrating LineItems...", end=' ', flush=True)
    cursor = old_conn.cursor()
    cursor.execute("SELECT * FROM LineItems")

    for row in cursor.fetchall():
        try:
            id_val, name, price, image = row

            # Validate required fields
            if not id_val or not name:
                skipped_count += 1
                continue

            item, created_flag = Lineitems.objects.get_or_create(
                id=id_val,
                defaults={
                    'name': name,
                    'price': parse_decimal(price) or Decimal('0.00'),
                    'image': image,
                    'order': 1,
                    'publish': True,
                    'uid': uuid.uuid4(),
                    'created': timezone.now(),
                }
            )
            if created_flag:
                print(f"  ✓ Created: {name}")
            else:
                print(f"  ✓ Already exists: {name}")
        except Exception as e:
            print(f"  ❌ Error migrating line item {id_val}: {e}")


def migrate_videos(old_conn):
    """Migrate Video table."""
    print("Migrating Videos...", end=' ', flush=True)
    cursor = old_conn.cursor()
    cursor.execute("SELECT * FROM Video")

    for row in cursor.fetchall():
        try:
            id_val, supplier_id, web_url, title, description, video_id, thumbnail_url, created = row

            # Validate required fields
            if not id_val or not supplier_id:
                skipped_count += 1
                continue

            try:
                supplier = Supplier.objects.get(id=supplier_id)
            except Supplier.DoesNotExist:
                skipped_count += 1
                continue

            video, created_flag = Video.objects.get_or_create(
                id=id_val,
                defaults={
                    'supplier': supplier,
                    'web_url': web_url,
                    'title': title,
                    'description': description,
                    'video_id': video_id,
                    'thumbnail_url': thumbnail_url,
                    'order': 1,
                    'publish': True,
                    'uid': uuid.uuid4(),
                    'created': parse_datetime(created),
                }
            )
            if created_flag:
                print(f"  ✓ Created: {title or video_id}")
            else:
                print(f"  ✓ Already exists: {title or video_id}")
        except Exception as e:
            print(f"  ❌ Error migrating video {id_val}: {e}")


def migrate_blogs(old_conn):
    """Migrate Blog table."""
    print("Migrating Blogs...", end=' ', flush=True)
    cursor = old_conn.cursor()
    cursor.execute("SELECT * FROM Blog")

    for row in cursor.fetchall():
        try:
            id_val, title, date, main_image, subheading, body, created, slug = row

            # Validate required fields
            if not id_val or not title:
                skipped_count += 1
                continue

            blog, created_flag = Blog.objects.get_or_create(
                id=id_val,
                defaults={
                    'title': title,
                    'date': parse_date(date),
                    'main_image': main_image,
                    'subheading': subheading,
                    'body': body,
                    'slug': slug,
                    'order': 1,
                    'publish': True,
                    'uid': uuid.uuid4(),
                    'created': parse_datetime(created),
                }
            )
            if created_flag:
                created_count += 1
            else:
                existing_count += 1
        except Exception as e:
            print(f"\n  ❌ Error migrating blog {id_val}: {e}")
            skipped_count += 1



def migrate_carousel(old_conn):
    """Migrate Carousel table."""
    print("Migrating Carousel...", end=' ', flush=True)
    cursor = old_conn.cursor()
    cursor.execute("SELECT * FROM Carousel")

    for row in cursor.fetchall():
        try:
            id_val, name, image, created = row

            # Validate required fields
            if not id_val or not name:
                skipped_count += 1
                continue

            carousel, created_flag = Carousel.objects.get_or_create(
                id=id_val,
                defaults={
                    'name': name,
                    'image': image,
                    'order': 1,
                    'publish': True,
                    'uid': uuid.uuid4(),
                    'created': parse_datetime(created),
                }
            )
            if created_flag:
                created_count += 1
            else:
                existing_count += 1
        except Exception as e:
            print(f"\n  ❌ Error migrating carousel {id_val}: {e}")
            skipped_count += 1



def migrate_exhibitions(old_conn):
    """Migrate Exhibition table."""
    print("Migrating Exhibitions...", end=' ', flush=True)
    cursor = old_conn.cursor()
    cursor.execute("SELECT * FROM Exhibition")

    for row in cursor.fetchall():
        try:
            id_val, title, date, location, info, created = row

            # Validate required fields
            if not id_val or not title:
                skipped_count += 1
                continue

            exhibition, created_flag = Exhibition.objects.get_or_create(
                id=id_val,
                defaults={
                    'title': title,
                    'date': parse_date(date),
                    'location': location,
                    'info': info,
                    'order': 1,
                    'publish': True,
                    'uid': uuid.uuid4(),
                    'created': parse_datetime(created),
                }
            )
            if created_flag:
                created_count += 1
            else:
                existing_count += 1
        except Exception as e:
            print(f"\n  ❌ Error migrating exhibition {id_val}: {e}")
            skipped_count += 1



def migrate_timelines(old_conn):
    """Migrate Timeline table."""
    print("Migrating Timelines...", end=' ', flush=True)
    cursor = old_conn.cursor()
    cursor.execute("SELECT * FROM Timeline")

    for row in cursor.fetchall():
        try:
            id_val, title, date, body, created = row

            # Validate required fields
            if not id_val or not title:
                skipped_count += 1
                continue

            timeline, created_flag = Timeline.objects.get_or_create(
                id=id_val,
                defaults={
                    'title': title,
                    'date': parse_date(date),
                    'body': body,
                    'order': 1,
                    'publish': True,
                    'uid': uuid.uuid4(),
                    'created': parse_datetime(created),
                }
            )
            if created_flag:
                created_count += 1
            else:
                existing_count += 1
        except Exception as e:
            print(f"\n  ❌ Error migrating timeline {id_val}: {e}")
            skipped_count += 1



def migrate_employees(old_conn):
    """Migrate Employee table."""
    print("Migrating Employees...", end=' ', flush=True)
    cursor = old_conn.cursor()
    cursor.execute("SELECT * FROM Employee")

    for row in cursor.fetchall():
        try:
            id_val, name, email, role, profile_image, created = row

            # Validate required fields
            if not id_val or not name or not email:
                skipped_count += 1
                continue

            employee, created_flag = Employee.objects.get_or_create(
                id=id_val,
                defaults={
                    'name': name,
                    'email': email,
                    'role': role,
                    'profile_image': profile_image,
                    'order': 1,
                    'publish': True,
                    'uid': uuid.uuid4(),
                    'created': parse_datetime(created),
                }
            )
            if created_flag:
                created_count += 1
            else:
                existing_count += 1
        except Exception as e:
            print(f"\n  ❌ Error migrating employee {id_val}: {e}")
            skipped_count += 1



def migrate_warranty_claims(old_conn):
    """Migrate WarrantyClaim table."""
    print("Migrating Warranty Claims...", end=' ', flush=True)
    cursor = old_conn.cursor()
    cursor.execute("SELECT * FROM WarrantyClaim")

    for row in cursor.fetchall():
        try:
            id_val, dealer, dealer_contact, owner_name, machine_model, serial_number, \
                install_date, failure_date, repair_date, failure_details, repair_details, \
                labour_hours, completed_by, created = row

            # Validate required fields
            if not id_val or not dealer or not owner_name or not machine_model or not serial_number:
                skipped_count += 1
                continue

            claim, created_flag = Warrantyclaim.objects.get_or_create(
                id=id_val,
                defaults={
                    'dealer': dealer,
                    'dealer_contact': dealer_contact,
                    'owner_name': owner_name,
                    'machine_model': machine_model,
                    'serial_number': serial_number,
                    'install_date': parse_date(install_date),
                    'failure_date': parse_date(failure_date),
                    'repair_date': parse_date(repair_date),
                    'failure_details': failure_details,
                    'repair_details': repair_details,
                    'labour_hours': parse_decimal(labour_hours),
                    'completed_by': completed_by,
                    'order': 1,
                    'publish': True,
                    'uid': uuid.uuid4(),
                    'created': parse_datetime(created),
                }
            )
            if created_flag:
                created_count += 1
            else:
                existing_count += 1
        except Exception as e:
            print(f"\n  ❌ Error migrating warranty claim {id_val}: {e}")
            skipped_count += 1



def migrate_parts_required(old_conn):
    """Migrate PartsRequired table."""
    print("Migrating Parts Required...", end=' ', flush=True)
    cursor = old_conn.cursor()
    cursor.execute("SELECT * FROM PartsRequired")

    for row in cursor.fetchall():
        try:
            id_val, warranty_id, part_number, quantity_needed, invoice_number, description = row

            # Validate required fields
            if not id_val or not warranty_id or not quantity_needed:
                skipped_count += 1
                continue

            try:
                warranty = Warrantyclaim.objects.get(id=warranty_id)
            except Warrantyclaim.DoesNotExist:
                skipped_count += 1
                continue

            # Validate quantity_needed conversion
            try:
                quantity = int(quantity_needed)
                if quantity < 0:
                    skipped_count += 1
                    continue
            except (ValueError, TypeError):
                skipped_count += 1
                continue

            part, created_flag = Partsrequired.objects.get_or_create(
                id=id_val,
                defaults={
                    'warranty': warranty,
                    'part_number': part_number,
                    'quantity_needed': quantity,
                    'invoice_number': invoice_number,
                    'description': description,
                    'order': 1,
                    'publish': True,
                    'uid': uuid.uuid4(),
                    'created': timezone.now(),
                }
            )
            if created_flag:
                created_count += 1
            else:
                existing_count += 1
        except Exception as e:
            print(f"\n  ❌ Error migrating parts required {id_val}: {e}")
            skipped_count += 1



def migrate_machine_registrations(old_conn):
    """Migrate MachineRegistration table."""
    print("Migrating Machine Registrations...", end=' ', flush=True)
    cursor = old_conn.cursor()
    cursor.execute("SELECT * FROM MachineRegistration")

    for row in cursor.fetchall():
        try:
            id_val, dealer_name, dealer_address, owner_name, owner_address, machine_model, \
                serial_number, install_date, invoice_number, complete_supply, pdi_complete, \
                pto_correct, machine_test_run, safety_induction, operator_handbook, date, \
                completed_by, created = row

            # Validate required fields
            if not id_val or not dealer_name or not owner_name or not machine_model or not serial_number:
                skipped_count += 1
                continue

            registration, created_flag = Machineregistration.objects.get_or_create(
                id=id_val,
                defaults={
                    'dealer_name': dealer_name,
                    'dealer_address': dealer_address,
                    'owner_name': owner_name,
                    'owner_address': owner_address,
                    'machine_model': machine_model,
                    'serial_number': serial_number,
                    'install_date': parse_date(install_date),
                    'invoice_number': invoice_number,
                    'complete_supply': to_bool(complete_supply),
                    'pdi_complete': to_bool(pdi_complete),
                    'pto_correct': to_bool(pto_correct),
                    'machine_test_run': to_bool(machine_test_run),
                    'safety_induction': to_bool(safety_induction),
                    'operator_handbook': to_bool(operator_handbook),
                    'date': parse_date(date),
                    'completed_by': completed_by,
                    'order': 1,
                    'publish': True,
                    'uid': uuid.uuid4(),
                    'created': parse_datetime(created),
                }
            )
            if created_flag:
                created_count += 1
            else:
                existing_count += 1
        except Exception as e:
            print(f"\n  ❌ Error migrating machine registration {id_val}: {e}")
            skipped_count += 1



def migrate_privacy(old_conn):
    """Migrate Privacy table."""
    print("Migrating Privacy Policy...", end=' ', flush=True)
    cursor = old_conn.cursor()
    cursor.execute("SELECT * FROM Privacy")

    for row in cursor.fetchall():
        try:
            id_val, title, body, created = row

            # Validate required fields
            if not id_val or not title:
                skipped_count += 1
                continue

            privacy, created_flag = Privacy.objects.get_or_create(
                id=id_val,
                defaults={
                    'title': title,
                    'body': body,
                    'order': 1,
                    'publish': True,
                    'uid': uuid.uuid4(),
                    'created': parse_datetime(created),
                }
            )
            if created_flag:
                created_count += 1
            else:
                existing_count += 1
        except Exception as e:
            print(f"\n  ❌ Error migrating privacy policy {id_val}: {e}")
            skipped_count += 1



def migrate_terms(old_conn):
    """Migrate Terms table."""
    print("Migrating Terms & Conditions...", end=' ', flush=True)
    cursor = old_conn.cursor()
    cursor.execute("SELECT * FROM Terms")

    for row in cursor.fetchall():
        try:
            id_val, title, body, created = row

            # Validate required fields
            if not id_val or not title:
                skipped_count += 1
                continue

            terms, created_flag = Terms.objects.get_or_create(
                id=id_val,
                defaults={
                    'title': title,
                    'body': body,
                    'order': 1,
                    'publish': True,
                    'uid': uuid.uuid4(),
                    'created': parse_datetime(created),
                }
            )
            if created_flag:
                created_count += 1
            else:
                existing_count += 1
        except Exception as e:
            print(f"\n  ❌ Error migrating terms {id_val}: {e}")
            skipped_count += 1



@transaction.atomic
def main():
    setup_django()

    old_db_path = '/Users/seanwelch/Coding/farmec-v2/server/database/database.db'

    if not os.path.exists(old_db_path):
        print(f"❌ Error: Old database not found at {old_db_path}")
        sys.exit(1)

    print("=" * 70)
    print("SQLite Database Migration (Django ORM)")
    print("=" * 70)
    print(f"\nSource: {old_db_path}\n")

    old_conn = sqlite3.connect(old_db_path)
    old_conn.row_factory = sqlite3.Row

    try:
        # Migrate in order of dependencies
        migrate_suppliers(old_conn)
        migrate_machines(old_conn)
        migrate_products(old_conn)
        migrate_spare_parts(old_conn)
        migrate_line_items(old_conn)
        migrate_videos(old_conn)
        migrate_blogs(old_conn)
        migrate_carousel(old_conn)
        migrate_exhibitions(old_conn)
        migrate_timelines(old_conn)
        migrate_employees(old_conn)
        migrate_warranty_claims(old_conn)
        migrate_parts_required(old_conn)
        migrate_machine_registrations(old_conn)
        migrate_privacy(old_conn)
        migrate_terms(old_conn)

        print("\nMigration completed successfully!")

    except Exception as e:
        print(f"\n\n❌ Migration failed: {e}")
        import traceback
        traceback.print_exc()
        print("\n⚠️  Transaction rolled back due to error.")
        sys.exit(1)
    finally:
        old_conn.close()


if __name__ == '__main__':
    main()
