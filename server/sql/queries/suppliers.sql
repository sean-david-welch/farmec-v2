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

-- name :GetSupplierWithResoueces :many
select s.id, s.name, s.logo_image, s.marketing_image, s.description, s.social_facebook, s.social_twitter, s.social_instagram, s.social_youtube, s.social_linkedin, s.social_website, s.created, v.id, v.supplier_id, v.web_url, v.title, v.description, v.video_id, v.thumbnail_url, v.created, m.id, m.supplier_id, m.name, m.machine_image, m.description, m.machine_link, m.created
from Supplier s
left join
    Video v on v.supplier_id = s.id
left join
    Machine m on m.supplier_id = s.id
order by s.created desc;

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