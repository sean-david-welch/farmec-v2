package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/sean-david-welch/farmec-v2/server/db"
)

type EmployeeRepo interface {
	GetEmployees(ctx context.Context) ([]db.Employee, error)
	GetEmployeeById(ctx context.Context, id string) (*db.Employee, error)
	CreateEmployee(ctx context.Context, employee *db.Employee) error
	UpdateEmployee(ctx context.Context, id string, employee *db.Employee) error
	DeleteEmployee(ctx context.Context, id string) error
}

type EmployeeRepoImpl struct {
	queries *db.Queries
}

func NewEmployeeRepo(sql *sql.DB) *EmployeeRepoImpl {
	queries := db.New(sql)
	return &EmployeeRepoImpl{queries: queries}
}

func (store *EmployeeRepoImpl) GetEmployees(ctx context.Context) ([]db.Employee, error) {
	employees, err := store.queries.GetEmployees(ctx)
	if err != nil {
		return nil, fmt.Errorf("error occurred while querying employees: %w", err)
	}

	var result []db.Employee
	for _, employee := range employees {
		result = append(result, db.Employee{
			ID:           employee.ID,
			Name:         employee.Name,
			Email:        employee.Email,
			Role:         employee.Role,
			ProfileImage: employee.ProfileImage,
			Created:      employee.Created,
		})
	}

	return result, nil
}

func (store *EmployeeRepoImpl) GetEmployeeById(ctx context.Context, id string) (*db.Employee, error) {
	employee, err := store.queries.GetEmployee(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("error occurred while querying the database for employee: %w", err)
	}

	return &employee, nil
}

func (store *EmployeeRepoImpl) CreateEmployee(ctx context.Context, employee *db.Employee) error {
	employee.ID = uuid.NewString()
	employee.Created = sql.NullString{
		String: time.Now().String(),
		Valid:  true,
	}
	params := db.CreateEmployeeParams{
		ID:           employee.ID,
		Name:         employee.Name,
		Email:        employee.Email,
		Role:         employee.Role,
		ProfileImage: employee.ProfileImage,
		Created:      employee.Created,
	}

	if err := store.queries.CreateEmployee(ctx, params); err != nil {
		return fmt.Errorf("error occurred while creating an employee: %w", err)
	}

	return nil
}

func (store *EmployeeRepoImpl) UpdateEmployee(ctx context.Context, id string, employee *db.Employee) error {
	if employee.ProfileImage.Valid {
		params := db.UpdateEmployeeParams{
			Name:         employee.Name,
			Email:        employee.Email,
			Role:         employee.Role,
			ProfileImage: employee.ProfileImage,
			ID:           id,
		}
		if err := store.queries.UpdateEmployee(ctx, params); err != nil {
			return fmt.Errorf("error occurred while updatig an employee: %w", err)
		}
	} else {
		params := db.UpdateEmployeeNoImageParams{
			Name:  employee.Name,
			Email: employee.Email,
			Role:  employee.Role,
			ID:    id,
		}
		if err := store.queries.UpdateEmployeeNoImage(ctx, params); err != nil {
			return fmt.Errorf("an error occurred while updating the employee: %w", err)
		}
	}
	return nil
}

func (store *EmployeeRepoImpl) DeleteEmployee(ctx context.Context, id string) error {
	if err := store.queries.DeleteEmployee(ctx, id); err != nil {
		return fmt.Errorf("error occurred while deleting employee: %w", err)
	}
	return nil
}
