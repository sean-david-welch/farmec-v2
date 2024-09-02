package stores

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/sean-david-welch/farmec-v2/server/db"
	"time"

	"github.com/google/uuid"
)

type RegistrationStore interface {
	GetRegistrations(ctx context.Context) ([]db.MachineRegistration, error)
	GetRegistrationById(ctx context.Context, id string) (*db.MachineRegistration, error)
	CreateRegistration(ctx context.Context, registration *db.MachineRegistration) error
	UpdateRegistration(ctx context.Context, id string, registration *db.MachineRegistration) error
	DeleteRegistration(ctx context.Context, id string) error
}

type RegistrationStoreImpl struct {
	queries *db.Queries
}

func NewRegistrationStore(sql *sql.DB) *RegistrationStoreImpl {
	queries := db.New(sql)
	return &RegistrationStoreImpl{queries: queries}
}

func (store *RegistrationStoreImpl) GetRegistrations(ctx context.Context) ([]db.MachineRegistration, error) {
	registrations, err := store.queries.GetRegistrations(ctx)
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

func (store *RegistrationStoreImpl) GetRegistrationById(ctx context.Context, id string) (*db.MachineRegistration, error) {
	registration, err := store.queries.GetRegistrationsByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("error occurred while getting a registration: %w", err)
	}
	return &registration, nil
}

func (store *RegistrationStoreImpl) CreateRegistration(ctx context.Context, registration *db.MachineRegistration) error {
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
	if err := store.queries.CreateRegistration(ctx, params); err != nil {
		return fmt.Errorf("error occurred while creating a registration: %w", err)
	}
	return nil
}

func (store *RegistrationStoreImpl) UpdateRegistration(ctx context.Context, id string, registration *db.MachineRegistration) error {
	query := `UPDATE "MachineRegistration" SET 
	"dealer_name" = ?, "dealer_address" = ?, "owner_name" = ?, "owner_address" = ?, 
	"machine_model" = ?, "serial_number" = ?, "install_date" = ?, "invoice_number" = ?, 
	"complete_supply" = ?, "pdi_complete" = ?, "pto_correct" = ?, "machine_test_run" = ?, 
	"safety_induction" = ?, "operator_handbook" = ?, "date" = ?, "completed_by" = ?
	WHERE "id" = ?`

	_, err := store.database.Exec(
		query,
		id, registration.DealerName, registration.DealerAddress, registration.OwnerName,
		registration.OwnerAddress, registration.MachineModel, registration.SerialNumber,
		registration.InstallDate, registration.InvoiceNumber, registration.CompleteSupply,
		registration.PdiComplete, registration.PtoCorrect, registration.MachineTestRun,
		registration.SafetyInduction, registration.OperatorHandbook, registration.Date,
		registration.CompletedBy,
	)
	if err != nil {
		return fmt.Errorf("error occurred while updating registration")
	}

	return nil
}

func (store *RegistrationStoreImpl) DeleteRegistration(ctx context.Context, id string) error {
	query := `DELETE FROM "MachineRegistration" WHERE "id" = ?`

	_, err := store.database.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error occurred while deleting registration")
	}

	return nil
}
