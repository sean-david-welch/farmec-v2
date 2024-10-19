package services

import (
	"context"
	"github.com/sean-david-welch/farmec-v2/server/db"
	"github.com/sean-david-welch/farmec-v2/server/lib"
	"github.com/sean-david-welch/farmec-v2/server/repository"
	"github.com/sean-david-welch/farmec-v2/server/types"
	"net/smtp"
)

type RegistrationService interface {
	GetRegistrations(ctx context.Context) ([]types.MachineRegistration, error)
	GetRegistrationById(ctx context.Context, id string) (*types.MachineRegistration, error)
	CreateRegistration(ctx context.Context, registration *db.MachineRegistration) error
	UpdateRegistration(ctx context.Context, id string, registration *db.MachineRegistration) error
	DeleteRegistration(ctx context.Context, id string) error
}

type RegistrationServiceImpl struct {
	smtpClient lib.SMTPClient
	store      repository.RegistrationRepo
}

func NewRegistrationService(store repository.RegistrationRepo, smtpClient lib.SMTPClient) *RegistrationServiceImpl {
	return &RegistrationServiceImpl{store: store, smtpClient: smtpClient}
}

func (service *RegistrationServiceImpl) GetRegistrations(ctx context.Context) ([]types.MachineRegistration, error) {
	registrations, err := service.store.GetRegistrations(ctx)
	if err != nil {
		return nil, err
	}
	var result []types.MachineRegistration
	for _, reg := range registrations {
		result = append(result, lib.SerializeMachineRegistration(reg))
	}
	return result, nil
}

func (service *RegistrationServiceImpl) GetRegistrationById(ctx context.Context, id string) (*types.MachineRegistration, error) {
	registration, err := service.store.GetRegistrationById(ctx, id)
	if err != nil {
		return nil, err
	}
	result := lib.SerializeMachineRegistration(*registration)
	return &result, nil
}

func (service *RegistrationServiceImpl) CreateRegistration(ctx context.Context, registration *db.MachineRegistration) error {
	if err := service.store.CreateRegistration(ctx, registration); err != nil {
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

func (service *RegistrationServiceImpl) UpdateRegistration(ctx context.Context, id string, registration *db.MachineRegistration) error {
	if err := service.store.UpdateRegistration(ctx, id, registration); err != nil {
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

func (service *RegistrationServiceImpl) DeleteRegistration(ctx context.Context, id string) error {
	if err := service.store.DeleteRegistration(ctx, id); err != nil {
		return err
	}

	return nil
}
