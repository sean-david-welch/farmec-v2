package repository

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/sean-david-welch/farmec-v2/server/models"
)

// type Product struct {
//     ID            string `json:"id"`
//     MachineID     string `json:"machineId"`
//     Name          string `json:"name"`
//     ProductImage  string `json:"product_image"`
//     Description   string `json:"description"`
//     ProductLink   string `json:"product_link"`
// }

type ProductRepository struct {
	database *sql.DB
}

func NewProductRepository(database *sql.DB) *ProductRepository {
	return &ProductRepository{database: database}
}

func ScanProduct(row interface{}, product *models.Product) error {
	var scanner interface {
		Scan(dest ...interface{}) error
	}

	switch value := row.(type) {
	case *sql.Row:
		scanner = value
	case *sql.Rows:
		scanner = value
	default: return fmt.Errorf("unsupported type: %T", value)
	}

	return scanner.Scan(&product.ID, &product.MachineID, &product.Name, &product.ProductImage, &product.Description, &product.ProductLink)
}

func (repository *ProductRepository) GetProducts(id string) ([]models.Product, error) {
	var products []models.Product

	query := `SELECT * FROM "Product" WHERE "machineId" = $1`
	rows, err := repository.database.Query(query, id); if err != nil {
		return nil, fmt.Errorf("error executing query: %w", err)
	}

	defer rows.Close()

	for rows.Next(){
		var product models.Product

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

func (repository *ProductRepository) GetProductById(id string) (*models.Product, error) {
	query := `SELECT * FROM "Product" WHERE "id" = $1`
	row := repository.database.QueryRow(query, id)


	var product models.Product

	if err := ScanProduct(row, &product); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, fmt.Errorf("error scanning row: %w", err)
	}

	return &product, nil
}

func (repository *ProductRepository) CreateProduct(product *models.Product) error {
	product.ID = uuid.NewString()

	query := `INSERT INTO "Product" (id, machineID, name, productImage, description, productLink)
	VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := repository.database.Exec(query, product.ID, product.MachineID, product.Name, product.ProductImage, product.Description, product.ProductLink)

	if err != nil {
		return fmt.Errorf("error creating product: %w", err)
	}
	
	return nil
}

func (repository *ProductRepository) UpdateMachine(id string, product *models.Product) error {
	query :=  `UPDATE "Product"
	SET machineID = $1, name = $2, productImage = $3, description = $4, productLink = $5
	WHERE id = $6`

	_, err := repository.database.Exec(query, product.MachineID, product.Name, product.ProductImage, product.Description, product.ProductLink, id)

	if err != nil {
		return fmt.Errorf("error updating product: %w", err)
	}

	return nil
}

func (repository *ProductRepository) DeleteProduct(id string) error {
	query := `DELETE FROM "Product" WHERE id = $1`

	_, err := repository.database.Exec(query, id); if err != nil {
		return fmt.Errorf("error deleting product: %w", err)
	}
	
	return nil
}