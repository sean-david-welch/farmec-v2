package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/sean-david-welch/farmec-v2/server/db"
)

type PartsRepo interface {
	GetParts(ctx context.Context, id string) ([]db.SparePart, error)
	GetPartById(ctx context.Context, id string) (*db.SparePart, error)
	CreatePart(ctx context.Context, part *db.SparePart) error
	UpdatePart(ctx context.Context, id string, part *db.SparePart) error
	DeletePart(ctx context.Context, id string) error
}

type PartsRepoImpl struct {
	queries *db.Queries
}

func NewPartsRepo(sql *sql.DB) *PartsRepoImpl {
	queries := db.New(sql)
	return &PartsRepoImpl{queries: queries}
}

func (repo *PartsRepoImpl) GetParts(ctx context.Context, id string) ([]db.SparePart, error) {
	parts, err := repo.queries.GetParts(ctx, id)
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

func (repo *PartsRepoImpl) GetPartById(ctx context.Context, id string) (*db.SparePart, error) {
	part, err := repo.queries.GetPartByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("error occurred while getting part from the db: %w", err)
	}
	return &part, err
}

func (repo *PartsRepoImpl) CreatePart(ctx context.Context, part *db.SparePart) error {
	part.ID = uuid.NewString()

	params := db.CreateSparePartParams{
		ID:             part.ID,
		SupplierID:     part.SupplierID,
		Name:           part.Name,
		PartsImage:     part.PartsImage,
		SparePartsLink: part.SparePartsLink,
	}
	if err := repo.queries.CreateSparePart(ctx, params); err != nil {
		return fmt.Errorf("error occurred while creating spare parts: %w", err)
	}
	return nil
}

func (repo *PartsRepoImpl) UpdatePart(ctx context.Context, id string, part *db.SparePart) error {
	if part.PartsImage.Valid {
		params := db.UpdateSparePartParams{
			SupplierID:     part.SupplierID,
			Name:           part.Name,
			PartsImage:     part.PartsImage,
			SparePartsLink: part.SparePartsLink,
			ID:             id,
		}
		if err := repo.queries.UpdateSparePart(ctx, params); err != nil {
			return fmt.Errorf("error occurred while updating a spare part: %w", err)
		}
	} else {
		params := db.UpdateSparePartNoImageParams{
			SupplierID:     part.SupplierID,
			Name:           part.Name,
			SparePartsLink: part.SparePartsLink,
			ID:             id,
		}
		if err := repo.queries.UpdateSparePartNoImage(ctx, params); err != nil {
			return fmt.Errorf("error occurred while updating a spare part: %w", err)
		}
	}
	return nil
}

func (repo *PartsRepoImpl) DeletePart(ctx context.Context, id string) error {
	if err := repo.queries.DeleteSparePart(ctx, id); err != nil {
		return fmt.Errorf("error occurred while deleting a spare part: %w", err)
	}
	return nil
}
