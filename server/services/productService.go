package services

import (
	"context"
	"database/sql"
	"errors"
	"github.com/sean-david-welch/farmec-v2/server/db"
	"github.com/sean-david-welch/farmec-v2/server/lib"
	"log"

	"github.com/sean-david-welch/farmec-v2/server/repository"
	"github.com/sean-david-welch/farmec-v2/server/types"
)

type ProductService interface {
	GetProducts(ctx context.Context, id string) ([]types.Product, error)
	CreateProduct(ctx context.Context, product *db.Product) (*types.ModelResult, error)
	UpdateProduct(ctx context.Context, id string, product *db.Product) (*types.ModelResult, error)
	DeleteProduct(ctx context.Context, id string) error
}

type ProductServiceImpl struct {
	folder   string
	s3Client lib.S3Client
	store    repository.ProductRepo
}

func NewProductService(store repository.ProductRepo, s3Client lib.S3Client, folder string) *ProductServiceImpl {
	return &ProductServiceImpl{
		store:    store,
		s3Client: s3Client,
		folder:   folder,
	}
}

func (service *ProductServiceImpl) GetProducts(ctx context.Context, id string) ([]types.Product, error) {
	products, err := service.store.GetProducts(ctx, id)
	if err != nil {
		return nil, err
	}
	var result []types.Product
	for _, product := range products {
		result = append(result, lib.SerializeProduct(product))
	}
	return result, nil
}

func (service *ProductServiceImpl) CreateProduct(ctx context.Context, product *db.Product) (*types.ModelResult, error) {
	productImage := product.ProductImage
	if !productImage.Valid {
		return nil, errors.New("image is empty")
	}

	presignedUrl, imageUrl, err := service.s3Client.GeneratePresignedUrl(service.folder, productImage.String)
	if err != nil {
		log.Printf("error occurred while generating presigned url: %v", err)
		return nil, err
	}

	product.ProductImage = sql.NullString{
		String: imageUrl,
		Valid:  true,
	}

	err = service.store.CreateProduct(ctx, product)
	if err != nil {
		return nil, err
	}

	result := &types.ModelResult{
		PresignedUrl: presignedUrl,
		ImageUrl:     imageUrl,
	}

	return result, nil
}

func (service *ProductServiceImpl) UpdateProduct(ctx context.Context, id string, product *db.Product) (*types.ModelResult, error) {
	productImage := product.ProductImage

	var presignedUrl, imageUrl string
	var err error

	if productImage.Valid {
		presignedUrl, imageUrl, err = service.s3Client.GeneratePresignedUrl(service.folder, productImage.String)
		if err != nil {
			log.Printf("error occurred while generating presigned url: %v", err)
			return nil, err
		}
		product.ProductImage = sql.NullString{
			String: imageUrl,
			Valid:  true,
		}
	}

	err = service.store.UpdateProduct(ctx, id, product)
	if err != nil {
		return nil, err
	}

	result := &types.ModelResult{
		PresignedUrl: presignedUrl,
		ImageUrl:     imageUrl,
	}

	return result, err
}

func (service *ProductServiceImpl) DeleteProduct(ctx context.Context, id string) error {
	product, err := service.store.GetProductById(ctx, id)
	if err != nil {
		return nil
	}

	if err := service.s3Client.DeleteImageFromS3(product.ProductImage.String); err != nil {
		return err
	}

	if err := service.store.DeleteProduct(ctx, id); err != nil {
		return err
	}

	return nil
}
