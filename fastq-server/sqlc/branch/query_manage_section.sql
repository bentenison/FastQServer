-- name: GetAllSections :many
select *
from manage_section
where company_code = ?;
-- name: GetSection :one
select *
from manage_section
where id = ?
    and company_code = ?;
-- name: AddSection :execresult
insert into manage_section (
        id,
        bench_wait,
        bench_process,
        description,
        branch_code,
        brach_name,
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
        ?
    );
-- name: UpdateSectionWait :exec
update manage_section
set bench_wait = ?,
    updated_at = ?,
    updated_by = ?
where id = ?;
-- name: UpdateSectionProcess :exec
update manage_section
set bench_process = ?,
    updated_at = ?,
    updated_by = ?
where id = ?;
-- name: UpdateSectionDescription :exec
update manage_section
set description = ?,
    updated_at = ?,
    updated_by = ?
where id = ?;
-- name: UpdateSection :exec
update manage_section
set name = ?,
    bench_wait=?,
    bench_process=?,
    description=?,
    updated_at = ?,
    updated_by = ?
where id = ?;
-- name: DeleteSection :execresult
delete from manage_section
where id = ?;