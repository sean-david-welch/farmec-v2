SELECT id, supplier
FROM Video
WHERE supplier NOT IN (SELECT id FROM Supplier);

DELETE
FROM Video
WHERE supplier NOT IN (SELECT id FROM Supplier);

SELECT id, machine
FROM Product
WHERE machine NOT IN (SELECT id FROM Machine);

DELETE
FROM Product
WHERE machine NOT IN (SELECT id FROM Machine);

SELECT id, warranty
FROM PartsRequired
WHERE warranty NOT IN (SELECT id FROM WarrantyClaim);

DELETE
FROM PartsRequired
WHERE warranty NOT IN (SELECT id FROM WarrantyClaim);
