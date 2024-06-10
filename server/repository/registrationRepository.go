package repository

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/sean-david-welch/farmec-v2/server/types"
)

type RegistrationRepository interface {
	GetRegistrations() ([]types.MachineRegistration, error)
	GetRegistrationById(id string) (*types.MachineRegistration, error)
	CreateRegistration(registration *types.MachineRegistration) error
	UpdateRegistration(id string, registration *types.MachineRegistration) error
	DeleteRegistration(id string) error
}

type RegistrationRepositoryImpl struct {
	database *sql.DB
}

func NewRegistrationRepository(database *sql.DB) *RegistrationRepositoryImpl {
	return &RegistrationRepositoryImpl{database: database}
}

func (repository *RegistrationRepositoryImpl) GetRegistrations() ([]types.MachineRegistration, error) {
	var registrations []types.MachineRegistration

	query := `SELECT * FROM "MachineRegistration" ORDER BY "created" DESC`
	rows, err := repository.database.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error while querying store: %w", err)
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Fatal("Failed to close store: ", err)
		}
	}()

	for rows.Next() {
		var registration types.MachineRegistration

		err := rows.Scan(
			&registration.ID, &registration.DealerName, &registration.DealerAddress,
			&registration.OwnerName, &registration.OwnerAddress, &registration.MachineModel,
			&registration.SerialNumber, &registration.InstallDate, &registration.InvoiceNumber,
			&registration.CompleteSupply, &registration.PdiComplete, &registration.PtoCorrect,
			&registration.MachineTestRun, &registration.SafetyInduction, &registration.OperatorHandbook,
			&registration.Date, &registration.CompletedBy, &registration.Created,
		)
		if err != nil {
			return nil, fmt.Errorf("error occurred while iterating over rows: %w", err)
		}

		registrations = append(registrations, registration)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error occurred after iterating over rows: %w", err)
	}

	return registrations, nil
}

func (repository *RegistrationRepositoryImpl) GetRegistrationById(id string) (*types.MachineRegistration, error) {
	var registration types.MachineRegistration

	query := `SELECT * FROM "MachineRegistration" WHERE "id" = $1`
	row := repository.database.QueryRow(query, id)

	if err := row.Scan(
		&registration.ID, &registration.DealerName, &registration.DealerAddress,
		&registration.OwnerName, &registration.OwnerAddress, &registration.MachineModel,
		&registration.SerialNumber, &registration.InstallDate, &registration.InvoiceNumber,
		&registration.CompleteSupply, &registration.PdiComplete, &registration.PtoCorrect,
		&registration.MachineTestRun, &registration.SafetyInduction, &registration.OperatorHandbook,
		&registration.Date, &registration.CompletedBy, &registration.Created,
	); err != nil {
		return nil, fmt.Errorf("error occurred while iterating over row: %w", err)
	}

	return &registration, nil
}

func (repository *RegistrationRepositoryImpl) CreateRegistration(registration *types.MachineRegistration) error {
	registration.ID = uuid.NewString()
	registration.Created = time.Now().String()

	query := `INSERT INTO "MachineRegistration" (
        "id", "dealer_name", "dealer_address", "owner_name", "owner_address", "machine_model", 
        "serial_number", "install_date", "invoice_number", "complete_supply", "pdi_complete", 
        "pto_correct", "machine_test_run", "safety_induction", "operator_handbook", "date", 
        "completed_by", "created"
    ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18)`

	_, err := repository.database.Exec(
		query,
		registration.ID, registration.DealerName, registration.DealerAddress, registration.OwnerName,
		registration.OwnerAddress, registration.MachineModel, registration.SerialNumber,
		registration.InstallDate, registration.InvoiceNumber, registration.CompleteSupply,
		registration.PdiComplete, registration.PtoCorrect, registration.MachineTestRun,
		registration.SafetyInduction, registration.OperatorHandbook, registration.Date,
		registration.CompletedBy, registration.Created,
	)
	if err != nil {
		return fmt.Errorf("error occurred while creating machine registration: %w", err)
	}

	return nil
}

func (repository *RegistrationRepositoryImpl) UpdateRegistration(id string, registration *types.MachineRegistration) error {
	query := `UPDATE "MachineRegistration" SET 
	"dealer_name" = $2, "dealer_address" = $3, "owner_name" = $4, "owner_address" = $5, 
	"machine_model" = $6, "serial_number" = $7, "install_date" = $8, "invoice_number" = $9, 
	"complete_supply" = $10, "pdi_complete" = $11, "pto_correct" = $12, "machine_test_run" = $13, 
	"safety_induction" = $14, "operator_handbook" = $15, "date" = $16, "completed_by" = $17, 
	WHERE "id" = $1`

	_, err := repository.database.Exec(
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

func (repository *RegistrationRepositoryImpl) DeleteRegistration(id string) error {
	query := `DELETE FROM "MachineRegistration" WHERE "id" = $1`

	_, err := repository.database.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error occurred while deleting registration")
	}

	return nil
}
