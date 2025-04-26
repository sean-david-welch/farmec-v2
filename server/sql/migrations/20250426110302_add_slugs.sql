-- +goose Up
-- +goose StatementBegin
-- Add the slug column
ALTER TABLE Supplier ADD COLUMN slug TEXT;

-- Update existing records to populate slugs
update Supplier
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
-- +goose StatementEnd