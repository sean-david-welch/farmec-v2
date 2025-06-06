package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/sean-david-welch/farmec-v2/server/db"
	"time"

	"github.com/google/uuid"
)

type RegistrationRepo interface {
	GetRegistrations(ctx context.Context) ([]db.MachineRegistration, error)
	GetRegistrationById(ctx context.Context, id string) (*db.MachineRegistration, error)
	CreateRegistration(ctx context.Context, registration *db.MachineRegistration) error
	UpdateRegistration(ctx context.Context, id string, registration *db.MachineRegistration) error
	DeleteRegistration(ctx context.Context, id string) error
}

type RegistrationRepoImpl struct {
	queries *db.Queries
}

func NewRegistrationRepo(sql *sql.DB) *RegistrationRepoImpl {
	queries := db.New(sql)
	return &RegistrationRepoImpl{queries: queries}
}

func (repo *RegistrationRepoImpl) GetRegistrations(ctx context.Context) ([]db.MachineRegistration, error) {
	registrations, err := repo.queries.GetRegistrations(ctx)
	if err != nil {
		return nil, fmt.Errorf("error occurred while getting registration from the db: %w", err)
	}

	var result []db.MachineRegistration
	for _, registration := range registrations {
		result = append(result, db.MachineRegistration{
			ID:               registration.ID,
			DealerName:       registration.DealerName,
			DealerAddress:    registration.DealerAddress,
			OwnerName:        registration.OwnerName,
			OwnerAddress:     registration.OwnerAddress,
			MachineModel:     registration.MachineModel,
			SerialNumber:     registration.SerialNumber,
			InstallDate:      registration.InstallDate,
			InvoiceNumber:    registration.InvoiceNumber,
			CompleteSupply:   registration.CompleteSupply,
			PdiComplete:      registration.PdiComplete,
			PtoCorrect:       registration.PtoCorrect,
			MachineTestRun:   registration.MachineTestRun,
			SafetyInduction:  registration.SafetyInduction,
			OperatorHandbook: registration.OperatorHandbook,
			Date:             registration.Date,
			CompletedBy:      registration.CompletedBy,
			Created:          registration.Created,
		})
	}
	return result, nil
}

func (repo *RegistrationRepoImpl) GetRegistrationById(ctx context.Context, id string) (*db.MachineRegistration, error) {
	registration, err := repo.queries.GetRegistrationsByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("error occurred while getting a registration: %w", err)
	}
	return &registration, nil
}

func (repo *RegistrationRepoImpl) CreateRegistration(ctx context.Context, registration *db.MachineRegistration) error {
	registration.ID = uuid.NewString()
	registration.Created = sql.NullString{
		String: time.Now().String(),
		Valid:  true,
	}

	params := db.CreateRegistrationParams{
		ID:               registration.ID,
		DealerName:       registration.DealerName,
		DealerAddress:    registration.DealerAddress,
		OwnerName:        registration.OwnerName,
		OwnerAddress:     registration.OwnerAddress,
		MachineModel:     registration.MachineModel,
		SerialNumber:     registration.SerialNumber,
		InstallDate:      registration.InstallDate,
		InvoiceNumber:    registration.InvoiceNumber,
		CompleteSupply:   registration.CompleteSupply,
		PdiComplete:      registration.PdiComplete,
		PtoCorrect:       registration.PtoCorrect,
		MachineTestRun:   registration.MachineTestRun,
		SafetyInduction:  registration.SafetyInduction,
		OperatorHandbook: registration.OperatorHandbook,
		Date:             registration.Date,
		CompletedBy:      registration.CompletedBy,
		Created:          registration.Created,
	}
	if err := repo.queries.CreateRegistration(ctx, params); err != nil {
		return fmt.Errorf("error occurred while creating a registration: %w", err)
	}
	return nil
}

func (repo *RegistrationRepoImpl) UpdateRegistration(ctx context.Context, id string, registration *db.MachineRegistration) error {
	params := db.UpdateRegistrationParams{
		DealerName:       registration.DealerName,
		DealerAddress:    registration.DealerAddress,
		OwnerName:        registration.OwnerName,
		OwnerAddress:     registration.OwnerAddress,
		MachineModel:     registration.MachineModel,
		SerialNumber:     registration.SerialNumber,
		InstallDate:      registration.InstallDate,
		InvoiceNumber:    registration.InvoiceNumber,
		CompleteSupply:   registration.CompleteSupply,
		PdiComplete:      registration.PdiComplete,
		PtoCorrect:       registration.PtoCorrect,
		MachineTestRun:   registration.MachineTestRun,
		SafetyInduction:  registration.SafetyInduction,
		OperatorHandbook: registration.OperatorHandbook,
		Date:             registration.Date,
		CompletedBy:      registration.CompletedBy,
		ID:               id,
	}
	if err := repo.queries.UpdateRegistration(ctx, params); err != nil {
		return fmt.Errorf("error occurred while updating registration: %w", err)
	}

	return nil
}

func (repo *RegistrationRepoImpl) DeleteRegistration(ctx context.Context, id string) error {
	if err := repo.queries.DeleteRegistration(ctx, id); err != nil {
		return fmt.Errorf("an error occurred while deleting a registration: %w", err)
	}
	return nil
}
