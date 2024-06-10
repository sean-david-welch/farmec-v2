package store

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/sean-david-welch/farmec-v2/server/types"
)

type LineItemStore interface {
	GetLineItems() ([]types.LineItem, error)
	GetLineItemById(id string) (*types.LineItem, error)
	CreateLineItem(lineItem *types.LineItem) error
	UpdateLineItem(id string, lineItem *types.LineItem) error
	DeleteLineItem(id string) error
}

type LineItemStoreImpl struct {
	database *sql.DB
}

func NewLineItemStore(database *sql.DB) *LineItemStoreImpl {
	return &LineItemStoreImpl{database: database}
}

func (store *LineItemStoreImpl) GetLineItems() ([]types.LineItem, error) {
	var lineItems []types.LineItem

	query := `SELECT * FROM "LineItems"`
	rows, err := store.database.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error occurred while querying database: %w", err)
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Fatal("Failed to close database: ", err)
		}
	}()

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

func (store *LineItemStoreImpl) GetLineItemById(id string) (*types.LineItem, error) {
	var lineItem types.LineItem

	query := `SELECT * FROM "LineItems" WHERE "id" = $1`
	row := store.database.QueryRow(query, id)

	if err := row.Scan(&lineItem.ID, &lineItem.Name, &lineItem.Price, &lineItem.Image); err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("error item found with the given id: %w", err)
		}

		return nil, fmt.Errorf("error occurrred while scanning line item: %w", err)
	}

	return &lineItem, nil
}

func (store *LineItemStoreImpl) CreateLineItem(lineItem *types.LineItem) error {
	lineItem.ID = uuid.NewString()

	query := `INSERT INTO "LineItems" (id, name, price, image) VALUES ($1, $2, $3, $4)`

	_, err := store.database.Exec(query, lineItem.ID, lineItem.Name, lineItem.Price, lineItem.Image)
	if err != nil {
		return fmt.Errorf("error occurred while creating line item: %w", err)
	}

	return nil
}

func (store *LineItemStoreImpl) UpdateLineItem(id string, lineItem *types.LineItem) error {
	query := `UPDATE "LineItems" SET "name" = $2, "price" = $3,  WHERE "id" = $1`
	args := []interface{}{id, lineItem.Name, lineItem.Price}

	if lineItem.Image != "" && lineItem.Image != "null" {
		query = `UPDATE "LineItems" SET "name" = $2, "price" = $3, "image" = $4 WHERE "id" = $1`
		args = []interface{}{id, lineItem.Name, lineItem.Price, lineItem.Image}
	}

	_, err := store.database.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("error occurred while updating line item: %w", err)
	}

	return nil
}

func (store *LineItemStoreImpl) DeleteLineItem(id string) error {
	query := `DELETE FROM "LineItems" WHERE "id" = $1`

	_, err := store.database.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error occurred while deleting line items: %w", err)
	}

	return nil
}
