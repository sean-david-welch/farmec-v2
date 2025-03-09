-- name: GetSuppliers :many
select id,
       name,
       logo_image,
       marketing_image,
       description,
       social_facebook,
       social_instagram,
       social_linkedin,
       social_twitter,
       social_youtube,
       social_website,
       created
from Supplier
order by created desc;

-- name: GetSupplierByID :one
select id,
       name,
       logo_image,
       marketing_image,
       description,
       social_facebook,
       social_instagram,
       social_linkedin,
       social_twitter,
       social_youtube,
       social_website,
       created
from Supplier
where id = ?;

-- name: GetVideosBySupplierID :many
SELECT id,
       supplier_id,
       web_url,
       title,
       description,
       video_id,
       thumbnail_url,
       created
FROM Video
WHERE supplier_id = ?;

-- name: GetMachinesBySupplierID :many
SELECT id, supplier_id, name, machine_image, description, machine_link, created
FROM Machine
WHERE supplier_id = ?;

-- name: CreateSupplier :exec
insert into Supplier (id, name, logo_image, marketing_image, description,
                      social_facebook, social_twitter, social_instagram,
                      social_youtube, social_linkedin, social_website, created)
values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: UpdateSupplierNoImage :exec
update Supplier
set name             = ?,
    description      = ?,
    social_facebook  = ?,
    social_twitter   = ?,
    social_instagram = ?,
    social_youtube   = ?,
    social_linkedin  = ?,
    social_website   = ?
where id = ?;

-- name: UpdateSupplier :exec
update Supplier
set name             = ?,
    logo_image       = ?,
    marketing_image  = ?,
    description      = ?,
    social_facebook  = ?,
    social_twitter   = ?,
    social_instagram = ?,
    social_youtube   = ?,
    social_linkedin  = ?,
    social_website   = ?
where id = ?;

-- name: DeleteSupplier :exec
delete
from Supplier
where id = ?