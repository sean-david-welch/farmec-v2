package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/sean-david-welch/farmec-v2/server/db"
)

type ProductRepo interface {
	GetProducts(ctx context.Context, id string) ([]db.Product, error)
	GetProductById(ctx context.Context, id string) (*db.Product, error)
	CreateProduct(ctx context.Context, product *db.Product) error
	UpdateProduct(ctx context.Context, id string, product *db.Product) error
	DeleteProduct(ctx context.Context, id string) error
}

type ProductRepoImpl struct {
	queries *db.Queries
}

func NewProductRepo(sql *sql.DB) *ProductRepoImpl {
	queries := db.New(sql)
	return &ProductRepoImpl{queries: queries}
}

func (store *ProductRepoImpl) GetProducts(ctx context.Context, id string) ([]db.Product, error) {
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

func (store *ProductRepoImpl) GetProductById(ctx context.Context, id string) (*db.Product, error) {
	product, err := store.queries.GetProductByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("an error occurred while getting product: %w", err)
	}
	return &product, nil
}

func (store *ProductRepoImpl) CreateProduct(ctx context.Context, product *db.Product) error {
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

func (store *ProductRepoImpl) UpdateProduct(ctx context.Context, id string, product *db.Product) error {
	if product.ProductImage.Valid {
		params := db.UpdateProductParams{
			MachineID:    product.MachineID,
			Name:         product.Name,
			ProductImage: product.ProductImage,
			Description:  product.Description,
			ProductLink:  product.ProductLink,
			ID:           id,
		}
		if err := store.queries.UpdateProduct(ctx, params); err != nil {
			return fmt.Errorf("an error occurred while updating product: %w", err)
		}
	} else {
		params := db.UpdateProductNoImageParams{
			MachineID:   product.MachineID,
			Name:        product.Name,
			Description: product.Description,
			ProductLink: product.ProductLink,
			ID:          id,
		}
		if err := store.queries.UpdateProductNoImage(ctx, params); err != nil {
			return fmt.Errorf("an error occurred while updating product: %w", err)
		}
	}
	return nil
}

func (store *ProductRepoImpl) DeleteProduct(ctx context.Context, id string) error {
	if err := store.queries.DeleteProduct(ctx, id); err != nil {
		return fmt.Errorf("error occurred while deleting a product: %w", err)
	}
	return nil
}
