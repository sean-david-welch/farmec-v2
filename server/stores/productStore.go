package stores

import (
	"context"
	"database/sql"
	"fmt"
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
			ID:           product.ID,
			MachineID:    product.MachineID,
			Name:         product.Name,
			ProductImage: product.ProductImage,
			Description:  product.Description,
			ProductLink:  product.ProductLink,
		})
	}
	return result, nil
}

func (store *ProductStoreImpl) GetProductById(ctx context.Context, id string) (*db.Product, error) {
	product, err := store.queries.GetProductByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("an error occurred while getting product: %w", err)
	}
	return &product, nil
}

func (store *ProductStoreImpl) CreateProduct(ctx context.Context, product *db.Product) error {
	product.ID = uuid.NewString()
	params := db.CreateProductParams{
		ID:           product.ID,
		MachineID:    product.MachineID,
		Name:         product.Name,
		ProductImage: product.ProductImage,
		Description:  product.Description,
		ProductLink:  product.ProductLink,
	}
	if err := store.queries.CreateProduct(ctx, params); err != nil {
		return fmt.Errorf("an error occurred while creating a product: %w", err)
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
