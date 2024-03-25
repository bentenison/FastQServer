package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/bentenison/fastq-server/models"
)

type sqlAdminRepo struct {
	db *sql.DB
}

func NewAdminRepo(db *sql.DB) models.AdminRepository {
	return &sqlAdminRepo{
		db: db,
	}
}

const addbranch = `-- name: Addbranch :execresult
insert into manage_branch (
        id,
        name,
        code,
        company_code,
        license,
        services,
        address,
        phone,
		email,
		password,
        license_key,
        check_in_url,
        ticket_page_url,
        display_url,
        created_at,
        updated_at,
        updated_by
    )
values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,?,?)
`

func (q *sqlAdminRepo) Addbranch(ctx context.Context, arg models.AddbranchParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, addbranch,
		arg.ID,
		arg.Name,
		arg.Code,
		arg.CompanyCode,
		arg.License,
		arg.Services,
		arg.Address,
		arg.Phone,
		arg.Email,
		arg.Password,
		arg.LicenseKey,
		arg.CheckInUrl,
		arg.TicketPageUrl,
		arg.DisplayUrl,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.UpdatedBy,
	)
}

const deleteBranch = `-- name: DeleteBranch :execresult
delete from manage_branch
where code = ?
`

func (q *sqlAdminRepo) DeleteBranch(ctx context.Context, code string) (sql.Result, error) {
	return q.db.ExecContext(ctx, deleteBranch, code)
}

const getAllBranches = `-- name: GetAllBranches :many
select id, code, name, license, services, address, phone, license_key, check_in_url, ticket_page_url, display_url, company_code, created_at, created_by, updated_at, updated_by, printer_details
from manage_branch
where company_code = ?
`

func (q *sqlAdminRepo) GetAllBranches(ctx context.Context, companyCode string) ([]*models.ManageBranch, error) {
	rows, err := q.db.QueryContext(ctx, getAllBranches, companyCode)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*models.ManageBranch
	for rows.Next() {
		var i models.ManageBranch
		if err := rows.Scan(
			&i.ID,
			&i.Code,
			&i.Name,
			&i.License,
			&i.Services,
			&i.Address,
			&i.Phone,
			&i.LicenseKey,
			&i.CheckInUrl,
			&i.TicketPageUrl,
			&i.DisplayUrl,
			&i.CompanyCode,
			&i.CreatedAt,
			&i.CreatedBy,
			&i.UpdatedAt,
			&i.UpdatedBy,
			&i.PrinterDetails,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getBranch = `-- name: GetBranch :one
select id, code, name, license, services, address, phone, license_key, check_in_url, ticket_page_url, display_url, company_code, created_at, created_by, updated_at, updated_by, printer_details
from manage_branch
where code = ?
    and company_code = ?
`

func (q *sqlAdminRepo) GetBranch(ctx context.Context, arg models.GetBranchParams) (*models.ManageBranch, error) {
	row := q.db.QueryRowContext(ctx, getBranch, arg.Code, arg.CompanyCode)
	var i models.ManageBranch
	err := row.Scan(
		&i.ID,
		&i.Code,
		&i.Name,
		&i.License,
		&i.Services,
		&i.Address,
		&i.Phone,
		&i.LicenseKey,
		&i.CheckInUrl,
		&i.TicketPageUrl,
		&i.DisplayUrl,
		&i.CompanyCode,
		&i.CreatedAt,
		&i.CreatedBy,
		&i.UpdatedAt,
		&i.UpdatedBy,
		&i.PrinterDetails,
	)
	return &i, err
}

func (q *sqlAdminRepo) GetBranchByEmail(ctx context.Context, email string) (*models.BranchAdmin, error) {
	user := &models.BranchAdmin{}
	stmt, err := q.db.PrepareContext(ctx, "select * from branch_admin where email = ?")
	if err != nil {
		log.Printf("cannot create a sql statement reason: %v ", stmt)
		return nil, err

	}
	defer stmt.Close()
	err = stmt.QueryRowContext(ctx, email).Scan(&user.BranchName, &user.BranchCode, &user.Email, &user.Password,
		&user.CompanyCode, &user.Company, &user.Isdeleted, &user.CreatedAt, &user.UpdatedAt, &user.UpdatedBy)
	switch {
	case err == sql.ErrNoRows:
		log.Printf("no user with email %s", email)
		return nil, sql.ErrNoRows
	case err != nil:
		log.Printf("cannot query database ,reson: %v", err.Error())
		return nil, err
	}
	// if err != nil {
	// 	log.Printf("could not query : %v ", stmt)
	// 	return nil, err

	// }
	return user, nil
}

func (q *sqlAdminRepo) GetUserByEmail(ctx context.Context, email string) (*models.UserAccount, error) {
	user := &models.UserAccount{}
	stmt, err := q.db.PrepareContext(ctx, "select * from company_info where email = ?")
	if err != nil {
		log.Printf("cannot create a sql statement reason: %v ", stmt)
		return nil, err

	}
	defer stmt.Close()
	err = stmt.QueryRowContext(ctx, email).Scan(&user.UserID, &user.Email, &user.FirstName, &user.LastName,
		&user.CompanyCode, &user.Company, &user.Country, &user.Timezone, &user.Password, &user.UpdatedBy, &user.CreatedAt, &user.UpdatedAt)
	switch {
	case err == sql.ErrNoRows:
		log.Printf("no user with email %s", email)
		return nil, sql.ErrNoRows
	case err != nil:
		log.Printf("cannot query database ,reson: %v", err.Error())
		return nil, err
	}
	// if err != nil {
	// 	log.Printf("could not query : %v ", stmt)
	// 	return nil, err

	// }
	return user, nil
}

const updateBranchAddress = `-- name: UpdateBranchAddress :exec
update manage_branch
set address = ?,
    updated_at = ?,
    updated_by = ?,
    name=?,
	license=?,
	services=?,
	phone=?,
	license_key=?,
	check_in_url=?,
	ticket_page_url=?,
	display_url=?
where code = ?
`

func (q *sqlAdminRepo) UpdateBranch(ctx context.Context, arg models.AddbranchParams) error {
	_, err := q.db.ExecContext(ctx, updateBranchAddress,
		arg.Address,
		arg.UpdatedAt,
		arg.UpdatedBy,
		arg.Name,
		arg.License,
		arg.Services,
		arg.Phone,
		arg.LicenseKey,
		arg.CheckInUrl,
		arg.TicketPageUrl,
		arg.DisplayUrl,
		arg.Code,
	)
	return err
}

// ?INFO counter repo methods
const addCounter = `-- name: AddSection :execresult
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
        ?
    )
`

func (q *sqlAdminRepo) AddCounter(ctx context.Context, arg models.AddCounterParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, addCounter,
		arg.ID,
		arg.Section,
		arg.SubSection,
		arg.UserID,
		arg.CounterName,
		arg.CounterNumber,
		arg.OverrideFifo,
		arg.TransferQ,
		arg.AsssignUser,
		arg.AssignService,
		arg.TransferPriority,
		arg.CancelQ,
		arg.BranchCode,
		arg.BranchName,
		arg.CompanyCode,
		arg.CompanyName,
		arg.CreatedAt,
		arg.CreatedBy,
		arg.UpdatedAt,
		arg.UpdatedBy,
	)
}

const deleteCounter = `-- name: DeleteCounter :execresult
update manage_counters set 
where id = ?
`

func (q *sqlAdminRepo) DeleteCounter(ctx context.Context, id string) (sql.Result, error) {
	return q.db.ExecContext(ctx, deleteCounter, id)
}

const getAllCounters = `-- name: GetAllCounters :many
select id, section, sub_section, user_id, counter_number, counter_name, override_fifo, transfer_q, asssign_user, assign_service, transfer_priority, cancel_q,cpuId ,diskId, branch_code, branch_name, company_name, company_code, created_at, updated_at, created_by, updated_by
from manage_counters
where company_code = ?
`

func (q *sqlAdminRepo) GetAllCounters(ctx context.Context, companyCode string) ([]*models.ManageCounter, error) {
	rows, err := q.db.QueryContext(ctx, getAllCounters, companyCode)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*models.ManageCounter
	for rows.Next() {
		var i models.ManageCounter
		if err := rows.Scan(
			&i.ID,
			&i.Section,
			&i.SubSection,
			&i.UserID,
			&i.CounterNumber,
			&i.CounterName,
			&i.OverrideFifo,
			&i.TransferQ,
			&i.AsssignUser,
			&i.AssignService,
			&i.TransferPriority,
			&i.CancelQ,
			// &i.Activated,
			&i.CpuId,
			&i.DiskId,
			&i.BranchCode,
			&i.BranchName,
			&i.CompanyName,
			&i.CompanyCode,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.CreatedBy,
			&i.UpdatedBy,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllinactiveCounters = `-- name: GetAllCounters :many
select id, section, sub_section, user_id, counter_number, counter_name, override_fifo, transfer_q, asssign_user, assign_service, transfer_priority, cancel_q, activated,cpuId ,diskId, branch_code, branch_name, company_name, company_code, created_at, updated_at, created_by, updated_by
from manage_counters
where company_code = ? and activated =0 or activated is null
`

func (q *sqlAdminRepo) GetAllInactiveCounters(ctx context.Context, companyCode string) ([]*models.ManageCounter, error) {
	rows, err := q.db.QueryContext(ctx, getAllinactiveCounters, companyCode)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*models.ManageCounter
	for rows.Next() {
		var i models.ManageCounter
		if err := rows.Scan(
			&i.ID,
			&i.Section,
			&i.SubSection,
			&i.UserID,
			&i.CounterNumber,
			&i.CounterName,
			&i.OverrideFifo,
			&i.TransferQ,
			&i.AsssignUser,
			&i.AssignService,
			&i.TransferPriority,
			&i.CancelQ,
			&i.Activated,
			&i.CpuId,
			&i.DiskId,
			&i.BranchCode,
			&i.BranchName,
			&i.CompanyName,
			&i.CompanyCode,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.CreatedBy,
			&i.UpdatedBy,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCounter = `-- name: GetSection :one
select id, section, sub_section, user_id, counter_number, counter_name, override_fifo, transfer_q, asssign_user, assign_service, transfer_priority, cancel_q, activated,cpuId ,diskId, branch_code, branch_name, company_name, company_code, created_at, updated_at, created_by, updated_by
from manage_counters
where id = ?
    and branch_code = ?
`

func (q *sqlAdminRepo) GetCounter(ctx context.Context, arg models.GetCounterParams) (*models.ManageCounter, error) {
	row := q.db.QueryRowContext(ctx, getCounter, arg.ID, arg.BranchCode)
	var i models.ManageCounter
	err := row.Scan(
		&i.ID,
		&i.Section,
		&i.SubSection,
		&i.UserID,
		&i.CounterNumber,
		&i.CounterName,
		&i.OverrideFifo,
		&i.TransferQ,
		&i.AsssignUser,
		&i.AssignService,
		&i.TransferPriority,
		&i.CancelQ,
		&i.Activated,
		&i.CpuId,
		&i.DiskId,
		&i.BranchCode,
		&i.BranchName,
		&i.CompanyName,
		&i.CompanyCode,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.CreatedBy,
		&i.UpdatedBy,
	)
	return &i, err
}

const updateCounter = `-- name: UpdateCounter :exec
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
    cancel_q = ?,
    updated_at = ?,
    updated_by = ?
where id = ?
`

func (q *sqlAdminRepo) UpdateCounter(ctx context.Context, arg models.UpdateCounterParams) error {
	_, err := q.db.ExecContext(ctx, updateCounter,
		arg.Section,
		arg.SubSection,
		arg.UserID,
		arg.CounterNumber,
		arg.CounterName,
		arg.OverrideFifo,
		arg.TransferQ,
		arg.AsssignUser,
		arg.AssignService,
		arg.TransferPriority,
		arg.CancelQ,
		arg.UpdatedAt,
		arg.UpdatedBy,
		arg.ID,
	)
	return err
}

const updateCounterSettings = `-- name: UpdateCounterSettings :exec
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
where id = ?
`

func (q *sqlAdminRepo) UpdateCounterSettings(ctx context.Context, arg models.UpdateCounterSettingsParams) error {
	_, err := q.db.ExecContext(ctx, updateCounterSettings,
		arg.OverrideFifo,
		arg.TransferQ,
		arg.AsssignUser,
		arg.AssignService,
		arg.TransferPriority,
		arg.TransferFinish,
		arg.CancelQ,
		arg.MergeSection,
		arg.EndQ,
		arg.UpdatedAt,
		arg.UpdatedBy,
		arg.ID,
	)
	return err
}

const addSection = `-- name: AddSection :execresult
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
    )
`

func (q *sqlAdminRepo) AddSection(ctx context.Context, arg models.AddSectionParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, addSection,
		arg.ID,
		arg.BenchWait,
		arg.BenchProcess,
		arg.Description,
		arg.BranchCode,
		arg.BrachName,
		arg.CompanyCode,
		arg.CompanyName,
		arg.CreatedAt,
		arg.CreatedBy,
		arg.UpdatedAt,
		arg.UpdatedBy,
	)
}

const deleteSection = `-- name: DeleteSection :execresult
delete from manage_section
where id = ?
`

func (q *sqlAdminRepo) DeleteSection(ctx context.Context, id string) (sql.Result, error) {
	return q.db.ExecContext(ctx, deleteSection, id)
}

const getAllSections = `-- name: GetAllSections :many
select id, name, bench_wait, bench_process, description, created_at, updated_at, created_by, updated_by, company_code, branch_code, company_name, brach_name
from manage_section
where company_code = ?
`

func (q *sqlAdminRepo) GetAllSections(ctx context.Context, companyCode string) ([]*models.ManageSection, error) {
	rows, err := q.db.QueryContext(ctx, getAllSections, companyCode)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*models.ManageSection
	for rows.Next() {
		var i models.ManageSection
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.BenchWait,
			&i.BenchProcess,
			&i.Description,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.CreatedBy,
			&i.UpdatedBy,
			&i.CompanyCode,
			&i.BranchCode,
			&i.CompanyName,
			&i.BrachName,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getSection = `-- name: GetSection :one
select id, name, bench_wait, bench_process, description, created_at, updated_at, created_by, updated_by, company_code, branch_code, company_name, brach_name
from manage_section
where id = ?
    and company_code = ?
`

func (q *sqlAdminRepo) GetSection(ctx context.Context, arg models.GetSectionParams) (*models.ManageSection, error) {
	row := q.db.QueryRowContext(ctx, getSection, arg.ID, arg.CompanyCode)
	var i models.ManageSection
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.BenchWait,
		&i.BenchProcess,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.CreatedBy,
		&i.UpdatedBy,
		&i.CompanyCode,
		&i.BranchCode,
		&i.CompanyName,
		&i.BrachName,
	)
	return &i, err
}

const updateSection = `-- name: UpdateSection :exec
update manage_section
set name = ?,
    bench_wait=?,
    bench_process=?,
    description=?,
    updated_at = ?,
    updated_by = ?
where id = ?
`

func (q *sqlAdminRepo) UpdateSection(ctx context.Context, arg models.UpdateSectionParams) error {
	_, err := q.db.ExecContext(ctx, updateSection,
		arg.Name,
		arg.BenchWait,
		arg.BenchProcess,
		arg.Description,
		arg.UpdatedAt,
		arg.UpdatedBy,
		arg.ID,
	)
	return err
}

const addService = `-- name: Addbranch :execresult
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
    )
`

func (q *sqlAdminRepo) AddService(ctx context.Context, arg models.GetServiceResult) (sql.Result, error) {
	return q.db.ExecContext(ctx, addService,
		arg.ID,
		arg.Name,
		arg.Code,
		arg.Prefix,
		arg.NumberStarts,
		arg.NumberEnds,
		arg.Hide,
		arg.ShowDisplay,
		arg.Description,
		arg.StartTime,
		arg.EndTime,
		arg.DefaultTime,
		arg.Workflow,
		arg.BranchCode,
		arg.BranchName,
		arg.CompanyCode,
		arg.CompanyName,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.UpdatedBy,
	)
}

const deleteService = `-- name: DeleteBranch :execresult
delete from manage_service
where id = ?
`

func (q *sqlAdminRepo) DeleteService(ctx context.Context, id string) (sql.Result, error) {
	return q.db.ExecContext(ctx, deleteService, id)
}

const getAllServices = `-- name: GetAllServices :many
select id, name, code, prefix, number_starts, number_ends, hide, show_display, description, start_time, end_time, default_time, workflow, branch_code, branch_name, company_code, company_name, updated_at, created_at, updated_by
from manage_service
where company_code = ?
`

func (q *sqlAdminRepo) GetAllServices(ctx context.Context, companyCode string) ([]*models.ManageService, error) {
	rows, err := q.db.QueryContext(ctx, getAllServices, companyCode)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*models.ManageService
	for rows.Next() {
		var i models.ManageService
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Code,
			&i.Prefix,
			&i.NumberStarts,
			&i.NumberEnds,
			&i.Hide,
			&i.ShowDisplay,
			&i.Description,
			&i.StartTime,
			&i.EndTime,
			&i.DefaultTime,
			&i.Workflow,
			&i.BranchCode,
			&i.BranchName,
			&i.CompanyCode,
			&i.CompanyName,
			&i.UpdatedAt,
			&i.CreatedAt,
			&i.UpdatedBy,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getService = `-- name: GetBranch :one
select id, name, code, prefix, number_starts, number_ends, hide, show_display, description, start_time, end_time, default_time, workflow, branch_code, branch_name, company_code, company_name, updated_at, created_at, updated_by
from manage_service
where id = ?
    and company_code = ?
`

type GetServiceParams struct {
	ID          string
	CompanyCode string
}

func (q *sqlAdminRepo) GetService(ctx context.Context, arg models.GetServiceParams) (*models.ManageService, error) {
	row := q.db.QueryRowContext(ctx, getService, arg.ID, arg.CompanyCode)
	var i models.ManageService
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Code,
		&i.Prefix,
		&i.NumberStarts,
		&i.NumberEnds,
		&i.Hide,
		&i.ShowDisplay,
		&i.Description,
		&i.StartTime,
		&i.EndTime,
		&i.DefaultTime,
		&i.Workflow,
		&i.BranchCode,
		&i.BranchName,
		&i.CompanyCode,
		&i.CompanyName,
		&i.UpdatedAt,
		&i.CreatedAt,
		&i.UpdatedBy,
	)
	return &i, err
}

// const updateService = `-- name: UpdateService :exec
// update manage_service
// set name = ?,
//
//	code = ?,
//	prefix = ?,
//	number_starts = ?,
//	number_ends = ?,
//	hide = ?,
//	show_display = ?,
//	start_time = ?,
//	end_time = ?,
//	default_time = ?,
//	workflow = ?,
//	number_ends = ?,
//	updated_at = ?,
//	updated_by = ?
//
// where id = ?
// `
const updateService = `-- name: UpdateService :exec
update manage_service
set name = ?,
    code = ?,
    prefix = ?,
    number_starts = ?,
    number_ends = ?,
    hide = ?,
    show_display = ?,
    default_time = ?,
    workflow = ?,
    updated_at = ?,
    updated_by = ?
where id = ?
`

func (q *sqlAdminRepo) UpdateService(ctx context.Context, arg models.UpdateServiceParams) error {
	_, err := q.db.ExecContext(ctx, updateService,
		arg.Name,
		arg.Code,
		arg.Prefix,
		arg.NumberStarts,
		arg.NumberEnds,
		arg.Hide,
		arg.ShowDisplay,
		// arg.StartTime,
		// arg.EndTime,
		arg.DefaultTime,
		arg.Workflow,
		// arg.NumberEnds_2,
		arg.UpdatedAt,
		arg.UpdatedBy,
		arg.ID,
	)
	return err
}

const addUser = `-- name: AddUser :execresult
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
    )
`

func (q *sqlAdminRepo) AddUser(ctx context.Context, arg models.AddUserParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, addUser,
		arg.ID,
		arg.Email,
		arg.Firstname,
		arg.Lastname,
		arg.UserType,
		arg.Suspended,
		arg.Title,
		arg.Password,
		arg.ImageUrl,
		arg.BranchCode,
		arg.BranchName,
		arg.CompanyCode,
		arg.CompanyName,
		arg.CreatedAt,
		arg.CreatedBy,
		arg.UpdatedAt,
		arg.UpdatedBy,
	)
}

const deleteUser = `-- name: DeleteUser :execresult
update manage_user set suspended = 1
where id = ?
`

func (q *sqlAdminRepo) DeleteUser(ctx context.Context, id string) (sql.Result, error) {
	return q.db.ExecContext(ctx, deleteUser, id)
}

const getAllUsers = `-- name: GetAllUsers :many
select id, email, firstname, lastname, created_at, user_type, suspended, branch_name, branch_code, company_name, company_code, title, password, image_url, updated_at, created_by, updated_by
from manage_user
where company_code = ?
`

func (q *sqlAdminRepo) GetAllUsers(ctx context.Context, arg models.GetAllUsersParams) ([]*models.ManageUser, error) {
	rows, err := q.db.QueryContext(ctx, getAllUsers, arg.CompanyCode)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*models.ManageUser
	for rows.Next() {
		var i models.ManageUser
		if err := rows.Scan(
			&i.ID,
			&i.Email,
			&i.Firstname,
			&i.Lastname,
			&i.CreatedAt,
			&i.UserType,
			&i.Suspended,
			&i.BranchName,
			&i.BranchCode,
			&i.CompanyName,
			&i.CompanyCode,
			&i.Title,
			&i.Password,
			&i.ImageUrl,
			&i.UpdatedAt,
			&i.CreatedBy,
			&i.UpdatedBy,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUser = `-- name: GetBranch :one
select id, email, firstname, lastname, created_at, user_type, suspended, branch_name, branch_code, company_name, company_code, title, password, image_url, updated_at, created_by, updated_by
from manage_user
where id = ? and company_code = ?
`

func (q *sqlAdminRepo) GetUser(ctx context.Context, arg models.GetUserParams) (*models.ManageUser, error) {
	row := q.db.QueryRowContext(ctx, getUser, arg.ID, arg.CompanyCode)
	var i models.ManageUser
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Firstname,
		&i.Lastname,
		&i.CreatedAt,
		&i.UserType,
		&i.Suspended,
		&i.BranchName,
		&i.BranchCode,
		&i.CompanyName,
		&i.CompanyCode,
		&i.Title,
		&i.Password,
		&i.ImageUrl,
		&i.UpdatedAt,
		&i.CreatedBy,
		&i.UpdatedBy,
	)
	return &i, err
}

const updateUserEmail = `-- name: UpdateUserEmail :exec
update manage_user
set email = ?,
    updated_at = ?,
    updated_by = ?
where id = ?
`

func (q *sqlAdminRepo) UpdateUserEmail(ctx context.Context, arg models.UpdateUserEmailParams) error {
	_, err := q.db.ExecContext(ctx, updateUserEmail,
		arg.Email,
		arg.UpdatedAt,
		arg.UpdatedBy,
		arg.ID,
	)
	return err
}

const updateUserPassword = `-- name: UpdateUserPassword :exec
update manage_user
set password = ?,
    updated_at = ?,
    updated_by = ?
where id = ?
`

func (q *sqlAdminRepo) UpdateUserPassword(ctx context.Context, arg models.UpdateUserPasswordParams) error {
	_, err := q.db.ExecContext(ctx, updateUserPassword,
		arg.Password,
		arg.UpdatedAt,
		arg.UpdatedBy,
		arg.ID,
	)
	return err
}

const updateUser = `-- name: UpdateUser :exec
update manage_user
set email = ?,
    firstname = ?,
    lastname = ?,
	password=?,
    suspended = ?,
    image_url = ?,
    title = ?,
    user_type = ?,
    updated_at = ?,
    updated_by = ?
where id = ?
`

func (q *sqlAdminRepo) UpdateUser(ctx context.Context, arg models.UpdateUserParams) error {
	_, err := q.db.ExecContext(ctx, updateUser,
		arg.Email,
		arg.Firstname,
		arg.Lastname,
		arg.Password,
		arg.Suspended,
		arg.ImageUrl,
		arg.Title,
		arg.UserType,
		arg.UpdatedAt,
		arg.UpdatedBy,
		arg.ID,
	)
	return err
}
func (q *sqlAdminRepo) ActivateSystemDAO(ip string, active bool) error {
	_, err := q.db.ExecContext(context.TODO(), UpdateSystemActive(ip, active))
	return err
}

func UpdateSystemActive(ip string, active bool) string {
	return fmt.Sprintf(`
		UPDATE systems SET 
			is_active = %t
		WHERE client_ip = '%s'`,
		active, ip,
	)
}
