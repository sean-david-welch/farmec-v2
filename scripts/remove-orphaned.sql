-- Find orphaned Videos
SELECT id, supplier
FROM Video
WHERE supplier NOT IN (SELECT id FROM Supplier);

-- Delete orphaned Videos
DELETE
FROM Video
WHERE supplier NOT IN (SELECT id FROM Supplier);

-- Find orphaned Products
SELECT id, machine
FROM Product
WHERE machine NOT IN (SELECT id FROM Machine);

-- Delete orphaned Products
DELETE
FROM Product
WHERE machine NOT IN (SELECT id FROM Machine);

-- Find orphaned PartsRequired
SELECT id, warranty
FROM PartsRequired
WHERE warranty NOT IN (SELECT id FROM WarrantyClaim);

-- Delete orphaned PartsRequired
DELETE
FROM PartsRequired
WHERE warranty NOT IN (SELECT id FROM WarrantyClaim);
