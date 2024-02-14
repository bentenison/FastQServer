-- name: GetAllTicketsForDay :many
select *
from ticket
where company_code = ?
    and branch_code = ?
    and date(created_at) = ?;
-- name: GetTicketsByStatus :many
select *
from ticket
where company_code = ?
    and branch_code = ?
    and ticket_status = ?;
-- name: GetTicketsByUser :many
select *
from ticket
where company_code = ?
    and branch_code = ?
    and counter_id = ?;
-- name: GetTicketsByBranch :many
select *
from ticket
where company_code = ?
    and branch_code = ?;
-- name: GetTicketsByTransfer :many
select *
from ticket
where company_code = ?
    and branch_code = ?
    and transfered_to = ?;
-- name: GetTicket :one
select *
from ticket
where id = ?
    and company_code = ?
    and branch_code = ?;
-- name: AddTicket :execresult
insert into ticket (
        id,
        service,
        counter_id,
        ticket_status,
        transfered_to,
        transfered_by,
        ticket_number,
        started_serving_at,
        end_serving_at,
        ticket_name,
        branch_code,
        branch_name,
        company_code,
        company_name,
        created_at,
        updated_at,
        updated_by
    )
values(
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?
    );
-- name: UpdateTicketStatus :exec
update ticket
set ticket_status = ?,
    updated_at = ?,
    updated_by = ?
where id = ?;
-- name: UpdateTicketUser :exec
update ticket
set counter_id = ?,
    updated_at = ?,
    updated_by = ?
where id = ?;
-- name: UpdateTicketStartTime :exec
update ticket
set started_serving_at = ?,
    updated_at = ?,
    updated_by = ?
where id = ?;
-- name: UpdateTicketTransferedTo :exec
update ticket
set transfered_to = ?,
    updated_at = ?,
    updated_by = ?
where id = ?;
-- name: DeleteSection :execresult
delete from ticket
where id = ?;