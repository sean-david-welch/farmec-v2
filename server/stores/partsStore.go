package stores

import (
	"context"
	"database/sql"
	"fmt"
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
			ID:             part.ID,
			SupplierID:     part.SupplierID,
			Name:           part.Name,
			PartsImage:     part.PartsImage,
			SparePartsLink: part.SparePartsLink,
		})
	}
	return result, nil
}

func (store *PartsStoreImpl) GetPartById(ctx context.Context, id string) (*db.SparePart, error) {
	part, err := store.queries.GetPartByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("error occurred while getting part from the db: %w", err)
	}
	return &part, err
}

func (store *PartsStoreImpl) CreatePart(ctx context.Context, part *db.SparePart) error {
	part.ID = uuid.NewString()

	params := db.CreateSparePartParams{
		ID:             part.ID,
		SupplierID:     part.SupplierID,
		Name:           part.Name,
		PartsImage:     part.PartsImage,
		SparePartsLink: part.SparePartsLink,
	}
	if err := store.queries.CreateSparePart(ctx, params); err != nil {
		return fmt.Errorf("error occurred while creating spare parts: %w", err)
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
