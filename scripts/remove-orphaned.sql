-- Disable foreign key constraints
PRAGMA foreign_keys = OFF;

-- Find orphaned Videos
SELECT id, supplier FROM Video WHERE supplier NOT IN (SELECT id FROM Supplier) AND supplier IS NOT NULL;

-- Delete orphaned Videos
DELETE FROM Video WHERE supplier NOT IN (SELECT id FROM Supplier) AND supplier IS NOT NULL;

-- Find orphaned Products
SELECT id, machine FROM Product WHERE machine NOT IN (SELECT id FROM Machine) AND machine IS NOT NULL;

-- Delete orphaned Products
DELETE FROM Product WHERE machine NOT IN (SELECT id FROM Machine) AND machine IS NOT NULL;

-- Find orphaned PartsRequired
SELECT id, warranty FROM PartsRequired WHERE warranty NOT IN (SELECT id FROM WarrantyClaim) AND warranty IS NOT NULL;

-- Delete orphaned PartsRequired
DELETE FROM PartsRequired WHERE warranty NOT IN (SELECT id FROM WarrantyClaim) AND warranty IS NOT NULL;

-- Find orphaned SpareParts
SELECT id, supplier FROM SpareParts WHERE supplier NOT IN (SELECT id FROM Supplier) AND supplier IS NOT NULL;

-- Delete orphaned SpareParts
DELETE FROM SpareParts WHERE supplier NOT IN (SELECT id FROM Supplier) AND supplier IS NOT NULL;

-- Find orphaned Machine
SELECT id, supplier FROM Machine WHERE supplier NOT IN (SELECT id FROM Supplier) AND supplier IS NOT NULL;

-- Delete orphaned Machine
DELETE FROM Machine WHERE supplier NOT IN (SELECT id FROM Supplier) AND supplier IS NOT NULL;

-- Re-enable foreign key constraints
PRAGMA foreign_keys = ON;
