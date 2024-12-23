// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: carousel.sql

package db

import (
	"context"
	"database/sql"
)

const createCarousel = `-- name: CreateCarousel :exec
insert into Carousel (id, name, image, created)
VALUES (?, ?, ?, ?)
`

type CreateCarouselParams struct {
	ID      string         `json:"id"`
	Name    string         `json:"name"`
	Image   sql.NullString `json:"image"`
	Created sql.NullString `json:"created"`
}

func (q *Queries) CreateCarousel(ctx context.Context, arg CreateCarouselParams) error {
	_, err := q.db.ExecContext(ctx, createCarousel,
		arg.ID,
		arg.Name,
		arg.Image,
		arg.Created,
	)
	return err
}

const deleteCarousel = `-- name: DeleteCarousel :exec
delete
from Carousel
where id = ?
`

func (q *Queries) DeleteCarousel(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteCarousel, id)
	return err
}

const getCarouselByID = `-- name: GetCarouselByID :one
select id, name, image, created
from Carousel
where id = ?
`

func (q *Queries) GetCarouselByID(ctx context.Context, id string) (Carousel, error) {
	row := q.db.QueryRowContext(ctx, getCarouselByID, id)
	var i Carousel
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Image,
		&i.Created,
	)
	return i, err
}

const getCarousels = `-- name: GetCarousels :many
select id, name, image, created
from Carousel
order by created desc
`

func (q *Queries) GetCarousels(ctx context.Context) ([]Carousel, error) {
	rows, err := q.db.QueryContext(ctx, getCarousels)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Carousel
	for rows.Next() {
		var i Carousel
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Image,
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

const updateCarousel = `-- name: UpdateCarousel :exec
update Carousel
set name  = ?,
    image = ?
where id = ?
`

type UpdateCarouselParams struct {
	Name  string         `json:"name"`
	Image sql.NullString `json:"image"`
	ID    string         `json:"id"`
}

func (q *Queries) UpdateCarousel(ctx context.Context, arg UpdateCarouselParams) error {
	_, err := q.db.ExecContext(ctx, updateCarousel, arg.Name, arg.Image, arg.ID)
	return err
}

const updateCarouselNoImage = `-- name: UpdateCarouselNoImage :exec
update Carousel
set name = ?
where id = ?
`

type UpdateCarouselNoImageParams struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

func (q *Queries) UpdateCarouselNoImage(ctx context.Context, arg UpdateCarouselNoImageParams) error {
	_, err := q.db.ExecContext(ctx, updateCarouselNoImage, arg.Name, arg.ID)
	return err
}
