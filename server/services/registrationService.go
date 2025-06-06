package services

import (
	"context"
	"github.com/sean-david-welch/farmec-v2/server/db"
	"github.com/sean-david-welch/farmec-v2/server/lib"
	"github.com/sean-david-welch/farmec-v2/server/repository"
	"github.com/sean-david-welch/farmec-v2/server/types"
	"log"
)

type RegistrationService interface {
	GetRegistrations(ctx context.Context) ([]types.MachineRegistration, error)
	GetRegistrationById(ctx context.Context, id string) (*types.MachineRegistration, error)
	CreateRegistration(ctx context.Context, registration *db.MachineRegistration) error
	UpdateRegistration(ctx context.Context, id string, registration *db.MachineRegistration) error
	DeleteRegistration(ctx context.Context, id string) error
}

type RegistrationServiceImpl struct {
	repo        repository.RegistrationRepo
	emailClient *lib.EmailClientImpl
}

func NewRegistrationService(repo repository.RegistrationRepo, emailClient *lib.EmailClientImpl) *RegistrationServiceImpl {
	return &RegistrationServiceImpl{repo: repo, emailClient: emailClient}
}

func (service *RegistrationServiceImpl) sendRegistrationEmail(registration *db.MachineRegistration) {
	data := &types.EmailData{
		Name:    registration.OwnerName,
		Email:   registration.DealerName,
		Message: registration.MachineModel,
	}

	if err := service.emailClient.SendFormNotification(data, "Machine Registration"); err != nil {
		log.Printf("Failed to send registration email: %v", err)
		return
	}
}

func (service *RegistrationServiceImpl) GetRegistrations(ctx context.Context) ([]types.MachineRegistration, error) {
	registrations, err := service.repo.GetRegistrations(ctx)
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
	registration, err := service.repo.GetRegistrationById(ctx, id)
	if err != nil {
		return nil, err
	}
	result := lib.SerializeMachineRegistration(*registration)
	return &result, nil
}

func (service *RegistrationServiceImpl) CreateRegistration(ctx context.Context, registration *db.MachineRegistration) error {
	if err := service.repo.CreateRegistration(ctx, registration); err != nil {
		return err
	}

	go service.sendRegistrationEmail(registration)

	return nil
}

func (service *RegistrationServiceImpl) UpdateRegistration(ctx context.Context, id string, registration *db.MachineRegistration) error {
	if err := service.repo.UpdateRegistration(ctx, id, registration); err != nil {
		return err
	}

	go service.sendRegistrationEmail(registration)

	return nil
}

func (service *RegistrationServiceImpl) DeleteRegistration(ctx context.Context, id string) error {
	if err := service.repo.DeleteRegistration(ctx, id); err != nil {
		return err
	}

	return nil
}
