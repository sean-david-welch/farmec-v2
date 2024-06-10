package services

import (
	"github.com/sean-david-welch/farmec-v2/server/store"
	"github.com/sean-david-welch/farmec-v2/server/types"
)

type RegistrationService interface {
	GetRegistrations() ([]types.MachineRegistration, error)
	GetRegistrationById(id string) (*types.MachineRegistration, error)
	CreateRegistration(registration *types.MachineRegistration) error
	UpdateRegistration(id string, registration *types.MachineRegistration) error
	DeleteRegistration(id string) error
}

type RegistrationServiceImpl struct {
	repository store.RegistrationRepository
}

func NewRegistrationService(repository store.RegistrationRepository) *RegistrationServiceImpl {
	return &RegistrationServiceImpl{repository: repository}
}

func (service *RegistrationServiceImpl) GetRegistrations() ([]types.MachineRegistration, error) {
	registrations, err := service.repository.GetRegistrations()
	if err != nil {
		return nil, err
	}

	return registrations, nil
}

func (service *RegistrationServiceImpl) GetRegistrationById(id string) (*types.MachineRegistration, error) {
	registration, err := service.repository.GetRegistrationById(id)
	if err != nil {
		return nil, err
	}

	return registration, nil
}

func (service *RegistrationServiceImpl) CreateRegistration(registration *types.MachineRegistration) error {
	if err := service.repository.CreateRegistration(registration); err != nil {
		return err
	}

	return nil
}

func (service *RegistrationServiceImpl) UpdateRegistration(id string, registration *types.MachineRegistration) error {
	if err := service.repository.UpdateRegistration(id, registration); err != nil {
		return err
	}

	return nil
}

func (service *RegistrationServiceImpl) DeleteRegistration(id string) error {
	if err := service.repository.DeleteRegistration(id); err != nil {
		return err
	}

	return nil
}
