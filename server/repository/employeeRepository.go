package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/sean-david-welch/farmec-v2/server/types"
)

type EmployeeRepository interface {
	GetEmployees() ([]types.Employee, error) 
	GetEmployeeById(id string) (*types.Employee, error) 
	CreateEmployee(employee *types.Employee) error 
	UpdateEmployee(id string, employee *types.Employee) error 
	DeleteEmployee(id string) error 
}

type EmployeeRepositoryImpl struct {
	database *sql.DB
}

func NewEmployeeRepository(database *sql.DB) *EmployeeRepositoryImpl {
	return &EmployeeRepositoryImpl{database: database}
}

func (repository *EmployeeRepositoryImpl) GetEmployees() ([]types.Employee, error) {
	var employees []types.Employee
	
	query := `SELECT * FROM "Employee"`
	rows, err := repository.database.Query(query); if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var employee types.Employee

		err := rows.Scan(&employee.ID, &employee.Name, &employee.Email, &employee.Role, &employee.Bio, &employee.ProfileImage, &employee.Created, &employee.Phone)
		if err != nil {
			return nil, fmt.Errorf("error occurred while scanning rows: %v", err)
		}
		employees = append(employees, employee)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error occurred after iterating over the rows: %w", err)
	}

	return employees, nil
}

func(repository *EmployeeRepositoryImpl) GetEmployeeById(id string) (*types.Employee, error) {
	query := `SELECT * FROM "Employee" WHERE "id" = $1`
	row := repository.database.QueryRow(query, id)

	var employee types.Employee

	err := row.Scan(&employee.ID, &employee.Name, &employee.Email, &employee.Role, &employee.Bio, &employee.ProfileImage, &employee.Created, &employee.Phone)
	if err != nil {
				if err == sql.ErrNoRows {
			return nil, fmt.Errorf("error item found with the given id: %w", err)
		}

		return nil, fmt.Errorf("error scanning row: %w", err)
	}

	return &employee, nil
}

func(repository *EmployeeRepositoryImpl) CreateEmployee(employee *types.Employee) error {
	employee.ID = uuid.NewString()
	employee.Created = time.Now()

	query := `INSERT INTO "Employee" (id, name, email, role, bio, profileImage, created, phone)
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	_, err := repository.database.Exec(query, employee.ID, employee.Name, employee.Email, employee.Role, employee.Bio, employee.ProfileImage, employee.Created, employee.Phone)
	if err != nil {
		return err
	}

	return nil
}

func(repository *EmployeeRepositoryImpl) UpdateEmployee(id string, employee *types.Employee) error {
	query := `UPDATE "Employee" SET name = $1, email = $2, role = $3, bio = $4, profileImage = $5, phone = $6 WHERE "id" = $7`

	
	_, err := repository.database.Exec(query, employee.Name, employee.Email, employee.Role, employee.Bio, employee.ProfileImage, employee.Phone, id)
	if err != nil {
		return err
	}
	
	return nil
}

func (repository *EmployeeRepositoryImpl) DeleteEmployee(id string) error {
	query := `DELETE FROM "Employee" WHERE "id" = $1`
	_, err := repository.database.Exec(query, id); if err != nil {
		return err
	}

	return nil
}