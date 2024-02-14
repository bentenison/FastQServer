package services

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/bentenison/fastq-server/models"
	"github.com/robfig/cron"
)

type Schedular struct {
	DB *sql.DB
}

func NewSchedular(db *sql.DB) Schedular {
	return Schedular{DB: db}
}
func GetTicketsbyBranchFromCloud() ([]models.AddTicketParams, error) {
	result := []models.AddTicketParams{}
	url := os.Getenv("CLOUDURL")
	url = url + models.COMAPNY_CODE
	method := "POST"

	// Create a new HTTP request
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Println("Error creating request:", err)
		return result, err
	}

	// Set the Content-Type header for JSON
	req.Header.Set("Content-Type", "application/json")

	// Make the request
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		log.Println("Error making the request:", err)
		return result, err
	}
	defer response.Body.Close()

	// Read and print the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println("Error reading the response body:", err)
		return result, err
	}

	log.Println(string(body))
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Println("Error unmarshaling the response body:", err)
		return result, err
	}
	return result, nil

}
func (s Schedular) AddCloudTicketsToLocal() {
	tickets, err := GetTicketsbyBranchFromCloud()
	if err != nil {
		log.Println("Error while getting cloud tickets", err)
		return
	}
	//add the last ticket number here
	number, err := GetLastTicketNumber(context.TODO(), s.DB)
	if err != nil {
		log.Println("error occured while getting last called ticket number", err)
	}
	for _, ticket := range tickets {
		ticket.TicketNumber = number.Number
		ticketPrefix := strings.Split(ticket.TicketName, "-")[0]
		ticket.TicketName = ticketPrefix + "-" + number.Number
	}
}
func GetLastTicketNumber(ctx context.Context, db *sql.DB) (models.TicketNumber, error) {
	row := db.QueryRowContext(ctx, `select ticket_number as number from ticket where date(created_at)=date(now()) and ticket_status = 0 order by created_at desc limit 1;`)
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
func AddBranchToCloud(branch models.UserAccount) {
	branchBytes, err := json.Marshal(branch)
	if err != nil {
		log.Println("Error marshaling request:", err)
		return
	}
	url := "https://jsonplaceholder.typicode.com/posts"
	method := "POST"

	// Create a new HTTP request
	req, err := http.NewRequest(method, url, bytes.NewBuffer(branchBytes))
	if err != nil {
		log.Println("Error creating request:", err)
		return
	}

	// Set the Content-Type header for JSON
	req.Header.Set("Content-Type", "application/json")

	// Make the request
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		log.Println("Error making the request:", err)
		return
	}
	defer response.Body.Close()

	// Read and print the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println("Error reading the response body:", err)
		return
	}

	log.Println(string(body))

}
func AddServiceToCloud(service models.AddServiceParams) {
	url := "https://jsonplaceholder.typicode.com/posts"
	method := "POST"

	// Create a sample JSON payload
	// payload := []byte(`{"title":"foo","body":"bar","userId":1}`)
	serviceBytes, err := json.Marshal(service)
	if err != nil {
		log.Println("Error marshaling request:", err)
		return
	}

	// Create a new HTTP request
	req, err := http.NewRequest(method, url, bytes.NewBuffer(serviceBytes))
	if err != nil {
		log.Println("Error creating request:", err)
		return
	}

	// Set the Content-Type header for JSON
	req.Header.Set("Content-Type", "application/json")

	// Make the request
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		log.Println("Error making the request:", err)
		return
	}
	defer response.Body.Close()

	// Read and print the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println("Error reading the response body:", err)
		return
	}

	log.Println(string(body))

}

func GetServerDetailsByIP(serverIP string, db *sql.DB) (models.ServerDetails, error) {

	// Prepare the query
	query := "SELECT server_ip, server_cpu, server_disk_id, id FROM server_details WHERE server_ip = ?"

	// Execute the query
	row := db.QueryRow(query, serverIP)

	// Scan the result into a ServerDetails struct
	var serverDetails models.ServerDetails
	err := row.Scan(&serverDetails.ServerIP, &serverDetails.ServerCPU, &serverDetails.ServerDiskId, &serverDetails.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			// No rows were returned
			return models.ServerDetails{}, fmt.Errorf("no server details found for server IP: %s", serverIP)
		}
		return models.ServerDetails{}, err
	}
	return serverDetails, nil
}

func RunCronJobs() {
	// 2
	s := cron.New()

	// 3
	s.AddFunc("@every 15s", func() {
		log.Println("Hello World!!!")
	})

	// 4
	s.Start()
}
