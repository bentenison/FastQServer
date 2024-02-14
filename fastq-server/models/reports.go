// models.go

package models

import "time"

// TicketsByService represents the result of Query 1.
type TicketsByService struct {
	Service      string
	TotalTickets int
}

// TicketsByStatus represents the result of Query 2.
type TicketsByStatus struct {
	TicketStatus int
	TotalTickets int
}

// AvgServiceTimeByService represents the result of Query 3.
type AvgServiceTimeByService struct {
	Service            string
	AvgServiceTimeSecs float64
}

// TicketsTransferredByCounter represents the result of Query 4.
type TicketsTransferredByCounter struct {
	TransferedBy       string
	TransferredTickets int
}

// TicketsByCompanyBranch represents the result of Query 5.
type TicketsByCompanyBranch struct {
	CompanyName  string
	BranchName   string
	TotalTickets int
}

// AvgWaitingTimeByService represents the result of Query 6.
type AvgWaitingTimeByService struct {
	Service            string
	AvgWaitingTimeSecs float64
}

// TicketsServedByCounter represents the result of Query 7.
type TicketsServedByCounter struct {
	CounterID     string
	ServedTickets int
}

// PercentageTicketsFinishedByService represents the result of Query 8.
type PercentageTicketsFinishedByService struct {
	Service            string
	TotalTickets       int
	FinishedTickets    int
	PercentageFinished float64
}

// TicketsBetweenDatesByStatusAndUser represents the result of Query 9.
type TicketsBetweenDatesByStatusAndUser struct {
	TicketStatus int
	User         string
	TicketCount  int
}

// TicketsWithUserInfo represents the result of Query 10.
type TicketsWithUserInfo struct {
	ID           string
	Service      string
	TicketStatus int
	UserName     string
}

// TicketsWithCounterInfo represents the result of Query 11.
type TicketsWithCounterInfo struct {
	ID           string
	Service      string
	TicketStatus int
	CounterName  string
}

// TicketsWithUserCounterInfo represents the result of Query 12.
type TicketsWithUserCounterInfo struct {
	ID           string
	Service      string
	TicketStatus int
	UserName     string
	CounterName  string
}

// TicketsServedByCounterWithUserInfo represents the result of Query 13.
type TicketsServedByCounterWithUserInfo struct {
	ID           string
	Service      string
	TicketStatus int
	CounterName  string
	UserName     string
}

// TicketsTransferredWithUserInfo represents the result of Query 14.
type TicketsTransferredWithUserInfo struct {
	ID               string
	Service          string
	TicketStatus     int
	TransferredTo    string
	TransferedByUser string
}

// TicketsByCompanyBranchCounterInfo represents the result of Query 15.
type TicketsByCompanyBranchCounterInfo struct {
	ID           string
	Service      string
	TicketStatus int
	Company      string
	Branch       string
	Counter      string
}

// TicketsWithAvgServiceTimeByCounter represents the result of Query 16.
type TicketsWithAvgServiceTimeByCounter struct {
	CounterName    string
	AvgServiceTime float64
}

// AverageActiveTimeOfUser represents the result of Query 17.
type AverageActiveTimeOfUser struct {
	UserName             string
	AvgActiveTimeSeconds float64
	AvgActiveTimeMinutes float64
	AvgActiveTimeHours   float64
}

// ActiveTimeOfUserPerDay represents the result of Query 18.
type ActiveTimeOfUserPerDay struct {
	UserName          string
	ServingDate       time.Time
	FirstTicketTime   time.Time
	LastTicketTime    time.Time
	ActiveTimeSeconds int
	ActiveTimeMinutes int
	ActiveTimeHours   int
}

// Continue with other structs as needed...
