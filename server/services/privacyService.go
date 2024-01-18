package services

import (
	"github.com/sean-david-welch/farmec-v2/server/repository"
	"github.com/sean-david-welch/farmec-v2/server/types"
)

type PrivacyService interface {
	GetPrivacys() ([]types.Privacy, error) 
	CreatePrivacy(privacy *types.Privacy) error 
	UpdatePrivacy(id string, privacy *types.Privacy) error 
	DeletePrivacy(id string) error 
}

type PrivacyServiceImpl struct {
	repository repository.PrivacyRepository
}

func NewPrivacyService(repository repository.PrivacyRepository) *PrivacyServiceImpl {
	return &PrivacyServiceImpl{repository: repository}
}

func(service *PrivacyServiceImpl) GetPrivacys() ([]types.Privacy, error) {
	privacys, err := service.repository.GetPrivacy(); if err != nil {
		return nil, err
	}
	
	return privacys, nil
}

func(service *PrivacyServiceImpl) CreatePrivacy(privacy *types.Privacy) error {
	if err := service.repository.CreatePrivacy(privacy); err != nil {
		return err
	}
	
	return nil
}

func(service *PrivacyServiceImpl) UpdatePrivacy(id string, privacy *types.Privacy) error {
	if err := service.repository.UpdatePrivacy(id, privacy); err != nil {
		return err
	}

	return nil
}

func(service *PrivacyServiceImpl) DeletePrivacy(id string) error {
	if err := service.repository.DeletePrivacy(id); err != nil {
		return err
	}

	return nil
}