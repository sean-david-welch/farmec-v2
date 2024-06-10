package services

import (
	"errors"
	"github.com/sean-david-welch/farmec-v2/server/lib"
	"log"

	"github.com/sean-david-welch/farmec-v2/server/stores"
	"github.com/sean-david-welch/farmec-v2/server/types"
)

type ProductService interface {
	GetProducts(id string) ([]types.Product, error)
	CreateProduct(product *types.Product) (*types.ModelResult, error)
	UpdateProduct(id string, product *types.Product) (*types.ModelResult, error)
	DeleteProduct(id string) error
}

type ProductServiceImpl struct {
	folder   string
	s3Client lib.S3Client
	store    stores.ProductStore
}

func NewProductService(store stores.ProductStore, s3Client lib.S3Client, folder string) *ProductServiceImpl {
	return &ProductServiceImpl{
		store:    store,
		s3Client: s3Client,
		folder:   folder,
	}
}

func (service *ProductServiceImpl) GetProducts(id string) ([]types.Product, error) {
	return service.store.GetProducts(id)
}

func (service *ProductServiceImpl) CreateProduct(product *types.Product) (*types.ModelResult, error) {
	productImage := product.ProductImage
	if productImage == "" || productImage == "null" {
		return nil, errors.New("image is empty")
	}

	presignedUrl, imageUrl, err := service.s3Client.GeneratePresignedUrl(service.folder, productImage)
	if err != nil {
		log.Printf("error occurred while generating presigned url: %v", err)
		return nil, err
	}

	product.ProductImage = imageUrl

	err = service.store.CreateProduct(product)
	if err != nil {
		return nil, err
	}

	result := &types.ModelResult{
		PresignedUrl: presignedUrl,
		ImageUrl:     imageUrl,
	}

	return result, nil
}

func (service *ProductServiceImpl) UpdateProduct(id string, product *types.Product) (*types.ModelResult, error) {
	productImage := product.ProductImage

	var presignedUrl, imageUrl string
	var err error

	if productImage != "" && productImage != "null" {
		presignedUrl, imageUrl, err = service.s3Client.GeneratePresignedUrl(service.folder, productImage)
		if err != nil {
			log.Printf("error occurred while generating presigned url: %v", err)
			return nil, err
		}
		product.ProductImage = imageUrl
	}

	err = service.store.UpdateMachine(id, product)
	if err != nil {
		return nil, err
	}

	result := &types.ModelResult{
		PresignedUrl: presignedUrl,
		ImageUrl:     imageUrl,
	}

	return result, err
}

func (service *ProductServiceImpl) DeleteProduct(id string) error {
	product, err := service.store.GetProductById(id)
	if err != nil {
		return nil
	}

	if err := service.s3Client.DeleteImageFromS3(product.ProductImage); err != nil {
		return err
	}

	if err := service.store.DeleteProduct(id); err != nil {
		return err
	}

	return nil
}
