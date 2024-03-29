package repository

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/sean-david-welch/farmec-v2/server/types"
)

type LineItemRepository interface {
	GetLineItems() ([]types.LineItem, error)
	GetLineItemById(id string) (*types.LineItem, error)
	CreateLineItem(lineItem *types.LineItem) error
	UpdateLineItem(id string, lineItem *types.LineItem) error
	DeleteLineItem(id string) error
}

type LineItemRepositoryImpl struct {
	database *sql.DB
}

func NewLineItemRepository(database *sql.DB) *LineItemRepositoryImpl {
	return &LineItemRepositoryImpl{database: database}
}

func (repository *LineItemRepositoryImpl) GetLineItems() ([]types.LineItem, error) {
	var lineItems []types.LineItem

	query := `SELECT * FROM "LineItems"`
	rows, err := repository.database.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error occurred while querying database: %w", err)
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

func (repository *LineItemRepositoryImpl) GetLineItemById(id string) (*types.LineItem, error) {
	var lineItem types.LineItem

	query := `SELECT * FROM "LineItems" WHERE "id" = $1`
	row := repository.database.QueryRow(query, id)

	if err := row.Scan(&lineItem.ID, &lineItem.Name, &lineItem.Price, &lineItem.Image); err != nil {

		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("error item found with the given id: %w", err)
		}

		return nil, fmt.Errorf("error occurrred while scanning line item: %w", err)
	}

	return &lineItem, nil
}

func (repository *LineItemRepositoryImpl) CreateLineItem(lineItem *types.LineItem) error {
	lineItem.ID = uuid.NewString()

	query := `INSERT INTO "LineItems" (id, name, price, image) VALUES ($1, $2, $3, $4)`

	_, err := repository.database.Exec(query, lineItem.ID, lineItem.Name, lineItem.Price, lineItem.Image)
	if err != nil {
		return fmt.Errorf("error occurred while creating line item: %w", err)
	}

	return nil
}

func (repository *LineItemRepositoryImpl) UpdateLineItem(id string, lineItem *types.LineItem) error {
	query := `UPDATE "LineItems" SET "name" = $2, "price" = $3,  WHERE "id" = $1`
	args := []interface{}{id, lineItem.Name, lineItem.Price}

	if lineItem.Image != "" && lineItem.Image != "null" {
		query = `UPDATE "LineItems" SET "name" = $2, "price" = $3, "image" = $4 WHERE "id" = $1`
		args = []interface{}{id, lineItem.Name, lineItem.Price, lineItem.Image}
	}

	_, err := repository.database.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("error occurred while updating line item: %w", err)
	}

	return nil
}

func (repository *LineItemRepositoryImpl) DeleteLineItem(id string) error {
	query := `DELETE FROM "LineItems" WHERE "id" = $1`

	_, err := repository.database.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error occurred while deleting line items: %w", err)
	}

	return nil
}
