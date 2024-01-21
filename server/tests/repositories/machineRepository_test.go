package repository

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/sean-david-welch/farmec-v2/server/repository"
	"github.com/sean-david-welch/farmec-v2/server/tests/mocks"
	"github.com/sean-david-welch/farmec-v2/server/types"
	"github.com/stretchr/testify/assert"
)

func TestGetMachine(test *testing.T) {
	db, mock, err := mocks.InitMockDatabase(test)
	if err != nil {
		test.Fatalf("Failed to init mock database: %s", err)
	}
	defer db.Close()

	machine1Name := "Machine 1"
	machine1Link := "www.google.com"
	machine2Name := "Machine 2"
	machine2Link := "www.google.com"

	machines := []types.Machine{
		{
			ID:           "1",
			SupplierID:   "12",
			Name:         "Machine 1",
			MachineImage: "image1.jpg",
			Description:  &machine1Name,
			MachineLink:  &machine1Link,
			Created:      time.Now(),
		},
		{
			ID:           "2",
			SupplierID:   "12",
			Name:         "machine 2",
			MachineImage: "image2.jpg",
			Description:  &machine2Name,
			MachineLink:  &machine2Link,
			Created:      time.Now(),
		},
	}

	supplierId := machines[0].SupplierID

	rows := sqlmock.NewRows([]string{"id", "supplierId", "name", "machine_image", "description", "machine_link", "created"})
	for _, machine := range machines {
		rows.AddRow(machine.ID, machine.SupplierID, machine.Name, machine.MachineImage, machine.Description, machine.MachineLink, machine.Created)
	}

	mock.ExpectQuery(`SELECT \* FROM "Machine"`).WillReturnRows(rows)

	repo := repository.NewMachineRepository(db)
	retrieved, err := repo.GetMachines(supplierId)
	if err != nil {
		test.Fatalf("error occurred while getting items: %s", err)
	}

	assert.NoError(test, err)
	if err == nil {
		assert.Len(test, retrieved, len(machines))
		assert.Equal(test, machines, retrieved)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		test.Errorf("unfulfilled expectations: %s", err)
	}
}

func TestCreateMachine(test *testing.T) {
	db, mock, err := mocks.InitMockDatabase(test)
	if err != nil {
		test.Fatalf("failed to init mock database: %s", err)
	}
	defer db.Close()

	machine1Name := "Machine 1"
	machine1Link := "www.google.com"

	machine := &types.Machine{
		ID:           "1",
		SupplierID:   "12",
		Name:         "Machine 1",
		MachineImage: "image1.jpg",
		Description:  &machine1Name,
		MachineLink:  &machine1Link,
		Created:      time.Now(),
	}

	mock.ExpectExec(`INSERT INTO "Machine" \(id, supplierId, name, machine_image, description, machine_link, created\) 
	VALUES \(\$1, \$2, \$3, \$4, \$5, \$6, \$7\)`).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), machine.Name, machine.MachineImage, machine.Description, machine.MachineLink, sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := repository.NewMachineRepository(db)
	err = repo.CreateMachine(machine)

	assert.NoError(test, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		test.Errorf("unfulfilled expectations: %s", err)
	}
}
