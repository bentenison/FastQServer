package repository

import (
	"context"
	"database/sql"

	"github.com/bentenison/fastq-server/models"
)

type Customer struct {
	ID            int
	Name          string
	NameAr        string
	ContactNumber string
}
type sqlTicketRepo struct {
	db *sql.DB
}

func NewTicketRepo(db *sql.DB) *sqlTicketRepo {
	return &sqlTicketRepo{
		db: db,
	}
}

const addTicket = `-- name: AddTicket :execresult
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
    )
`

func (q *sqlTicketRepo) AddTicket(ctx context.Context, arg models.AddTicketParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, addTicket,
		arg.ID,
		arg.Service,
		arg.CounterID,
		arg.TicketStatus,
		arg.TransferedTo,
		arg.TransferedBy,
		arg.TicketNumber,
		arg.StartedServingAt,
		arg.EndServingAt,
		arg.TicketName,
		arg.BranchCode,
		arg.BranchName,
		arg.CompanyCode,
		arg.CompanyName,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.UpdatedBy,
	)
}

const deleteTicket = `-- name: DeleteSection :execresult
delete from ticket
where id = ?
`

func (q *sqlTicketRepo) DeleteTicket(ctx context.Context, id string) (sql.Result, error) {
	return q.db.ExecContext(ctx, deleteTicket, id)
}

const getAllTicketsForDay = `-- name: GetAllTicketsForDay :many
select id, service, ticket_status, counter_id, transfered_to, transfered_by, customer_id, ticket_number, created_at, updated_at, updated_by, started_serving_at, end_serving_at, ticket_name, company_name, company_code, branch_code, branch_name
from ticket
where company_code = ?
    and date(created_at) = ?
`

func (q *sqlTicketRepo) GetAllTicketsForDay(ctx context.Context, arg models.GetAllTicketsForDayParams) ([]*models.Ticket, error) {
	rows, err := q.db.QueryContext(ctx, getAllTicketsForDay, arg.CompanyCode, arg.BranchCode, arg.CreatedAt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*models.Ticket
	for rows.Next() {
		var i models.Ticket
		if err := rows.Scan(
			&i.ID,
			&i.Service,
			&i.TicketStatus,
			&i.CounterID,
			&i.TransferedTo,
			&i.TransferedBy,
			&i.CustomerID,
			&i.TicketNumber,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.UpdatedBy,
			&i.StartedServingAt,
			&i.EndServingAt,
			&i.TicketName,
			&i.CompanyName,
			&i.CompanyCode,
			&i.BranchCode,
			&i.BranchName,
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

const getTicket = `-- name: GetTicket :one
select id, service, ticket_status, counter_id, transfered_to, transfered_by, customer_id, ticket_number, created_at, updated_at, updated_by, started_serving_at, end_serving_at, ticket_name, company_name, company_code, branch_code, branch_name
from ticket
where id = ?
    and company_code = ?
`

func (q *sqlTicketRepo) GetTicket(ctx context.Context, arg models.GetTicketParams) (*models.Ticket, error) {
	row := q.db.QueryRowContext(ctx, getTicket, arg.ID, arg.CompanyCode, arg.BranchCode)
	var i models.Ticket
	err := row.Scan(
		&i.ID,
		&i.Service,
		&i.TicketStatus,
		&i.CounterID,
		&i.TransferedTo,
		&i.TransferedBy,
		&i.CustomerID,
		&i.TicketNumber,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UpdatedBy,
		&i.StartedServingAt,
		&i.EndServingAt,
		&i.TicketName,
		&i.CompanyName,
		&i.CompanyCode,
		&i.BranchCode,
		&i.BranchName,
	)
	return &i, err
}

const getTicketToProcess = `-- name: GetTicket :one
select id, service, ticket_status, counter_id, transfered_to, transfered_by, customer_id, ticket_number, created_at, updated_at, updated_by, started_serving_at, end_serving_at, ticket_name, company_name, company_code, branch_code, branch_name
from ticket
where company_code = ?
	and date(created_at)=date(now()) and ticket_status = 'CREATED' order by created_at asc limit 1;
`

func (q *sqlTicketRepo) GetTicketToProcess(ctx context.Context, arg models.GetTicketParams) (*models.Ticket, error) {
	row := q.db.QueryRowContext(ctx, getTicketToProcess, arg.CompanyCode)
	var i models.Ticket
	err := row.Scan(
		&i.ID,
		&i.Service,
		&i.TicketStatus,
		&i.CounterID,
		&i.TransferedTo,
		&i.TransferedBy,
		&i.CustomerID,
		&i.TicketNumber,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UpdatedBy,
		&i.StartedServingAt,
		&i.EndServingAt,
		&i.TicketName,
		&i.CompanyName,
		&i.CompanyCode,
		&i.BranchCode,
		&i.BranchName,
	)
	return &i, err
}

func (q *sqlTicketRepo) GetLastTicketNumber(ctx context.Context, service string) (models.TicketNumber, error) {
	row := q.db.QueryRowContext(ctx, `select ticket_number as number from ticket where date(created_at)=date(now()) and service = ? and ticket_status = 0 order by created_at desc limit 1;`, service)
	var i models.TicketNumber
	err := row.Scan(
		&i.Number,
	)
	switch {
	case err == sql.ErrNoRows:
		return models.TicketNumber{
			Number: "0",
		}, nil
	case err != nil:
		return models.TicketNumber{}, err
	}
	return i, err
}

const getTicketsByBranch = `-- name: GetTicketsByBranch :many
select id, service, ticket_status, counter_id, transfered_to, transfered_by, customer_id, ticket_number, created_at, updated_at, updated_by, started_serving_at, end_serving_at, ticket_name, company_name, company_code, branch_code, branch_name
from ticket
where company_code = ?
`

func (q *sqlTicketRepo) GetTicketsByBranch(ctx context.Context, arg models.GetTicketsByBranchParams) ([]*models.Ticket, error) {
	rows, err := q.db.QueryContext(ctx, getTicketsByBranch, arg.CompanyCode)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*models.Ticket
	for rows.Next() {
		var i models.Ticket
		if err := rows.Scan(
			&i.ID,
			&i.Service,
			&i.TicketStatus,
			&i.CounterID,
			&i.TransferedTo,
			&i.TransferedBy,
			&i.CustomerID,
			&i.TicketNumber,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.UpdatedBy,
			&i.StartedServingAt,
			&i.EndServingAt,
			&i.TicketName,
			&i.CompanyName,
			&i.CompanyCode,
			&i.BranchCode,
			&i.BranchName,
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

const getTicketsByStatus = `-- name: GetTicketsByStatus :many
select id, service, ticket_status, counter_id, transfered_to, transfered_by, customer_id, ticket_number, created_at, updated_at, updated_by, started_serving_at, end_serving_at, ticket_name, company_name, company_code, branch_code, branch_name
from ticket
where company_code = ?
    and ticket_status = ?
`

func (q *sqlTicketRepo) GetTicketsByStatus(ctx context.Context, arg models.GetTicketsByStatusParams) ([]*models.Ticket, error) {
	rows, err := q.db.QueryContext(ctx, getTicketsByStatus, arg.CompanyCode, arg.BranchCode, arg.TicketStatus)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*models.Ticket
	for rows.Next() {
		var i models.Ticket
		if err := rows.Scan(
			&i.ID,
			&i.Service,
			&i.TicketStatus,
			&i.CounterID,
			&i.TransferedTo,
			&i.TransferedBy,
			&i.CustomerID,
			&i.TicketNumber,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.UpdatedBy,
			&i.StartedServingAt,
			&i.EndServingAt,
			&i.TicketName,
			&i.CompanyName,
			&i.CompanyCode,
			&i.BranchCode,
			&i.BranchName,
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

const getTicketsByTransfer = `-- name: GetTicketsByTransfer :many
select id, service, ticket_status, counter_id, transfered_to, transfered_by, customer_id, ticket_number, created_at, updated_at, updated_by, started_serving_at, end_serving_at, ticket_name, company_name, company_code, branch_code, branch_name
from ticket
where company_code = ?
    and transfered_to = ?
`

func (q *sqlTicketRepo) GetTicketsByTransfer(ctx context.Context, arg models.GetTicketsByTransferParams) ([]*models.Ticket, error) {
	rows, err := q.db.QueryContext(ctx, getTicketsByTransfer, arg.CompanyCode, arg.BranchCode, arg.TransferedTo)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*models.Ticket
	for rows.Next() {
		var i models.Ticket
		if err := rows.Scan(
			&i.ID,
			&i.Service,
			&i.TicketStatus,
			&i.CounterID,
			&i.TransferedTo,
			&i.TransferedBy,
			&i.CustomerID,
			&i.TicketNumber,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.UpdatedBy,
			&i.StartedServingAt,
			&i.EndServingAt,
			&i.TicketName,
			&i.CompanyName,
			&i.CompanyCode,
			&i.BranchCode,
			&i.BranchName,
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

const getTicketsByUser = `-- name: GetTicketsByUser :many
select id, service, ticket_status, counter_id, transfered_to, transfered_by, customer_id, ticket_number, created_at, updated_at, updated_by, started_serving_at, end_serving_at, ticket_name, company_name, company_code, branch_code, branch_name
from ticket
where company_code = ?
    and counter_id = ?
`

func (q *sqlTicketRepo) GetTicketsByUser(ctx context.Context, arg models.GetTicketsByUserParams) ([]*models.Ticket, error) {
	rows, err := q.db.QueryContext(ctx, getTicketsByUser, arg.CompanyCode, arg.BranchCode, arg.CounterID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*models.Ticket
	for rows.Next() {
		var i models.Ticket
		if err := rows.Scan(
			&i.ID,
			&i.Service,
			&i.TicketStatus,
			&i.CounterID,
			&i.TransferedTo,
			&i.TransferedBy,
			&i.CustomerID,
			&i.TicketNumber,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.UpdatedBy,
			&i.StartedServingAt,
			&i.EndServingAt,
			&i.TicketName,
			&i.CompanyName,
			&i.CompanyCode,
			&i.BranchCode,
			&i.BranchName,
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

const updateTicketStartTime = `-- name: UpdateTicketStartTime :exec
update ticket
set started_serving_at = ?,
    updated_at = ?,
    updated_by = ?,
	counter_id =?
where id = ?
`

func (q *sqlTicketRepo) UpdateTicketStartTime(ctx context.Context, arg models.UpdateTicketStartTimeParams) error {
	_, err := q.db.ExecContext(ctx, updateTicketStartTime,
		arg.StartedServingAt,
		arg.UpdatedAt,
		arg.UpdatedBy,
		arg.CounterId,
		arg.ID,
	)
	return err
}

const updateTicketEndTime = `-- name: UpdateTicketStartTime :exec
update ticket
set end_serving_at = ?,
    updated_at = ?,
    updated_by = ?
where id = ?
`

func (q *sqlTicketRepo) UpdateTicketEndTime(ctx context.Context, arg models.UpdateTicketEndTimeParams) error {
	_, err := q.db.ExecContext(ctx, updateTicketEndTime,
		arg.EndServingAt,
		arg.UpdatedAt,
		arg.UpdatedBy,
		arg.ID,
	)
	return err
}

const updateTicketStatus = `-- name: UpdateTicketStatus :exec
update ticket
set ticket_status = ?,
    updated_at = ?,
    updated_by = ?
where id = ?
`

func (q *sqlTicketRepo) UpdateTicketStatus(ctx context.Context, arg models.UpdateTicketStatusParams) error {
	_, err := q.db.ExecContext(ctx, updateTicketStatus,
		arg.TicketStatus,
		arg.UpdatedAt,
		arg.UpdatedBy,
		arg.ID,
	)
	return err
}

const updateTicketTransferedTo = `-- name: UpdateTicketTransferedTo :exec
update ticket
set transfered_to = ?,
    updated_at = ?,
    updated_by = ?
where id = ?
`

func (q *sqlTicketRepo) UpdateTicketTransferedTo(ctx context.Context, arg models.UpdateTicketTransferedToParams) error {
	_, err := q.db.ExecContext(ctx, updateTicketTransferedTo,
		arg.TransferedTo,
		arg.UpdatedAt,
		arg.UpdatedBy,
		arg.ID,
	)
	return err
}

const updateTicketUser = `-- name: UpdateTicketUser :exec
update ticket
set counter_id = ?,
    updated_at = ?,
    updated_by = ?
where id = ?
`

func (q *sqlTicketRepo) UpdateTicketUser(ctx context.Context, arg models.UpdateTicketUserParams) error {
	_, err := q.db.ExecContext(ctx, updateTicketUser,
		arg.CounterID,
		arg.UpdatedAt,
		arg.UpdatedBy,
		arg.ID,
	)
	return err
}

func (q *sqlTicketRepo) CreateCustomer(customer *Customer) (int64, error) {
	result, err := q.db.Exec("INSERT INTO customer (name, name_ar, contact_number) VALUES (?, ?, ?)", customer.Name, customer.NameAr, customer.ContactNumber)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (q *sqlTicketRepo) GetCustomer(id int) (*Customer, error) {
	customer := &Customer{}
	err := q.db.QueryRow("SELECT id, name, name_ar, contact_number FROM customer WHERE id = ?", id).
		Scan(&customer.ID, &customer.Name, &customer.NameAr, &customer.ContactNumber)
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (q *sqlTicketRepo) UpdateCustomer(customer *Customer) error {
	_, err := q.db.Exec("UPDATE customer SET name = ?, name_ar = ?, contact_number = ? WHERE id = ?",
		customer.Name, customer.NameAr, customer.ContactNumber, customer.ID)
	return err
}

func (q *sqlTicketRepo) DeleteCustomer(id int) error {
	_, err := q.db.Exec("DELETE FROM customer WHERE id = ?", id)
	return err
}
