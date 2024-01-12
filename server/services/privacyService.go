package services

import (
	"github.com/sean-david-welch/farmec-v2/server/models"
	"github.com/sean-david-welch/farmec-v2/server/repository"
)

type PrivacyService struct {
	repository *repository.PrivacyRepository
}

func NewPrivacyService(repository *repository.PrivacyRepository) *PrivacyService {
	return &PrivacyService{repository: repository}
}

func(service *PrivacyService) GetPrivacys() ([]models.Privacy, error) {
	privacys, err := service.repository.GetPrivacy(); if err != nil {
		return nil, err
	}
	
	return privacys, nil
}

func(service *PrivacyService) CreatePrivacy(privacy *models.Privacy) error {
	if err := service.repository.CreatePrivacy(privacy); err != nil {
		return err
	}
	
	return nil
}

func(service *PrivacyService) UpdatePrivacy(id string, privacy *models.Privacy) error {
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