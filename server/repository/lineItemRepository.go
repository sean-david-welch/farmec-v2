package repository

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/sean-david-welch/farmec-v2/server/models"
)

// type LineItem struct {
//     ID     string `json:"id"`
//     Name   string `json:"name"`
//     Price  int    `json:"price"`
//     Image  string `json:"image"`
// }

type LineItemRepository struct {
	database *sql.DB
}

func NewLineItemRepository(database *sql.DB) *LineItemRepository {
	return &LineItemRepository{database: database}
}

func(repository *LineItemRepository) GetLineItems() ([]models.LineItem, error) {
	var lineItems []models.LineItem

	query := `SELECT * FROM "LineItems"`
	rows, err := repository.database.Query(query); if err != nil {
		return nil, fmt.Errorf("error occurred while querying databse: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var lineItem models.LineItem

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

func(repository *LineItemRepository) GetLineItemById(id string) (*models.LineItem, error) {
	var lineItem *models.LineItem

	query := `SELECT * FROM "LineItems" WHERE "id" = $1`
	row := repository.database.QueryRow(query, id)

	if err := row.Scan(&lineItem.ID, &lineItem.Name, &lineItem.Price, &lineItem.Image); err != nil {
		return nil, fmt.Errorf("error occurrred while scanning line item: %w", err)
	}

	return lineItem, nil
}

func(repository *LineItemRepository) CreateLineItem(lineItem *models.LineItem) error {
	lineItem.ID = uuid.NewString()

	query := `INSERT INTO "Lineitems" (id, name, price, image) VALUES ($1, $2, $3, $4)`

	_, err := repository.database.Exec(query, lineItem.ID, lineItem.Name, lineItem.Price, lineItem.Image); if err != nil {
		return fmt.Errorf("error occurred while creating line item: %w", err)
	}
	
	return nil 
}

func(repository *LineItemRepository) UpdateLineItem(id string, lineItem *models.LineItem) error {
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