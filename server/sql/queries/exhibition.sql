-- name: GetExhibitions :many
select id, title, date, location, info, created
from Exhibition
order by created desc;

-- name: GetExhibitionByID :one
select id, title, date, location, info, created
from Exhibition
where id = ?;

-- name: CreateExhibition :exec
insert into Exhibition (id, title, date, location, info, created)
values (?, ?, ?, ?, ?, ?);

-- name: UpdateExhibition :exec
update Exhibition
set title    = ?,
    date     = ?,
    location = ?,
    info     = ?
where id = ?;

-- name: DeleteExhibition :exec
delete
from Exhibition
where id = ?;