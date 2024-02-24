package models

import (
	"context"
	"database/sql"
	"time"
)

//? this interface holds only those functionn which interacts with sql that means Repository is for DB related operations

type UserRepository interface {
	GetUserById(ctx context.Context, id string) (*UserAccount, error)
	GetUserByEmail(ctx context.Context, email string) (*UserAccount, error)
	Create(ctx context.Context, u *UserAccount) error
	// Update(ctx context.Context, user *User) error
	// Delete(ctx context.Context, id string) (*User, error)
	GetTimezone(ctx context.Context, country string) ([]*TimeZones, error)
	GetCountries(ctx context.Context) ([]*Countries, error)
	GetBranchByEmail(ctx context.Context, email string) (*BranchAdmin, error)
	UpdateSystem(ctx context.Context, arg UpdateSystemParams) error
	AddServer(ctx context.Context, arg ServerDetails) error
	// InsetSystem(ctx context.Context, arg UpdateSystemParams) error
	GetServerDetailsByIP(serverIP string) (ServerDetails, error)
}

//* service holds the user related operations which may involave DB or not

type UserService interface {
	SignUp(ctx context.Context, user *UserAccount) error
	SignIn(ctx context.Context, user *UserAccount) error
	ForgotPassword(ctx context.Context, email string) error
	ResetPassword(ctx context.Context, password, token string) error
	// UpdateProfile() error
	GetMasters(ctx context.Context) (*MasterResult, error)
	GetTimezonesByCountry(ctx context.Context, country string) ([]*TimeZones, error)
	GetUserByID(ctx context.Context, country string) (*UserAccount, error)
	BranchLogin(ctx context.Context, user *UserAccount) error
	GetServerByID(ctx context.Context, id string) (ServerDetails, error)
}

type TokenRepository interface {
	SetRefreshToken(ctx context.Context, userId, tokenId string, expiresIn time.Duration) error
	DeleteRefrashToken(ctx context.Context, userId, tokenId string) error
}

type TokenService interface {
	GenerateIDToken(u *UserAccount)
	GeneratePairFromUser(ctx context.Context, user *UserAccount, prevToken string) (*TokenPair, error)
	ValidateIDToken(tokenString string) (*UserAccount, error)
	ValidateRefreshToken(tokenString string) (*RefreshToken, error)
}

type MailService interface {
	SendMail(to, subject string) error
	ParseTemplate(templateFileName string, data interface{}) error
	// UpdateProfile() error
}

type AdminService interface {
	AddbranchService(ctx context.Context, arg AddbranchParams) (sql.Result, error)
	DeleteBranchService(ctx context.Context, code string) (sql.Result, error)
	GetAllBranchesService(ctx context.Context, companyCode string) ([]*ManageBranch, error)
	GetBranchService(ctx context.Context, arg GetBranchParams) (*ManageBranch, error)
	UpdateBranchService(ctx context.Context, arg AddbranchParams) error
	//? counter methods
	AddCounterService(ctx context.Context, arg AddCounterParams) (sql.Result, error)
	DeleteCounterService(ctx context.Context, id string) (sql.Result, error)
	GetAllCountersService(ctx context.Context, companyCode string) ([]*ManageCounter, error)
	GetCounterService(ctx context.Context, arg GetCounterParams) (*ManageCounter, error)
	UpdateCounterService(ctx context.Context, arg UpdateCounterParams) error
	GetAllUnActivatedCountersService(ctx context.Context, companyCode string) ([]*ManageCounter, error)
	//? section methods
	AddSectionService(ctx context.Context, arg AddSectionParams) (sql.Result, error)
	DeleteSectionService(ctx context.Context, id string) (sql.Result, error)
	GetAllSectionsService(ctx context.Context, companyCode string) ([]*ManageSection, error)
	GetSectionService(ctx context.Context, arg GetSectionParams) (*ManageSection, error)
	UpdateSectionService(ctx context.Context, arg UpdateSectionParams) error
	//? service methods
	AddService(ctx context.Context, arg GetServiceResult) (sql.Result, error)
	DeleteService(ctx context.Context, id string) (sql.Result, error)
	GetAllServices(ctx context.Context, companyCode string) ([]*ManageService, error)
	GetService(ctx context.Context, arg GetServiceParams) (*ManageService, error)
	UpdateService(ctx context.Context, arg UpdateServiceParams) error
	//? user service
	AddUserService(ctx context.Context, arg AddUserParams) (sql.Result, error)
	DeleteUserService(ctx context.Context, id string) (sql.Result, error)
	GetAllUsers(ctx context.Context, arg GetAllUsersParams) ([]*ManageUser, error)
	GetUser(ctx context.Context, arg GetUserParams) (*ManageUser, error)
	UpdateUserEmailService(ctx context.Context, arg UpdateUserEmailParams) error
	UpdateUserPassword(ctx context.Context, arg UpdateUserPasswordParams) error
	UpdateUserService(ctx context.Context, arg UpdateUserParams) error
	SetActiveService(ip string, active bool) error
}

type AdminRepository interface {
	Addbranch(ctx context.Context, arg AddbranchParams) (sql.Result, error)
	DeleteBranch(ctx context.Context, code string) (sql.Result, error)
	GetAllBranches(ctx context.Context, companyCode string) ([]*ManageBranch, error)
	GetBranch(ctx context.Context, arg GetBranchParams) (*ManageBranch, error)
	UpdateBranch(ctx context.Context, arg AddbranchParams) error
	GetBranchByEmail(ctx context.Context, email string) (*BranchAdmin, error)
	//? counter repo methods
	AddCounter(ctx context.Context, arg AddCounterParams) (sql.Result, error)
	DeleteCounter(ctx context.Context, id string) (sql.Result, error)
	GetAllCounters(ctx context.Context, companyCode string) ([]*ManageCounter, error)
	GetCounter(ctx context.Context, arg GetCounterParams) (*ManageCounter, error)
	UpdateCounter(ctx context.Context, arg UpdateCounterParams) error
	GetAllInactiveCounters(ctx context.Context, companyCode string) ([]*ManageCounter, error)
	//? section repo methods
	AddSection(ctx context.Context, arg AddSectionParams) (sql.Result, error)
	DeleteSection(ctx context.Context, id string) (sql.Result, error)
	GetAllSections(ctx context.Context, companyCode string) ([]*ManageSection, error)
	GetSection(ctx context.Context, arg GetSectionParams) (*ManageSection, error)
	UpdateSection(ctx context.Context, arg UpdateSectionParams) error
	//? service repo methods
	AddService(ctx context.Context, arg GetServiceResult) (sql.Result, error)
	DeleteService(ctx context.Context, id string) (sql.Result, error)
	GetAllServices(ctx context.Context, companyCode string) ([]*ManageService, error)
	GetService(ctx context.Context, arg GetServiceParams) (*ManageService, error)
	UpdateService(ctx context.Context, arg UpdateServiceParams) error
	//? user repo methods
	GetUserByEmail(ctx context.Context, email string) (*UserAccount, error)
	AddUser(ctx context.Context, arg AddUserParams) (sql.Result, error)
	DeleteUser(ctx context.Context, id string) (sql.Result, error)
	GetAllUsers(ctx context.Context, arg GetAllUsersParams) ([]*ManageUser, error)
	GetUser(ctx context.Context, arg GetUserParams) (*ManageUser, error)
	UpdateUserEmail(ctx context.Context, arg UpdateUserEmailParams) error
	UpdateUserPassword(ctx context.Context, arg UpdateUserPasswordParams) error
	UpdateUser(ctx context.Context, arg UpdateUserParams) error
	ActivateSystemDAO(ip string, active bool) error
}

type TicketRepository interface {
	AddTicket(ctx context.Context, arg AddTicketParams) (sql.Result, error)
	// DeleteBranch(ctx context.Context, code string) (sql.Result, error)
	DeleteTicket(ctx context.Context, id string) (sql.Result, error)
	GetAllTicketsForDay(ctx context.Context, arg GetAllTicketsForDayParams) ([]*Ticket, error)
	GetTicket(ctx context.Context, arg GetTicketParams) (*Ticket, error)
	GetTicketsByBranch(ctx context.Context, arg GetTicketsByBranchParams) ([]*Ticket, error)
	GetTicketsByStatus(ctx context.Context, arg GetTicketsByStatusParams) ([]*Ticket, error)
	GetTicketsByTransfer(ctx context.Context, arg GetTicketsByTransferParams) ([]*Ticket, error)
	GetTicketsByUser(ctx context.Context, arg GetTicketsByUserParams) ([]*Ticket, error)
	UpdateTicketStartTime(ctx context.Context, arg UpdateTicketStartTimeParams) error
	UpdateTicketStatus(ctx context.Context, arg UpdateTicketStatusParams) error
	UpdateTicketTransferedTo(ctx context.Context, arg UpdateTicketTransferedToParams) error
	UpdateTicketUser(ctx context.Context, arg UpdateTicketUserParams) error
	GetLastTicketNumber(ctx context.Context) (TicketNumber, error)
	UpdateTicketEndTime(ctx context.Context, arg UpdateTicketEndTimeParams) error
	GetTicketToProcess(ctx context.Context, arg GetTicketParams) (*Ticket, error)
}

type TicketService interface {
	AddTicketService(ctx context.Context, arg AddTicketParams) (sql.Result, error)
	// DeleteBranchService(ctx context.Context, code string) (sql.Result, error)
	DeleteTicketService(ctx context.Context, id string) (sql.Result, error)
	GetAllTicketsForDayService(ctx context.Context, arg GetAllTicketsForDayParams) ([]*Ticket, error)
	GetTicketService(ctx context.Context, arg GetTicketParams) (*Ticket, error)
	GetTicketsByBranchService(ctx context.Context, arg GetTicketsByBranchParams) ([]*Ticket, error)
	GetTicketsByStatusService(ctx context.Context, arg GetTicketsByStatusParams) ([]*Ticket, error)
	GetTicketsByTransferService(ctx context.Context, arg GetTicketsByTransferParams) ([]*Ticket, error)
	GetTicketsByUserService(ctx context.Context, arg GetTicketsByUserParams) ([]*Ticket, error)
	UpdateTicketStartTimeService(ctx context.Context, arg UpdateTicketStartTimeParams) error
	UpdateTicketStatusService(ctx context.Context, arg UpdateTicketStatusParams) error
	UpdateTicketTransferedToService(ctx context.Context, arg UpdateTicketTransferedToParams) error
	UpdateTicketUserService(ctx context.Context, arg UpdateTicketUserParams) error
	GetLastTicketNumberService(ctx context.Context) (TicketNumber, error)
	UpdateTicketEndTimeService(ctx context.Context, arg UpdateTicketEndTimeParams) error
	GetTicketToProcessService(ctx context.Context, arg GetTicketParams) (*Ticket, error)
	GetWaitingTicketsService(ctx context.Context) ([]*Ticket, error)
	GetNoShowTicketsService(ctx context.Context, id string) ([]*Ticket, error)
	GetFinishedTicketsService(ctx context.Context, id string) ([]*Ticket, error)
	GetLastCalledService(ctx context.Context) (*Ticket, error)
	GetLastStartedTicket(ctx context.Context, id string) (*Ticket, error)
	GetSystemService(ctx context.Context) (*ServerDetails, error)
	UpdateSystemIp(ctx context.Context, client UpdateSystemParams) (sql.Result, error)
	GetAllSystemService(ctx context.Context) ([]*UpdateSystemParams, error)
}

type LicenseService interface {
	UpdateLicenseService(lsc LicensePayload) (bool, error)
	CheckCounterActivationService(ctr ActiveCounter) (bool, error)
	UpdateCounterDetails(ctr ActiveCounter) (bool, error)
	ActivateCounterService(ctr ActiveCounter) (bool, error)
	CheckLicensePresent() (bool, error)
	CheckAndGetActiveCounter(ActiveCounter) (ManageCounter, error)
	AuthCounterUserService(ctr AuthCounterUser) (*ManageUser, error)
	UpdateCounterUserService(arg string) error
}
type ConfigService interface {
	AddVideoService(ctx context.Context, v Video) (sql.Result, error)
	GetAllVideoService(ctx context.Context, id string) ([]VideoRes, error)
	GetVideoService(ctx context.Context, id string) (VideoRes, error)
	AddSchedularConfigService(ctx context.Context, v VideoSchedular) (sql.Result, error)
	GetSchedularConfService(ctx context.Context, id string) (VideoSchedularRes, error)
	GetAllSchedularConfService(ctx context.Context, id string) ([]VideoSchedularRes, error)
	AddAudioConfigService(a AudioConfig) error
	GetAudionConfigService(id string) (AudioConfigRes, error)
	UpdateAudioConfigService(a AudioConfig) error
	GetAllAudioConfigService(id string) ([]AudioConfigRes, error)
	AddDSConfigService(a DSConfig) error
	GetDSConfigService(id string) (DSConfigRes, error)
	UpdateDsConfigService(a DSConfig) error
	AddAnnouncementService(a AnnouncementConf) error
	GetAnnouncement(a string) (AnnouncementConfRes, error)
	UpdateAnnouncement(a string, speed int, text string) error
	AddTicketConf(tc TicketConf) error
	GetTicketConf(a string) (TicketConfRes, error)
	UpdateTicketConf(code string, tc TicketConf) error
	GetAllConfService(id string) (AllConfig, error)
	GetAllAnnouncement(a string) ([]AnnouncementConfRes, error)
	SelectAnnouncementToDisplay(id, code string) error
	GetServerByIDService(id string) (ServerDetails, error)
}
