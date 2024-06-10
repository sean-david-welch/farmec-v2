package services

import (
	"github.com/sean-david-welch/farmec-v2/server/stores"
	"github.com/sean-david-welch/farmec-v2/server/types"
)

type WarrantyService interface {
	GetWarranties() ([]types.DealerOwnerInfo, error)
	GetWarrantyById(id string) (*types.WarrantyClaim, []types.PartsRequired, error)
	CreateWarranty(warranty *types.WarrantyClaim, parts []types.PartsRequired) error
	UpdateWarranty(id string, warranty *types.WarrantyClaim, parts []types.PartsRequired) error
	DeleteWarranty(id string) error
}

type WarrantyServiceImpl struct {
	store stores.WarrantyStore
}

func NewWarrantyService(store stores.WarrantyStore) *WarrantyServiceImpl {
	return &WarrantyServiceImpl{store: store}
}

func (service *WarrantyServiceImpl) GetWarranties() ([]types.DealerOwnerInfo, error) {
	warranties, err := service.store.GetWarranties()
	if err != nil {
		return nil, err
	}

	return warranties, nil
}

func (service *WarrantyServiceImpl) GetWarrantyById(id string) (*types.WarrantyClaim, []types.PartsRequired, error) {
	warranty, partsRequired, err := service.store.GetWarrantyById(id)
	if err != nil {
		return nil, nil, err
	}

	return warranty, partsRequired, nil
}

func (service *WarrantyServiceImpl) CreateWarranty(warranty *types.WarrantyClaim, parts []types.PartsRequired) error {
	if err := service.store.CreateWarranty(warranty, parts); err != nil {
		return err
	}

	return nil
}

func (service *WarrantyServiceImpl) UpdateWarranty(id string, warranty *types.WarrantyClaim, parts []types.PartsRequired) error {
	if err := service.store.UpdateWarranty(id, warranty, parts); err != nil {
		return err
	}

	return nil
}

func (service *WarrantyServiceImpl) DeleteWarranty(id string) error {
	if err := service.store.DeleteWarranty(id); err != nil {
		return err
	}

	return nil
}
