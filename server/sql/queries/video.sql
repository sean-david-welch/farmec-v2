-- name: SelectVideos :many
select id,
       supplier_id,
       web_url,
       title,
       description,
       video_id,
       thumbnail_url,
       created
from Video
where supplier_id = ?;

-- name: SelectVideoByID :one
select id,
       supplier_id,
       web_url,
       title,
       description,
       video_id,
       thumbnail_url,
       created
from Video
where id = ?;

-- name: CreateVideo :exec
insert into Video (id, supplier_id, web_url, title, description, video_id, thumbnail_url, created)
values (?, ?, ?, ?, ?, ?, ?, ?);

-- name: UpdateVideo :exec
update Video
set supplier_id   = ?,
    web_url       = ?,
    title         = ?,
    description   = ?,
    video_id      = ?,
    thumbnail_url = ?
where id = ?;

-- name: DeleteVideo :exec
delete
from Video
where id = ?;
