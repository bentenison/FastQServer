package services

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/bentenison/fastq-server/models"
	"github.com/google/uuid"
)

type adminService struct {
	AdminRepository models.AdminRepository
	// ImageRepository model.ImageRepository
}

/*
* USConfig will hold repositories that will eventually be injected into this service layer
* now it means the USConfig is used to hide the implementation of service layer
 */
type ASConfig struct {
	AdminRepository models.AdminRepository
	// ImageRepository model.ImageRepository
}

// NewUserService is a factory function for
// initializing a UserService with its repository layer dependencies
func NewAdminService(c *ASConfig) models.AdminService {
	return &adminService{
		AdminRepository: c.AdminRepository,
		// ImageRepository: c.ImageRepository,
	}
}
func (as *adminService) AddbranchService(ctx context.Context, arg models.AddbranchParams) (sql.Result, error) {
	ad, err := as.AdminRepository.GetBranchByEmail(ctx, arg.Email)
	if err != nil && err != sql.ErrNoRows {
		log.Println("error while adding branch in DB", err)
		return nil, err
	}
	if ad != nil {
		return nil, errors.New("branch already exist")
	}
	if err == sql.ErrNoRows {
		arg.ID = uuid.NewString()
		const createdFormat = "2006-01-02 15:04:05"
		arg.CreatedAt = time.Now().Format(createdFormat)
		arg.UpdatedAt = time.Now().Format(createdFormat)
		result, err := as.AdminRepository.Addbranch(ctx, arg)
		if err != nil {
			log.Println("error while adding branch in DB", err)
			return nil, err
		}
		return result, nil

	}
	return nil, err
}

func (as *adminService) DeleteBranchService(ctx context.Context, code string) (sql.Result, error) {
	result, err := as.AdminRepository.DeleteBranch(ctx, code)
	if err != nil {
		log.Println("error while deleting branch from DB", err)
		return nil, err
	}
	return result, nil
}

func (as *adminService) GetAllBranchesService(ctx context.Context, companyCode string) ([]*models.ManageBranch, error) {
	result, err := as.AdminRepository.GetAllBranches(ctx, companyCode)
	if err != nil {
		log.Println("error while getting all branches from DB", err)
		return nil, err
	}
	return result, nil
}

func (as *adminService) GetBranchService(ctx context.Context, arg models.GetBranchParams) (*models.ManageBranch, error) {
	result, err := as.AdminRepository.GetBranch(ctx, arg)
	if err != nil {
		log.Println("error while adding branch in DB", err)
		return nil, err
	}
	return result, nil
}

func (as *adminService) UpdateBranchService(ctx context.Context, arg models.AddbranchParams) error {
	err := as.AdminRepository.UpdateBranch(ctx, arg)
	if err != nil {
		log.Println("error while updating branch in DB", err)
		return err
	}
	return nil
}

func (as *adminService) AddCounterService(ctx context.Context, arg models.AddCounterParams) (sql.Result, error) {
	arg.ID = uuid.NewString()
	const createdFormat = "2006-01-02 15:04:05"
	arg.CreatedAt = time.Now().Format(createdFormat)
	arg.UpdatedAt = time.Now().Format(createdFormat)
	result, err := as.AdminRepository.AddCounter(ctx, arg)
	if err != nil {
		log.Println("error while adding counter in DB", err)
		return nil, err
	}
	return result, nil
}

func (as *adminService) DeleteCounterService(ctx context.Context, id string) (sql.Result, error) {
	result, err := as.AdminRepository.DeleteCounter(ctx, id)
	if err != nil {
		log.Println("error while adding counter in DB", err)
		return nil, err
	}
	return result, nil
}

func (as *adminService) GetAllCountersService(ctx context.Context, companyCode string) ([]*models.ManageCounter, error) {
	result, err := as.AdminRepository.GetAllCounters(ctx, companyCode)
	if err != nil {
		log.Println("error while getting all counters from DB", err)
		return nil, err
	}
	return result, nil
}
func (as *adminService) GetAllUnActivatedCountersService(ctx context.Context, companyCode string) ([]*models.ManageCounter, error) {
	result, err := as.AdminRepository.GetAllInactiveCounters(ctx, companyCode)
	if err != nil {
		log.Println("error while getting all counters from DB", err)
		return nil, err
	}
	return result, nil
}

func (as *adminService) GetCounterService(ctx context.Context, arg models.GetCounterParams) (*models.ManageCounter, error) {
	result, err := as.AdminRepository.GetCounter(ctx, arg)
	if err != nil {
		log.Println("error while getting counter from DB", err)
		return nil, err
	}
	return result, nil
}
func (as *adminService) UpdateCounterService(ctx context.Context, arg models.UpdateCounterParams) error {
	// arg.ID = uuid.NewString()
	const createdFormat = "2006-01-02 15:04:05"
	// arg.CreatedAt = time.Now().Format(createdFormat)
	arg.UpdatedAt = time.Now().Format(createdFormat)
	err := as.AdminRepository.UpdateCounter(ctx, arg)
	if err != nil {
		log.Println("error while adding counter in DB", err)
		return err
	}
	return nil
}

// ? section service
func (as *adminService) AddSectionService(ctx context.Context, arg models.AddSectionParams) (sql.Result, error) {
	result, err := as.AdminRepository.AddSection(ctx, arg)
	if err != nil {
		log.Println("error while adding section in DB", err)
		return nil, err
	}
	return result, nil
}
func (as *adminService) DeleteSectionService(ctx context.Context, id string) (sql.Result, error) {
	result, err := as.AdminRepository.DeleteSection(ctx, id)
	if err != nil {
		log.Println("error while deleting section in DB", err)
		return nil, err
	}
	return result, nil
}
func (as *adminService) GetAllSectionsService(ctx context.Context, companyCode string) ([]*models.ManageSection, error) {
	result, err := as.AdminRepository.GetAllSections(ctx, companyCode)
	if err != nil {
		log.Println("error while getting all section in DB", err)
		return nil, err
	}
	return result, nil
}
func (as *adminService) GetSectionService(ctx context.Context, arg models.GetSectionParams) (*models.ManageSection, error) {
	result, err := as.AdminRepository.GetSection(ctx, arg)
	if err != nil {
		log.Println("error while getting section in DB", err)
		return nil, err
	}
	return result, nil
}
func (as *adminService) UpdateSectionService(ctx context.Context, arg models.UpdateSectionParams) error {
	err := as.AdminRepository.UpdateSection(ctx, arg)
	if err != nil {
		log.Println("error while updating section in DB", err)
		return err
	}
	return nil
}

// ? Service services section
func (as *adminService) AddService(ctx context.Context, arg models.GetServiceResult) (sql.Result, error) {
	arg.ID = uuid.NewString()
	t := time.Now()
	ts := t.Format("2006-01-02 15:04:05")
	arg.CreatedAt = ts
	arg.UpdatedAt = ts
	result, err := as.AdminRepository.AddService(ctx, arg)
	if err != nil {
		log.Println("error while adding service in DB", err)
		return nil, err
	}
	return result, nil
}

func (as *adminService) DeleteService(ctx context.Context, id string) (sql.Result, error) {
	result, err := as.AdminRepository.DeleteService(ctx, id)
	if err != nil {
		log.Println("error while deleting service in DB", err)
		return nil, err
	}
	return result, nil
}

func (as *adminService) GetAllServices(ctx context.Context, companyCode string) ([]*models.ManageService, error) {
	result, err := as.AdminRepository.GetAllServices(ctx, companyCode)
	if err != nil {
		log.Println("error while getting all services in DB", err)
		return nil, err
	}
	return result, nil
}

func (as *adminService) GetService(ctx context.Context, arg models.GetServiceParams) (*models.ManageService, error) {
	result, err := as.AdminRepository.GetService(ctx, arg)
	if err != nil {
		log.Println("error while getting service in DB", err)
		return nil, err
	}
	return result, nil
}
func (as *adminService) UpdateService(ctx context.Context, arg models.UpdateServiceParams) error {
	err := as.AdminRepository.UpdateService(ctx, arg)
	if err != nil {
		log.Println("error while updating service in DB", err)
		return err
	}
	return nil
}

// ? User services section
func (as *adminService) AddUserService(ctx context.Context, arg models.AddUserParams) (sql.Result, error) {
	ad, err := as.AdminRepository.GetUserByEmail(ctx, arg.Email)
	if err != nil && err != sql.ErrNoRows {
		log.Println("error while getting user by email", err)
		return nil, err
	}
	if ad != nil {
		log.Println("user already exists.")
		return nil, errors.New("error user already esists")
	}
	if err == sql.ErrNoRows {
		arg.ID = uuid.NewString()
		const createdFormat = "2006-01-02 15:04:05"
		arg.CreatedAt = time.Now().Format(createdFormat)
		arg.UpdatedAt = time.Now().Format(createdFormat)
		result, err := as.AdminRepository.AddUser(ctx, arg)
		if err != nil {
			log.Println("error while adding user in DB", err)
			return nil, err
		}
		return result, nil

	}
	return nil, err
}
func (as *adminService) DeleteUserService(ctx context.Context, id string) (sql.Result, error) {
	result, err := as.AdminRepository.DeleteUser(ctx, id)
	if err != nil {
		log.Println("error while deleting user in DB", err)
		return nil, err
	}
	return result, nil
}
func (as *adminService) GetAllUsers(ctx context.Context, arg models.GetAllUsersParams) ([]*models.ManageUser, error) {
	result, err := as.AdminRepository.GetAllUsers(ctx, arg)
	if err != nil {
		log.Println("error while getting all user in DB", err)
		return nil, err
	}
	return result, nil
}
func (as *adminService) GetUser(ctx context.Context, arg models.GetUserParams) (*models.ManageUser, error) {
	result, err := as.AdminRepository.GetUser(ctx, arg)
	if err != nil {
		log.Println("error while getting user in DB", err)
		return nil, err
	}
	return result, nil
}
func (as *adminService) UpdateUserEmailService(ctx context.Context, arg models.UpdateUserEmailParams) error {
	err := as.AdminRepository.UpdateUserEmail(ctx, arg)
	if err != nil {
		log.Println("error while updating user email in DB", err)
		return err
	}
	return nil
}

func (as *adminService) UpdateUserPassword(ctx context.Context, arg models.UpdateUserPasswordParams) error {
	err := as.AdminRepository.UpdateUserPassword(ctx, arg)
	if err != nil {
		log.Println("error while updating password in DB", err)
		return err
	}
	return nil
}

func (as *adminService) UpdateUserService(ctx context.Context, arg models.UpdateUserParams) error {
	const createdFormat = "2006-01-02 15:04:05"
	// arg.CreatedAt = time.Now().Format(createdFormat)
	arg.UpdatedAt = time.Now().Format(createdFormat)
	err := as.AdminRepository.UpdateUser(ctx, arg)
	if err != nil {
		log.Println("error while updating user in DB", err)
		return err
	}
	return nil
}
