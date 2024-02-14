package models

import (
	"time"

	"github.com/gocraft/dbr/v2"
)

var COMAPNY_CODE string

type ManageBranch struct {
	ID             dbr.NullString
	Code           dbr.NullString
	Name           dbr.NullString
	License        dbr.NullString
	Services       dbr.NullString
	Address        dbr.NullString
	Phone          dbr.NullString
	LicenseKey     dbr.NullString
	CheckInUrl     dbr.NullString
	TicketPageUrl  dbr.NullString
	DisplayUrl     dbr.NullString
	CompanyCode    dbr.NullString
	CreatedAt      time.Time
	CreatedBy      dbr.NullString
	UpdatedAt      time.Time
	UpdatedBy      dbr.NullString
	PrinterDetails dbr.NullString
}

type ManageService struct {
	ID           dbr.NullString
	Name         dbr.NullString
	Code         dbr.NullString
	Prefix       dbr.NullString
	NumberStarts int
	NumberEnds   int
	Hide         bool
	ShowDisplay  bool
	Description  dbr.NullString
	StartTime    dbr.NullString
	EndTime      dbr.NullString
	DefaultTime  int
	Workflow     dbr.NullString
	BranchCode   dbr.NullString
	BranchName   dbr.NullString
	CompanyCode  dbr.NullString
	CompanyName  dbr.NullString
	UpdatedAt    time.Time
	CreatedAt    time.Time
	UpdatedBy    dbr.NullString
}

type ManageUser struct {
	ID          dbr.NullString
	Email       dbr.NullString
	Firstname   dbr.NullString
	Lastname    dbr.NullString
	CreatedAt   time.Time
	UserType    dbr.NullString
	Suspended   bool
	BranchName  dbr.NullString
	BranchCode  dbr.NullString
	CompanyName dbr.NullString
	CompanyCode dbr.NullString
	Title       dbr.NullString
	Password    dbr.NullString
	ImageUrl    dbr.NullString
	UpdatedAt   time.Time
	CreatedBy   dbr.NullString
	UpdatedBy   dbr.NullString
}

type ManageCounter struct {
	ID               dbr.NullString
	Section          dbr.NullString
	SubSection       dbr.NullString
	UserID           dbr.NullString
	CounterNumber    dbr.NullString
	CounterName      dbr.NullString
	OverrideFifo     bool
	TransferQ        bool
	AsssignUser      bool
	AssignService    bool
	TransferPriority bool
	CancelQ          bool
	Activated        dbr.NullBool
	CpuId            dbr.NullString
	DiskId           dbr.NullString
	BranchCode       dbr.NullString
	BranchName       dbr.NullString
	CompanyName      dbr.NullString
	CompanyCode      dbr.NullString
	CreatedAt        time.Time
	UpdatedAt        time.Time
	CreatedBy        dbr.NullString
	UpdatedBy        dbr.NullString
}
type ManageSection struct {
	ID           dbr.NullString
	Name         dbr.NullString
	BenchWait    int
	BenchProcess int
	Description  dbr.NullString
	CreatedAt    time.Time
	UpdatedAt    time.Time
	CreatedBy    dbr.NullString
	UpdatedBy    dbr.NullString
	CompanyCode  dbr.NullString
	BranchCode   dbr.NullString
	CompanyName  dbr.NullString
	BrachName    dbr.NullString
}

type Ticket struct {
	ID               dbr.NullString
	Service          dbr.NullString
	TicketStatus     string
	CounterID        dbr.NullString
	TransferedTo     dbr.NullString
	TransferedBy     dbr.NullString
	CustomerID       dbr.NullString
	TicketNumber     dbr.NullString
	CreatedAt        time.Time
	UpdatedAt        time.Time
	UpdatedBy        dbr.NullString
	StartedServingAt time.Time
	EndServingAt     time.Time
	TicketName       dbr.NullString
	CompanyName      dbr.NullString
	CompanyCode      dbr.NullString
	BranchCode       dbr.NullString
	BranchName       dbr.NullString
	ServingTime      dbr.NullString
	WaitingTime      dbr.NullString
}

type TicketNumber struct {
	Number string `json:"number,omitempty"`
}
