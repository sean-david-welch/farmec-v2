package services

import (
	"errors"
	"log"

	"github.com/sean-david-welch/farmec-v2/server/repository"
	"github.com/sean-david-welch/farmec-v2/server/types"
	"github.com/sean-david-welch/farmec-v2/server/utils"
)

type ProductService interface {
	GetProducts(id string) ([]types.Product, error)
	CreateProduct(product *types.Product) (*types.ModelResult, error)
	UpdateProduct(id string, product *types.Product) (*types.ModelResult, error)
	DeleteProduct(id string) error
}

type ProductServiceImpl struct {
	folder     string
	s3Client   utils.S3Client
	repository repository.ProductRepository
}

func NewProductService(repository repository.ProductRepository, s3Client utils.S3Client, folder string) *ProductServiceImpl {
	return &ProductServiceImpl{
		repository: repository,
		s3Client:   s3Client,
		folder:     folder,
	}
}

func (service *ProductServiceImpl) GetProducts(id string) ([]types.Product, error) {
	return service.repository.GetProducts(id)
}

func (service *ProductServiceImpl) CreateProduct(product *types.Product) (*types.ModelResult, error) {
	productImage := product.ProductImage
	if productImage == "" {
		return nil, errors.New("machine image is empty")
	}

	presignedUrl, imageUrl, err := service.s3Client.GeneratePresignedUrl(service.folder, productImage)
	if err != nil {
		log.Printf("error occurred while generating presigned url: %v", err)
		return nil, err
	}

	product.ProductImage = imageUrl

	service.repository.CreateProduct(product)
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

	service.repository.UpdateMachine(id, product)
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
	product, err := service.repository.GetProductById(id)
	if err != nil {
		return nil
	}

	if err := service.s3Client.DeleteImageFromS3(product.ProductImage); err != nil {
		return err
	}

	if err := service.repository.DeleteProduct(id); err != nil {
		return err
	}

	return nil
}
