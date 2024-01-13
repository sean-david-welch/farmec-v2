package services

import (
	"github.com/sean-david-welch/farmec-v2/server/models"
	"github.com/sean-david-welch/farmec-v2/server/repository"
)

type RegistrationService struct {
	repository*repository.RegistrationRepository
}

func NewRegistrationService(repository*repository.RegistrationRepository) *RegistrationService {
	return &RegistrationService{repository: repository}
}

func(service *RegistrationService) GetRegistrations() ([]models.MachineRegistration, error) {
	registrations, err := service.repository.GetRegistrations(); if err != nil {
		return nil, err
	}

	return registrations, nil
}

func(service *RegistrationService) GetRegistrationById(id string) (*models.MachineRegistration, error) {
	registration, err := service.repository.GetRegistrationById(id); if err != nil {
		return nil, err
	}

	return registration, nil
}

func(service *RegistrationService) CreateRegistration(registration *models.MachineRegistration) error {
	if err := service.repository.CreateRegistration(registration); err != nil {
		return err
	}
	
	return nil 
}

func(service *RegistrationService) UpdateRegistration(id string, registration *models.MachineRegistration) error {
	if err := service.repository.UpdateRegistration(id, registration); err != nil {
		return err
	}

	return nil
}

func(service *RegistrationService) DeleteRegistration(id string) error {
	if err := service.repository.DeleteRegistration(id); err != nil {
		return err
	}

	return nil
}

