-- name: GetParts :many
select id, supplier_id, name, parts_image, spare_parts_link
from SpareParts
where supplier_id = ?;

-- name: GetPartByID :one
select id, supplier_id, name, parts_image, spare_parts_link
from SpareParts
where id = ?;

-- name: GetPartsSupplier :one
select id, name, logo_image, marketing_image, description, social_facebook, social_twitter, social_instagram, social_youtube, social_linkedin, social_website, created
from Supplier
where id = ?;

-- name: CreateSparePart :exec
insert into SpareParts (id, supplier_id, name, parts_image, spare_parts_link)
values (?, ?, ?, ?, ?);

-- name: UpdateSparePartNoImage :exec
update SpareParts
set supplier_id      = ?,
    name             = ?,
    spare_parts_link = ?
where id = ?;

-- name: UpdateSparePart :exec
update SpareParts
set supplier_id      = ?,
    name             = ?,
    parts_image      = ?,
    spare_parts_link = ?
where id = ?;

-- name: DeleteSparePart :exec
delete
from SpareParts
where id = ?;