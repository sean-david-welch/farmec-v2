-- name: GetWarranties :many
select id, dealer, owner_name
from WarrantyClaim
order by created desc;

-- name: GetWarrantyByID :many
select wc.id,
       wc.dealer,
       wc.dealer_contact,
       wc.owner_name,
       wc.machine_model,
       wc.serial_number,
       wc.install_date,
       wc.failure_date,
       wc.repair_date,
       wc.failure_details,
       wc.repair_details,
       wc.labour_hours,
       wc.completed_by,
       wc.created,
       pr.id as part_id,
       pr.part_number,
       pr.quantity_needed,
       pr.invoice_number,
       pr.description
from WarrantyClaim wc
         left join
     PartsRequired pr on wc.id = pr.warranty_id
where wc.id = ?;

-- name: CreateWarranty :exec
insert into WarrantyClaim (id, dealer, dealer_contact, owner_name, machine_model, serial_number, install_date,
                           failure_date, repair_date, failure_details, repair_details, labour_hours, completed_by,
                           created)
values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: CreatePartsRequired :exec
insert into PartsRequired (id, warranty_id, part_number, quantity_needed, invoice_number, description)
values (?, ?, ?, ?, ?, ?);

-- name: UpdateWarranty :exec
update WarrantyClaim
set dealer          = ?,
    dealer_contact  = ?,
    owner_name      = ?,
    machine_model   = ?,
    serial_number   = ?,
    install_date    = ?,
    failure_date    = ?,
    repair_date     = ?,
    failure_details = ?,
    repair_details  = ?,
    labour_hours    = ?,
    completed_by    = ?
where id = ?;

-- name: DeleteWarranty :exec
delete
from WarrantyClaim
where id = ?;

-- name: DeletePartsRequired :exec
delete
from PartsRequired
where warranty_id = ?;
