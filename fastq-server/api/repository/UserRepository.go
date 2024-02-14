package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/bentenison/fastq-server/models"
)

type sqlUserRepo struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *sqlUserRepo {
	return &sqlUserRepo{
		DB: db,
	}
}

// InsertQuery generates an insert query for the UpdateSystemParams struct
func InsertQuery(params models.UpdateSystemParams) string {
	return fmt.Sprintf(`
		INSERT INTO your_table_name (
			server_ip, server_cpu, server_disk_id, 
			client_ip, client_cpu, client_disk_id, 
			is_deleted, is_active
		) VALUES (
			'%s', '%s', '%s', 
			'%s', '%s', '%s', 
			%t, %t
		)`,
		params.ServerIP, params.ServerCPU, params.ServerDiskId,
		params.ClientIP, params.ClientCPU, params.ClientDiskId,
		params.IsDeleted, params.IsActive,
	)
}

// UpdateQuery generates an update query for the UpdateSystemParams struct
func UpdateQuery(params models.UpdateSystemParams) string {
	return fmt.Sprintf(`
		UPDATE your_table_name SET 
			server_ip = '%s',
			server_cpu = '%s',
			client_ip = '%s',
			client_cpu = '%s',
			client_disk_id = '%s',
			is_deleted = %t,
			is_active = %t
		WHERE server_disk_id = '%s'`,
		params.ServerIP, params.ServerCPU,
		params.ClientIP, params.ClientCPU, params.ClientDiskId,
		params.IsDeleted, params.IsActive, params.ServerDiskId,
	)
}
func (s *sqlUserRepo) UpdateSystem(ctx context.Context, arg models.UpdateSystemParams) error {
	_, err := s.DB.ExecContext(ctx, UpdateQuery(arg),
		arg.ServerCPU,
		arg.ServerDiskId,
		arg.ServerCPU,
		arg.ClientCPU,
		arg.ClientDiskId,
		arg.ClientIP,
		arg.IsActive,
		arg.IsDeleted,
	)
	return err
}
func (s *sqlUserRepo) GetServerDetailsByIP(serverIP string) (models.ServerDetails, error) {

	// Prepare the query
	query := "SELECT server_ip, server_cpu, server_disk_id, id FROM server_details WHERE server_ip = ?"

	// Execute the query
	row := s.DB.QueryRow(query, serverIP)

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
func (s *sqlUserRepo) InsetSystem(ctx context.Context, arg models.UpdateSystemParams) error {
	_, err := s.DB.ExecContext(ctx, InsertQuery(arg),
		arg.ServerCPU,
		arg.ServerDiskId,
		arg.ServerCPU,
		arg.ClientCPU,
		arg.ClientDiskId,
		arg.ClientIP,
		arg.IsActive,
		arg.IsDeleted,
	)
	return err
}

func (s *sqlUserRepo) AddServer(ctx context.Context, arg models.ServerDetails) error {
	query := "INSERT INTO server_details (server_ip, server_cpu, server_disk_id, id) VALUES (?, ?, ?, ?)"
	_, err := s.DB.Exec(query, arg.ServerIP, arg.ServerCPU, arg.ServerDiskId, arg.Id)
	if err != nil {
		return err
	}
	return nil
}
func (s *sqlUserRepo) GetUserById(ctx context.Context, id string) (*models.UserAccount, error) {
	user := &models.UserAccount{}
	stmt, err := s.DB.PrepareContext(ctx, "select * from company_info where company_code = ?")
	if err != nil {
		log.Printf("cannot create a sql statement reason: %v ", stmt)
		return nil, err

	}
	defer stmt.Close()
	err = stmt.QueryRowContext(ctx, id).Scan(&user.UserID, &user.Email, &user.FirstName, &user.LastName,
		&user.CompanyCode, &user.Company, &user.Country, &user.Timezone, &user.Password, &user.UpdatedBy, &user.CreatedAt, &user.UpdatedAt)
	switch {
	case err == sql.ErrNoRows:
		log.Printf("no user with id %s", id)
		return nil, sql.ErrNoRows
	case err != nil:
		log.Printf("cannot query database ,reson: %v", err.Error())
	}
	if err != nil {
		log.Printf("could not query : %v ", stmt)
		return nil, err

	}
	return user, nil
}
func (s *sqlUserRepo) GetUserByEmail(ctx context.Context, email string) (*models.UserAccount, error) {
	user := &models.UserAccount{}
	stmt, err := s.DB.PrepareContext(ctx, "select * from company_info where email = ?")
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
func (s *sqlUserRepo) GetBranchByEmail(ctx context.Context, email string) (*models.BranchAdmin, error) {
	user := &models.BranchAdmin{}
	stmt, err := s.DB.PrepareContext(ctx, "select * from branch_admin where email = ?")
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

func (s *sqlUserRepo) Create(ctx context.Context, u *models.UserAccount) error {
	query := "INSERT INTO company_info (id,email,password,firstname,lastname,company_code,company,country,timezone,created_at,updated_at,updated_by) VALUES (?,?,?,?,?,?,?,?,?,?,?,?)"

	rows, err := s.DB.QueryContext(ctx, query, u.UserID, u.Email, u.Password, u.FirstName, u.LastName, u.CompanyCode,
		u.Company, u.Country, u.Timezone, u.CreatedAt, u.UpdatedAt, u.UpdatedBy)
	if err != nil {
		// check unique constraint
		log.Printf("Could not create a user with email: %v. Reason: %v\n", u.Email, err)
		return err
	}
	log.Println(rows)

	return nil
}

func (s *sqlUserRepo) GetTimezone(ctx context.Context, country string) ([]*models.TimeZones, error) {
	timezones := []*models.TimeZones{}
	stmt, err := s.DB.PrepareContext(ctx, "select zone_name,country_code,abbreviation,time_start,gmt_offset from time_zone where country_code = ?")
	if err != nil {
		log.Printf("cannot create a sql statement reason: %v ", stmt)
		return nil, err

	}
	defer stmt.Close()
	rows, err := stmt.QueryContext(ctx, country)
	if err != nil {
		log.Printf("could not query : %v ", stmt)
		return nil, err

	}
	defer rows.Close()
	for rows.Next() {
		var tz models.TimeZones
		if err := rows.Scan(&tz.ZoneName, &tz.CountryCode, &tz.Abbreviation, &tz.TimeStart, &tz.GMT); err != nil {
			// Check for a scan error.
			// Query rows will be closed with defer.
			log.Println("error scanning rows into type", err)
		}
		timezones = append(timezones, &tz)
	}
	return timezones, nil
}
func (s *sqlUserRepo) GetCountries(ctx context.Context) ([]*models.Countries, error) {
	countries := []*models.Countries{}
	stmt, err := s.DB.PrepareContext(ctx, "select * from country")
	if err != nil {
		log.Printf("cannot create a sql statement reason: %v ", stmt)
		return nil, err

	}
	defer stmt.Close()
	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		log.Printf("could not query : %v ", stmt)
		return nil, err

	}
	defer rows.Close()
	for rows.Next() {
		var country models.Countries
		if err := rows.Scan(&country.CountryCode, &country.CountryName); err != nil {
			// Check for a scan error.
			// Query rows will be closed with defer.
			log.Println("error scanning rows into type", err)
		}
		countries = append(countries, &country)
	}
	return countries, nil
}

// func (h *sqlUserRepo) Update() {

// }
// func (h *sqlUserRepo) Delete() {

// }
