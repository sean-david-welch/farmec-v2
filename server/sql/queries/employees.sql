-- name: GetEmployees :many
select id, name, email, role, profile_image, created
from Employee
order by created desc;

-- name: GetEmployee :one
select id, name, email, role, profile_image, created
from Employee
where id = ?;

-- name: CreateEmployee :exec
insert into Employee (id, name, email, role, profile_image, created)
values (?, ?, ?, ?, ?, ?);

-- name: UpdateEmployeeNoImage :exec
update Employee
set name  = ?,
    email = ?,
    role  = ?
where id = ?;

-- name: UpdateEmployee :exec
update Employee
set name          = ?,
    email         = ?,
    role          = ?,
    profile_image = ?
where id = ?;

-- name: DeleteEmployee :exec
delete
from Employee
where id = ?;
