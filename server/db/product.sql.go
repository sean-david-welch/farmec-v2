// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: product.sql

package db

import (
	"context"
	"database/sql"
)

const createProduct = `-- name: CreateProduct :exec
insert into Product (id, machine_id, name, product_image, description, product_link) values (?, ?, ?, ?, ?, ?)
`

type CreateProductParams struct {
	ID           string         `json:"id"`
	MachineID    string         `json:"machine_id"`
	Name         string         `json:"name"`
	ProductImage sql.NullString `json:"product_image"`
	Description  sql.NullString `json:"description"`
	ProductLink  sql.NullString `json:"product_link"`
}

func (q *Queries) CreateProduct(ctx context.Context, arg CreateProductParams) error {
	_, err := q.db.ExecContext(ctx, createProduct,
		arg.ID,
		arg.MachineID,
		arg.Name,
		arg.ProductImage,
		arg.Description,
		arg.ProductLink,
	)
	return err
}

const deleteProduct = `-- name: DeleteProduct :exec
delete from Product where id = ?
`

func (q *Queries) DeleteProduct(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, deleteProduct, id)
	return err
}

const getProductByID = `-- name: GetProductByID :one
select id, machine_id, name, product_image, description, product_link from Product where id = ?
`

func (q *Queries) GetProductByID(ctx context.Context, id string) (Product, error) {
	row := q.db.QueryRowContext(ctx, getProductByID, id)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.MachineID,
		&i.Name,
		&i.ProductImage,
		&i.Description,
		&i.ProductLink,
	)
	return i, err
}

const getProducts = `-- name: GetProducts :many
select id, machine_id, name, product_image, description, product_link from Product
`

func (q *Queries) GetProducts(ctx context.Context) ([]Product, error) {
	rows, err := q.db.QueryContext(ctx, getProducts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Product
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.ID,
			&i.MachineID,
			&i.Name,
			&i.ProductImage,
			&i.Description,
			&i.ProductLink,
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

const updateProduct = `-- name: UpdateProduct :exec
update Product set machine_id = ?, name = ?, product_image = ?, description = ?, product_link = ? where id = ?
`

type UpdateProductParams struct {
	MachineID    string         `json:"machine_id"`
	Name         string         `json:"name"`
	ProductImage sql.NullString `json:"product_image"`
	Description  sql.NullString `json:"description"`
	ProductLink  sql.NullString `json:"product_link"`
	ID           string         `json:"id"`
}

func (q *Queries) UpdateProduct(ctx context.Context, arg UpdateProductParams) error {
	_, err := q.db.ExecContext(ctx, updateProduct,
		arg.MachineID,
		arg.Name,
		arg.ProductImage,
		arg.Description,
		arg.ProductLink,
		arg.ID,
	)
	return err
}

const updateProductNoImage = `-- name: UpdateProductNoImage :exec
update Product set machine_id = ?, name = ?, description = ?, product_link = ? where id = ?
`

type UpdateProductNoImageParams struct {
	MachineID   string         `json:"machine_id"`
	Name        string         `json:"name"`
	Description sql.NullString `json:"description"`
	ProductLink sql.NullString `json:"product_link"`
	ID          string         `json:"id"`
}

func (q *Queries) UpdateProductNoImage(ctx context.Context, arg UpdateProductNoImageParams) error {
	_, err := q.db.ExecContext(ctx, updateProductNoImage,
		arg.MachineID,
		arg.Name,
		arg.Description,
		arg.ProductLink,
		arg.ID,
	)
	return err
}
