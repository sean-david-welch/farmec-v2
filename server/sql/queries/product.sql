-- name: GetProducts :many
select id, machine_id, name, product_image, description, product_link
from Product
where machine_id = ?;

-- name: GetProductByID :one
select id, machine_id, name, product_image, description, product_link
from Product
where id = ?;

-- name: CreateProduct :exec
insert into Product (id, machine_id, name, product_image, description, product_link)
values (?, ?, ?, ?, ?, ?);

-- name: UpdateProductNoImage :exec
update Product
set machine_id   = ?,
    name         = ?,
    description  = ?,
    product_link = ?
where id = ?;

-- name: UpdateProduct :exec
update Product
set machine_id    = ?,
    name          = ?,
    product_image = ?,
    description   = ?,
    product_link  = ?
where id = ?;

-- name: DeleteProduct :exec
delete
from Product
where id = ?;