package services

import (
	"github.com/sean-david-welch/farmec-v2/server/repository"
	"github.com/sean-david-welch/farmec-v2/server/types"
)

type RegistrationService struct {
	repository*repository.RegistrationRepository
}

func NewRegistrationService(repository*repository.RegistrationRepository) *RegistrationService {
	return &RegistrationService{repository: repository}
}

func(service *RegistrationService) GetRegistrations() ([]types.MachineRegistration, error) {
	registrations, err := service.repository.GetRegistrations(); if err != nil {
		return nil, err
	}

	return registrations, nil
}

func(service *RegistrationService) GetRegistrationById(id string) (*types.MachineRegistration, error) {
	registration, err := service.repository.GetRegistrationById(id); if err != nil {
		return nil, err
	}

	return registration, nil
}

func(service *RegistrationService) CreateRegistration(registration *types.MachineRegistration) error {
	if err := service.repository.CreateRegistration(registration); err != nil {
		return err
	}
	
	return nil 
}

func(service *RegistrationService) UpdateRegistration(id string, registration *types.MachineRegistration) error {
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

