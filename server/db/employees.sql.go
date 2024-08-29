// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: employees.sql

package db

import (
	"context"
	"database/sql"
)

const createEmployee = `-- name: CreateEmployee :exec
insert into Employee (id, name, email, role, profile_image, created)
values (?, ?, ?, ?, ?, ?)
`

type CreateEmployeeParams struct {
	ID           string         `json:"id"`
	Name         string         `json:"name"`
	Email        string         `json:"email"`
	Role         string         `json:"role"`
	ProfileImage sql.NullString `json:"profile_image"`
	Created      sql.NullString `json:"created"`
}

func (q *Queries) CreateEmployee(ctx context.Context, arg CreateEmployeeParams) error {
	_, err := q.db.ExecContext(ctx, createEmployee,
		arg.ID,
		arg.Name,
		arg.Email,
		arg.Role,
		arg.ProfileImage,
		arg.Created,
	)
	return err
}

const deleteEmployee = `-- name: DeleteEmployee :exec
delete from Employee where id = ?
`

func (q *Queries) DeleteEmployee(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteEmployee, id)
	return err
}

const getEmployee = `-- name: GetEmployee :one
select id, name, email, role, profile_image, created from Employee where id = ?
`

func (q *Queries) GetEmployee(ctx context.Context, id string) (Employee, error) {
	row := q.db.QueryRowContext(ctx, getEmployee, id)
	var i Employee
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Role,
		&i.ProfileImage,
		&i.Created,
	)
	return i, err
}

const getEmployees = `-- name: GetEmployees :many
select id, name, email, role, profile_image, created
from Employee order by created desc
`

func (q *Queries) GetEmployees(ctx context.Context) ([]Employee, error) {
	rows, err := q.db.QueryContext(ctx, getEmployees)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Employee
	for rows.Next() {
		var i Employee
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Email,
			&i.Role,
			&i.ProfileImage,
			&i.Created,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateEmployee = `-- name: UpdateEmployee :exec
update Employee
set name = ?, email = ?, role = ?, profile_image = ?
where id = ?
`

type UpdateEmployeeParams struct {
	Name         string         `json:"name"`
	Email        string         `json:"email"`
	Role         string         `json:"role"`
	ProfileImage sql.NullString `json:"profile_image"`
	ID           string         `json:"id"`
}

func (q *Queries) UpdateEmployee(ctx context.Context, arg UpdateEmployeeParams) error {
	_, err := q.db.ExecContext(ctx, updateEmployee,
		arg.Name,
		arg.Email,
		arg.Role,
		arg.ProfileImage,
		arg.ID,
	)
	return err
}

const updateEmployeeNoImage = `-- name: UpdateEmployeeNoImage :exec
update Employee
set name = ?, email = ?, role = ?
where id = ?
`

type UpdateEmployeeNoImageParams struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
	ID    string `json:"id"`
}

func (q *Queries) UpdateEmployeeNoImage(ctx context.Context, arg UpdateEmployeeNoImageParams) error {
	_, err := q.db.ExecContext(ctx, updateEmployeeNoImage,
		arg.Name,
		arg.Email,
		arg.Role,
		arg.ID,
	)
	return err
}
