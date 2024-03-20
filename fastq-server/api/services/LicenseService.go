package services

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/bentenison/fastq-server/api/helpers"
	"github.com/bentenison/fastq-server/models"
)

type licenseService struct {
	db *sql.DB
}

const filpathstr = "license_" + "FASTQ Solutions" + ".json"

/*
* LSConfig will hold repositories that will eventually be injected into this service layer
* now it means the USConfig is used to hide the implementation of service layer
 */
type LSConfig struct {
	DB *sql.DB
}

// NewUserService is a factory function for
// initializing a UserService with its repository layer dependencies
func NewLicenseService(c *LSConfig) models.LicenseService {
	return &licenseService{
		db: c.DB,
	}
}

func (l *licenseService) UpdateLicenseService(lsc models.LicensePayload) (bool, error) {
	info, err := os.Stat(filpathstr)
	if err != nil {
		log.Println("license file not present, activate system", err)
		return false, nil
	}
	licenseBytes, err := os.ReadFile(filpathstr)
	if err != nil {
		log.Println("unable to read activation file, activate system", err)
		return false, nil
	}
	decryptedData, err := helpers.DecryptData(licenseBytes, info.Name())
	if err != nil {
		log.Println("unable to decrypt activation file, activate system", err)
		return false, nil
	}
	var license models.LicensePayload
	err = json.Unmarshal(decryptedData, &license)
	if err != nil {
		log.Println("unmarshal err, activate system", err)
		return false, nil
	}
	license.AllowedCountersPerBranch = lsc.AllowedCountersPerBranch
	Lbytes, err := json.Marshal(license)
	if err != nil {
		log.Println("marshal err, activate system", err)
		return false, nil
	}
	helpers.SaveDataToFile(filpathstr, Lbytes, true)
	if err != nil {
		log.Println("error updating license", err)
		return false, nil
	}
	return true, nil
}
func (l *licenseService) CheckCounterActivationService(ctr models.ActiveCounter) (bool, error) {
	info, err := os.Stat(filpathstr)
	if err != nil {
		log.Println("license file not present, activate system", err)
		return false, err
	}
	licenseBytes, err := os.ReadFile(filpathstr)
	if err != nil {
		log.Println("unable to read activation file, activate system", err)
		return false, err
	}
	decryptedData, err := helpers.DecryptData(licenseBytes, info.Name())
	if err != nil {
		log.Println("unable to decrypt activation file, activate system", err)
		return false, nil
	}
	var license models.LicensePayload
	err = json.Unmarshal(decryptedData, &license)
	if err != nil {
		log.Println("unmarshal err, activate system", err)
		return false, err
	}
	if license.Firm != ctr.BranchName {
		return false, nil
	}
	for i, _ := range license.ActiveCounterPerBranch {
		// license.ActiveCounterPerBranch[i].BranchId = ctr.BranchId
		// license.ActiveCounterPerBranch[i].CounterId = ctr.CounterId
		if license.ActiveCounterPerBranch[i].CpuId == ctr.CpuId && license.ActiveCounterPerBranch[i].HardDiskId == ctr.HardDiskId {
			return true, nil
		}

	}
	// Lbytes, err := json.Marshal(license)
	// if err != nil {
	// 	log.Println("marshal err, activate system", err)
	// 	return false, nil
	// }
	// helpers.SaveDataToFile(filpathstr, Lbytes, true)
	// if err != nil {
	// 	log.Println("error activating counter", err)
	// 	return false, nil
	// }
	return false, nil
}
func (l *licenseService) UpdateCounterDetails(ctr models.ActiveCounter) (bool, error) {
	return true, nil
}
func (l *licenseService) AuthCounterUserService(ctr models.AuthCounterUser) (*models.ManageUser, error) {
	user, err := AuthCounterUser(context.TODO(), ctr, l.db)
	if err != nil {
		log.Println("user not present,", err)
		return &models.ManageUser{}, errors.New("User is not present")
	}
	ok, err := l.CheckFirmNameInLicense(user.CompanyName.String)
	if err != nil {
		log.Println("firm not present,", err)
		return &models.ManageUser{}, errors.New("firm validation failed")
	}
	if !ok {
		log.Println("firm not present,", err)
		return &models.ManageUser{}, errors.New("firm validation failed")
	}
	counter, err := checkUserAlreayLoggedIn(context.TODO(), *user, l.db)
	if err == sql.ErrNoRows || counter.UserID.String == "" {
		log.Println("user not present,", err)
		err := updateCounter(context.TODO(), *user, ctr.CounterId, l.db)
		if err != nil {
			log.Print("error occured while updating counter", err)
			return user, err
		}
		return user, nil
		// return &models.ManageUser{}, nil
	}
	if counter.UserID.String != "" {
		return &models.ManageUser{}, errors.New("user is already logged into another counter")
	}
	if err != nil {
		log.Println("user not present,", err)
		return &models.ManageUser{}, errors.New("user is not present")
	}
	return &models.ManageUser{}, errors.New("error user is logged in to other counter")
}
func (l *licenseService) CounterLogoutService(ctrId string, user models.ManageUser) error {
	user.ID.String = ""
	err := updateCounter(context.TODO(), user, ctrId, l.db)
	if err != nil {
		log.Print("error occured while updating counter", err)
		return err
	}
	return nil
	// return &models.ManageUser{}, nil
}
func (l *licenseService) ActivateCounterService(ctr models.ActiveCounter) (bool, error) {
	info, err := os.Stat(filpathstr)
	if err != nil {
		log.Println("license file not present, activate system", err)
		return false, nil
	}
	licenseBytes, err := os.ReadFile(filpathstr)
	if err != nil {
		log.Println("unable to read activation file, activate system", err)
		return false, nil
	}
	decryptedData, err := helpers.DecryptData(licenseBytes, info.Name())
	if err != nil {
		log.Println("unable to decrypt activation file, activate system", err)
		return false, nil
	}
	var license models.LicensePayload
	err = json.Unmarshal(decryptedData, &license)
	if err != nil {
		log.Println("unmarshal err, activate system", err)
		return false, nil
	}

	if license.Firm != ctr.BranchName {
		return false, errors.New("error while checking firm in license")
	}
	var activeCounter models.ActiveCounter
	activeCounter.BranchId = ctr.BranchId
	activeCounter.CounterId = ctr.CounterId
	activeCounter.CpuId = ctr.CpuId
	activeCounter.HardDiskId = ctr.HardDiskId
	if len(license.ActiveCounterPerBranch) >= license.AllowedCountersPerBranch {
		return false, errors.New("maximum allowed counters are already activated")
	}

	// for i, _ := range license.ActiveCounterPerBranch {
	// 	license.ActiveCounterPerBranch[i].BranchId = ctr.BranchId
	// 	license.ActiveCounterPerBranch[i].CounterId = ctr.CounterId
	// 	license.ActiveCounterPerBranch[i].CpuId = ctr.CpuId
	// 	license.ActiveCounterPerBranch[i].HardDiskId = ctr.HardDiskId
	// }
	license.ActiveCounterPerBranch = append(license.ActiveCounterPerBranch, activeCounter)
	Lbytes, err := json.Marshal(license)
	if err != nil {
		log.Println("marshal err, activate system", err)
		return false, nil
	}
	err = helpers.SaveDataToFile(filpathstr, Lbytes, true)
	if err != nil {
		log.Println("error activating counter", err)
		return false, nil
	}
	//insert syatem info in systemns
	system := models.UpdateSystemParams{}
	system.ServerIP = helpers.DiscoverService()
	if err != nil {
		log.Println("error getting IP", err)
		return false, err
	}
	system.ServerDiskId, err = helpers.GetHDDId()
	if err != nil {
		log.Println("error getting harddisk Id", err)
		return false, err
	}
	system.ServerCPU, err = helpers.GetCPUId()
	if err != nil {
		log.Println("error getting cpu Id", err)
		return false, err
	}
	system.ClientCPU = ctr.CpuId
	system.ClientDiskId = ctr.HardDiskId
	system.CounterId = ctr.CounterId
	system.CounterName = ctr.CounterName
	err = InsetSystem(context.TODO(), system, l.db)
	if err != nil {
		log.Println("error inserting system", err)
		return false, nil
	}
	activation := models.CounterActivation{
		ClientCPU:    system.ClientCPU,
		ClientDiskId: system.ClientDiskId,
		Id:           ctr.CounterId,
	}
	err = UpdateCounterActivation(context.TODO(), activation, l.db)
	if err != nil {
		log.Println("error updating counter", err)
		return false, nil
	}
	return true, nil
}
func (l *licenseService) UpdateCompanyNameInLicense(companyName string) (bool, error) {
	info, err := os.Stat(filpathstr)
	if err != nil {
		log.Println("license file not present, activate system", err)
		return false, nil
	}
	licenseBytes, err := os.ReadFile(filpathstr)
	if err != nil {
		log.Println("unable to read activation file, activate system", err)
		return false, nil
	}
	decryptedData, err := helpers.DecryptData(licenseBytes, info.Name())
	if err != nil {
		log.Println("unable to decrypt activation file, activate system", err)
		return false, nil
	}
	var license models.LicensePayload
	err = json.Unmarshal(decryptedData, &license)
	if err != nil {
		log.Println("unmarshal err, activate system", err)
		return false, nil
	}
	license.Firm = companyName
	Lbytes, err := json.Marshal(license)
	if err != nil {
		log.Println("marshal err, activate system", err)
		return false, nil
	}
	err = helpers.SaveDataToFile(filpathstr, Lbytes, true)
	if err != nil {
		log.Println("error activating counter", err)
		return false, nil
	}
	return true, nil
}
func (l *licenseService) DeleteCounterFromLicense(id string) (bool, error) {
	info, err := os.Stat(filpathstr)
	if err != nil {
		log.Println("license file not present, activate system", err)
		return false, nil
	}
	licenseBytes, err := os.ReadFile(filpathstr)
	if err != nil {
		log.Println("unable to read activation file, activate system", err)
		return false, nil
	}
	decryptedData, err := helpers.DecryptData(licenseBytes, info.Name())
	if err != nil {
		log.Println("unable to decrypt activation file, activate system", err)
		return false, nil
	}
	var license models.LicensePayload
	err = json.Unmarshal(decryptedData, &license)
	if err != nil {
		log.Println("unmarshal err, activate system", err)
		return false, nil
	}
	for i, v := range license.ActiveCounterPerBranch {
		if v.CounterId == id {
			license.ActiveCounterPerBranch = append(license.ActiveCounterPerBranch[:i], license.ActiveCounterPerBranch[i+1:]...)
			break
		}
	}
	Lbytes, err := json.Marshal(license)
	if err != nil {
		log.Println("marshal err, activate system", err)
		return false, nil
	}
	err = helpers.SaveDataToFile(filpathstr, Lbytes, true)
	if err != nil {
		log.Println("error activating counter", err)
		return false, nil
	}
	err = deletesystem(context.Background(), id, l.db)
	if err != nil {
		log.Println("error deleting system", err)
		return false, nil
	}
	return true, nil
}
func (l *licenseService) CheckFirmNameInLicense(companyName string) (bool, error) {
	info, err := os.Stat(filpathstr)
	if err != nil {
		log.Println("license file not present, activate system", err)
		return false, nil
	}
	licenseBytes, err := os.ReadFile(filpathstr)
	if err != nil {
		log.Println("unable to read activation file, activate system", err)
		return false, nil
	}
	decryptedData, err := helpers.DecryptData(licenseBytes, info.Name())
	if err != nil {
		log.Println("unable to decrypt activation file, activate system", err)
		return false, nil
	}
	var license models.LicensePayload
	err = json.Unmarshal(decryptedData, &license)
	if err != nil {
		log.Println("unmarshal err, activate system", err)
		return false, nil
	}
	if license.Firm == companyName {
		return true, nil
	}
	return false, nil
}
func (l *licenseService) CheckLicensePresent() (bool, error) {
	info, err := os.Stat(filpathstr)
	if err != nil {
		log.Println("license file not present, activate system", err)
		return false, nil
	}
	licenseBytes, err := os.ReadFile(filpathstr)
	if err != nil {
		log.Println("unable to read activation file, activate system", err)
		return false, nil
	}
	_, err = helpers.DecryptData(licenseBytes, info.Name())
	if err != nil {
		log.Println("unable to decrypt activation file, activate system", err)
		return false, nil
	}
	// log.Println("decrypted Data", decryptedData)

	return true, nil
}
func (l *licenseService) CheckAndGetActiveCounter(arg models.ActiveCounter) (models.ManageCounter, error) {
	info, err := CheckAndGetCounter(context.TODO(), arg, l.db)
	if err != nil {
		log.Println("license file not present, activate system", err)
		return models.ManageCounter{}, nil
	}

	return info, nil
}
func (l *licenseService) UpdateCounterUserService(arg string) error {
	err := updateCounterUser(context.TODO(), arg, l.db)
	if err != nil {
		log.Println("update counter user error ", err)
		return nil
	}
	return nil
}

// InsertQuery generates an insert query for the UpdateSystemParams struct
func InsertQuery(params models.UpdateSystemParams) string {
	return fmt.Sprintf(`
		INSERT INTO systems (
			server_ip, server_cpu, server_disk_id, 
			client_ip, client_cpu, client_disk_id, 
			is_deleted, is_active,counter_name,counter_id
		) VALUES (
			'%s', '%s', '%s', 
			'%s', '%s', '%s', 
			%t, %t,'%s','%s'
		)`,
		params.ServerIP, params.ServerCPU, params.ServerDiskId,
		params.ClientIP, params.ClientCPU, params.ClientDiskId,
		params.IsDeleted, params.IsActive, params.CounterName, params.CounterId,
	)
}

// UpdateQuery generates an update query for the UpdateSystemParams struct
func UpdateQuery(params models.UpdateSystemParams) string {
	return fmt.Sprintf(`
		UPDATE systems SET 
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
func UpdateSystem(ctx context.Context, arg models.UpdateSystemParams, db *sql.DB) error {
	_, err := db.ExecContext(ctx, UpdateQuery(arg))
	return err
}
func InsetSystem(ctx context.Context, arg models.UpdateSystemParams, db *sql.DB) error {
	_, err := db.ExecContext(ctx, InsertQuery(arg))
	return err
}
func deletesystem(ctx context.Context, id string, db *sql.DB) error {
	_, err := db.ExecContext(ctx, fmt.Sprintf("delete from systems where counter_id = %s;", id))
	return err
}
func UpdateCounterActivation(ctx context.Context, arg models.CounterActivation, db *sql.DB) error {
	_, err := db.ExecContext(ctx, `update manage_counters set cpuId =?,diskId = ?,activated=1 where id =?`,
		arg.ClientCPU,
		arg.ClientDiskId,
		arg.Id,
	)
	return err
}

const getCounter = `-- name: GetSection :one
select id, section, sub_section, user_id, counter_number, counter_name, override_fifo, transfer_q, asssign_user, assign_service, transfer_priority, cancel_q, activated,cpuId ,diskId, branch_code, branch_name, company_name, company_code, created_at, updated_at, created_by, updated_by
from manage_counters
where cpuId = ?
    and diskId = ?
`

func CheckAndGetCounter(ctx context.Context, arg models.ActiveCounter, db *sql.DB) (models.ManageCounter, error) {
	row := db.QueryRowContext(ctx, getCounter, arg.CpuId, arg.HardDiskId)
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
	return i, err
}

const getUser = `-- name: GetBranch :one
select id, email, firstname, lastname, created_at, user_type, suspended, branch_name, branch_code, company_name, company_code, title, password, image_url, updated_at, created_by, updated_by
from manage_user
where email = ? and password = ?
`

func AuthCounterUser(ctx context.Context, arg models.AuthCounterUser, db *sql.DB) (*models.ManageUser, error) {
	row := db.QueryRowContext(ctx, getUser, arg.Email, arg.Password)
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
func checkUserAlreayLoggedIn(ctx context.Context, arg models.ManageUser, db *sql.DB) (*models.ManageCounter, error) {
	var i models.ManageCounter
	row := db.QueryRowContext(ctx, "select * from manage_counters where user_id = ? ", arg.ID.String)
	if row.Err() == sql.ErrNoRows {
		return &i, sql.ErrNoRows
	}
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
func updateCounter(ctx context.Context, arg models.ManageUser, ctrId string, db *sql.DB) error {
	// var i models.ManageCounter
	res, err := db.ExecContext(ctx, "update manage_counters set user_id = ? where id = ? ", arg.ID.String, ctrId)
	if err != nil {
		log.Println("error while updating counters", err)
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		log.Println("error while updating counters", err)
		return err
	}
	if affected <= 0 {
		log.Println("error while updating counters", err)
		return err
	}
	return nil
}

func updateCounterUser(ctx context.Context, arg string, db *sql.DB) error {
	// var i models.ManageCounter
	res, err := db.ExecContext(ctx, "update manage_counters set user_id = '' where id = ? ", arg)
	if err != nil {
		log.Println("error while updating counters", err)
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		log.Println("error while updating counters", err)
		return err
	}
	if affected <= 0 {
		log.Println("error while updating counters", err)
		return err
	}
	return nil
}
