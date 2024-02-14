-- name: GetAllUsers :many
select *
from manage_user
where company_code = ?
    and branch_code = ?;
-- name: GetBranch :one
select *
from manage_user
where id = ?
    and company_code = ?
    and branch_code = ?;
-- name: AddUser :execresult
insert into manage_user (
        id,
        email,
        firstname,
        lastname,
        user_type,
        suspended,
        title,
        password,
        image_url,
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
        ?
    );
-- name: UpdateUserPassword :exec
update manage_user
set password = ?,
    updated_at = ?,
    updated_by = ?
where id = ?;
-- name: UpdateUserNames :exec
update manage_user
set firstname = ?,
    lastname = ?,
    updated_at = ?,
    updated_by = ?
where id = ?;
-- name: UpdateUserEmail :exec
update manage_user
set email = ?,
    updated_at = ?,
    updated_by = ?
where id = ?;
-- name: UpdateUserSuspention :exec
update manage_user
set suspended = ?,
    updated_at = ?,
    updated_by = ?
where id = ?;
-- name: UpdateUser :exec
update manage_user
set email = ?,
    firstname = ?,
    lastname = ?,
    suspended = ?,
    image_url = ?,
    title = ?,
    user_type = ?,
    updated_at = ?,
    updated_by = ?
where id = ?;
-- name: DeleteUser :execresult
delete from manage_user
where id = ?;