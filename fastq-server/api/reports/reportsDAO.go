package reports

import (
	"database/sql"
	"log"
	"strconv"
	"time"

	"github.com/bentenison/fastq-server/models"
)

type ReportsDAO struct {
	DB *sql.DB
}

func NewReportsDAO(db *sql.DB) *ReportsDAO {
	return &ReportsDAO{DB: db}
}

// Query 1: Total Tickets by Service
func (dao *ReportsDAO) GetTotalTicketsByService(companyCode string) (map[string]int, error) {
	query := "SELECT service, COUNT(*) as total_tickets FROM ticket WHERE company_code = ? GROUP BY service"
	rows, err := dao.DB.Query(query, companyCode)
	if err != nil {
		log.Println("Error executing query:", err)
		return nil, err
	}
	defer rows.Close()

	result := make(map[string]int)
	for rows.Next() {
		var service string
		var totalTickets int
		err := rows.Scan(&service, &totalTickets)
		if err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}
		result[service] = totalTickets
	}

	return result, nil
}

// Query 2: Tickets by Status
func (dao *ReportsDAO) GetTicketsByStatus(companyCode string) (map[string]int, error) {
	query := "SELECT ticket_status, COUNT(*) as total_tickets FROM ticket WHERE company_code = ? GROUP BY ticket_status"
	rows, err := dao.DB.Query(query, companyCode)
	if err != nil {
		log.Println("Error executing query:", err)
		return nil, err
	}
	defer rows.Close()

	result := make(map[string]int)
	for rows.Next() {
		var ticketStatus string
		var totalTickets int
		err := rows.Scan(&ticketStatus, &totalTickets)
		if err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}
		result[ticketStatus] = totalTickets
	}

	return result, nil
}
func (dao *ReportsDAO) GetAllWaitingTickets(companyCode string) (string, error) {
	query := "SELECT COUNT(*) as total_tickets FROM ticket WHERE company_code = ? AND DATE(created_at) = curdate()"
	rows, err := dao.DB.Query(query, companyCode)
	if err != nil {
		log.Println("Error executing query:", err)
		return "", err
	}
	defer rows.Close()

	result := ""
	for rows.Next() {
		var ticketStatus string
		// var totalTickets int
		err := rows.Scan(&ticketStatus)
		if err != nil {
			log.Println("Error scanning row:", err)
			return "nil", err
		}
		result = ticketStatus

	}

	return result, nil
}

// Query 3: Average Service Time by Service
func (dao *ReportsDAO) GetAverageServiceTimeByService(companyCode string) (map[string]float64, error) {
	query := "SELECT service,DATE(started_serving_at) as served_on, AVG(TIMESTAMPDIFF(MINUTE, started_serving_at, end_serving_at)) as avg_service_time_seconds FROM ticket WHERE started_serving_at IS NOT NULL AND end_serving_at IS NOT NULL AND company_code = ? GROUP BY service, DATE(started_serving_at)"
	rows, err := dao.DB.Query(query, companyCode)
	if err != nil {
		log.Println("Error executing query:", err)
		return nil, err
	}
	defer rows.Close()

	result := make(map[string]float64)
	for rows.Next() {
		var service string
		var avgServiceTimeSeconds float64
		var servedOn string
		err := rows.Scan(&service, &servedOn, &avgServiceTimeSeconds)
		if err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}
		result[service] = avgServiceTimeSeconds
	}

	return result, nil
}

// Continue with other queries...
// Query 4: Tickets Transferred by Counter
func (dao *ReportsDAO) GetTicketsTransferredByCounter(companyCode string) (map[string]int, error) {
	query := "SELECT transfered_by, COUNT(*) as transferred_tickets FROM ticket WHERE transfered_by IS NOT NULL AND company_code = ? GROUP BY transfered_by"
	rows, err := dao.DB.Query(query, companyCode)
	if err != nil {
		log.Println("Error executing query:", err)
		return nil, err
	}
	defer rows.Close()

	result := make(map[string]int)
	for rows.Next() {
		var transferredBy string
		var transferredTickets int
		err := rows.Scan(&transferredBy, &transferredTickets)
		if err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}
		result[transferredBy] = transferredTickets
	}

	return result, nil
}

// Query 5: Tickets by Company and Branch
func (dao *ReportsDAO) GetTicketsByCompanyBranch(companyCode string) (map[string]map[string]int, error) {
	query := "SELECT company_name, COUNT(*) as total_tickets FROM ticket WHERE company_code = ? GROUP BY company_name"
	rows, err := dao.DB.Query(query, companyCode)
	if err != nil {
		log.Println("Error executing query:", err)
		return nil, err
	}
	defer rows.Close()

	result := make(map[string]map[string]int)
	for rows.Next() {
		var companyName, branchName string
		var totalTickets int
		err := rows.Scan(&companyName, &totalTickets)
		if err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}
		if result[companyName] == nil {
			result[companyName] = make(map[string]int)
		}
		result[companyName][branchName] = totalTickets
	}

	return result, nil
}

// Query 6: Average Waiting Time
func (dao *ReportsDAO) GetAverageWaitingTime(companyCode string) (map[string]float64, error) {
	query := "SELECT service,DATE(created_at) as created_on,AVG(TIMESTAMPDIFF(MINUTE, created_at, started_serving_at)) as avg_waiting_time_seconds FROM ticket WHERE started_serving_at IS NOT NULL AND company_code = ? GROUP BY service,DATE(created_at)"
	rows, err := dao.DB.Query(query, companyCode)
	if err != nil {
		log.Println("Error executing query:", err)
		return nil, err
	}
	defer rows.Close()

	result := make(map[string]float64)
	for rows.Next() {
		var service string
		var served_on string
		var avgWaitingTimeSeconds float64
		err := rows.Scan(&service, &served_on, &avgWaitingTimeSeconds)
		if err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}
		result[service] = avgWaitingTimeSeconds
	}

	return result, nil
}

// Query 7: Tickets Served by Counter
func (dao *ReportsDAO) GetTicketsServedByCounter(companyCode string) (map[string]int, error) {
	query := "SELECT counter_id, COUNT(*) as served_tickets FROM ticket WHERE started_serving_at IS NOT NULL AND company_code = ? GROUP BY counter_id"
	rows, err := dao.DB.Query(query, companyCode)
	if err != nil {
		log.Println("Error executing query:", err)
		return nil, err
	}
	defer rows.Close()

	result := make(map[string]int)
	for rows.Next() {
		var counterID string
		var servedTickets int
		err := rows.Scan(&counterID, &servedTickets)
		if err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}
		result[counterID] = servedTickets
	}

	return result, nil
}

// Query 8: Percentage of tickets finished by service
func (dao *ReportsDAO) GetPercentageTicketsFinishedByService(companyCode string) (map[string]map[string]interface{}, error) {
	query := "SELECT service, COUNT(*) AS total_tickets, SUM(CASE WHEN end_serving_at IS NOT NULL THEN 1 ELSE 0 END) AS finished_tickets, (SUM(CASE WHEN end_serving_at IS NOT NULL THEN 1 ELSE 0 END) / COUNT(*)) * 100 AS percentage_finished FROM ticket WHERE company_code = ? GROUP BY service"
	rows, err := dao.DB.Query(query, companyCode)
	if err != nil {
		log.Println("Error executing query:", err)
		return nil, err
	}
	defer rows.Close()

	result := make(map[string]map[string]interface{})
	for rows.Next() {
		var service string
		var totalTickets, finishedTickets int
		var percentageFinished float64
		err := rows.Scan(&service, &totalTickets, &finishedTickets, &percentageFinished)
		if err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}
		result[service] = map[string]interface{}{
			"total_tickets":       totalTickets,
			"finished_tickets":    finishedTickets,
			"percentage_finished": percentageFinished,
		}
	}

	return result, nil
}

// Query 9: Tickets between 2 dates by ticket_status and user
func (dao *ReportsDAO) GetTicketsBetweenDatesByStatusUser(companyCode string, startDate, endDate string) (map[string]map[string]int, error) {
	query := "SELECT ticket_status, updated_by AS user, COUNT(*) AS ticket_count FROM ticket WHERE created_at BETWEEN ? AND ? AND company_code = ? GROUP BY ticket_status, updated_by"
	rows, err := dao.DB.Query(query, startDate, endDate, companyCode)
	if err != nil {
		log.Println("Error executing query:", err)
		return nil, err
	}
	defer rows.Close()

	result := make(map[string]map[string]int)
	for rows.Next() {
		var ticketStatus int
		var user string
		var ticketCount int
		err := rows.Scan(&ticketStatus, &user, &ticketCount)
		if err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}
		if result[user] == nil {
			result[user] = make(map[string]int)
		}
		result[user][strconv.Itoa(ticketStatus)] = ticketCount
	}

	return result, nil
}

// Query 10: Tickets with User Information
func (dao *ReportsDAO) GetTicketsWithUserInfo(companyCode string) ([]map[string]interface{}, error) {
	// query := "SELECT t.id, t.service, t.ticket_status, u.email AS user_name FROM ticket t INNER JOIN manage_user u ON t.updated_by = u.id WHERE t.company_code = ?"
	// SELECT concat(u.firstname," ",u.lastname) AS user_name,c.counter_name,t.ticket_name, t.service, t.ticket_status,c.company_name FROM ticket t INNER JOIN manage_user u ON t.updated_by = u.id inner join manage_counters c on t.counter_id=c.id WHERE t.company_code = "nDmmyD0" and date(t.created_at)=date(now())
	query := `SELECT concat(u.firstname," ",u.lastname) AS user_name,
	c.counter_name,t.ticket_name, t.service, t.ticket_status,c.company_name,TIMESTAMPDIFF(SECOND, t.started_serving_at, t.end_serving_at) as time_taken
	FROM ticket t INNER JOIN manage_user u ON t.updated_by = u.id inner join manage_counters c on t.counter_id=c.id WHERE t.company_code = ? `
	rows, err := dao.DB.Query(query, companyCode)
	if err != nil {
		log.Println("Error executing query:", err)
		return nil, err
	}
	defer rows.Close()

	result := make([]map[string]interface{}, 0)
	for rows.Next() {
		var tickeName, CounterName, service, company, ticketStatus, userName, time_taken string
		err := rows.Scan(&userName, &CounterName, &tickeName, &service, &ticketStatus, &company, &time_taken)
		if err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}
		data := map[string]interface{}{
			"user_name":     userName,
			"counter_name":  CounterName,
			"ticket_name":   tickeName,
			"service":       service,
			"ticket_status": ticketStatus,
			"company_name":  userName,
			"time_taken":    time_taken,
		}
		result = append(result, data)
	}

	return result, nil
}

// Query 11: Tickets with Counter Information
func (dao *ReportsDAO) GetTicketsWithCounterInfo(companyCode string) ([]map[string]interface{}, error) {
	query := "SELECT t.id, t.service, t.ticket_status, c.counter_name FROM ticket t INNER JOIN manage_counters c ON t.counter_id = c.id WHERE t.company_code = ?"
	rows, err := dao.DB.Query(query, companyCode)
	if err != nil {
		log.Println("Error executing query:", err)
		return nil, err
	}
	defer rows.Close()

	result := make([]map[string]interface{}, 0)
	for rows.Next() {
		var ticketID, service, ticketStatus, counterName string
		err := rows.Scan(&ticketID, &service, &ticketStatus, &counterName)
		if err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}
		data := map[string]interface{}{
			"id":            ticketID,
			"service":       service,
			"ticket_status": ticketStatus,
			"counter_name":  counterName,
		}
		result = append(result, data)
	}

	return result, nil
}

// Query 12: Tickets with User and Counter Information
func (dao *ReportsDAO) GetTicketsWithUserCounterInfo(companyCode string) ([]map[string]interface{}, error) {
	query := "SELECT t.id, t.service, t.ticket_status, u.username AS user_name, c.counter_name FROM ticket t INNER JOIN manage_user u ON t.updated_by = u.id INNER JOIN manage_counters c ON t.counter_id = c.id WHERE t.company_code = ?"
	rows, err := dao.DB.Query(query, companyCode)
	if err != nil {
		log.Println("Error executing query:", err)
		return nil, err
	}
	defer rows.Close()

	result := make([]map[string]interface{}, 0)
	for rows.Next() {
		var ticketID, service, ticketStatus, userName, counterName string
		err := rows.Scan(&ticketID, &service, &ticketStatus, &userName, &counterName)
		if err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}
		data := map[string]interface{}{
			"id":            ticketID,
			"service":       service,
			"ticket_status": ticketStatus,
			"user_name":     userName,
			"counter_name":  counterName,
		}
		result = append(result, data)
	}

	return result, nil
}

// Query 13: Tickets Served by Counter with User Information
func (dao *ReportsDAO) GetTicketsServedByCounterWithUserInfo(companyCode string) ([]map[string]interface{}, error) {
	query := "SELECT t.id, t.service, t.ticket_status, c.counter_name, u.username AS user_name FROM ticket t INNER JOIN manage_counters c ON t.counter_id = c.id INNER JOIN manage_user u ON t.updated_by = u.id WHERE t.started_serving_at IS NOT NULL AND t.company_code = ?"
	rows, err := dao.DB.Query(query, companyCode)
	if err != nil {
		log.Println("Error executing query:", err)
		return nil, err
	}
	defer rows.Close()

	result := make([]map[string]interface{}, 0)
	for rows.Next() {
		var ticketID, service, ticketStatus, counterName, userName string
		err := rows.Scan(&ticketID, &service, &ticketStatus, &counterName, &userName)
		if err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}
		data := map[string]interface{}{
			"id":            ticketID,
			"service":       service,
			"ticket_status": ticketStatus,
			"counter_name":  counterName,
			"user_name":     userName,
		}
		result = append(result, data)
	}

	return result, nil
}

// Query 14: Tickets Transferred with User Information
func (dao *ReportsDAO) GetTicketsTransferredWithUserInfo(companyCode string) ([]map[string]interface{}, error) {
	query := "SELECT t.id, t.service, t.ticket_status, t.transfered_to, u.username AS transfered_by_user FROM ticket t INNER JOIN manage_user u ON t.transfered_by = u.id WHERE t.transfered_to IS NOT NULL AND t.company_code = ?"
	rows, err := dao.DB.Query(query, companyCode)
	if err != nil {
		log.Println("Error executing query:", err)
		return nil, err
	}
	defer rows.Close()

	result := make([]map[string]interface{}, 0)
	for rows.Next() {
		var ticketID, service, ticketStatus, transferedTo, transferedByUser string
		err := rows.Scan(&ticketID, &service, &ticketStatus, &transferedTo, &transferedByUser)
		if err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}
		data := map[string]interface{}{
			"id":                 ticketID,
			"service":            service,
			"ticket_status":      ticketStatus,
			"transfered_to":      transferedTo,
			"transfered_by_user": transferedByUser,
		}
		result = append(result, data)
	}

	return result, nil
}

// Query 15: Tickets by Company, Branch, and Counter Information
func (dao *ReportsDAO) GetTicketsByCompanyBranchCounter(companyCode string) ([]models.TicketsByCompanyBranchCounterInfo, error) {
	query := "SELECT t.id, t.service, t.ticket_status, t.company_name, t.branch_name, c.counter_name FROM ticket t INNER JOIN manage_counters c ON t.counter_id = c.id WHERE t.company_code = ?"
	rows, err := dao.DB.Query(query, companyCode)
	if err != nil {
		log.Println("Error executing query:", err)
		return nil, err
	}
	defer rows.Close()

	var result []models.TicketsByCompanyBranchCounterInfo
	for rows.Next() {
		var ticket models.TicketsByCompanyBranchCounterInfo
		err := rows.Scan(&ticket.ID, &ticket.Service, &ticket.TicketStatus, &ticket.Company, &ticket.Branch, &ticket.Counter)
		if err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}
		result = append(result, ticket)
	}

	return result, nil
}

// Query 16: Tickets with Average Service Time by Counter
func (dao *ReportsDAO) GetAverageServiceTimeByCounter(companyCode string) ([]models.TicketsWithAvgServiceTimeByCounter, error) {
	query := "SELECT c.counter_name, AVG(TIMESTAMPDIFF(SECOND, t.started_serving_at, t.end_serving_at)) as avg_service_time_seconds FROM ticket t INNER JOIN manage_counters c ON t.counter_id = c.id WHERE t.started_serving_at IS NOT NULL AND t.end_serving_at IS NOT NULL AND t.company_code = ? and t.ticket_status = 'FINISHED' and date(t.created_at) = date(now()) GROUP BY c.counter_name"
	rows, err := dao.DB.Query(query, companyCode)
	if err != nil {
		log.Println("Error executing query:", err)
		return nil, err
	}
	defer rows.Close()

	var result []models.TicketsWithAvgServiceTimeByCounter
	for rows.Next() {
		var ticket models.TicketsWithAvgServiceTimeByCounter
		err := rows.Scan(&ticket.CounterName, &ticket.AvgServiceTime)
		if err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}
		result = append(result, ticket)
	}

	return result, nil
}

// Query 17: Average Active Time of User in Minutes and Hours
func (dao *ReportsDAO) GetAverageActiveTimeOfUser(companyCode string) ([]models.AverageActiveTimeOfUser, error) {
	query := "SELECT u.username AS user_name, AVG(TIMESTAMPDIFF(SECOND, t.created_at, NOW())) as avg_active_time_seconds, AVG(TIMESTAMPDIFF(MINUTE, t.created_at, NOW())) as avg_active_time_minutes, AVG(TIMESTAMPDIFF(HOUR, t.created_at, NOW())) as avg_active_time_hours FROM ticket t INNER JOIN manage_user u ON t.updated_by = u.id WHERE t.company_code = ? GROUP BY u.username"
	rows, err := dao.DB.Query(query, companyCode)
	if err != nil {
		log.Println("Error executing query:", err)
		return nil, err
	}
	defer rows.Close()

	var result []models.AverageActiveTimeOfUser
	for rows.Next() {
		var activeTime models.AverageActiveTimeOfUser
		err := rows.Scan(&activeTime.UserName, &activeTime.AvgActiveTimeSeconds, &activeTime.AvgActiveTimeMinutes, &activeTime.AvgActiveTimeHours)
		if err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}
		result = append(result, activeTime)
	}

	return result, nil
}

// Query 18: Active Time of User per Day
func (dao *ReportsDAO) GetActiveTimeOfUserPerDay(companyCode string) ([]models.ActiveTimeOfUserPerDay, error) {
	query := "SELECT u.email AS user_name, DATE(t.started_serving_at) AS serving_date, MIN(t.started_serving_at) AS first_ticket_time, MAX(t.started_serving_at) AS last_ticket_time, TIMESTAMPDIFF(SECOND, MIN(t.started_serving_at), MAX(t.started_serving_at)) AS active_time_seconds, TIMESTAMPDIFF(MINUTE, MIN(t.started_serving_at), MAX(t.started_serving_at)) AS active_time_minutes, TIMESTAMPDIFF(HOUR, MIN(t.started_serving_at), MAX(t.started_serving_at)) AS active_time_hours FROM ticket t INNER JOIN manage_user u ON t.updated_by = u.id WHERE t.started_serving_at IS NOT NULL AND t.company_code = ? GROUP BY u.email, DATE(t.started_serving_at)"
	rows, err := dao.DB.Query(query, companyCode)
	if err != nil {
		log.Println("Error executing query:", err)
		return nil, err
	}
	defer rows.Close()

	var result []models.ActiveTimeOfUserPerDay
	for rows.Next() {
		var activeTime models.ActiveTimeOfUserPerDay
		err := rows.Scan(&activeTime.UserName, &activeTime.ServingDate, &activeTime.FirstTicketTime, &activeTime.LastTicketTime, &activeTime.ActiveTimeSeconds, &activeTime.ActiveTimeMinutes, &activeTime.ActiveTimeHours)
		if err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}
		result = append(result, activeTime)
	}

	return result, nil
}

// Continue with other queries...
type TicketStats struct {
	Date            string `json:"date"`
	FinishedTickets int    `json:"finishedTickets"`
	TotalTickets    int    `json:"totalTickets"`
}

// GetTicketStatsForLast8Days fetches finished and created ticket counts for the last 8 days.
func (dao *ReportsDAO) GetTicketStatsForLast8Days(code string) ([]TicketStats, error) {
	currentTime := time.Now()
	startDate := currentTime.AddDate(0, 0, -8) // last 8 days

	rows, err := dao.DB.Query(`
		SELECT
			DATE(created_at) as date,
			SUM(CASE WHEN ticket_status = 'FINISHED' THEN 1 ELSE 0 END) as finishedTickets,
			COUNT(*) as totalTickets
		FROM
			ticket
		WHERE
			DATE(created_at) >= ? AND DATE(created_at) <= ? AND company_code =?
		GROUP BY
			date
		ORDER BY
			date;
	`, startDate.Format("2006-01-02"), currentTime.Format("2006-01-02"), code)

	if err != nil {
		log.Println("Error executing query:", err)
		return nil, err
	}
	defer rows.Close()

	var ticketStatsList []TicketStats

	for rows.Next() {
		var ticketStats TicketStats
		err := rows.Scan(&ticketStats.Date, &ticketStats.FinishedTickets, &ticketStats.TotalTickets)
		if err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}
		ticketStatsList = append(ticketStatsList, ticketStats)
	}

	return ticketStatsList, nil
}

type HourlyFinishedTickets struct {
	Hour          int `json:"hour"`
	FinishedCount int `json:"finishedCount"`
}

// GetHourlyTicketStatsForDay fetches the count of tickets created every hour in a day.
func (dao *ReportsDAO) GetHourlyTicketFinishedForToday(code string) ([]HourlyFinishedTickets, error) {
	// Get the current date
	currentDate := time.Now().Format("2006-01-02")

	rows, err := dao.DB.Query(`
		SELECT
			HOUR(created_at) as hour,
			COUNT(*) as createdCount
		FROM
			ticket
		WHERE
			DATE(created_at) = ? AND company_code = ? AND ticket_status = 'FINISHED'
		GROUP BY
			hour
		ORDER BY
			hour;
	`, currentDate, code)

	if err != nil {
		log.Println("Error executing query:", err)
		return nil, err
	}
	defer rows.Close()

	var hourlyTicketStatsList []HourlyFinishedTickets

	for rows.Next() {
		var hourlyTicketStats HourlyFinishedTickets
		err := rows.Scan(&hourlyTicketStats.Hour, &hourlyTicketStats.FinishedCount)
		if err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}
		hourlyTicketStatsList = append(hourlyTicketStatsList, hourlyTicketStats)
	}

	return hourlyTicketStatsList, nil
}

type HourlyTicketStats struct {
	Hour         int `json:"hour"`
	CreatedCount int `json:"createdCount"`
}

// GetHourlyTicketStatsForDay fetches the count of tickets created every hour in a day.
func (dao *ReportsDAO) GetHourlyTicketStatsForToday(code string) ([]HourlyTicketStats, error) {
	// Get the current date
	currentDate := time.Now().Format("2006-01-02")

	rows, err := dao.DB.Query(`
		SELECT
			HOUR(created_at) as hour,
			COUNT(*) as createdCount
		FROM
			ticket
		WHERE
			DATE(created_at) = ? AND company_code = ? AND ticket_status = 'CREATED'
		GROUP BY
			hour
		ORDER BY
			hour;
	`, currentDate, code)

	if err != nil {
		log.Println("Error executing query:", err)
		return nil, err
	}
	defer rows.Close()

	var hourlyTicketStatsList []HourlyTicketStats

	for rows.Next() {
		var hourlyTicketStats HourlyTicketStats
		err := rows.Scan(&hourlyTicketStats.Hour, &hourlyTicketStats.CreatedCount)
		if err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}
		hourlyTicketStatsList = append(hourlyTicketStatsList, hourlyTicketStats)
	}

	return hourlyTicketStatsList, nil
}

type HourlyNoShow struct {
	Hour        int `json:"hour"`
	NoShowCount int `json:"noshowCount"`
}

// GetHourlyTicketStatsForDay fetches the count of tickets created every hour in a day.
func (dao *ReportsDAO) GetHourlyTicketNoShowForToday(code string) ([]HourlyNoShow, error) {
	// Get the current date
	currentDate := time.Now().Format("2006-01-02")

	rows, err := dao.DB.Query(`
		SELECT
			HOUR(created_at) as hour,
			COUNT(*) as createdCount
		FROM
			ticket
		WHERE
			DATE(created_at) = ? AND company_code = ? AND ticket_status = 'NO_SHOW'
		GROUP BY
			hour
		ORDER BY
			hour;
	`, currentDate, code)

	if err != nil {
		log.Println("Error executing query:", err)
		return nil, err
	}
	defer rows.Close()

	var hourlyTicketStatsList []HourlyNoShow

	for rows.Next() {
		var hourlyTicketStats HourlyNoShow
		err := rows.Scan(&hourlyTicketStats.Hour, &hourlyTicketStats.NoShowCount)
		if err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}
		hourlyTicketStatsList = append(hourlyTicketStatsList, hourlyTicketStats)
	}

	return hourlyTicketStatsList, nil
}
