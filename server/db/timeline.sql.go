// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: timeline.sql

package db

import (
	"context"
	"database/sql"
)

const createTimeline = `-- name: CreateTimeline :exec
insert into Timeline (id, title, date, body, created)
values (?, ?, ?, ?, ?)
`

type CreateTimelineParams struct {
	ID      string         `json:"id"`
	Title   string         `json:"title"`
	Date    sql.NullString `json:"date"`
	Body    sql.NullString `json:"body"`
	Created sql.NullString `json:"created"`
}

func (q *Queries) CreateTimeline(ctx context.Context, arg CreateTimelineParams) error {
	_, err := q.db.ExecContext(ctx, createTimeline,
		arg.ID,
		arg.Title,
		arg.Date,
		arg.Body,
		arg.Created,
	)
	return err
}

const deleteTimeline = `-- name: DeleteTimeline :exec
delete
from Timeline
where id = ?
`

func (q *Queries) DeleteTimeline(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteTimeline, id)
	return err
}

const getTimelineByID = `-- name: GetTimelineByID :one
select id, title, date, body, created
from Timeline
where id = ?
`

func (q *Queries) GetTimelineByID(ctx context.Context, id string) (Timeline, error) {
	row := q.db.QueryRowContext(ctx, getTimelineByID, id)
	var i Timeline
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Date,
		&i.Body,
		&i.Created,
	)
	return i, err
}

const getTimelines = `-- name: GetTimelines :many
select id, title, date, body, created
from Timeline
`

func (q *Queries) GetTimelines(ctx context.Context) ([]Timeline, error) {
	rows, err := q.db.QueryContext(ctx, getTimelines)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Timeline
	for rows.Next() {
		var i Timeline
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Date,
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

const updateTimeline = `-- name: UpdateTimeline :exec
update Timeline
set title = ?,
    date  = ?,
    body  = ?
where id = ?
`

type UpdateTimelineParams struct {
	Title string         `json:"title"`
	Date  sql.NullString `json:"date"`
	Body  sql.NullString `json:"body"`
	ID    string         `json:"id"`
}

func (q *Queries) UpdateTimeline(ctx context.Context, arg UpdateTimelineParams) error {
	_, err := q.db.ExecContext(ctx, updateTimeline,
		arg.Title,
		arg.Date,
		arg.Body,
		arg.ID,
	)
	return err
}
