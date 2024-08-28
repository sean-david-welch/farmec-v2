package services

import (
	"github.com/sean-david-welch/farmec-v2/server/lib"
	"github.com/sean-david-welch/farmec-v2/server/store"
	"github.com/sean-david-welch/farmec-v2/server/types"
	"net/smtp"
)

type RegistrationService interface {
	GetRegistrations() ([]types.MachineRegistration, error)
	GetRegistrationById(id string) (*types.MachineRegistration, error)
	CreateRegistration(registration *types.MachineRegistration) error
	UpdateRegistration(id string, registration *types.MachineRegistration) error
	DeleteRegistration(id string) error
}

type RegistrationServiceImpl struct {
	smtpClient lib.SMTPClient
	store      store.RegistrationStore
}

func NewRegistrationService(store store.RegistrationStore, smtpClient lib.SMTPClient) *RegistrationServiceImpl {
	return &RegistrationServiceImpl{store: store, smtpClient: smtpClient}
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

	client, err := service.smtpClient.SetupSMTPClient()
	if err != nil {
		return err
	}
	defer func(client *smtp.Client) {
		err := client.Close()
		if err != nil {
			return
		}
	}(client)

	data := &types.EmailData{
		Name:    registration.OwnerName,
		Email:   registration.DealerName,
		Message: registration.MachineModel,
	}

	if err := service.smtpClient.SendFormNotification(client, data, "Machine Registration"); err != nil {
		return err
	}

	return nil
}

func (service *RegistrationServiceImpl) UpdateRegistration(id string, registration *types.MachineRegistration) error {
	if err := service.store.UpdateRegistration(id, registration); err != nil {
		return err
	}

	client, err := service.smtpClient.SetupSMTPClient()
	if err != nil {
		return err
	}
	defer func(client *smtp.Client) {
		err := client.Close()
		if err != nil {
			return
		}
	}(client)

	data := &types.EmailData{
		Name:    registration.OwnerName,
		Email:   registration.DealerName,
		Message: registration.MachineModel,
	}

	if err := service.smtpClient.SendFormNotification(client, data, "Machine Registration"); err != nil {
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
