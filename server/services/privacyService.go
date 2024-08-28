package services

import (
	"github.com/sean-david-welch/farmec-v2/server/store"
	"github.com/sean-david-welch/farmec-v2/server/types"
)

type PrivacyService interface {
	GetPrivacys() ([]types.Privacy, error)
	CreatePrivacy(privacy *types.Privacy) error
	UpdatePrivacy(id string, privacy *types.Privacy) error
	DeletePrivacy(id string) error
}

type PrivacyServiceImpl struct {
	store store.PrivacyStore
}

func NewPrivacyService(store store.PrivacyStore) *PrivacyServiceImpl {
	return &PrivacyServiceImpl{store: store}
}

func (service *PrivacyServiceImpl) GetPrivacys() ([]types.Privacy, error) {
	privacys, err := service.store.GetPrivacy()
	if err != nil {
		return nil, err
	}

	return privacys, nil
}

func (service *PrivacyServiceImpl) CreatePrivacy(privacy *types.Privacy) error {
	if err := service.store.CreatePrivacy(privacy); err != nil {
		return err
	}

	return nil
}

func (service *PrivacyServiceImpl) UpdatePrivacy(id string, privacy *types.Privacy) error {
	if err := service.store.UpdatePrivacy(id, privacy); err != nil {
		return err
	}

	return nil
}

func (service *PrivacyServiceImpl) DeletePrivacy(id string) error {
	if err := service.store.DeletePrivacy(id); err != nil {
		return err
	}

	return nil
}
