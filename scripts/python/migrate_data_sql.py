#!/usr/bin/env python
"""
SQL-Based Database Migration Script

Migrates data from old SQLite database to new Django database using:
- Bulk SQL operations for performance
- Python-generated UUIDs for proper RFC 4122 compliance
- Single transaction for atomicity and rollback capability

Schema transformations:
- TEXT dates → DATE/DATETIME
- INTEGER (0/1) → BOOLEAN
- REAL → DECIMAL
- Generates BaseModel fields: order, publish, uid, created, modified
"""

import os
import sys
import sqlite3
import uuid
from datetime import datetime
from decimal import Decimal

# Database paths
OLD_DB = '/Users/seanwelch/Coding/farmec-v2/server/database/database.db'
NEW_DB = '/Users/seanwelch/Coding/farmec-v2/database/database.db'


def parse_date(date_str):
    """Parse TEXT date to ISO date string."""
    if not date_str:
        return None
    try:
        dt = datetime.fromisoformat(date_str)
        return dt.date().isoformat()
    except (ValueError, AttributeError, TypeError):
        return None


def parse_datetime(datetime_str):
    """Parse TEXT datetime to ISO datetime string."""
    if not datetime_str:
        return datetime.now().isoformat()
    try:
        dt = datetime.fromisoformat(datetime_str)
        return dt.isoformat()
    except (ValueError, AttributeError, TypeError):
        return datetime.now().isoformat()


def to_bool(value):
    """Convert to boolean (0 or 1)."""
    if value is None:
        return 0
    return 1 if value else 0


def clear_tables(cursor):
    """Clear all application tables in reverse dependency order."""
    print("\n" + "=" * 60)
    print("Phase 2: Clearing existing data")
    print("=" * 60)

    tables_in_order = [
        'PartsRequired',
        'Product',
        'WarrantyClaim',
        'MachineRegistration',
        'SpareParts',
        'Video',
        'Machine',
        'Blog',
        'Carousel',
        'Exhibition',
        'Timeline',
        'Employee',
        'LineItems',
        'Supplier',
        'Privacy',
        'Terms',
    ]

    for table in tables_in_order:
        cursor.execute(f"DELETE FROM {table}")
        print(f"  ✓ Cleared {table}")


def get_old_data(old_cursor, table_name):
    """Fetch all data from old database table."""
    old_cursor.execute(f"SELECT * FROM {table_name}")
    columns = [description[0] for description in old_cursor.description]
    rows = []
    for row in old_cursor.fetchall():
        rows.append(dict(zip(columns, row)))
    return rows


def migrate_supplier(old_cursor, new_cursor):
    """Migrate Supplier table."""
    print("  Migrating Supplier...", end=' ', flush=True)
    rows = get_old_data(old_cursor, 'Supplier')

    data = []
    for row in rows:
        data.append((
            row['id'],
            row['name'],
            row['logo_image'],
            row['marketing_image'],
            row['description'],
            row['social_facebook'],
            row['social_twitter'],
            row['social_instagram'],
            row['social_youtube'],
            row['social_linkedin'],
            row['social_website'],
            row['slug'],
            1,  # order
            1,  # publish
            str(uuid.uuid4()),  # uid
            parse_datetime(row['created']),  # created
            datetime.now().isoformat(),  # modified
        ))

    new_cursor.executemany("""
        INSERT INTO Supplier (
            id, name, logo_image, marketing_image, description,
            social_facebook, social_twitter, social_instagram,
            social_youtube, social_linkedin, social_website, slug,
            "order", publish, uid, created, modified
        ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
    """, data)

    print(f"✓ ({len(data)} records)")
    return len(data)


def migrate_machine(old_cursor, new_cursor):
    """Migrate Machine table."""
    print("  Migrating Machine...", end=' ', flush=True)
    rows = get_old_data(old_cursor, 'Machine')

    data = []
    for row in rows:
        data.append((
            row['id'],
            row['supplier_id'],
            row['name'],
            row['machine_image'],
            row['description'],
            row['machine_link'],
            row['slug'],
            1,  # order
            1,  # publish
            str(uuid.uuid4()),  # uid
            parse_datetime(row['created']),  # created
            datetime.now().isoformat(),  # modified
        ))

    new_cursor.executemany("""
        INSERT INTO Machine (
            id, supplier_id, name, machine_image, description,
            machine_link, slug, "order", publish, uid, created, modified
        ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
    """, data)

    print(f"✓ ({len(data)} records)")
    return len(data)


def migrate_product(old_cursor, new_cursor):
    """Migrate Product table."""
    print("  Migrating Product...", end=' ', flush=True)
    rows = get_old_data(old_cursor, 'Product')

    data = []
    for row in rows:
        data.append((
            row['id'],
            row['machine_id'],
            row['name'],
            row['product_image'],
            row['description'],
            row['product_link'],
            row['slug'],
            1,  # order
            1,  # publish
            str(uuid.uuid4()),  # uid
            datetime.now().isoformat(),  # created
            datetime.now().isoformat(),  # modified
        ))

    new_cursor.executemany("""
        INSERT INTO Product (
            id, machine_id, name, product_image, description,
            product_link, slug, "order", publish, uid, created, modified
        ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
    """, data)

    print(f"✓ ({len(data)} records)")
    return len(data)


def migrate_spareparts(old_cursor, new_cursor):
    """Migrate SpareParts table."""
    print("  Migrating SpareParts...", end=' ', flush=True)
    rows = get_old_data(old_cursor, 'SpareParts')

    data = []
    for row in rows:
        data.append((
            row['id'],
            row['supplier_id'],
            row['name'],
            row['parts_image'],
            row['spare_parts_link'],
            row['slug'],
            1,  # order
            1,  # publish
            str(uuid.uuid4()),  # uid
            datetime.now().isoformat(),  # created
            datetime.now().isoformat(),  # modified
        ))

    new_cursor.executemany("""
        INSERT INTO SpareParts (
            id, supplier_id, name, parts_image, spare_parts_link, slug,
            "order", publish, uid, created, modified
        ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
    """, data)

    print(f"✓ ({len(data)} records)")
    return len(data)


def migrate_lineitems(old_cursor, new_cursor):
    """Migrate LineItems table."""
    print("  Migrating LineItems...", end=' ', flush=True)
    rows = get_old_data(old_cursor, 'LineItems')

    data = []
    for row in rows:
        data.append((
            row['id'],
            row['name'],
            str(Decimal(str(row['price']))) if row['price'] else '0.00',
            row['image'],
            1,  # order
            1,  # publish
            str(uuid.uuid4()),  # uid
            datetime.now().isoformat(),  # created
            datetime.now().isoformat(),  # modified
        ))

    new_cursor.executemany("""
        INSERT INTO LineItems (
            id, name, price, image, "order", publish, uid, created, modified
        ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
    """, data)

    print(f"✓ ({len(data)} records)")
    return len(data)


def migrate_video(old_cursor, new_cursor):
    """Migrate Video table."""
    print("  Migrating Video...", end=' ', flush=True)
    rows = get_old_data(old_cursor, 'Video')

    data = []
    for row in rows:
        data.append((
            row['id'],
            row['supplier_id'],
            row['web_url'],
            row['title'],
            row['description'],
            row['video_id'],
            row['thumbnail_url'],
            1,  # order
            1,  # publish
            str(uuid.uuid4()),  # uid
            parse_datetime(row['created']),  # created
            datetime.now().isoformat(),  # modified
        ))

    new_cursor.executemany("""
        INSERT INTO Video (
            id, supplier_id, web_url, title, description, video_id,
            thumbnail_url, "order", publish, uid, created, modified
        ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
    """, data)

    print(f"✓ ({len(data)} records)")
    return len(data)


def migrate_blog(old_cursor, new_cursor):
    """Migrate Blog table."""
    print("  Migrating Blog...", end=' ', flush=True)
    rows = get_old_data(old_cursor, 'Blog')

    data = []
    for row in rows:
        data.append((
            row['id'],
            row['title'],
            parse_date(row['date']),
            row['main_image'],
            row['subheading'],
            row['body'],
            row['slug'],
            1,  # order
            1,  # publish
            str(uuid.uuid4()),  # uid
            parse_datetime(row['created']),  # created
            datetime.now().isoformat(),  # modified
        ))

    new_cursor.executemany("""
        INSERT INTO Blog (
            id, title, date, main_image, subheading, body, slug,
            "order", publish, uid, created, modified
        ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
    """, data)

    print(f"✓ ({len(data)} records)")
    return len(data)


def migrate_carousel(old_cursor, new_cursor):
    """Migrate Carousel table."""
    print("  Migrating Carousel...", end=' ', flush=True)
    rows = get_old_data(old_cursor, 'Carousel')

    data = []
    for row in rows:
        data.append((
            row['id'],
            row['name'],
            row['image'],
            1,  # order
            1,  # publish
            str(uuid.uuid4()),  # uid
            parse_datetime(row['created']),  # created
            datetime.now().isoformat(),  # modified
        ))

    new_cursor.executemany("""
        INSERT INTO Carousel (
            id, name, image, "order", publish, uid, created, modified
        ) VALUES (?, ?, ?, ?, ?, ?, ?, ?)
    """, data)

    print(f"✓ ({len(data)} records)")
    return len(data)


def migrate_exhibition(old_cursor, new_cursor):
    """Migrate Exhibition table."""
    print("  Migrating Exhibition...", end=' ', flush=True)
    rows = get_old_data(old_cursor, 'Exhibition')

    data = []
    for row in rows:
        data.append((
            row['id'],
            row['title'],
            parse_date(row['date']),
            row['location'],
            row['info'],
            1,  # order
            1,  # publish
            str(uuid.uuid4()),  # uid
            parse_datetime(row['created']),  # created
            datetime.now().isoformat(),  # modified
        ))

    new_cursor.executemany("""
        INSERT INTO Exhibition (
            id, title, date, location, info, "order", publish, uid, created, modified
        ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
    """, data)

    print(f"✓ ({len(data)} records)")
    return len(data)


def migrate_timeline(old_cursor, new_cursor):
    """Migrate Timeline table."""
    print("  Migrating Timeline...", end=' ', flush=True)
    rows = get_old_data(old_cursor, 'Timeline')

    data = []
    for row in rows:
        data.append((
            row['id'],
            row['title'],
            parse_date(row['date']),
            row['body'],
            1,  # order
            1,  # publish
            str(uuid.uuid4()),  # uid
            parse_datetime(row['created']),  # created
            datetime.now().isoformat(),  # modified
        ))

    new_cursor.executemany("""
        INSERT INTO Timeline (
            id, title, date, body, "order", publish, uid, created, modified
        ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
    """, data)

    print(f"✓ ({len(data)} records)")
    return len(data)


def migrate_employee(old_cursor, new_cursor):
    """Migrate Employee table."""
    print("  Migrating Employee...", end=' ', flush=True)
    rows = get_old_data(old_cursor, 'Employee')

    data = []
    for row in rows:
        data.append((
            row['id'],
            row['name'],
            row['email'],
            row['role'],
            row['profile_image'],
            1,  # order
            1,  # publish
            str(uuid.uuid4()),  # uid
            parse_datetime(row['created']),  # created
            datetime.now().isoformat(),  # modified
        ))

    new_cursor.executemany("""
        INSERT INTO Employee (
            id, name, email, role, profile_image, "order", publish, uid, created, modified
        ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
    """, data)

    print(f"✓ ({len(data)} records)")
    return len(data)


def migrate_warrantyclaim(old_cursor, new_cursor):
    """Migrate WarrantyClaim table."""
    print("  Migrating WarrantyClaim...", end=' ', flush=True)
    rows = get_old_data(old_cursor, 'WarrantyClaim')

    data = []
    for row in rows:
        labour_hours = None
        if row['labour_hours']:
            try:
                labour_hours = str(Decimal(str(row['labour_hours'])))
            except:
                labour_hours = None

        data.append((
            row['id'],
            row['dealer'],
            row['dealer_contact'],
            row['owner_name'],
            row['machine_model'],
            row['serial_number'],
            parse_date(row['install_date']),
            parse_date(row['failure_date']),
            parse_date(row['repair_date']),
            row['failure_details'],
            row['repair_details'],
            labour_hours,
            row['completed_by'],
            1,  # order
            1,  # publish
            str(uuid.uuid4()),  # uid
            parse_datetime(row['created']),  # created
            datetime.now().isoformat(),  # modified
        ))

    new_cursor.executemany("""
        INSERT INTO WarrantyClaim (
            id, dealer, dealer_contact, owner_name, machine_model, serial_number,
            install_date, failure_date, repair_date, failure_details, repair_details,
            labour_hours, completed_by, "order", publish, uid, created, modified
        ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
    """, data)

    print(f"✓ ({len(data)} records)")
    return len(data)


def migrate_partsrequired(old_cursor, new_cursor):
    """Migrate PartsRequired table."""
    print("  Migrating PartsRequired...", end=' ', flush=True)
    rows = get_old_data(old_cursor, 'PartsRequired')

    data = []
    for row in rows:
        data.append((
            row['id'],
            row['warranty_id'],
            row['part_number'],
            int(row['quantity_needed']),
            row['invoice_number'],
            row['description'],
            1,  # order
            1,  # publish
            str(uuid.uuid4()),  # uid
            datetime.now().isoformat(),  # created
            datetime.now().isoformat(),  # modified
        ))

    new_cursor.executemany("""
        INSERT INTO PartsRequired (
            id, warranty_id, part_number, quantity_needed, invoice_number,
            description, "order", publish, uid, created, modified
        ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
    """, data)

    print(f"✓ ({len(data)} records)")
    return len(data)


def migrate_machineregistration(old_cursor, new_cursor):
    """Migrate MachineRegistration table."""
    print("  Migrating MachineRegistration...", end=' ', flush=True)
    rows = get_old_data(old_cursor, 'MachineRegistration')

    data = []
    for row in rows:
        data.append((
            row['id'],
            row['dealer_name'],
            row['dealer_address'],
            row['owner_name'],
            row['owner_address'],
            row['machine_model'],
            row['serial_number'],
            parse_date(row['install_date']),
            row['invoice_number'],
            to_bool(row['complete_supply']),
            to_bool(row['pdi_complete']),
            to_bool(row['pto_correct']),
            to_bool(row['machine_test_run']),
            to_bool(row['safety_induction']),
            to_bool(row['operator_handbook']),
            parse_date(row['date']),
            row['completed_by'],
            1,  # order
            1,  # publish
            str(uuid.uuid4()),  # uid
            parse_datetime(row['created']),  # created
            datetime.now().isoformat(),  # modified
        ))

    new_cursor.executemany("""
        INSERT INTO MachineRegistration (
            id, dealer_name, dealer_address, owner_name, owner_address,
            machine_model, serial_number, install_date, invoice_number,
            complete_supply, pdi_complete, pto_correct, machine_test_run,
            safety_induction, operator_handbook, date, completed_by,
            "order", publish, uid, created, modified
        ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
    """, data)

    print(f"✓ ({len(data)} records)")
    return len(data)


def migrate_privacy(old_cursor, new_cursor):
    """Migrate Privacy table."""
    print("  Migrating Privacy...", end=' ', flush=True)
    rows = get_old_data(old_cursor, 'Privacy')

    data = []
    for row in rows:
        data.append((
            row['id'],
            row['title'],
            row['body'],
            1,  # order
            1,  # publish
            str(uuid.uuid4()),  # uid
            parse_datetime(row['created']),  # created
            datetime.now().isoformat(),  # modified
        ))

    new_cursor.executemany("""
        INSERT INTO Privacy (
            id, title, body, "order", publish, uid, created, modified
        ) VALUES (?, ?, ?, ?, ?, ?, ?, ?)
    """, data)

    print(f"✓ ({len(data)} records)")
    return len(data)


def migrate_terms(old_cursor, new_cursor):
    """Migrate Terms table."""
    print("  Migrating Terms...", end=' ', flush=True)
    rows = get_old_data(old_cursor, 'Terms')

    data = []
    for row in rows:
        data.append((
            row['id'],
            row['title'],
            row['body'],
            1,  # order
            1,  # publish
            str(uuid.uuid4()),  # uid
            parse_datetime(row['created']),  # created
            datetime.now().isoformat(),  # modified
        ))

    new_cursor.executemany("""
        INSERT INTO Terms (
            id, title, body, "order", publish, uid, created, modified
        ) VALUES (?, ?, ?, ?, ?, ?, ?, ?)
    """, data)

    print(f"✓ ({len(data)} records)")
    return len(data)


def verify_migration(old_conn, new_conn):
    """Verify migration success."""
    print("\n" + "=" * 60)
    print("Phase 4: Verification")
    print("=" * 60)

    tables = [
        'Supplier', 'Machine', 'Product', 'SpareParts', 'LineItems', 'Video',
        'Blog', 'Carousel', 'Exhibition', 'Timeline', 'Employee',
        'WarrantyClaim', 'PartsRequired', 'MachineRegistration',
        'Privacy', 'Terms'
    ]

    old_cursor = old_conn.cursor()
    new_cursor = new_conn.cursor()

    all_match = True
    for table in tables:
        old_cursor.execute(f"SELECT COUNT(*) FROM {table}")
        old_count = old_cursor.fetchone()[0]

        new_cursor.execute(f"SELECT COUNT(*) FROM {table}")
        new_count = new_cursor.fetchone()[0]

        match = "✓" if old_count == new_count else "⚠"
        print(f"  {match} {table:20} | Old: {old_count:4} | New: {new_count:4}")

        if old_count != new_count:
            all_match = False

    return all_match


def main():
    if not os.path.exists(OLD_DB):
        print(f"❌ Error: Old database not found at {OLD_DB}")
        sys.exit(1)

    if not os.path.exists(NEW_DB):
        print(f"❌ Error: New database not found at {NEW_DB}")
        sys.exit(1)

    print("=" * 60)
    print("SQL-Based Database Migration")
    print("=" * 60)
    print(f"\nSource: {OLD_DB}")
    print(f"Target: {NEW_DB}")

    old_conn = sqlite3.connect(OLD_DB)
    old_conn.row_factory = sqlite3.Row
    old_cursor = old_conn.cursor()

    new_conn = sqlite3.connect(NEW_DB)
    new_cursor = new_conn.cursor()

    try:
        # Phase 1: Setup
        print("\n" + "=" * 60)
        print("Phase 1: Setup")
        print("=" * 60)
        new_cursor.execute("PRAGMA foreign_keys = ON")
        print("  ✓ Foreign keys enabled")
        new_cursor.execute("BEGIN TRANSACTION")
        print("  ✓ Transaction started")

        # Phase 2: Clear existing data
        clear_tables(new_cursor)

        # Phase 3: Migrate data
        print("\n" + "=" * 60)
        print("Phase 3: Migrating Data")
        print("=" * 60)

        total_migrated = 0
        total_migrated += migrate_supplier(old_cursor, new_cursor)
        total_migrated += migrate_machine(old_cursor, new_cursor)
        total_migrated += migrate_product(old_cursor, new_cursor)
        total_migrated += migrate_spareparts(old_cursor, new_cursor)
        total_migrated += migrate_lineitems(old_cursor, new_cursor)
        total_migrated += migrate_video(old_cursor, new_cursor)
        total_migrated += migrate_blog(old_cursor, new_cursor)
        total_migrated += migrate_carousel(old_cursor, new_cursor)
        total_migrated += migrate_exhibition(old_cursor, new_cursor)
        total_migrated += migrate_timeline(old_cursor, new_cursor)
        total_migrated += migrate_employee(old_cursor, new_cursor)
        total_migrated += migrate_warrantyclaim(old_cursor, new_cursor)
        total_migrated += migrate_partsrequired(old_cursor, new_cursor)
        total_migrated += migrate_machineregistration(old_cursor, new_cursor)
        total_migrated += migrate_privacy(old_cursor, new_cursor)
        total_migrated += migrate_terms(old_cursor, new_cursor)

        # Phase 4: Verification
        all_match = verify_migration(old_conn, new_conn)

        # Phase 5: Commit
        print("\n" + "=" * 60)
        print("Phase 5: Commit")
        print("=" * 60)

        if all_match:
            new_cursor.execute("COMMIT")
            print("  ✓ Transaction committed")
            print(f"\n✅ Migration completed successfully!")
            print(f"   Total records migrated: {total_migrated}")
        else:
            new_cursor.execute("ROLLBACK")
            print("  ⚠ Transaction rolled back due to verification failure")
            sys.exit(1)

    except Exception as e:
        print(f"\n❌ Migration failed: {e}")
        try:
            new_cursor.execute("ROLLBACK")
            print("  ✓ Transaction rolled back")
        except:
            pass
        import traceback
        traceback.print_exc()
        sys.exit(1)
    finally:
        old_conn.close()
        new_conn.close()

    print("=" * 60)


if __name__ == '__main__':
    main()
