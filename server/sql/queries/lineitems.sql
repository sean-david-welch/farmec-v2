-- name: GetLineItems :many
select id, name, price, image
from LineItems;

-- name: GetLineItemByID :one
select id, name, price, image
from LineItems
where id = ?;

-- name: CreateLineItem :exec
insert into LineItems (id, name, price, image)
values (?, ?, ?, ?);

-- name: UpdateLineItemNoImage :exec
update LineItems
set name  = ?,
    price = ?
where id = ?;

-- name: UpdateLineItem :exec
update LineItems
set name  = ?,
    price = ?,
    image = ?
where id = ?;

-- name: DeleteLineItem :exec
delete
from LineItems
where id = ?;