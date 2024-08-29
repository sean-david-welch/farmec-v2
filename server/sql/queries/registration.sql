-- name: GetRegistrations :many
select id,
       dealer_name,
       dealer_address,
       owner_name,
       owner_address,
       machine_model,
       serial_number,
       install_date,
       invoice_number,
       complete_supply,
       pdi_complete,
       pto_correct,
       machine_test_run,
       safety_induction,
       operator_handbook,
       date,
       completed_by,
       created
from MachineRegistration;

-- name: GetRegistrationsByID :one
select id,
       dealer_name,
       dealer_address,
       owner_name,
       owner_address,
       machine_model,
       serial_number,
       install_date,
       invoice_number,
       complete_supply,
       pdi_complete,
       pto_correct,
       machine_test_run,
       safety_induction,
       operator_handbook,
       date,
       completed_by,
       created
from MachineRegistration
where id = ?;

-- name: CreateRegistration :exec
insert into MachineRegistration (id, dealer_name, dealer_address, owner_name, owner_address, machine_model,
                                 serial_number, install_date, invoice_number, complete_supply, pdi_complete,
                                 pto_correct, machine_test_run, safety_induction, operator_handbook, date, completed_by,
                                 created)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: UpdateRegistration :exec
update MachineRegistration
set dealer_name       = ?,
    dealer_address    = ?,
    owner_name        = ?,
    owner_address     = ?,
    machine_model     = ?,
    serial_number     = ?,
    install_date      = ?,
    invoice_number    = ?,
    complete_supply   = ?,
    pdi_complete      = ?,
    pto_correct       = ?,
    machine_test_run  = ?,
    safety_induction  = ?,
    operator_handbook = ?,
    date              = ?,
    completed_by      = ?
where id = ?;

-- name: DeleteRegistration :exec
delete
from MachineRegistration
where id = ?;