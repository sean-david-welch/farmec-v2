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

func TestGetEmployee(test *testing.T) {
	db, mock, err := mocks.InitMockDatabase(test)
	if err != nil {
		test.Fatalf("Failed to initialize mock database: %s", err)
	}

	defer db.Close()

	employees := []types.Employee{
		{
			ID:           "1",
			Name:         "John Doe",
			Email:        "johndoe@example.com",
			Role:         "Developer",
			ProfileImage: "johndoe.jpg",
			Created:      time.Now(),
		},
		{
			ID:           "2",
			Name:         "Jane Doe",
			Email:        "janedoe@example.com",
			Role:         "Manager",
			ProfileImage: "janedoe.jpg",
			Created:      time.Now(),
		},
	}

	rows := sqlmock.NewRows([]string{"id", "name", "email", "role", "profile_image", "created"})
	for _, employee := range employees {
		rows.AddRow(employee.ID, test.Name, employee.Email, employee.Role, employee.ProfileImage, employee.Created)
	}

	mock.ExpectQuery(`SELECT \* FROM "Employee"`).WillReturnRows(rows)

	repo := repository.NewEmployeeRepository(db)
	retrieved, err := repo.GetEmployees()
	if err != nil {
		test.Fatalf("error occurred when getting items: %s", err)
	}

	assert.NoError(test, err)
	if err == nil {
		assert.Len(test, retrieved, len(employees))
		assert.Equal(test, employees, retrieved)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		test.Errorf("expectations unfullfillled: %s", err)
	}
}

func TestCreateEmployee(test *testing.T) {
	db, mock, err := mocks.InitMockDatabase(test)
	if err != nil {
		test.Fatalf("failed to init mock database: %s", err)
	}
	defer db.Close()

	employee := &types.Employee{
		ID:           "1",
		Name:         "John Doe",
		Email:        "johndoe@example.com",
		Role:         "Developer",
		ProfileImage: "johndoe.jpg",
		Created:      time.Now(),
	}

	mock.ExpectExec(`INSERT INTO "Employee" 
		\("id", "name", "email", "role", "profile_image", "created"\) VALUES \(\$1, \$2, \$3, \$4, \$5, \$6\)`).
		WithArgs(sqlmock.AnyArg(), employee.Name, employee.Email, employee.Role, employee.ProfileImage, sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := repository.NewEmployeeRepository(db)
	err = repo.CreateEmployee(employee)

	assert.NoError(test, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		test.Errorf("unfullfilled expectations: %s", err)
	}
}
