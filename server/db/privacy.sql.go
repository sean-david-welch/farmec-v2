// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: privacy.sql

package db

import (
	"context"
	"database/sql"
)

const createPrivacy = `-- name: CreatePrivacy :exec
insert into Privacy (id, title, body, created)
VALUES (?, ?, ?, ?)
`

type CreatePrivacyParams struct {
	ID      string         `json:"id"`
	Title   string         `json:"title"`
	Body    sql.NullString `json:"body"`
	Created sql.NullString `json:"created"`
}

func (q *Queries) CreatePrivacy(ctx context.Context, arg CreatePrivacyParams) error {
	_, err := q.db.ExecContext(ctx, createPrivacy,
		arg.ID,
		arg.Title,
		arg.Body,
		arg.Created,
	)
	return err
}

const deletePrivacy = `-- name: DeletePrivacy :exec
delete
from Privacy
where id = ?
`

func (q *Queries) DeletePrivacy(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deletePrivacy, id)
	return err
}

const getPrivacies = `-- name: GetPrivacies :many
select id, title, body, created
from Privacy
`

func (q *Queries) GetPrivacies(ctx context.Context) ([]Privacy, error) {
	rows, err := q.db.QueryContext(ctx, getPrivacies)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Privacy
	for rows.Next() {
		var i Privacy
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Body,
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

const getPrivacyByID = `-- name: GetPrivacyByID :one
select id, title, body, created
from Privacy
where id = ?
`

func (q *Queries) GetPrivacyByID(ctx context.Context, id string) (Privacy, error) {
	row := q.db.QueryRowContext(ctx, getPrivacyByID, id)
	var i Privacy
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Body,
		&i.Created,
	)
	return i, err
}

const updatePrivacy = `-- name: UpdatePrivacy :exec
update Privacy
set title = ?,
    body  = ?
where id = ?
`

type UpdatePrivacyParams struct {
	Title string         `json:"title"`
	Body  sql.NullString `json:"body"`
	ID    string         `json:"id"`
}

func (q *Queries) UpdatePrivacy(ctx context.Context, arg UpdatePrivacyParams) error {
	_, err := q.db.ExecContext(ctx, updatePrivacy, arg.Title, arg.Body, arg.ID)
	return err
}
