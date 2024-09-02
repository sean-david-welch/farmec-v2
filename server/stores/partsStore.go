package stores

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/sean-david-welch/farmec-v2/server/db"
)

type PartsStore interface {
	GetParts(ctx context.Context, id string) ([]db.SparePart, error)
	GetPartById(ctx context.Context, id string) (*db.SparePart, error)
	CreatePart(ctx context.Context, part *db.SparePart) error
	UpdatePart(ctx context.Context, id string, part *db.SparePart) error
	DeletePart(ctx context.Context, id string) error
}

type PartsStoreImpl struct {
	queries *db.Queries
}

func NewPartsStore(sql *sql.DB) *PartsStoreImpl {
	queries := db.New(sql)
	return &PartsStoreImpl{queries: queries}
}

func (store *PartsStoreImpl) GetParts(ctx context.Context, id string) ([]db.SparePart, error) {
	parts, err := store.queries.GetParts(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("error occurred while getting spare parts: %w", err)
	}
	var result []db.SparePart
	for _, part := range parts {
		result = append(result, db.SparePart{
			ID: part.ID,
			SupplierID: part.SupplierID,
			Name: part.Name,
			PartsImage: part.PartsImage,
			SparePartsLink: part.SparePartsLink,
		})
	}
	return result, nil
}

func (store *PartsStoreImpl) GetPartById(ctx context.Context, id string) (*db.SparePart, error) {
	query := `SELECT * FROM "SpareParts" WHERE id = ?`
	row := store.database.QueryRow(query, id)

	var part dbSparePart.

	if err := ScanParts(row, &part); err != nil {
		if errors.Is(err, sqlctx context.Context, .ErrNoRows) {
			return nil, fmt.Errorf("error item found with the given id: %w", err)
		}
		return nil, fmt.Errorf("error scanning row: %w", err)
	}

	return &part, nil
}

func (store *PartsStoreImpl) CreatePart(ctx context.Context, part *db.SparePart) error {
	part.ID = uuid.NewString()

	query := `INSERT INTO "SpareParts" (id, supplier_id, name, parts_image, spare_parts_link)
	VALUES (?, ?, ?, ?, ?)`

	_, err := store.database.Exec(query, part.ID, part.SupplierID, part.Name, part.PartsImage, part.SparePartsLink)

	if err != nil {
		return fmt.Errorf("error creating spare part: %w", err)
	}

	return nil
}

func (store *PartsStoreImpl) UpdatePart(ctx context.Context, id string, part *db.SparePart) error {
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

func (store *PartsStoreImpl) DeletePart(ctx context.Context, id string) error {
	query := `DELETE FROM "SpareParts" WHERE id = ?`

	_, err := store.database.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error deleting part: %w", err)
	}

	return nil
}
