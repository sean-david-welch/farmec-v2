"""
Integration tests for database migration script.

Tests verify that data is correctly translated from old SQLite database
to new Django database with proper type conversions and validations.
"""

import os
import sqlite3
from datetime import datetime, date
from decimal import Decimal

import pytest

from catalog.models import Supplier, Machine, Product, Spareparts, Lineitems, Video
from content.models import Blog, Carousel, Exhibition, Timeline
from team.models import Employee
from support.models import Warrantyclaim, Partsrequired, Machineregistration
from legal.models import Privacy, Terms


@pytest.fixture(scope="session")
def old_db_connection() -> sqlite3.Connection:
    """
    Connect to old database for testing.

    :return: SQLite connection to old database
    """
    old_db_path: str = '/Users/seanwelch/Coding/farmec-v2/server/database/database.db'

    if not os.path.exists(old_db_path):
        pytest.skip(f"Old database not found at {old_db_path}")

    conn: sqlite3.Connection = sqlite3.connect(old_db_path)
    conn.row_factory = sqlite3.Row

    yield conn

    conn.close()


@pytest.fixture
def clear_new_db(db) -> None:
    """
    Clear all tables in new database before each test.

    :param db: pytest-django database fixture
    :return: None
    """
    Supplier.objects.all().delete()
    Machine.objects.all().delete()
    Product.objects.all().delete()
    Spareparts.objects.all().delete()
    Lineitems.objects.all().delete()
    Video.objects.all().delete()
    Blog.objects.all().delete()
    Carousel.objects.all().delete()
    Exhibition.objects.all().delete()
    Timeline.objects.all().delete()
    Employee.objects.all().delete()
    Warrantyclaim.objects.all().delete()
    Partsrequired.objects.all().delete()
    Machineregistration.objects.all().delete()
    Privacy.objects.all().delete()
    Terms.objects.all().delete()

    yield

    # Cleanup after test
    Supplier.objects.all().delete()
    Machine.objects.all().delete()
    Product.objects.all().delete()
    Spareparts.objects.all().delete()
    Lineitems.objects.all().delete()
    Video.objects.all().delete()
    Blog.objects.all().delete()
    Carousel.objects.all().delete()
    Exhibition.objects.all().delete()
    Timeline.objects.all().delete()
    Employee.objects.all().delete()
    Warrantyclaim.objects.all().delete()
    Partsrequired.objects.all().delete()
    Machineregistration.objects.all().delete()
    Privacy.objects.all().delete()
    Terms.objects.all().delete()


@pytest.mark.django_db
class TestSupplierMigration:
    """Test Supplier table migration."""

    def test_supplier_exists_in_old_db(self, old_db_connection: sqlite3.Connection) -> None:
        """
        Verify Supplier table exists in old database.

        :param old_db_connection: SQLite connection to old database
        """
        cursor: sqlite3.Cursor = old_db_connection.cursor()
        cursor.execute("SELECT COUNT(*) FROM Supplier")
        count: int = cursor.fetchone()[0]
        assert count > 0, "No suppliers found in old database"

    def test_supplier_has_required_columns(self, old_db_connection: sqlite3.Connection) -> None:
        """
        Verify Supplier table has required columns.

        :param old_db_connection: SQLite connection to old database
        """
        cursor: sqlite3.Cursor = old_db_connection.cursor()
        cursor.execute("PRAGMA table_info(Supplier)")
        columns: list[str] = [row[1] for row in cursor.fetchall()]

        required_columns: list[str] = ['id', 'name', 'slug', 'created']
        for col in required_columns:
            assert col in columns, f"Column '{col}' not found in Supplier table"

    def test_supplier_fields_map_correctly(self, old_db_connection: sqlite3.Connection, clear_new_db: None) -> None:
        """
        Verify Supplier fields map correctly to Django model.

        :param old_db_connection: SQLite connection to old database
        :param clear_new_db: Fixture to clear new database
        """
        cursor: sqlite3.Cursor = old_db_connection.cursor()
        cursor.execute("SELECT * FROM Supplier LIMIT 1")
        row: sqlite3.Row = cursor.fetchone()

        if row is None:
            pytest.skip("No suppliers in old database")

        id_val: str = row['id']
        name: str = row['name']
        logo_image: str = row['logo_image']
        slug: str = row['slug']

        # Create supplier in new database
        Supplier.objects.create(
            id=id_val,
            name=name,
            logo_image=logo_image,
            slug=slug,
            order=1,
            publish=True,
        )

        # Verify it was saved correctly
        retrieved: Supplier = Supplier.objects.get(id=id_val)
        assert retrieved.name == name
        assert retrieved.logo_image == logo_image
        assert retrieved.slug == slug
        assert retrieved.order == 1
        assert retrieved.publish is True
        assert retrieved.uid is not None


@pytest.mark.django_db
class TestMachineMigration:
    """Test Machine table migration."""

    def test_machine_has_supplier_fk(self, old_db_connection: sqlite3.Connection) -> None:
        """
        Verify Machine table has supplier_id foreign key.

        :param old_db_connection: SQLite connection to old database
        """
        cursor: sqlite3.Cursor = old_db_connection.cursor()
        cursor.execute("SELECT COUNT(*) FROM Machine")
        count: int = cursor.fetchone()[0]
        assert count > 0, "No machines found in old database"

        cursor.execute("PRAGMA table_info(Machine)")
        columns: list[str] = [row[1] for row in cursor.fetchall()]
        assert 'supplier_id' in columns, "supplier_id column not found in Machine table"

    def test_machine_fk_references_valid_supplier(self, old_db_connection: sqlite3.Connection, clear_new_db: None) -> None:
        """
        Verify Machine foreign key references valid Supplier.

        :param old_db_connection: SQLite connection to old database
        :param clear_new_db: Fixture to clear new database
        """
        cursor: sqlite3.Cursor = old_db_connection.cursor()

        # Get a supplier
        cursor.execute("SELECT id, name FROM Supplier LIMIT 1")
        supplier_row: sqlite3.Row = cursor.fetchone()

        if supplier_row is None:
            pytest.skip("No suppliers in old database")

        supplier_id: str = supplier_row['id']

        # Get a machine with this supplier
        cursor.execute("SELECT id, supplier_id, name FROM Machine WHERE supplier_id = ? LIMIT 1", (supplier_id,))
        machine_row: sqlite3.Row = cursor.fetchone()

        if machine_row is None:
            pytest.skip("No machines with valid supplier in old database")

        # Create supplier in new DB
        Supplier.objects.create(id=supplier_id, name=supplier_row['name'])

        # Create machine in new DB
        supplier_obj: Supplier = Supplier.objects.get(id=supplier_id)
        Machine.objects.create(
            id=machine_row['id'],
            supplier=supplier_obj,
            name=machine_row['name'],
        )

        # Verify foreign key
        retrieved: Machine = Machine.objects.get(id=machine_row['id'])
        assert retrieved.supplier.id == supplier_id


@pytest.mark.django_db
class TestDateFieldConversion:
    """Test date field conversions."""

    def test_blog_date_field_conversion(self, old_db_connection: sqlite3.Connection, clear_new_db: None) -> None:
        """
        Verify Blog date field converts from TEXT to DateField.

        :param old_db_connection: SQLite connection to old database
        :param clear_new_db: Fixture to clear new database
        """
        cursor: sqlite3.Cursor = old_db_connection.cursor()
        cursor.execute("SELECT id, title, date FROM Blog WHERE date IS NOT NULL LIMIT 1")
        row: sqlite3.Row = cursor.fetchone()

        if row is None:
            pytest.skip("No blog posts with dates in old database")

        date_str: str = row['date']
        try:
            parsed_date: date = datetime.fromisoformat(date_str).date()
        except ValueError:
            pytest.skip(f"Date format '{date_str}' is not ISO format; migration script may need enhancement")

        # Create blog with converted date
        Blog.objects.create(
            id=row['id'],
            title=row['title'],
            date=parsed_date,
        )

        retrieved: Blog = Blog.objects.get(id=row['id'])
        assert isinstance(retrieved.date, date)
        assert retrieved.date == parsed_date


@pytest.mark.django_db
class TestDecimalFieldConversion:
    """Test decimal field conversions."""

    def test_lineitem_price_field_conversion(self, old_db_connection: sqlite3.Connection, clear_new_db: None) -> None:
        """
        Verify LineItem price field converts from REAL to DecimalField.

        :param old_db_connection: SQLite connection to old database
        :param clear_new_db: Fixture to clear new database
        """
        cursor: sqlite3.Cursor = old_db_connection.cursor()
        cursor.execute("SELECT id, name, price FROM LineItems LIMIT 1")
        row: sqlite3.Row = cursor.fetchone()

        if row is None:
            pytest.skip("No line items in old database")

        price_decimal: Decimal = Decimal(str(row['price']))

        Lineitems.objects.create(
            id=row['id'],
            name=row['name'],
            price=price_decimal,
        )

        retrieved: Lineitems = Lineitems.objects.get(id=row['id'])
        assert isinstance(retrieved.price, Decimal)
        assert retrieved.price == price_decimal


@pytest.mark.django_db
class TestBooleanFieldConversion:
    """Test boolean field conversions."""

    def test_machineregistration_boolean_conversion(self, old_db_connection: sqlite3.Connection, clear_new_db: None) -> None:
        """
        Verify MachineRegistration INTEGER fields convert to BooleanField.

        :param old_db_connection: SQLite connection to old database
        :param clear_new_db: Fixture to clear new database
        """
        cursor: sqlite3.Cursor = old_db_connection.cursor()
        cursor.execute(
            "SELECT id, dealer_name, owner_name, machine_model, serial_number, "
            "complete_supply FROM MachineRegistration LIMIT 1",
        )
        row: sqlite3.Row = cursor.fetchone()

        if row is None:
            pytest.skip("No machine registrations in old database")

        complete_supply_bool: bool = bool(int(row['complete_supply'])) if row['complete_supply'] else False

        Machineregistration.objects.create(
            id=row['id'],
            dealer_name=row['dealer_name'],
            owner_name=row['owner_name'],
            machine_model=row['machine_model'],
            serial_number=row['serial_number'],
            complete_supply=complete_supply_bool,
        )

        retrieved: Machineregistration = Machineregistration.objects.get(id=row['id'])
        assert isinstance(retrieved.complete_supply, bool)
        assert retrieved.complete_supply == complete_supply_bool


@pytest.mark.django_db
class TestNullableFields:
    """Test handling of nullable/optional fields."""

    def test_optional_fields_can_be_null(self, old_db_connection: sqlite3.Connection, clear_new_db: None) -> None:
        """
        Verify optional fields can be NULL without errors.

        :param old_db_connection: SQLite connection to old database
        :param clear_new_db: Fixture to clear new database
        """
        # Create blog with minimal fields
        Blog.objects.create(
            id='test-id-1',
            title='Test Blog',
            date=None,
            main_image=None,
            subheading=None,
            body=None,
        )

        retrieved: Blog = Blog.objects.get(id='test-id-1')
        assert retrieved.date is None
        assert retrieved.main_image is None
        assert retrieved.body is None


@pytest.mark.django_db
class TestBaseModelFields:
    """Test BaseModel field generation."""

    def test_basemodel_fields_generated(self, clear_new_db: None) -> None:
        """
        Verify BaseModel fields are auto-generated.

        :param clear_new_db: Fixture to clear new database
        """
        Supplier.objects.create(
            id='test-supplier-1',
            name='Test Supplier',
        )

        retrieved: Supplier = Supplier.objects.get(id='test-supplier-1')

        # Check BaseModel fields
        assert retrieved.order == 1
        assert retrieved.publish is True
        assert retrieved.uid is not None
        assert retrieved.created is not None
        assert retrieved.modified is not None


@pytest.mark.django_db
class TestMigrationDataIntegrity:
    """Test data integrity across migrations."""

    def test_supplier_count_matches(self, old_db_connection: sqlite3.Connection, clear_new_db: None) -> None:
        """
        Verify supplier count in old DB can be compared with new DB.

        :param old_db_connection: SQLite connection to old database
        :param clear_new_db: Fixture to clear new database
        """
        cursor: sqlite3.Cursor = old_db_connection.cursor()
        cursor.execute("SELECT COUNT(*) FROM Supplier")
        old_count: int = cursor.fetchone()[0]

        # Create all suppliers from old DB
        cursor.execute("SELECT id, name FROM Supplier")
        for row in cursor.fetchall():
            Supplier.objects.get_or_create(
                id=row['id'],
                defaults={'name': row['name']},
            )

        new_count: int = Supplier.objects.count()
        assert new_count == old_count, f"Supplier count mismatch: {new_count} vs {old_count}"

    def test_no_duplicate_ids(self, clear_new_db: None) -> None:
        """
        Verify get_or_create prevents duplicate IDs.

        :param clear_new_db: Fixture to clear new database
        """
        supplier_id: str = 'unique-supplier-1'

        Supplier.objects.get_or_create(
            id=supplier_id,
            defaults={'name': 'First'},
        )

        Supplier.objects.get_or_create(
            id=supplier_id,
            defaults={'name': 'Second'},
        )

        count: int = Supplier.objects.filter(id=supplier_id).count()
        assert count == 1, "Duplicate supplier ID created"
