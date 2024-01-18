package repository

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/sean-david-welch/farmec-v2/server/types"
)

type LineItemRepository struct {
	database *sql.DB
}

func NewLineItemRepository(database *sql.DB) *LineItemRepository {
	return &LineItemRepository{database: database}
}

func(repository *LineItemRepository) GetLineItems() ([]types.LineItem, error) {
	var lineItems []types.LineItem

	query := `SELECT * FROM "LineItems"`
	rows, err := repository.database.Query(query); if err != nil {
		return nil, fmt.Errorf("error occurred while querying databse: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var lineItem types.LineItem

		if err := rows.Scan(&lineItem.ID, &lineItem.Name, &lineItem.Price, &lineItem.Image); err != nil {
			return nil, fmt.Errorf("error occurred while scanning line item: %w", err)
		}

		lineItems = append(lineItems, lineItem)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error occurred after iterating over rows: %w", err)
	}

	return lineItems, nil
}

func(repository *LineItemRepository) GetLineItemById(id string) (*types.LineItem, error) {
	var lineItem types.LineItem

	query := `SELECT * FROM "LineItems" WHERE "id" = $1`
	row := repository.database.QueryRow(query, id)

	if err := row.Scan(&lineItem.ID, &lineItem.Name, &lineItem.Price, &lineItem.Image); err != nil {
		return nil, fmt.Errorf("error occurrred while scanning line item: %w", err)
	}

	return &lineItem, nil
}

func(repository *LineItemRepository) CreateLineItem(lineItem *types.LineItem) error {
	lineItem.ID = uuid.NewString()

	query := `INSERT INTO "Lineitems" (id, name, price, image) VALUES ($1, $2, $3, $4)`

	_, err := repository.database.Exec(query, lineItem.ID, lineItem.Name, lineItem.Price, lineItem.Image); if err != nil {
		return fmt.Errorf("error occurred while creating line item: %w", err)
	}
	
	return nil 
}

func(repository *LineItemRepository) UpdateLineItem(id string, lineItem *types.LineItem) error {
	query := `UPDATE "Lineitems" SET "name" = $2", "price" = $3, "image" = $4 WHERE "id" = $1`

	_, err := repository.database.Exec(query, lineItem.ID, lineItem.Name, lineItem.Price, lineItem.Image); if err != nil {
		return fmt.Errorf("error occurred while updating line item: %w", err)
	}
	
	return nil 
}

func(repository *LineItemRepository) DeleteLineItem(id string) error {
	query := `DELETE FROM "LineItems" WHERE "id" = $1`

	_, err := repository.database.Exec(query, id); if err != nil {
		return fmt.Errorf("error occurred while deleting line items: %w", err)
	}
	
	return nil
}