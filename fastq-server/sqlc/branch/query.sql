-- name: GetAllBranchAdmins :many
select *
from branch_admin
where company_code = ?;
-- name: GetBranchAdmin :one
select *
from branch_admin
where branch_code = ?
    and company_code = ?;
-- name: AddbranchAdmin :execresult
insert into branch_admin (
        id,
        branch_name,
        branch_code,
        company_code,
        company_name,
        username,
        firstname,
        lastname,
        email,
        created_at,
        updated_at,
        updated_by
    )
values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);
-- name: UpdateBranchUserEmail :exec
update branch_admin
set email = ?,
    updated_at = ?,
    updated_by = ?
where id = ?;
-- name: UpdateUserEmailByEmail :exec
update branch_admin
set email = ?,
    updated_at = ?,
    updated_by = ?
where email = ?;