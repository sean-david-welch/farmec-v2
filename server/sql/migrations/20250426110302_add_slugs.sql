-- +goose Up
-- +goose StatementBegin
-- Add the slug column to all tables
ALTER TABLE Supplier ADD COLUMN slug TEXT;
ALTER TABLE Machine ADD COLUMN slug TEXT;
ALTER TABLE Product ADD COLUMN slug TEXT;
ALTER TABLE Blog ADD COLUMN slug TEXT;
ALTER TABLE SpareParts ADD COLUMN slug TEXT;

-- Update existing records to populate slugs
UPDATE Supplier
SET slug = LOWER(
    REPLACE(
        REPLACE(
            REPLACE(
                REPLACE(
                    REPLACE(name, ' ', '-'),
                    '.', ''),
                '&', 'and'),
            ',', ''),
        '--', '-')
    );

UPDATE Machine
SET slug = LOWER(
    REPLACE(
        REPLACE(
            REPLACE(
                REPLACE(
                    REPLACE(name, ' ', '-'),
                    '.', ''),
                '&', 'and'),
            ',', ''),
        '--', '-')
    );

UPDATE Product
SET slug = LOWER(
    REPLACE(
        REPLACE(
            REPLACE(
                REPLACE(
                    REPLACE(name, ' ', '-'),
                    '.', ''),
                '&', 'and'),
            ',', ''),
        '--', '-')
    );

UPDATE Blog
SET slug = LOWER(
    REPLACE(
        REPLACE(
            REPLACE(
                REPLACE(
                    REPLACE(title, ' ', '-'),
                    '.', ''),
                '&', 'and'),
            ',', ''),
        '--', '-')
    );

UPDATE SpareParts
SET slug = LOWER(
    REPLACE(
        REPLACE(
            REPLACE(
                REPLACE(
                    REPLACE(name, ' ', '-'),
                    '.', ''),
                '&', 'and'),
            ',', ''),
        '--', '-')
    );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE Supplier DROP COLUMN slug;
ALTER TABLE Machine DROP COLUMN slug;
ALTER TABLE Product DROP COLUMN slug;
ALTER TABLE Blog DROP COLUMN slug;
ALTER TABLE SpareParts DROP COLUMN slug;
-- +goose StatementEnd