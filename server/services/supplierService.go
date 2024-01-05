package services

import (
	"github.com/sean-david-welch/farmec-v2/server/models"
	"github.com/sean-david-welch/farmec-v2/server/repository"
)

type SupplierService struct {
    repository *repository.SupplierRepository
}

func NewSupplierService(repository *repository.SupplierRepository) *SupplierService {
    return &SupplierService{repository: repository}
}

func (service *SupplierService) GetSuppliers() ([]models.Supplier, error) {
    return service.repository.GetSuppliers()
}

func (service *SupplierService) CreateSupplier(supplier *models.Supplier) error {
    return service.repository.CreateSupplier(supplier)
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