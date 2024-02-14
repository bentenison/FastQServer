-- name: GetAllServices :many
select *
from manage_service
where company_code = ?;
-- name: GetBranch :one
select *
from manage_service
where id = ?
    and company_code = ?;
-- name: Addbranch :execresult
insert into manage_service (
        id,
        name,
        code,
        prefix,
        number_starts,
        number_ends,
        hide,
        show_display,
        description,
        start_time,
        end_time,
        default_time,
        workflow,
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
        ?,
        ?,
        ?,
        ?
    );
-- name: UpdateServicePrefix :exec
update manage_service
set prefix = ?,
    updated_at = ?,
    updated_by = ?
where id = ?;
-- name: UpdateServiceStartTime :exec
update manage_service
set start_time = ?,
    updated_at = ?,
    updated_by = ?
where code = ?;
-- name: UpdateServiceEndTime :exec
update manage_service
set end_time = ?,
    updated_at = ?,
    updated_by = ?
where id = ?;
-- name: UpdateServiceDefaultTime :exec
update manage_service
set default_time = ?,
    updated_at = ?,
    updated_by = ?
where id = ?;
-- name: UpdateServiceName :exec
update manage_service
set name = ?,
    updated_at = ?,
    updated_by = ?
where id = ?;
-- name: UpdateServiceNumberStart :exec
update manage_service
set number_starts = ?,
    updated_at = ?,
    updated_by = ?
where id = ?;
-- name: UpdateService :exec
update manage_service
set name = ?,
    code = ?,
    prefix = ?,
    number_starts = ?,
    number_ends = ?,
    hide = ?,
    show_display = ?,
    start_time = ?,
    end_time = ?,
    default_time = ?,
    workflow = ?,
    number_ends = ?,
    updated_at = ?,
    updated_by = ?
where id = ?;
-- name: DeleteBranch :execresult
delete from manage_service
where id = ?;