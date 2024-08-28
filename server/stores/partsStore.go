package stores

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/sean-david-welch/farmec-v2/server/types"
)

type PartsStore interface {
	GetParts(id string) ([]types.Sparepart, error)
	GetPartById(id string) (*types.Sparepart, error)
	CreatePart(part *types.Sparepart) error
	UpdatePart(id string, part *types.Sparepart) error
	DeletePart(id string) error
}

type PartsStoreImpl struct {
	database *sql.DB
}

func NewPartsStore(database *sql.DB) *PartsStoreImpl {
	return &PartsStoreImpl{database: database}
}

func ScanParts(row interface{}, part *types.Sparepart) error {
	var scanner interface {
		Scan(dest ...interface{}) error
	}

	switch value := row.(type) {
	case *sql.Row:
		scanner = value
	case *sql.Rows:
		scanner = value
	default:
		return fmt.Errorf("unsupported type: %T", value)
	}

	return scanner.Scan(&part.ID, &part.SupplierID, &part.Name, &part.PartsImage, &part.SparePartsLink)
}

func (store *PartsStoreImpl) GetParts(id string) ([]types.Sparepart, error) {
	var parts []types.Sparepart

	query := `SELECT * FROM "SpareParts" WHERE "supplier_id" = ?`
	rows, err := store.database.Query(query, id)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %w", err)
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Fatal("Failed to close database: ", err)
		}
	}()

	for rows.Next() {
		var part types.Sparepart

		if err := ScanParts(rows, &part); err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}
		parts = append(parts, part)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error after iterating over rows: %w", err)
	}

	return parts, nil
}

func (store *PartsStoreImpl) GetPartById(id string) (*types.Sparepart, error) {
	query := `SELECT * FROM "SpareParts" WHERE id = ?`
	row := store.database.QueryRow(query, id)

	var part types.Sparepart

	if err := ScanParts(row, &part); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("error item found with the given id: %w", err)
		}
		return nil, fmt.Errorf("error scanning row: %w", err)
	}

	return &part, nil
}

func (store *PartsStoreImpl) CreatePart(part *types.Sparepart) error {
	part.ID = uuid.NewString()

	query := `INSERT INTO "SpareParts" (id, supplier_id, name, parts_image, spare_parts_link)
	VALUES (?, ?, ?, ?, ?)`

	_, err := store.database.Exec(query, part.ID, part.SupplierID, part.Name, part.PartsImage, part.SparePartsLink)

	if err != nil {
		return fmt.Errorf("error creating spare part: %w", err)
	}

	return nil
}

func (store *PartsStoreImpl) UpdatePart(id string, part *types.Sparepart) error {
	query := `UPDATE "SpareParts" SET supplier_id = ?, name = ?, spare_parts_link  = ? WHERE ID = ?`
	args := []interface{}{id, part.SupplierID, part.Name, part.SparePartsLink}

	if part.PartsImage != "" && part.PartsImage != "null" {
		query = `UPDATE "SpareParts" SET supplier_id = ?, name = ?, parts_image = ?, spare_parts_link  = ? WHERE ID = ?`
		args = []interface{}{id, part.SupplierID, part.Name, part.PartsImage, part.SparePartsLink}
	}

	_, err := store.database.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("error updating part: %w", err)
	}

	return nil
}

func (store *PartsStoreImpl) DeletePart(id string) error {
	query := `DELETE FROM "SpareParts" WHERE id = ?`

	_, err := store.database.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error deleting part: %w", err)
	}

	return nil
}
