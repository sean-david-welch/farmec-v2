package services

import (
	"errors"

	"github.com/sean-david-welch/farmec-v2/server/models"
	"github.com/sean-david-welch/farmec-v2/server/repository"
	"github.com/sean-david-welch/farmec-v2/server/types"
	"github.com/sean-david-welch/farmec-v2/server/utils"
)

type ProductService struct {
	folder string
	s3Client *utils.S3Client
	repository *repository.ProductRepository
}

func NewProductService(repository *repository.ProductRepository, s3Client *utils.S3Client, folder string) *ProductService {
	return &ProductService{
		repository: repository,
		s3Client: s3Client,
		folder: folder,
	}
}

func (service *ProductService) GetProducts(id string) ([]models.Product, error) {
	return service.repository.GetProducts(id)
}

func (service *ProductService) CreateProduct(product *models.Product) (*types.ModelResult, error) {
 	productImage := product.ProductImage; if productImage == "" {
		return nil, errors.New("machine image is empty")
	}

	presignedUrl, imageUrl, err := service.s3Client.GeneratePresignedUrl(service.folder, productImage)
	if err != nil {
		return nil, err
	}

	product.ProductImage = productImage

	service.repository.CreateProduct(product); if err != nil {
		return nil, err
	}

	result := &types.ModelResult{
		PresginedUrl: presignedUrl,
		ImageUrl: imageUrl,
	}

	return result, nil
}

func (serice *ProductService) UpdateProduct(id string, product *models.Product) (*types.ModelResult, error) {
	productImage := product.ProductImage

	var presignedUrl, imageUrl string
	var err error

	if productImage != "" {
		presignedUrl, imageUrl, err = serice.s3Client.GeneratePresignedUrl(serice.folder, productImage)
		if err != nil {
			return nil, err
		}
		product.ProductImage = imageUrl
	}

	serice.repository.UpdateMachine(id, product); if err != nil {
		return nil, err
	}

	result := &types.ModelResult{
		PresginedUrl: presignedUrl,
		ImageUrl: imageUrl,
	}

	return result, err
}


func (service *ProductService) DeleteProduct(id string) error {
	product, err := service.repository.GetProductById(id); if err != nil {
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

