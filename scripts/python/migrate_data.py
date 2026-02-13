import os
import sys
import sqlite3
import uuid
import logging
from datetime import datetime, date
from decimal import Decimal
from typing import Optional, Any
from django.utils import timezone
from django.db import transaction

logger: logging.Logger = logging.getLogger(__name__)


def setup_django() -> None:
    """
    Initialize Django before importing models.

    Sets up Django environment by adding the project path to sys.path,
    configuring the Django settings module, and calling django.setup().

    :raises ImportError: If Django cannot be imported or initialized
    """
    sys.path.insert(0, '/Users/seanwelch/Coding/farmec-v2')
    os.environ.setdefault('DJANGO_SETTINGS_MODULE', 'farmec.settings')
    import django
    django.setup()


def parse_date(date_str: Optional[str]) -> Optional[date]:
    """
    Parse TEXT date to date object.

    Handles ISO format dates (YYYY-MM-DD) and returns None for empty or invalid inputs.

    :param date_str: String representation of date in ISO format
    :return: Parsed date object or None if parsing fails
    """
    if not date_str:
        return None
    try:
        return datetime.fromisoformat(date_str).date()
    except (ValueError, AttributeError, TypeError):
        return None


def parse_datetime(datetime_str: Optional[str]) -> datetime:
    """
    Parse TEXT datetime to datetime object.

    Handles ISO format datetimes and returns current time for empty or invalid inputs.

    :param datetime_str: String representation of datetime in ISO format
    :return: Parsed datetime object or current timezone-aware datetime if parsing fails
    """
    if not datetime_str:
        return timezone.now()
    try:
        dt: datetime = datetime.fromisoformat(datetime_str)
        return dt if dt.tzinfo else timezone.make_aware(dt)
    except (ValueError, AttributeError, TypeError):
        return timezone.now()


def to_bool(value: Optional[Any]) -> bool:
    """
    Convert INTEGER to boolean.

    Converts integer values (0 or 1) to boolean, with None defaulting to False.

    :param value: Integer value to convert
    :return: Boolean representation of the value
    """
    if value is None:
        return False
    try:
        return bool(int(value))
    except (ValueError, TypeError):
        return False


def parse_decimal(value: Optional[Any]) -> Optional[Decimal]:
    """
    Convert value to Decimal with error handling.

    Safely converts numeric values to Decimal, returning None for empty or invalid inputs.

    :param value: Numeric value to convert
    :return: Decimal representation or None if conversion fails
    """
    if not value:
        return None
    try:
        return Decimal(str(value))
    except (ValueError, TypeError, ArithmeticError):
        return None


def migrate_suppliers(old_conn: sqlite3.Connection) -> None:
    """
    Migrate Supplier table from old database to new Django database.

    :param old_conn: SQLite connection to old database
    :raises Exception: If error occurs during individual supplier migration
    """
    from catalog.models import Supplier

    cursor: sqlite3.Cursor = old_conn.cursor()
    cursor.execute("SELECT * FROM Supplier")

    for row in cursor.fetchall():
        try:
            id_val: str
            name: str
            logo_image: Optional[str]
            marketing_image: Optional[str]
            description: Optional[str]
            social_facebook: Optional[str]
            social_twitter: Optional[str]
            social_instagram: Optional[str]
            social_youtube: Optional[str]
            social_linkedin: Optional[str]
            social_website: Optional[str]
            created: Optional[str]
            slug: Optional[str]

            id_val, name, logo_image, marketing_image, description, social_facebook, social_twitter, social_instagram, social_youtube, social_linkedin, social_website, created, slug = row
            if not id_val or not name:
                logger.warning(f"Skipped Supplier (missing required fields): {id_val}")
                continue
            Supplier.objects.get_or_create(
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
                },
            )
        except Exception as e:
            logger.error(f"Error migrating supplier {id_val}: {e}")


def migrate_machines(old_conn: sqlite3.Connection) -> None:
    """
    Migrate Machine table from old database to new Django database.

    :param old_conn: SQLite connection to old database
    :raises Exception: If error occurs during individual machine migration
    """
    from catalog.models import Machine, Supplier

    cursor: sqlite3.Cursor = old_conn.cursor()
    cursor.execute("SELECT * FROM Machine")
    for row in cursor.fetchall():
        try:
            id_val: str
            supplier_id: str
            name: str
            machine_image: Optional[str]
            description: Optional[str]
            machine_link: Optional[str]
            created: Optional[str]
            slug: Optional[str]

            id_val, supplier_id, name, machine_image, description, machine_link, created, slug = row
            if not id_val or not name or not supplier_id:
                continue
            try:
                supplier: Supplier = Supplier.objects.get(id=supplier_id)
            except Supplier.DoesNotExist:
                logger.warning(f"Skipped Machine {id_val}: Supplier {supplier_id} not found")
                continue

            Machine.objects.get_or_create(
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
                },
            )
        except Exception as e:
            logger.error(f"Error migrating machine {id_val}: {e}")


def migrate_products(old_conn: sqlite3.Connection) -> None:
    """
    Migrate Product table from old database to new Django database.

    :param old_conn: SQLite connection to old database
    :raises Exception: If error occurs during individual product migration
    """
    from catalog.models import Product, Machine

    cursor: sqlite3.Cursor = old_conn.cursor()
    cursor.execute("SELECT * FROM Product")
    for row in cursor.fetchall():
        try:
            id_val: str
            machine_id: str
            name: str
            product_image: Optional[str]
            description: Optional[str]
            product_link: Optional[str]
            slug: Optional[str]

            id_val, machine_id, name, product_image, description, product_link, slug = row
            if not id_val or not name or not machine_id:
                continue

            try:
                machine: Machine = Machine.objects.get(id=machine_id)
            except Machine.DoesNotExist:
                logger.warning(f"Skipped Product {id_val}: Machine {machine_id} not found")
                continue

            Product.objects.get_or_create(
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
                },
            )
        except Exception as e:
            logger.error(f"Error migrating product {id_val}: {e}")


def migrate_spare_parts(old_conn: sqlite3.Connection) -> None:
    """
    Migrate SpareParts table from old database to new Django database.

    :param old_conn: SQLite connection to old database
    :raises Exception: If error occurs during individual spare part migration
    """
    from catalog.models import Spareparts, Supplier

    cursor: sqlite3.Cursor = old_conn.cursor()
    cursor.execute("SELECT * FROM SpareParts")

    for row in cursor.fetchall():
        try:
            id_val: str
            supplier_id: str
            name: str
            parts_image: Optional[str]
            spare_parts_link: Optional[str]
            slug: Optional[str]

            id_val, supplier_id, name, parts_image, spare_parts_link, slug = row
            if not id_val or not name or not supplier_id:
                continue
            try:
                supplier: Supplier = Supplier.objects.get(id=supplier_id)
            except Supplier.DoesNotExist:
                logger.warning(f"Skipped SpareParts {id_val}: Supplier {supplier_id} not found")
                continue
            Spareparts.objects.get_or_create(
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
                },
            )
        except Exception as e:
            logger.error(f"Error migrating spare part {id_val}: {e}")


def migrate_line_items(old_conn: sqlite3.Connection) -> None:
    """
    Migrate LineItems table from old database to new Django database.

    :param old_conn: SQLite connection to old database
    :raises Exception: If error occurs during individual line item migration
    """
    from catalog.models import Lineitems

    cursor: sqlite3.Cursor = old_conn.cursor()
    cursor.execute("SELECT * FROM LineItems")

    for row in cursor.fetchall():
        try:
            id_val: str
            name: str
            price: float
            image: Optional[str]

            id_val, name, price, image = row

            if not id_val or not name:
                continue

            Lineitems.objects.get_or_create(
                id=id_val,
                defaults={
                    'name': name,
                    'price': parse_decimal(price) or Decimal('0.00'),
                    'image': image,
                    'order': 1,
                    'publish': True,
                    'uid': uuid.uuid4(),
                    'created': timezone.now(),
                },
            )
        except Exception as e:
            logger.error(f"Error migrating line item {id_val}: {e}")


def migrate_videos(old_conn: sqlite3.Connection) -> None:
    """
    Migrate Video table from old database to new Django database.

    :param old_conn: SQLite connection to old database
    :raises Exception: If error occurs during individual video migration
    """
    from catalog.models import Video, Supplier

    cursor: sqlite3.Cursor = old_conn.cursor()
    cursor.execute("SELECT * FROM Video")

    for row in cursor.fetchall():
        try:
            id_val: str
            supplier_id: str
            web_url: Optional[str]
            title: Optional[str]
            description: Optional[str]
            video_id: Optional[str]
            thumbnail_url: Optional[str]
            created: Optional[str]

            id_val, supplier_id, web_url, title, description, video_id, thumbnail_url, created = row

            if not id_val or not supplier_id:
                continue

            try:
                supplier: Supplier = Supplier.objects.get(id=supplier_id)
            except Supplier.DoesNotExist:
                logger.warning(f"Skipped Video {id_val}: Supplier {supplier_id} not found")
                continue

            Video.objects.get_or_create(
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
                },
            )
        except Exception as e:
            logger.error(f"Error migrating video {id_val}: {e}")


def migrate_blogs(old_conn: sqlite3.Connection) -> None:
    """
    Migrate Blog table from old database to new Django database.

    :param old_conn: SQLite connection to old database
    :raises Exception: If error occurs during individual blog migration
    """
    from content.models import Blog

    cursor: sqlite3.Cursor = old_conn.cursor()
    cursor.execute("SELECT * FROM Blog")

    for row in cursor.fetchall():
        try:
            id_val: str
            title: str
            date_str: Optional[str]
            main_image: Optional[str]
            subheading: Optional[str]
            body: Optional[str]
            created: Optional[str]
            slug: Optional[str]

            id_val, title, date_str, main_image, subheading, body, created, slug = row

            if not id_val or not title:
                continue

            Blog.objects.get_or_create(
                id=id_val,
                defaults={
                    'title': title,
                    'date': parse_date(date_str),
                    'main_image': main_image,
                    'subheading': subheading,
                    'body': body,
                    'slug': slug,
                    'order': 1,
                    'publish': True,
                    'uid': uuid.uuid4(),
                    'created': parse_datetime(created),
                },
            )
        except Exception as e:
            logger.error(f"Error migrating blog {id_val}: {e}")


def migrate_carousel(old_conn: sqlite3.Connection) -> None:
    """
    Migrate Carousel table from old database to new Django database.

    :param old_conn: SQLite connection to old database
    :raises Exception: If error occurs during individual carousel migration
    """
    from content.models import Carousel

    cursor: sqlite3.Cursor = old_conn.cursor()
    cursor.execute("SELECT * FROM Carousel")

    for row in cursor.fetchall():
        try:
            id_val: str
            name: str
            image: Optional[str]
            created: Optional[str]

            id_val, name, image, created = row

            if not id_val or not name:
                continue

            Carousel.objects.get_or_create(
                id=id_val,
                defaults={
                    'name': name,
                    'image': image,
                    'order': 1,
                    'publish': True,
                    'uid': uuid.uuid4(),
                    'created': parse_datetime(created),
                },
            )
        except Exception as e:
            logger.error(f"Error migrating carousel {id_val}: {e}")


def migrate_exhibitions(old_conn: sqlite3.Connection) -> None:
    """
    Migrate Exhibition table from old database to new Django database.

    :param old_conn: SQLite connection to old database
    :raises Exception: If error occurs during individual exhibition migration
    """
    from content.models import Exhibition

    cursor: sqlite3.Cursor = old_conn.cursor()
    cursor.execute("SELECT * FROM Exhibition")

    for row in cursor.fetchall():
        try:
            id_val: str
            title: str
            date_str: Optional[str]
            location: Optional[str]
            info: Optional[str]
            created: Optional[str]

            id_val, title, date_str, location, info, created = row

            if not id_val or not title:
                continue

            Exhibition.objects.get_or_create(
                id=id_val,
                defaults={
                    'title': title,
                    'date': parse_date(date_str),
                    'location': location,
                    'info': info,
                    'order': 1,
                    'publish': True,
                    'uid': uuid.uuid4(),
                    'created': parse_datetime(created),
                },
            )
        except Exception as e:
            logger.error(f"Error migrating exhibition {id_val}: {e}")


def migrate_timelines(old_conn: sqlite3.Connection) -> None:
    """
    Migrate Timeline table from old database to new Django database.

    :param old_conn: SQLite connection to old database
    :raises Exception: If error occurs during individual timeline migration
    """
    from content.models import Timeline

    cursor: sqlite3.Cursor = old_conn.cursor()
    cursor.execute("SELECT * FROM Timeline")

    for row in cursor.fetchall():
        try:
            id_val: str
            title: str
            date_str: Optional[str]
            body: Optional[str]
            created: Optional[str]

            id_val, title, date_str, body, created = row

            if not id_val or not title:
                continue

            Timeline.objects.get_or_create(
                id=id_val,
                defaults={
                    'title': title,
                    'date': parse_date(date_str),
                    'body': body,
                    'order': 1,
                    'publish': True,
                    'uid': uuid.uuid4(),
                    'created': parse_datetime(created),
                },
            )
        except Exception as e:
            logger.error(f"Error migrating timeline {id_val}: {e}")


def migrate_employees(old_conn: sqlite3.Connection) -> None:
    """
    Migrate Employee table from old database to new Django database.

    :param old_conn: SQLite connection to old database
    :raises Exception: If error occurs during individual employee migration
    """
    from team.models import Employee

    cursor: sqlite3.Cursor = old_conn.cursor()
    cursor.execute("SELECT * FROM Employee")

    for row in cursor.fetchall():
        try:
            id_val: str
            name: str
            email: str
            role: str
            profile_image: Optional[str]
            created: Optional[str]

            id_val, name, email, role, profile_image, created = row

            if not id_val or not name or not email:
                continue

            Employee.objects.get_or_create(
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
                },
            )
        except Exception as e:
            logger.error(f"Error migrating employee {id_val}: {e}")


def migrate_warranty_claims(old_conn: sqlite3.Connection) -> None:
    """
    Migrate WarrantyClaim table from old database to new Django database.

    :param old_conn: SQLite connection to old database
    :raises Exception: If error occurs during individual warranty claim migration
    """
    from support.models import Warrantyclaim

    cursor: sqlite3.Cursor = old_conn.cursor()
    cursor.execute("SELECT * FROM WarrantyClaim")

    for row in cursor.fetchall():
        try:
            id_val: str
            dealer: str
            dealer_contact: Optional[str]
            owner_name: str
            machine_model: str
            serial_number: str
            install_date: Optional[str]
            failure_date: Optional[str]
            repair_date: Optional[str]
            failure_details: Optional[str]
            repair_details: Optional[str]
            labour_hours: Optional[str]
            completed_by: Optional[str]
            created: Optional[str]

            id_val, dealer, dealer_contact, owner_name, machine_model, serial_number, \
                install_date, failure_date, repair_date, failure_details, repair_details, \
                labour_hours, completed_by, created = row

            if not id_val or not dealer or not owner_name or not machine_model or not serial_number:
                continue

            Warrantyclaim.objects.get_or_create(
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
                },
            )
        except Exception as e:
            logger.error(f"Error migrating warranty claim {id_val}: {e}")


def migrate_parts_required(old_conn: sqlite3.Connection) -> None:
    """
    Migrate PartsRequired table from old database to new Django database.

    :param old_conn: SQLite connection to old database
    :raises Exception: If error occurs during individual parts required migration
    """
    from support.models import Partsrequired, Warrantyclaim

    cursor: sqlite3.Cursor = old_conn.cursor()
    cursor.execute("SELECT * FROM PartsRequired")

    for row in cursor.fetchall():
        try:
            id_val: str
            warranty_id: str
            part_number: Optional[str]
            quantity_needed: str
            invoice_number: Optional[str]
            description: Optional[str]

            id_val, warranty_id, part_number, quantity_needed, invoice_number, description = row

            if not id_val or not warranty_id or not quantity_needed:
                continue

            try:
                warranty: Warrantyclaim = Warrantyclaim.objects.get(id=warranty_id)
            except Warrantyclaim.DoesNotExist:
                logger.warning(f"Skipped PartsRequired {id_val}: Warranty {warranty_id} not found")
                continue

            try:
                quantity: int = int(quantity_needed)
                if quantity < 0:
                    continue
            except (ValueError, TypeError):
                continue

            Partsrequired.objects.get_or_create(
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
                },
            )
        except Exception as e:
            logger.error(f"Error migrating parts required {id_val}: {e}")


def migrate_machine_registrations(old_conn: sqlite3.Connection) -> None:
    """
    Migrate MachineRegistration table from old database to new Django database.

    :param old_conn: SQLite connection to old database
    :raises Exception: If error occurs during individual machine registration migration
    """
    from support.models import Machineregistration

    cursor: sqlite3.Cursor = old_conn.cursor()
    cursor.execute("SELECT * FROM MachineRegistration")

    for row in cursor.fetchall():
        try:
            id_val: str
            dealer_name: str
            dealer_address: Optional[str]
            owner_name: str
            owner_address: Optional[str]
            machine_model: str
            serial_number: str
            install_date: Optional[str]
            invoice_number: Optional[str]
            complete_supply: Optional[int]
            pdi_complete: Optional[int]
            pto_correct: Optional[int]
            machine_test_run: Optional[int]
            safety_induction: Optional[int]
            operator_handbook: Optional[int]
            date_str: Optional[str]
            completed_by: Optional[str]
            created: Optional[str]

            id_val, dealer_name, dealer_address, owner_name, owner_address, machine_model, \
                serial_number, install_date, invoice_number, complete_supply, pdi_complete, \
                pto_correct, machine_test_run, safety_induction, operator_handbook, date_str, \
                completed_by, created = row

            if not id_val or not dealer_name or not owner_name or not machine_model or not serial_number:
                continue

            Machineregistration.objects.get_or_create(
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
                    'date': parse_date(date_str),
                    'completed_by': completed_by,
                    'order': 1,
                    'publish': True,
                    'uid': uuid.uuid4(),
                    'created': parse_datetime(created),
                },
            )
        except Exception as e:
            logger.error(f"Error migrating machine registration {id_val}: {e}")


def migrate_privacy(old_conn: sqlite3.Connection) -> None:
    """
    Migrate Privacy table from old database to new Django database.

    :param old_conn: SQLite connection to old database
    :raises Exception: If error occurs during individual privacy migration
    """
    from legal.models import Privacy

    cursor: sqlite3.Cursor = old_conn.cursor()
    cursor.execute("SELECT * FROM Privacy")

    for row in cursor.fetchall():
        try:
            id_val: str
            title: str
            body: Optional[str]
            created: Optional[str]

            id_val, title, body, created = row

            if not id_val or not title:
                continue

            Privacy.objects.get_or_create(
                id=id_val,
                defaults={
                    'title': title,
                    'body': body,
                    'order': 1,
                    'publish': True,
                    'uid': uuid.uuid4(),
                    'created': parse_datetime(created),
                },
            )
        except Exception as e:
            logger.error(f"Error migrating privacy policy {id_val}: {e}")


def migrate_terms(old_conn: sqlite3.Connection) -> None:
    """
    Migrate Terms table from old database to new Django database.

    :param old_conn: SQLite connection to old database
    :raises Exception: If error occurs during individual terms migration
    """
    from legal.models import Terms

    cursor: sqlite3.Cursor = old_conn.cursor()
    cursor.execute("SELECT * FROM Terms")

    for row in cursor.fetchall():
        try:
            id_val: str
            title: str
            body: Optional[str]
            created: Optional[str]

            id_val, title, body, created = row

            if not id_val or not title:
                continue

            Terms.objects.get_or_create(
                id=id_val,
                defaults={
                    'title': title,
                    'body': body,
                    'order': 1,
                    'publish': True,
                    'uid': uuid.uuid4(),
                    'created': parse_datetime(created),
                },
            )
        except Exception as e:
            logger.error(f"Error migrating terms {id_val}: {e}")


def main() -> None:
    """
    Main migration function with transactional support.

    Orchestrates the migration of all tables from old to new database.
    All operations are wrapped in a transaction for atomicity.

    :raises SystemExit: If database not found or migration fails
    """
    setup_django()

    old_db_path: str = '/Users/seanwelch/Coding/farmec-v2/server/database/database.db'

    if not os.path.exists(old_db_path):
        logger.error(f"Old database not found at {old_db_path}")
        sys.exit(1)

    logger.info("Starting migration")

    old_conn: sqlite3.Connection = sqlite3.connect(old_db_path)
    old_conn.row_factory = sqlite3.Row

    try:
        with transaction.atomic():
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

            logger.info("Migration completed successfully")

    except Exception as e:
        logger.exception(f"Migration failed: {e}")
        sys.exit(1)
    finally:
        old_conn.close()


if __name__ == '__main__':
    logging.basicConfig(level=logging.INFO, format='%(asctime)s - %(levelname)s - %(message)s')
    main()
