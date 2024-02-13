package repository_test

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/sean-david-welch/farmec-v2/server/repository"
	"github.com/sean-david-welch/farmec-v2/server/tests/mocks"
	"github.com/sean-david-welch/farmec-v2/server/types"
	"github.com/stretchr/testify/assert"
)

func TestGetWarrantyById(test *testing.T) {
	db, mock, err := mocks.InitMockDatabase(test)
	if err != nil {
		test.Fatalf("Failed to init mock database: %s", err)
	}
	defer db.Close()

	DealerContact := "John Doe"
	OwnerName := "Alice Johnson"
	MachineModel := "Model ABC"
	SerialNumber := "SN123456"
	InstallDate := "2024-01-15"
	FailureDate := "2024-02-10"
	RepairDate := "2024-02-20"
	FailureDetails := "Engine malfunction"
	RepairDetails := "Replaced engine"
	LabourHours := "5"
	CompletedBy := "Technician A"

	warranty := types.WarrantyClaim{
		ID:             "warranty123",
		Dealer:         "Dealer XYZ",
		DealerContact:  &DealerContact,
		OwnerName:      &OwnerName,
		MachineModel:   &MachineModel,
		SerialNumber:   &SerialNumber,
		InstallDate:    &InstallDate,
		FailureDate:    &FailureDate,
		RepairDate:     &RepairDate,
		FailureDetails: &FailureDetails,
		RepairDetails:  &RepairDetails,
		LabourHours:    &LabourHours,
		CompletedBy:    &CompletedBy,
		Created:        time.Now(),
	}

	parts := []types.PartsRequired{
		{
			ID:             "part1",
			WarrantyID:     "warranty123",
			PartNumber:     "PN123",
			QuantityNeeded: "1",
			InvoiceNumber:  "INV123",
			Description:    "Engine",
		},
		{
			ID:             "part2",
			WarrantyID:     "warranty123",
			PartNumber:     "PN124",
			QuantityNeeded: "2",
			InvoiceNumber:  "INV124",
			Description:    "Oil Filter",
		},
	}

	mock.ExpectQuery(`SELECT \* FROM "WarrantyClaim" WHERE "id" = \$1`).
		WithArgs(warranty.ID).
		WillReturnRows(sqlmock.NewRows([]string{
			"id", "dealer", "dealer_contact", "owner_name", "machine_model",
			"serial_number", "install_date", "failure_date", "repair_date",
			"failure_details", "repair_details", "labour_hours", "completed_by", "created",
		}).AddRow(
			warranty.ID, warranty.Dealer, warranty.DealerContact, warranty.OwnerName, warranty.MachineModel,
			warranty.SerialNumber, warranty.InstallDate, warranty.FailureDate, warranty.RepairDate,
			warranty.FailureDetails, warranty.RepairDetails, warranty.LabourHours, warranty.CompletedBy, warranty.Created,
		))

	rows := sqlmock.NewRows([]string{
		"id", "warrantyId", "part_number", "quantity_needed", "invoice_number", "description",
	})
	for _, p := range parts {
		rows.AddRow(
			p.ID, p.WarrantyID, p.PartNumber, p.QuantityNeeded, p.InvoiceNumber, p.Description,
		)
	}
	mock.ExpectQuery(`SELECT \* FROM "PartsRequired" WHERE "warrantyId" = \$1`).
		WithArgs(warranty.ID).
		WillReturnRows(rows)

	repo := repository.NewWarrantyRepository(db)
	retrievedWarranty, retrievedParts, err := repo.GetWarrantyById(warranty.ID)

	assert.NoError(test, err)
	assert.Equal(test, &warranty, retrievedWarranty)
	assert.ElementsMatch(test, parts, retrievedParts)

	if err := mock.ExpectationsWereMet(); err != nil {
		test.Errorf("unfulfilled expectations: %s", err)
	}
}

func TestCreateWarranty(test *testing.T) {
	db, mock, err := mocks.InitMockDatabase(test)
	if err != nil {
		test.Fatalf("failed to init mock database: %s", err)
	}
	defer db.Close()

	DealerContact := "John Doe"
	OwnerName := "Alice Johnson"
	MachineModel := "Model ABC"
	SerialNumber := "SN123456"
	InstallDate := "2024-01-15"
	FailureDate := "2024-02-10"
	RepairDate := "2024-02-20"
	FailureDetails := "Engine malfunction"
	RepairDetails := "Replaced engine"
	LabourHours := "5"
	CompletedBy := "Technician A"

	warranty := &types.WarrantyClaim{
		Dealer:         "Dealer XYZ",
		DealerContact:  &DealerContact,
		OwnerName:      &OwnerName,
		MachineModel:   &MachineModel,
		SerialNumber:   &SerialNumber,
		InstallDate:    &InstallDate,
		FailureDate:    &FailureDate,
		RepairDate:     &RepairDate,
		FailureDetails: &FailureDetails,
		RepairDetails:  &RepairDetails,
		LabourHours:    &LabourHours,
		CompletedBy:    &CompletedBy,
	}

	parts := []types.PartsRequired{
		{
			WarrantyID:     "",
			PartNumber:     "PN101",
			QuantityNeeded: "1",
			InvoiceNumber:  "INV1001",
			Description:    "Engine Block",
		},
		{
			WarrantyID:     "",
			PartNumber:     "PN102",
			QuantityNeeded: "2",
			InvoiceNumber:  "INV1002",
			Description:    "Oil Filter",
		},
	}

	mock.ExpectBegin()
	mock.ExpectExec(`INSERT INTO "WarrantyClaim"`).
		WithArgs(sqlmock.AnyArg(), warranty.Dealer, warranty.DealerContact, warranty.OwnerName,
			warranty.MachineModel, warranty.SerialNumber, warranty.InstallDate, warranty.FailureDate, warranty.RepairDate,
			warranty.FailureDetails, warranty.RepairDetails, warranty.LabourHours, warranty.CompletedBy, sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))

	for range parts {
		mock.ExpectExec(`INSERT INTO "PartsRequired"`).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
			WillReturnResult(sqlmock.NewResult(1, 1))
	}

	mock.ExpectCommit()

	repo := repository.NewWarrantyRepository(db)
	err = repo.CreateWarranty(warranty, parts)

	assert.NoError(test, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		test.Errorf("unfulfilled expectations: %s", err)
	}
}
