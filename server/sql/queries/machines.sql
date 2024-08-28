-- name: GetMachines :many
select id, supplier_id, name, machine_image, description, machine_link, created
from Machine
where supplier_id = ?;

-- name :GetMachineByID :one
select id, supplier_id, name, machine_image, description, machine_link, created
from Machine
where id = ?;

-- name :CreateMachine :exec
insert into Machine (id, supplier_id, name, machine_image, description, machine_link, created)
values (?, ?, ?, ?, ?, ?, ?);

--name :UpdateMachineNoImage :exec
update Machine
set supplier_id  = ?,
    name         = ?,
    description  = ?,
    machine_link = ?
where id = ?;

--name :UpdateMachine :exec
update Machine
set supplier_id   = ?,
    name          = ?,
    machine_image = ?,
    description   = ?,
    machine_link  = ?
where id = ?;

--name :DeleteMachine :exec
delete from Machine where id = ?;