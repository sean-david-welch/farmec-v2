-- +goose Up
-- +goose StatementBegin
CREATE TABLE Blog
(
    id         TEXT PRIMARY KEY,
    title      TEXT NOT NULL,
    date       TEXT,
    main_image TEXT,
    subheading TEXT,
    body       TEXT,
    created    TEXT
) STRICT;
CREATE TABLE Carousel
(
    id      TEXT PRIMARY KEY,
    name    TEXT NOT NULL,
    image   TEXT,
    created TEXT
) STRICT;
CREATE TABLE Employee
(
    id            TEXT PRIMARY KEY,
    name          TEXT NOT NULL,
    email         TEXT NOT NULL,
    role          TEXT NOT NULL,
    profile_image TEXT,
    created       TEXT
) STRICT;
CREATE TABLE Exhibition
(
    id       TEXT PRIMARY KEY,
    title    TEXT NOT NULL,
    date     TEXT,
    location TEXT,
    info     TEXT,
    created  TEXT
) STRICT;
CREATE TABLE LineItems
(
    id    TEXT PRIMARY KEY,
    name  TEXT NOT NULL,
    price REAL NOT NULL,
    image TEXT
) STRICT;
CREATE TABLE Supplier
(
    id               TEXT PRIMARY KEY,
    name             TEXT NOT NULL,
    logo_image       TEXT,
    marketing_image  TEXT,
    description      TEXT,
    social_facebook  TEXT,
    social_twitter   TEXT,
    social_instagram TEXT,
    social_youtube   TEXT,
    social_linkedin  TEXT,
    social_website   TEXT,
    created          TEXT
) STRICT;
CREATE TABLE Machine
(
    id            TEXT PRIMARY KEY,
    supplier_id   TEXT NOT NULL,
    name          TEXT NOT NULL,
    machine_image TEXT,
    description   TEXT,
    machine_link  TEXT,
    created       TEXT,
    FOREIGN KEY (supplier_id) REFERENCES Supplier (id) ON DELETE CASCADE
) STRICT;
CREATE TABLE MachineRegistration
(
    id                TEXT PRIMARY KEY,
    dealer_name       TEXT NOT NULL,
    dealer_address    TEXT,
    owner_name        TEXT NOT NULL,
    owner_address     TEXT,
    machine_model     TEXT NOT NULL,
    serial_number     TEXT NOT NULL,
    install_date      TEXT,
    invoice_number    TEXT,
    complete_supply   INTEGER,
    pdi_complete      INTEGER,
    pto_correct       INTEGER,
    machine_test_run  INTEGER,
    safety_induction  INTEGER,
    operator_handbook INTEGER,
    date              TEXT,
    completed_by      TEXT,
    created           TEXT
) STRICT;
CREATE TABLE WarrantyClaim
(
    id              TEXT PRIMARY KEY,
    dealer          TEXT NOT NULL,
    dealer_contact  TEXT,
    owner_name      TEXT NOT NULL,
    machine_model   TEXT NOT NULL,
    serial_number   TEXT NOT NULL,
    install_date    TEXT,
    failure_date    TEXT,
    repair_date     TEXT,
    failure_details TEXT,
    repair_details  TEXT,
    labour_hours    TEXT,
    completed_by    TEXT,
    created         TEXT
) STRICT;
CREATE TABLE PartsRequired
(
    id              TEXT PRIMARY KEY,
    warranty_id     TEXT NOT NULL,
    part_number     TEXT,
    quantity_needed TEXT NOT NULL,
    invoice_number  TEXT,
    description     TEXT,
    FOREIGN KEY (warranty_id) REFERENCES WarrantyClaim (id) ON DELETE CASCADE
) STRICT;
CREATE TABLE Privacy
(
    id      TEXT PRIMARY KEY,
    title   TEXT NOT NULL,
    body    TEXT,
    created TEXT
) STRICT;
CREATE TABLE Product
(
    id            TEXT PRIMARY KEY,
    machine_id    TEXT NOT NULL,
    name          TEXT NOT NULL,
    product_image TEXT,
    description   TEXT,
    product_link  TEXT,
    FOREIGN KEY (machine_id) REFERENCES Machine (id) ON DELETE CASCADE
) STRICT;
CREATE TABLE SpareParts
(
    id               TEXT PRIMARY KEY,
    supplier_id      TEXT NOT NULL,
    name             TEXT NOT NULL,
    parts_image      TEXT,
    spare_parts_link TEXT,
    FOREIGN KEY (supplier_id) REFERENCES Supplier (id) ON DELETE CASCADE
) STRICT;
CREATE TABLE Terms
(
    id      TEXT PRIMARY KEY,
    title   TEXT NOT NULL,
    body    TEXT,
    created TEXT
) STRICT;
CREATE TABLE Timeline
(
    id      TEXT PRIMARY KEY,
    title   TEXT NOT NULL,
    date    TEXT,
    body    TEXT,
    created TEXT
) STRICT;
CREATE TABLE Video
(
    id            TEXT PRIMARY KEY,
    supplier_id   TEXT NOT NULL,
    web_url       TEXT,
    title         TEXT,
    description   TEXT,
    video_id      TEXT,
    thumbnail_url TEXT,
    created       TEXT,
    FOREIGN KEY (supplier_id) REFERENCES Supplier (id) ON DELETE CASCADE
) STRICT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
