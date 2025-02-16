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
	GetRegistrations(request context.Context) ([]types.MachineRegistration, error)
	GetRegistrationById(request context.Context, id string) (*types.MachineRegistration, error)
	CreateRegistration(request context.Context, registration *db.MachineRegistration) error
	UpdateRegistration(request context.Context, id string, registration *db.MachineRegistration) error
	DeleteRegistration(request context.Context, id string) error
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

func (service *RegistrationServiceImpl) GetRegistrations(request context.Context) ([]types.MachineRegistration, error) {
	registrations, err := service.repo.GetRegistrations(request)
	if err != nil {
		return nil, err
	}
	var result []types.MachineRegistration
	for _, reg := range registrations {
		result = append(result, lib.SerializeMachineRegistration(reg))
	}
	return result, nil
}

func (service *RegistrationServiceImpl) GetRegistrationById(request context.Context, id string) (*types.MachineRegistration, error) {
	registration, err := service.repo.GetRegistrationById(request, id)
	if err != nil {
		return nil, err
	}
	result := lib.SerializeMachineRegistration(*registration)
	return &result, nil
}

func (service *RegistrationServiceImpl) CreateRegistration(request context.Context, registration *db.MachineRegistration) error {
	if err := service.repo.CreateRegistration(request, registration); err != nil {
		return err
	}

	go service.sendRegistrationEmail(registration)

	return nil
}

func (service *RegistrationServiceImpl) UpdateRegistration(request context.Context, id string, registration *db.MachineRegistration) error {
	if err := service.repo.UpdateRegistration(request, id, registration); err != nil {
		return err
	}

	go service.sendRegistrationEmail(registration)

	return nil
}

func (service *RegistrationServiceImpl) DeleteRegistration(request context.Context, id string) error {
	if err := service.repo.DeleteRegistration(request, id); err != nil {
		return err
	}

	return nil
}
