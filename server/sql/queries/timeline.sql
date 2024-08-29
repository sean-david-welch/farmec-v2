-- name: GetTimelines :many
select id, title, date, body, created from Timeline;

-- name: GetTimelineByID :one
select id, title, date, body, created from Timeline where id = ?;

-- name: CreateTimeline :exec
insert into Timeline (id, title, date, body, created) values (?, ?, ?, ?, ?);

-- name: UpdateTimeline :exec
update Timeline set title = ?, date = ?, body = ? where id = ?;

-- name: DeleteTimeline :exec
delete from Timeline where id = ?;
