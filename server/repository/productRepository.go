package repository

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/sean-david-welch/farmec-v2/server/types"
)

type ProductRepository interface {
	GetProducts(id string) ([]types.Product, error)
	GetProductById(id string) (*types.Product, error)
	CreateProduct(product *types.Product) error
	UpdateMachine(id string, product *types.Product) error
	DeleteProduct(id string) error
}

type ProductRepositoryImpl struct {
	database *sql.DB
}

func NewProductRepository(database *sql.DB) *ProductRepositoryImpl {
	return &ProductRepositoryImpl{database: database}
}

func ScanProduct(row interface{}, product *types.Product) error {
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

	return scanner.Scan(&product.ID, &product.MachineID, &product.Name, &product.ProductImage, &product.Description, &product.ProductLink)
}

func (repository *ProductRepositoryImpl) GetProducts(id string) ([]types.Product, error) {
	var products []types.Product

	query := `SELECT * FROM "Product" WHERE "machine_id" = $1`
	rows, err := repository.database.Query(query, id)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		var product types.Product

		if err := ScanProduct(rows, &product); err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}

		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error after iterating over rows: %w", err)
	}

	return products, nil
}

func (repository *ProductRepositoryImpl) GetProductById(id string) (*types.Product, error) {
	query := `SELECT * FROM "Product" WHERE "id" = $1`
	row := repository.database.QueryRow(query, id)

	var product types.Product

	if err := ScanProduct(row, &product); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("error item found with the given id: %w", err)
		}

		return nil, fmt.Errorf("error scanning row: %w", err)
	}

	return &product, nil
}

func (repository *ProductRepositoryImpl) CreateProduct(product *types.Product) error {
	product.ID = uuid.NewString()

	query := `INSERT INTO "Product" (id, machine_id, name, product_image, description, product_link)
	VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := repository.database.Exec(query, product.ID, product.MachineID, product.Name, product.ProductImage, product.Description, product.ProductLink)

	if err != nil {
		return fmt.Errorf("error creating product: %w", err)
	}

	return nil
}

func (repository *ProductRepositoryImpl) UpdateMachine(id string, product *types.Product) error {
	query := `UPDATE "Product" SET machine_id = $1, name = $2, description = $3, product_link = $4 WHERE id = $6`
	args := []interface{}{product.MachineID, product.Name, product.Description, product.ProductLink, id}

	if product.ProductImage != "" && product.ProductImage != "null" {
		query = `UPDATE "Product" SET machine_id = $1, name = $2, product_image = $3, description = $4, product_link = $5 WHERE id = $6`
		args = []interface{}{product.MachineID, product.Name, product.ProductImage, product.Description, product.ProductLink, id}

	}

	_, err := repository.database.Exec(query, args...)

	if err != nil {
		return fmt.Errorf("error updating product: %w", err)
	}

	return nil
}

func (repository *ProductRepositoryImpl) DeleteProduct(id string) error {
	query := `DELETE FROM "Product" WHERE id = $1`

	_, err := repository.database.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error deleting product: %w", err)
	}

	return nil
}
