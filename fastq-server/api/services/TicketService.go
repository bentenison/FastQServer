package services

import (
	"context"
	"database/sql"
	"time"

	"github.com/bentenison/fastq-server/models"
	"github.com/google/uuid"
)

type ticketService struct {
	TicketRepo models.TicketRepository
	db         *sql.DB

	// ImageRepository model.ImageRepository
}

/*
* USConfig will hold repositories that will eventually be injected into this service layer
* now it means the USConfig is used to hide the implementation of service layer
 */
type TConfig struct {
	// AdminRepository models.AdminRepository
	// ImageRepository model.ImageRepository
	TicketRepository models.TicketRepository
	DB               *sql.DB
}

// NewUserService is a factory function for
// initializing a UserService with its repository layer dependencies
func NewTicketService(c *TConfig) models.TicketService {
	return &ticketService{
		TicketRepo: c.TicketRepository,
		db:         c.DB,
		// ImageRepository: c.ImageRepository,
	}
}

func (ts *ticketService) AddTicketService(ctx context.Context, arg models.AddTicketParams) (sql.Result, error) {
	arg.ID = uuid.NewString()
	t := time.Now()
	tm := t.Format("2006-01-02 15:04:05")
	arg.CreatedAt = tm
	arg.UpdatedAt = tm
	return ts.TicketRepo.AddTicket(ctx, arg)
}

func (ts *ticketService) DeleteTicketService(ctx context.Context, id string) (sql.Result, error) {
	return ts.TicketRepo.DeleteTicket(ctx, id)
}
func (ts *ticketService) GetLastTicketNumberService(ctx context.Context) (models.TicketNumber, error) {
	return ts.TicketRepo.GetLastTicketNumber(ctx)
}
func (ts *ticketService) GetAllTicketsForDayService(ctx context.Context, arg models.GetAllTicketsForDayParams) ([]*models.Ticket, error) {
	return ts.TicketRepo.GetAllTicketsForDay(ctx, arg)
}
func (ts *ticketService) GetTicketService(ctx context.Context, arg models.GetTicketParams) (*models.Ticket, error) {
	return ts.TicketRepo.GetTicket(ctx, arg)
}
func (ts *ticketService) GetTicketToProcessService(ctx context.Context, arg models.GetTicketParams) (*models.Ticket, error) {
	return ts.TicketRepo.GetTicketToProcess(ctx, arg)
}
func (ts *ticketService) GetTicketsByBranchService(ctx context.Context, arg models.GetTicketsByBranchParams) ([]*models.Ticket, error) {
	return ts.TicketRepo.GetTicketsByBranch(ctx, arg)
}
func (ts *ticketService) GetTicketsByStatusService(ctx context.Context, arg models.GetTicketsByStatusParams) ([]*models.Ticket, error) {
	return ts.TicketRepo.GetTicketsByStatus(ctx, arg)
}
func (ts *ticketService) GetTicketsByTransferService(ctx context.Context, arg models.GetTicketsByTransferParams) ([]*models.Ticket, error) {
	return ts.TicketRepo.GetTicketsByTransfer(ctx, arg)
}
func (ts *ticketService) GetTicketsByUserService(ctx context.Context, arg models.GetTicketsByUserParams) ([]*models.Ticket, error) {
	return ts.TicketRepo.GetTicketsByUser(ctx, arg)
}
func (ts *ticketService) UpdateTicketStartTimeService(ctx context.Context, arg models.UpdateTicketStartTimeParams) error {
	t := time.Now()
	tm := t.Format("2006-01-02 15:04:05")
	arg.StartedServingAt = tm
	arg.UpdatedAt = tm
	return ts.TicketRepo.UpdateTicketStartTime(ctx, arg)
}
func (ts *ticketService) UpdateTicketEndTimeService(ctx context.Context, arg models.UpdateTicketEndTimeParams) error {
	t := time.Now()
	tm := t.Format("2006-01-02 15:04:05")
	arg.EndServingAt = tm
	arg.UpdatedAt = tm
	return ts.TicketRepo.UpdateTicketEndTime(ctx, arg)
}
func (ts *ticketService) UpdateTicketStatusService(ctx context.Context, arg models.UpdateTicketStatusParams) error {
	t := time.Now()
	tm := t.Format("2006-01-02 15:04:05")
	arg.UpdatedAt = tm
	return ts.TicketRepo.UpdateTicketStatus(ctx, arg)
}
func (ts *ticketService) UpdateTicketTransferedToService(ctx context.Context, arg models.UpdateTicketTransferedToParams) error {
	t := time.Now()
	tm := t.Format("2006-01-02 15:04:05")
	arg.UpdatedAt = tm
	return ts.TicketRepo.UpdateTicketTransferedTo(ctx, arg)
}
func (ts *ticketService) UpdateTicketUserService(ctx context.Context, arg models.UpdateTicketUserParams) error {
	t := time.Now()
	tm := t.Format("2006-01-02 15:04:05")
	arg.UpdatedAt = tm
	return ts.TicketRepo.UpdateTicketUser(ctx, arg)
}
func (ts *ticketService) GetWaitingTicketsService(ctx context.Context) ([]*models.Ticket, error) {
	return GetWaitingTickets(ctx, ts.db)
}
func (ts *ticketService) GetNoShowTicketsService(ctx context.Context, id string) ([]*models.Ticket, error) {
	return GetNoShowTickets(ctx, ts.db, id)
}
func (ts *ticketService) GetFinishedTicketsService(ctx context.Context, id string) ([]*models.Ticket, error) {
	return GetFinishedTickets(ctx, ts.db, id)
}
func (ts *ticketService) GetLastCalledService(ctx context.Context) (*models.Ticket, error) {
	return GetLastCalled(ctx, ts.db)
}
func (ts *ticketService) GetLastStartedTicket(ctx context.Context, id string) (*models.Ticket, error) {
	return GetLastStarted(ctx, ts.db, id)
}
func (ts *ticketService) GetDoneTicketsByUserService(ctx context.Context, id string) ([]*models.Ticket, error) {
	return GetDoneTicketsByUser(ctx, ts.db, id)
}
func (ts *ticketService) GetSystemService(ctx context.Context) (*models.ServerDetails, error) {
	return GetSystems(ctx, ts.db)
}
func (ts *ticketService) UpdateSystemIp(ctx context.Context, client models.UpdateSystemParams) (sql.Result, error) {
	return UpdateSystemsIP(ctx, ts.db, client)
}
func (ts *ticketService) GetAllSystemService(ctx context.Context) ([]*models.UpdateSystemParams, error) {
	return GetAllSystems(ctx, ts.db)
}
func GetDoneTicketsByUser(ctx context.Context, db *sql.DB, id string) ([]*models.Ticket, error) {
	return nil, nil
}
func GetWaitingTickets(ctx context.Context, db *sql.DB) ([]*models.Ticket, error) {
	rows, err := db.QueryContext(ctx, `select * from ticket where date(created_at)=date(now()) and ticket_status = 'CREATED' order by created_at asc limit 5`)
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
			&i.ServingTime,
			&i.WaitingTime,
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
func GetNoShowTickets(ctx context.Context, db *sql.DB, id string) ([]*models.Ticket, error) {
	rows, err := db.QueryContext(ctx, `select * from ticket where date(created_at)=date(now()) and ticket_status = 'NO_SHOW' and counter_id =? order by created_at asc limit 5`, id)
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
			&i.ServingTime,
			&i.WaitingTime,
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
func GetFinishedTickets(ctx context.Context, db *sql.DB, id string) ([]*models.Ticket, error) {
	rows, err := db.QueryContext(ctx, `select * from ticket where date(created_at)=date(now()) and ticket_status = 'FINISHED' and counter_id =? order by created_at asc limit 5`, id)
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
			&i.ServingTime,
			&i.WaitingTime,
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
func GetLastCalled(ctx context.Context, db *sql.DB) (*models.Ticket, error) {
	rows := db.QueryRowContext(ctx, `select * from ticket where date(created_at)=date(now()) and ticket_status = 'FINISHED' order by created_at desc limit 1`)
	var i models.Ticket
	err := rows.Scan(
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
		&i.ServingTime,
		&i.WaitingTime,
	)
	if err != nil {
		return nil, err
	}
	return &i, nil
}
func GetLastStarted(ctx context.Context, db *sql.DB, id string) (*models.Ticket, error) {
	rows := db.QueryRowContext(ctx, `select * from ticket where date(created_at)=date(now()) and ticket_status = 'STARTED' and updated_by =? order by created_at asc limit 1`, id)
	var i models.Ticket
	err := rows.Scan(
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
		&i.ServingTime,
		&i.WaitingTime,
	)
	if err != nil {
		return nil, err
	}
	return &i, nil
}
func GetSystems(ctx context.Context, db *sql.DB) (*models.ServerDetails, error) {
	rows := db.QueryRowContext(ctx, `select * from server_details where type = 'DS'`)
	var i models.ServerDetails
	err := rows.Scan(
		&i.ServerIP,
		&i.ServerCPU,
		&i.ServerDiskId,
		&i.Id,
		&i.Type,
	)
	if err != nil {
		return nil, err
	}
	return &i, nil
}
func UpdateSystemsIP(ctx context.Context, db *sql.DB, client models.UpdateSystemParams) (sql.Result, error) {
	rows, err := db.ExecContext(ctx, `update systems set client_ip = ? where client_cpu = ? and client_disk_id =?`, client.ClientIP, client.ClientCPU, client.ClientDiskId)
	if err != nil {
		return nil, err
	}
	return rows, nil
}
func GetAllSystems(ctx context.Context, db *sql.DB) ([]*models.UpdateSystemParams, error) {
	rows, err := db.QueryContext(ctx, `select * from systems`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*models.UpdateSystemParams
	for rows.Next() {
		var i models.UpdateSystemParams
		if err := rows.Scan(
			&i.ServerCPU,
			&i.ServerDiskId,
			&i.ServerIP,
			&i.ClientIP,
			&i.ClientCPU,
			&i.ClientDiskId,
			&i.IsActive,
			&i.IsDeleted,
			&i.CounterName,
			&i.CounterId,
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
