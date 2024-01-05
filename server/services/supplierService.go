package services

import (
	"github.com/sean-david-welch/farmec-v2/server/models"
	"github.com/sean-david-welch/farmec-v2/server/repository"
	"github.com/sean-david-welch/farmec-v2/server/utils"
)

type SupplierService struct {
	folder string
	s3Client *utils.S3Client
    repository *repository.SupplierRepository
}

func NewSupplierService(repository *repository.SupplierRepository, s3Client *utils.S3Client, folder string) *SupplierService {
    return &SupplierService{
		repository: repository, 
		s3Client: s3Client, 
		folder: folder,
	}
}

func (service *SupplierService) GetSuppliers() ([]models.Supplier, error) {
    return service.repository.GetSuppliers()
}

func (service *SupplierService) CreateSupplier(supplier *models.Supplier) (string, string, error) {
	logoImage := supplier.LogoImage
	// marketingImage := supplier.MarketingImage

	presignedUrl, imageUrl, err := service.s3Client.GeneratePresignedUrl(service.folder, logoImage); if err != nil {
        return "", "", err
    }

    service.repository.CreateSupplier(supplier); if err != nil {
		return "", "", err
	}

	return presignedUrl, imageUrl, nil
}

func (service *SupplierService) GetSupplierById(id string) (*models.Supplier, error) {
    return service.repository.GetSupplierById(id)
}

func (service *SupplierService) UpdateSupplier(id string, supplier *models.Supplier) error {
    return service.repository.UpdateSupplier(id, supplier)
}

func (service *SupplierService) DeleteSupplier(id string) error {
    return service.repository.DeleteSupplier(id)
}