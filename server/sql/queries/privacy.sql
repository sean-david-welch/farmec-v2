-- name: GetPrivacies :many
select id, title, body, created
from Privacy;

-- name: GetPrivacyByID :one
select id, title, body, created
from Privacy
where id = ?;

-- name: CreatePrivacy :exec
insert into Privacy (id, title, body, created)
VALUES (?, ?, ?, ?);

-- name: UpdatePrivacy :exec
update Privacy
set title = ?,
    body  = ?
where id = ?;

-- name: DeletePrivacy :exec
delete
from Privacy
where id = ?;