-- name: GetAllCounters :many
select *
from manage_counters
where company_code = ?;
-- name: GetSection :one
select *
from manage_counters
where id = ?
    and branch_code = ?;
-- name: AddSection :execresult
insert into manage_counters (
        id,
        section,
        sub_section,
        user_id,
        counter_name,
        counter_number,
        override_fifo,
        transfer_q,
        asssign_user,
        assign_service,
        transfer_priority,
        cancel_q,
        transfer_finish,
        merge_section,
        end_q,
        branch_code,
        branch_name,
        company_code,
        company_name,
        created_at,
        created_by,
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
        ?,
        ?,
        ?,
        ?,
        ?,
        ?,
        ?
    );
-- name: UpdateCounterNumber :exec
update manage_counters
set counter_number = ?,
    updated_at = ?,
    updated_by = ?
where id = ?;
-- name: UpdateCounterName :exec
update manage_counters
set counter_name = ?,
    updated_at = ?,
    updated_by = ?
where id = ?;
-- name: UpdateCounterSection :exec
update manage_counters
set section = ?,
    updated_at = ?,
    updated_by = ?
where id = ?;
-- name: UpdateCounterSubSection :exec
update manage_counters
set sub_section = ?,
    updated_at = ?,
    updated_by = ?
where id = ?;
-- name: UpdateCounterSettings :exec
update manage_counters
set override_fifo = ?,
    transfer_q = ?,
    asssign_user = ?,
    assign_service = ?,
    transfer_priority = ?,
    transfer_finish = ?,
    cancel_q = ?,
    merge_section = ?,
    end_q = ?,
    updated_at = ?,
    updated_by = ?
where id = ?;
-- name: DeleteCounter :execresult
delete from manage_counters
where id = ?;
-- name: UpdateCounter :exec
update manage_counters
set section = ?,
    sub_section = ?,
    user_id = ?,
    counter_number = ?,
    counter_name = ?,
    override_fifo = ?,
    transfer_q = ?,
    asssign_user = ?,
    assign_service = ?,
    transfer_priority = ?,
    transfer_finish = ?,
    cancel_q = ?,
    merge_section = ?,
    end_q = ?,
    updated_at = ?,
    updated_by = ?
where id = ?;