package services

import (
	"github.com/sean-david-welch/farmec-v2/server/repository"
	"github.com/sean-david-welch/farmec-v2/server/types"
)

type WarrantyService struct {
	repository *repository.WarrantyRepository
}

func NewWarrantyService(repository *repository.WarrantyRepository) *WarrantyService {
	return &WarrantyService{repository: repository}
}

func(service *WarrantyService) GetWarranties() ([]types.DealerOwnerInfo, error) {
	warranties, err := service.repository.GetWarranties(); if err != nil {
		return nil, err
	}

	return warranties, nil
}

func(service *WarrantyService) GetWarrantyById(id string) (*types.WarrantyClaim, []types.PartsRequired, error) {
	warranty, partsRequired, err := service.repository.GetWarrantyById(id); if err != nil {
		return nil, nil, err
	}

	return warranty, partsRequired, nil
}

func(service *WarrantyService) CreateWarranty(warranty *types.WarrantyClaim, parts []types.PartsRequired) error {
	if err := service.repository.CreateWarranty(warranty, parts); err != nil {
		return err
	}

	return nil
}

func(service *WarrantyService) UpdateWarranty(id string, warranty *types.WarrantyClaim, parts []types.PartsRequired) error {
	if err := service.repository.UpdateWarranty(id, warranty, parts); err != nil {
		return err
	}

	return nil
}

func (service *WarrantyService) DeleteWarranty(id string) error {
	if err := service.repository.DeleteWarranty(id); err != nil {
		return err
	}

	return nil
}
