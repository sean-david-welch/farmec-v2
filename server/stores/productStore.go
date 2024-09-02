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

type ProductStore interface {
	GetProducts(ctx context.Context, id string) ([]db.Product, error)
	GetProductById(ctx context.Context, id string) (*db.Product, error)
	CreateProduct(ctx context.Context, product *db.Product) error
	UpdateProduct(ctx context.Context, id string, product *db.Product) error
	DeleteProduct(ctx context.Context, id string) error
}

type ProductStoreImpl struct {
	queries *db.Queries
}

func NewProductStore(sql *sql.DB) *ProductStoreImpl {
	queries := db.New(sql)
	return &ProductStoreImpl{queries: queries}
}

func (store *ProductStoreImpl) GetProducts(ctx context.Context, id string) ([]db.Product, error) {
	products, err := store.queries.GetProducts(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("an error occurred while getting products: %w", err)
	}

	var result []db.Product
	for _, product := range products {
		result = append(result, db.Product{
			ID: product.ID,
			MachineID: product.MachineID,
			Name: product.Name,
			ProductImage: product.ProductImage,
			Description: product.Description,
			ProductLink: product.ProductLink,
		})
	}
	return result, nil
}

func (store *ProductStoreImpl) GetProductById(ctx context.Context, id string) (*db.Product, error) {
	query := `SELECT * FROM "Product" WHERE "id" = ?`
	row := store.database.QueryRow(query, id)

	var product db.Product

	if err := ScanProduct(row, &product); err != nil {
		if errors.Is(err, ctx context.Context, sql.ErrNoRows) {
			return nil, fmt.Errorf("error item found with the given id: %w", err)
		}

		return nil, fmt.Errorf("error scanning row: %w", err)
	}

	return &product, nil
}

func (store *ProductStoreImpl) CreateProduct(ctx context.Context, product *db.Product) error {
	product.ID = uuid.NewString()

	query := `INSERT INTO "Product" (id, machine_id, name, product_image, description, product_link)
	VALUES (?, ?, ?, ?, ?, ?)`

	_, err := store.database.Exec(query, product.ID, product.MachineID, product.Name, product.ProductImage, product.Description, product.ProductLink)

	if err != nil {
		return fmt.Errorf("error creating product: %w", err)
	}

	return nil
}

func (store *ProductStoreImpl) UpdateProduct(ctx context.Context, id string, product *db.Product) error {
	query := `UPDATE "Product" SET machine_id = ?, name = ?, description = ?, product_link = ? WHERE id = ?`
	args := []interface{}{product.MachineID, product.Name, product.Description, product.ProductLink, id}

	if product.ProductImage != "" && product.ProductImage != "null" {
		query = `UPDATE "Product" SET machine_id = ?, name = ?, product_image = ?, description = ?, product_link = ? WHERE id = ?`
		args = []interface{}{product.MachineID, product.Name, product.ProductImage, product.Description, product.ProductLink, id}

	}

	_, err := store.database.Exec(query, args...)

	if err != nil {
		return fmt.Errorf("error updating product: %w", err)
	}

	return nil
}

func (store *ProductStoreImpl) DeleteProduct(ctx context.Context, id string) error {
	query := `DELETE FROM "Product" WHERE id = ?`

	_, err := store.database.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error deleting product: %w", err)
	}

	return nil
}
