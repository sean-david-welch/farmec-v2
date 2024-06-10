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

func TestGetRegistration(test *testing.T) {
	db, mock, err := mocks.InitMockDatabase(test)
	if err != nil {
		test.Fatalf("Failed to init mock store: %s", err)
	}
	defer db.Close()

	registrations := []types.MachineRegistration{
		{
			ID:               "1",
			DealerName:       "Dealer One",
			DealerAddress:    "123 Dealer St, CityOne",
			OwnerName:        "Owner One",
			OwnerAddress:     "456 Owner Ave, CityTwo",
			MachineModel:     "ModelOne",
			SerialNumber:     "SN001",
			InstallDate:      "2024-01-01",
			InvoiceNumber:    "INV001",
			CompleteSupply:   true,
			PdiComplete:      true,
			PtoCorrect:       true,
			MachineTestRun:   true,
			SafetyInduction:  true,
			OperatorHandbook: true,
			Date:             "2024-01-02",
			CompletedBy:      "TechnicianOne",
			Created:          time.Now(),
		},
		{
			ID:               "2",
			DealerName:       "Dealer Two",
			DealerAddress:    "789 Dealer Blvd, CityThree",
			OwnerName:        "Owner Two",
			OwnerAddress:     "101 Owner Road, CityFour",
			MachineModel:     "ModelTwo",
			SerialNumber:     "SN002",
			InstallDate:      "2024-01-10",
			InvoiceNumber:    "INV002",
			CompleteSupply:   false,
			PdiComplete:      false,
			PtoCorrect:       false,
			MachineTestRun:   false,
			SafetyInduction:  false,
			OperatorHandbook: false,
			Date:             "2024-01-11",
			CompletedBy:      "TechnicianTwo",
			Created:          time.Now(),
		},
	}

	rows := sqlmock.NewRows([]string{
		"id", "dealer_name", "dealer_address", "owner_name", "owner_address",
		"machine_model", "serial_number", "install_date", "invoice_number",
		"complete_supply", "pdi_complete", "pto_correct", "machine_test_run",
		"safety_induction", "operator_handbook", "date", "completed_by", "created",
	})

	for _, reg := range registrations {
		rows.AddRow(
			reg.ID, reg.DealerName, reg.DealerAddress, reg.OwnerName, reg.OwnerAddress,
			reg.MachineModel, reg.SerialNumber, reg.InstallDate, reg.InvoiceNumber,
			reg.CompleteSupply, reg.PdiComplete, reg.PtoCorrect, reg.MachineTestRun,
			reg.SafetyInduction, reg.OperatorHandbook, reg.Date, reg.CompletedBy, reg.Created,
		)
	}
	mock.ExpectQuery(`SELECT \* FROM "MachineRegistration"`).WillReturnRows(rows)

	repo := repository.NewRegistrationRepository(db)
	retrieved, err := repo.GetRegistrations()
	if err != nil {
		test.Fatalf("error occurred while getting items: %s", err)
	}

	assert.NoError(test, err)
	if err == nil {
		assert.Len(test, retrieved, len(registrations))
		assert.Equal(test, registrations, retrieved)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		test.Errorf("unfullfilled expectations: %s", err)
	}
}

func TestCreateRegistration(test *testing.T) {
	db, mock, err := mocks.InitMockDatabase(test)
	if err != nil {
		test.Fatalf("failed to init mock store: %s", err)
	}
	defer db.Close()

	registration := &types.MachineRegistration{
		ID:               "1",
		DealerName:       "Dealer One",
		DealerAddress:    "123 Dealer St, CityOne",
		OwnerName:        "Owner One",
		OwnerAddress:     "456 Owner Ave, CityTwo",
		MachineModel:     "ModelOne",
		SerialNumber:     "SN001",
		InstallDate:      "2024-01-01",
		InvoiceNumber:    "INV001",
		CompleteSupply:   true,
		PdiComplete:      true,
		PtoCorrect:       true,
		MachineTestRun:   true,
		SafetyInduction:  true,
		OperatorHandbook: true,
		Date:             "2024-01-02",
		CompletedBy:      "TechnicianOne",
		Created:          time.Now(),
	}

	mock.ExpectExec(`INSERT INTO "MachineRegistration" \(\s*`+
		`"id",\s*"dealer_name",\s*"dealer_address",\s*"owner_name",\s*"owner_address",\s*"machine_model",\s*`+
		`"serial_number",\s*"install_date",\s*"invoice_number",\s*"complete_supply",\s*"pdi_complete",\s*`+
		`"pto_correct",\s*"machine_test_run",\s*"safety_induction",\s*"operator_handbook",\s*"date",\s*`+
		`"completed_by",\s*"created"\s*\)`+
		`\s*VALUES\s*\(\s*\$1,\s*\$2,\s*\$3,\s*\$4,\s*\$5,\s*\$6,\s*\$7,\s*\$8,\s*\$9,\s*\$10,\s*\$11,\s*`+
		`\$12,\s*\$13,\s*\$14,\s*\$15,\s*\$16,\s*\$17,\s*\$18\s*\)`).
		WithArgs(
			sqlmock.AnyArg(), registration.DealerName, registration.DealerAddress, registration.OwnerName, registration.OwnerAddress,
			registration.MachineModel, registration.SerialNumber, registration.InstallDate, registration.InvoiceNumber,
			registration.CompleteSupply, registration.PdiComplete, registration.PtoCorrect, registration.MachineTestRun,
			registration.SafetyInduction, registration.OperatorHandbook, registration.Date, registration.CompletedBy, sqlmock.AnyArg(),
		).WillReturnResult(sqlmock.NewResult(1, 1))

	repo := repository.NewRegistrationRepository(db)
	err = repo.CreateRegistration(registration)

	assert.NoError(test, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		test.Errorf("unfulfilled expectations: %s", err)
	}
}
