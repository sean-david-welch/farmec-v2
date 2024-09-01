-- name: GetTerms :many
select id, title, body, created
from Terms;

-- name: GetTermByID :one
select id, title, body, created
from Terms
where id = ?;

-- name: CreateTerm :exec
insert into Terms (id, title, body, created)
values (?, ?, ?, ?);

-- name: UpdateTerm :exec
update Terms
set title   = ?,
    body    = ?,
    created = ?
where id = ?;

-- name: DeleteTerm :exec
delete
from Terms
where id = ?;