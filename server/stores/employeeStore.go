package stores

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/sean-david-welch/farmec-v2/server/db"
)

type EmployeeStore interface {
	GetEmployees(ctx context.Context) ([]db.Employee, error)
	GetEmployeeById(ctx context.Context, id string) (*db.Employee, error)
	CreateEmployee(ctx context.Context, employee *db.Employee) error
	UpdateEmployee(ctx context.Context, id string, employee *db.Employee) error
	DeleteEmployee(ctx context.Context, id string) error
}

type EmployeeStoreImpl struct {
	queries *db.Queries
}

func NewEmployeeStore(sql *sql.DB) *EmployeeStoreImpl {
	queries := db.New(sql)
	return &EmployeeStoreImpl{queries: queries}
}

func (store *EmployeeStoreImpl) GetEmployees(ctx context.Context) ([]db.Employee, error) {

}

func (store *EmployeeStoreImpl) GetEmployeeById(id string) (*db.Employee, error) {
	query := `SELECT * FROM "Employee" WHERE "id" = ?`
	row := store.database.QueryRow(query, id)

	var employee db.Employee

	err := row.Scan(&employee.ID, &employee.Name, &employee.Email, &employee.Role, &employee.ProfileImage, &employee.Created)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("error item found with the given id: %w", err)
		}

		return nil, fmt.Errorf("error scanning row: %w", err)
	}

	return &employee, nil
}

func (store *EmployeeStoreImpl) CreateEmployee(employee *db.Employee) error {
	employee.ID = uuid.NewString()
	employee.Created = time.Now().String()

	query := `INSERT INTO "Employee" (id, name, email, role, profile_image, created)
				VALUES (?, ?, ?, ?, ?, ?)`

	_, err := store.database.Exec(query, employee.ID, employee.Name, employee.Email, employee.Role, employee.ProfileImage, employee.Created)
	if err != nil {
		return err
	}

	return nil
}

func (store *EmployeeStoreImpl) UpdateEmployee(id string, employee *db.Employee) error {
	query := `UPDATE "Employee" SET name = ?, email = ?, role = ? WHERE id = ?`
	args := []interface{}{id, employee.Name, employee.Email, employee.Role}

	if employee.ProfileImage != "" && employee.ProfileImage != "null" {
		query = `UPDATE "Employee" SET name = ?, email = ?, role = ?, profile_image = ? WHERE "id" = ?`
		args = []interface{}{id, employee.Name, employee.Email, employee.Role, employee.ProfileImage}
	}

	_, err := store.database.Exec(query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (store *EmployeeStoreImpl) DeleteEmployee(id string) error {
	query := `DELETE FROM "Employee" WHERE "id" = ?`
	_, err := store.database.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
