-- name: GetAllBranches :many
select *
from manage_branch
where company_code = ?;
-- name: GetBranch :one
select *
from manage_branch
where code = ?
    and company_code = ?;
-- name: Addbranch :execresult
insert into manage_branch (
        id,
        name,
        code,
        company_code,
        license,
        services,
        address,
        phone,
        license_key,
        check_in_url,
        ticket_page_url,
        display_url,
        created_at,
        updated_at,
        updated_by
    )
values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);
-- name: UpdateBranchName :exec
update manage_branch
set name = ?,
    updated_at = ?,
    updated_by = ?
where code = ?;
-- name: UpdateBranchServices :exec
update manage_branch
set services = ?,
    updated_at = ?,
    updated_by = ?
where code = ?;
-- name: UpdateBranchAddress :exec
update manage_branch
set address = ?,
    updated_at = ?,
    updated_by = ?
where code = ?;
-- name: UpdateBranchLicenseKey :exec
update manage_branch
set license_key = ?,
    updated_at = ?,
    updated_by = ?
where code = ?;
-- name: UpdateBranchLicense :exec
update manage_branch
set license = ?,
    updated_at = ?,
    updated_by = ?
where code = ?;
-- name: UpdateBranchPhone :exec
update manage_branch
set phone = ?,
    updated_at = ?,
    updated_by = ?
where code = ?;
-- name: DeleteBranch :execresult
delete from manage_branch
where code = ?;