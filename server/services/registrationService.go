package services

import (
	"github.com/sean-david-welch/farmec-v2/server/stores"
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
	store stores.RegistrationStore
}

func NewRegistrationService(store stores.RegistrationStore) *RegistrationServiceImpl {
	return &RegistrationServiceImpl{store: store}
}

func (service *RegistrationServiceImpl) GetRegistrations() ([]types.MachineRegistration, error) {
	registrations, err := service.store.GetRegistrations()
	if err != nil {
		return nil, err
	}

	return registrations, nil
}

func (service *RegistrationServiceImpl) GetRegistrationById(id string) (*types.MachineRegistration, error) {
	registration, err := service.store.GetRegistrationById(id)
	if err != nil {
		return nil, err
	}

	return registration, nil
}

func (service *RegistrationServiceImpl) CreateRegistration(registration *types.MachineRegistration) error {
	if err := service.store.CreateRegistration(registration); err != nil {
		return err
	}

	return nil
}

func (service *RegistrationServiceImpl) UpdateRegistration(id string, registration *types.MachineRegistration) error {
	if err := service.store.UpdateRegistration(id, registration); err != nil {
		return err
	}

	return nil
}

func (service *RegistrationServiceImpl) DeleteRegistration(id string) error {
	if err := service.store.DeleteRegistration(id); err != nil {
		return err
	}

	return nil
}
