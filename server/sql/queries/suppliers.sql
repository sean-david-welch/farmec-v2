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
       created,
       slug
from Supplier
where id = ?;

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