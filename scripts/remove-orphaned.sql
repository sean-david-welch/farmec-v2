-- Disable foreign key constraints
PRAGMA foreign_keys = OFF;

-- Find orphaned Videos
SELECT id, supplier_id FROM Video WHERE supplier_id NOT IN (SELECT id FROM Supplier);

-- Delete orphaned Videos
DELETE FROM Video WHERE supplier_id NOT IN (SELECT id FROM Supplier);

-- Find orphaned Products
SELECT id, machine_id FROM Product WHERE machine_id NOT IN (SELECT id FROM Machine);

-- Delete orphaned Products
DELETE FROM Product WHERE machine_id NOT IN (SELECT id FROM Machine);

-- Find orphaned PartsRequired
SELECT id, warranty_id FROM PartsRequired WHERE warranty_id NOT IN (SELECT id FROM WarrantyClaim);

-- Delete orphaned PartsRequired
DELETE FROM PartsRequired WHERE warranty_id NOT IN (SELECT id FROM WarrantyClaim);

-- Find orphaned SpareParts
SELECT id, supplier_id FROM SpareParts WHERE supplier_id NOT IN (SELECT id FROM Supplier);

-- Delete orphaned SpareParts
DELETE FROM SpareParts WHERE supplier_id NOT IN (SELECT id FROM Supplier);

-- Find orphaned Machine
SELECT id, supplier_id FROM Machine WHERE supplier_id NOT IN (SELECT id FROM Supplier);

-- Delete orphaned Machine
DELETE FROM Machine WHERE supplier_id NOT IN (SELECT id FROM Supplier);

-- Re-enable foreign key constraints
PRAGMA foreign_keys = ON;
