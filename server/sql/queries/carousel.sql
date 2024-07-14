-- name: GetCarousels :many
select id, name, image, created from Carousel order by created desc;

-- name: GetCarouselByID :one
select id, name, image, created from Carousel where id = ?;

-- name: CreateCarousel :exec
insert into Carousel (id, name, image, created) VALUES (?, ?, ?, ?);

-- name: UpdateCarousel :exec
update Carousel set name = ?, image = ? where id = ?;

-- name: DeleteCarousel :exec
delete from Carousel where id = ?;