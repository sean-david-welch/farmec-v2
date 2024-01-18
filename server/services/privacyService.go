package services

import (
	"github.com/sean-david-welch/farmec-v2/server/repository"
	"github.com/sean-david-welch/farmec-v2/server/types"
)

type PrivacyService struct {
	repository *repository.PrivacyRepository
}

func NewPrivacyService(repository *repository.PrivacyRepository) *PrivacyService {
	return &PrivacyService{repository: repository}
}

func(service *PrivacyService) GetPrivacys() ([]types.Privacy, error) {
	privacys, err := service.repository.GetPrivacy(); if err != nil {
		return nil, err
	}
	
	return privacys, nil
}

func(service *PrivacyService) CreatePrivacy(privacy *types.Privacy) error {
	if err := service.repository.CreatePrivacy(privacy); err != nil {
		return err
	}
	
	return nil
}

func(service *PrivacyService) UpdatePrivacy(id string, privacy *types.Privacy) error {
	if err := service.repository.UpdatePrivacy(id, privacy); err != nil {
		return err
	}

	return nil
}

func(service *PrivacyService) DeletePrivacy(id string) error {
	if err := service.repository.DeletePrivacy(id); err != nil {
		return err
	}

	return nil
}