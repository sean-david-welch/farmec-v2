package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/sean-david-welch/farmec-v2/server/db"
)

type LineItemRepo interface {
	GetLineItems(ctx context.Context) ([]db.LineItem, error)
	GetLineItemById(ctx context.Context, id string) (*db.LineItem, error)
	CreateLineItem(ctx context.Context, lineItem *db.LineItem) error
	UpdateLineItem(ctx context.Context, id string, lineItem *db.LineItem) error
	DeleteLineItem(ctx context.Context, id string) error
}

type LineItemRepoImpl struct {
	queries *db.Queries
}

func NewLineItemRepo(sql *sql.DB) *LineItemRepoImpl {
	queries := db.New(sql)
	return &LineItemRepoImpl{queries: queries}
}

func (repo *LineItemRepoImpl) GetLineItems(ctx context.Context) ([]db.LineItem, error) {
	lineItems, err := repo.queries.GetLineItems(ctx)
	if err != nil {
		return nil, fmt.Errorf("error occurred while getting line items: %w", err)
	}
	var result []db.LineItem
	for _, lineItem := range lineItems {
		result = append(result, db.LineItem{
			ID:    lineItem.ID,
			Name:  lineItem.Name,
			Price: lineItem.Price,
			Image: lineItem.Image,
		})
	}
	return result, nil
}

func (repo *LineItemRepoImpl) GetLineItemById(ctx context.Context, id string) (*db.LineItem, error) {
	lineItem, err := repo.queries.GetLineItemByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("error occurred while getting line items from db: %w", err)
	}
	return &lineItem, nil
}

func (repo *LineItemRepoImpl) CreateLineItem(ctx context.Context, lineItem *db.LineItem) error {
	lineItem.ID = uuid.NewString()

	params := db.CreateLineItemParams{
		ID:    lineItem.ID,
		Name:  lineItem.Name,
		Price: lineItem.Price,
		Image: lineItem.Image,
	}
	if err := repo.queries.CreateLineItem(ctx, params); err != nil {
		return fmt.Errorf("error occurred while creating line items: %w", err)
	}
	return nil
}

func (repo *LineItemRepoImpl) UpdateLineItem(ctx context.Context, id string, lineItem *db.LineItem) error {
	if lineItem.Image.Valid {
		params := db.UpdateLineItemParams{
			Name:  lineItem.Name,
			Price: lineItem.Price,
			Image: lineItem.Image,
			ID:    id,
		}
		if err := repo.queries.UpdateLineItem(ctx, params); err != nil {
			return fmt.Errorf("error occurred while updating a line item: %w", err)
		}
	} else {
		params := db.UpdateLineItemNoImageParams{
			Name:  lineItem.Name,
			Price: lineItem.Price,
			ID:    id,
		}
		if err := repo.queries.UpdateLineItemNoImage(ctx, params); err != nil {
			return fmt.Errorf("error occurred while updating a line item: %w", err)
		}
	}
	return nil
}

func (repo *LineItemRepoImpl) DeleteLineItem(ctx context.Context, id string) error {
	if err := repo.queries.DeleteLineItem(ctx, id); err != nil {
		return fmt.Errorf("error occurred while deleting a line item: %w", err)
	}
	return nil
}
