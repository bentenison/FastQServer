package models

type AddbranchParams struct {
	ID            string `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	Code          string `json:"code,omitempty"`
	CompanyCode   string `json:"company_code,omitempty"`
	License       string `json:"license,omitempty"`
	Email         string `json:"email,omitempty"`
	Password      string `json:"password,omitempty"`
	Services      string `json:"services,omitempty"`
	Address       string `json:"address,omitempty"`
	Phone         string `json:"phone,omitempty"`
	LicenseKey    string `json:"license_key,omitempty"`
	CheckInUrl    string `json:"check_in_url,omitempty"`
	TicketPageUrl string `json:"ticket_page_url,omitempty"`
	DisplayUrl    string `json:"display_url,omitempty"`
	CreatedAt     string `json:"created_at,omitempty"`
	UpdatedAt     string `json:"updated_at,omitempty"`
	UpdatedBy     string `json:"updated_by,omitempty"`
}

type GetBranchParams struct {
	Code        string `json:"code,omitempty"`
	CompanyCode string `json:"company_code,omitempty"`
	Email       string `json:"email,omitempty"`
}

type UpdateBranchAddressParams struct {
	Address   string `json:"address,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	UpdatedBy string `json:"updated_by,omitempty"`
	Code      string `json:"code,omitempty"`
}

type UpdateBranchLicenseParams struct {
	License   string `json:"license,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	UpdatedBy string `json:"updated_by,omitempty"`
	Code      string `json:"code,omitempty"`
}

type UpdateBranchLicenseKeyParams struct {
	LicenseKey string `json:"license_key,omitempty"`
	UpdatedAt  string `json:"updated_at,omitempty"`
	UpdatedBy  string `json:"updated_by,omitempty"`
	Code       string `json:"code,omitempty"`
}

type UpdateBranchNameParams struct {
	Name      string `json:"name,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	UpdatedBy string `json:"updated_by,omitempty"`
	Code      string `json:"code,omitempty"`
}

type UpdateBranchPhoneParams struct {
	Phone     string `json:"phone,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	UpdatedBy string `json:"updated_by,omitempty"`
	Code      string `json:"code,omitempty"`
}

type UpdateBranchServicesParams struct {
	Services  string `json:"services,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	UpdatedBy string `json:"updated_by,omitempty"`
	Code      string `json:"code,omitempty"`
}

type AddCounterParams struct {
	ID               string `json:"id,omitempty"`
	Section          string `json:"section,omitempty"`
	SubSection       string `json:"sub_section,omitempty"`
	UserID           string `json:"user_id,omitempty"`
	CounterName      string `json:"counter_name,omitempty"`
	CounterNumber    string `json:"counter_number,omitempty"`
	OverrideFifo     bool   `json:"override_fifo,omitempty"`
	TransferQ        bool   `json:"transfer_q,omitempty"`
	AsssignUser      bool   `json:"asssign_user,omitempty"`
	AssignService    bool   `json:"assign_service,omitempty"`
	TransferPriority bool   `json:"transfer_priority,omitempty"`
	CancelQ          bool   `json:"cancel_q,omitempty"`
	TransferFinish   bool   `json:"transfer_finish,omitempty"`
	MergeSection     bool   `json:"merge_section,omitempty"`
	EndQ             bool   `json:"end_q,omitempty"`
	BranchCode       string `json:"branch_code,omitempty"`
	BranchName       string `json:"branch_name,omitempty"`
	CompanyCode      string `json:"company_code,omitempty"`
	CompanyName      string `json:"company_name,omitempty"`
	CreatedAt        string `json:"created_at,omitempty"`
	CreatedBy        string `json:"created_by,omitempty"`
	UpdatedAt        string `json:"updated_at,omitempty"`
	UpdatedBy        string `json:"updated_by,omitempty"`
}

type GetCounterParams struct {
	ID         string `json:"id,omitempty"`
	BranchCode string `json:"branch_code,omitempty"`
}

type UpdateCounterNameParams struct {
	CounterName string `json:"counter_name,omitempty"`
	UpdatedAt   string `json:"updated_at,omitempty"`
	UpdatedBy   string `json:"updated_by,omitempty"`
	ID          string `json:"id,omitempty"`
}

type UpdateCounterNumberParams struct {
	CounterNumber string `json:"counter_number,omitempty"`
	UpdatedAt     string `json:"updated_at,omitempty"`
	UpdatedBy     string `json:"updated_by,omitempty"`
	ID            string `json:"id,omitempty"`
}

type UpdateCounterSectionParams struct {
	Section   string `json:"section,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	UpdatedBy string `json:"updated_by,omitempty"`
	ID        string `json:"id,omitempty"`
}

type UpdateCounterSettingsParams struct {
	OverrideFifo     bool   `json:"override_fifo,omitempty"`
	TransferQ        bool   `json:"transfer_q,omitempty"`
	AsssignUser      bool   `json:"asssign_user,omitempty"`
	AssignService    bool   `json:"assign_service,omitempty"`
	TransferPriority bool   `json:"transfer_priority,omitempty"`
	TransferFinish   bool   `json:"transfer_finish,omitempty"`
	CancelQ          bool   `json:"cancel_q,omitempty"`
	MergeSection     bool   `json:"merge_section,omitempty"`
	EndQ             bool   `json:"end_q,omitempty"`
	UpdatedAt        string `json:"updated_at,omitempty"`
	UpdatedBy        string `json:"updated_by,omitempty"`
	ID               string `json:"id,omitempty"`
}

type UpdateCounterSubSectionParams struct {
	SubSection string `json:"sub_section,omitempty"`
	UpdatedAt  string `json:"updated_at,omitempty"`
	UpdatedBy  string `json:"updated_by,omitempty"`
	ID         string `json:"id,omitempty"`
}

type AddSectionParams struct {
	ID           string `json:"id,omitempty"`
	BenchWait    int    `json:"bench_wait,omitempty"`
	BenchProcess int    `json:"bench_process,omitempty"`
	Description  string `json:"description,omitempty"`
	BranchCode   string `json:"branch_code,omitempty"`
	BrachName    string `json:"brach_name,omitempty"`
	CompanyCode  string `json:"company_code,omitempty"`
	CompanyName  string `json:"company_name,omitempty"`
	CreatedAt    string `json:"created_at,omitempty"`
	CreatedBy    string `json:"created_by,omitempty"`
	UpdatedAt    string `json:"updated_at,omitempty"`
	UpdatedBy    string `json:"updated_by,omitempty"`
}

type GetSectionParams struct {
	ID          string `json:"id,omitempty"`
	CompanyCode string `json:"company_code,omitempty"`
}

type UpdateSectionDescriptionParams struct {
	Description string `json:"description,omitempty"`
	UpdatedAt   string `json:"updated_at,omitempty"`
	UpdatedBy   string `json:"updated_by,omitempty"`
	ID          string `json:"id,omitempty"`
}

type UpdateSectionNameParams struct {
	Name      string `json:"name,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	UpdatedBy string `json:"updated_by,omitempty"`
	ID        string `json:"id,omitempty"`
}

type UpdateSectionProcessParams struct {
	BenchProcess int    `json:"bench_process,omitempty"`
	UpdatedAt    string `json:"updated_at,omitempty"`
	UpdatedBy    string `json:"updated_by,omitempty"`
	ID           string `json:"id,omitempty"`
}

type UpdateSectionWaitParams struct {
	BenchWait int    `json:"bench_wait,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	UpdatedBy string `json:"updated_by,omitempty"`
	ID        string `json:"id,omitempty"`
}

type AddServiceParams struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Code         string `json:"code"`
	Prefix       string `json:"prefix"`
	NumberStarts int    `json:"number_starts"`
	NumberEnds   int    `json:"number_ends"`
	Hide         bool   `json:"hide"`
	ShowDisplay  bool   `json:"show_display"`
	Description  string `json:"description"`
	StartTime    string `json:"start_time"`
	EndTime      string `json:"end_time"`
	DefaultTime  int    `json:"default_time"`
	Workflow     string `json:"workflow"`
	BranchCode   string `json:"branch_code"`
	BranchName   string `json:"branch_name"`
	CompanyCode  string `json:"company_code"`
	CompanyName  string `json:"company_name"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
	UpdatedBy    string `json:"updated_by"`
}
type GetServiceResult struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Code         string `json:"code"`
	Prefix       string `json:"prefix"`
	NumberStarts int    `json:"number_starts"`
	NumberEnds   int    `json:"number_ends"`
	Hide         bool   `json:"hide"`
	ShowDisplay  bool   `json:"show_display"`
	Description  string `json:"description"`
	StartTime    string `json:"start_time"`
	EndTime      string `json:"end_time"`
	DefaultTime  int    `json:"default_time"`
	Workflow     string `json:"workflow"`
	BranchCode   string `json:"branch_code"`
	BranchName   string `json:"branch_name"`
	CompanyCode  string `json:"company_code"`
	CompanyName  string `json:"company_name"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
	UpdatedBy    string `json:"updated_by"`
}

type GetServiceParams struct {
	ID          string `json:"id,omitempty"`
	CompanyCode string `json:"company_code,omitempty"`
}

type UpdateServiceDefaultTimeParams struct {
	DefaultTime int    `json:"default_time,omitempty"`
	UpdatedAt   string `json:"updated_at,omitempty"`
	UpdatedBy   string `json:"updated_by,omitempty"`
	ID          string `json:"id,omitempty"`
}

type UpdateServiceEndTimeParams struct {
	EndTime   string `json:"end_time,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	UpdatedBy string `json:"updated_by,omitempty"`
	ID        string `json:"id,omitempty"`
}

type UpdateServiceNameParams struct {
	Name      string `json:"name,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	UpdatedBy string `json:"updated_by,omitempty"`
	ID        string `json:"id,omitempty"`
}

type UpdateServiceNumberEndParams struct {
	NumberEnds int    `json:"number_ends,omitempty"`
	UpdatedAt  string `json:"updated_at,omitempty"`
	UpdatedBy  string `json:"updated_by,omitempty"`
	ID         string `json:"id,omitempty"`
}

type UpdateServiceNumberStartParams struct {
	NumberStarts int    `json:"number_starts,omitempty"`
	UpdatedAt    string `json:"updated_at,omitempty"`
	UpdatedBy    string `json:"updated_by,omitempty"`
	ID           string `json:"id,omitempty"`
}

type UpdateServicePrefixParams struct {
	Prefix    string `json:"prefix,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	UpdatedBy string `json:"updated_by,omitempty"`
	ID        string `json:"id,omitempty"`
}
type UpdateServiceStartTimeParams struct {
	StartTime string `json:"start_time,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	UpdatedBy string `json:"updated_by,omitempty"`
	Code      string `json:"code,omitempty"`
}

type AddUserParams struct {
	ID          string `json:"id,omitempty"`
	Email       string `json:"email,omitempty"`
	Firstname   string `json:"firstname,omitempty"`
	Lastname    string `json:"lastname,omitempty"`
	UserType    string `json:"user_type,omitempty"`
	Suspended   bool   `json:"suspended,omitempty"`
	Title       string `json:"title,omitempty"`
	Password    string `json:"password,omitempty"`
	ImageUrl    string `json:"image_url,omitempty"`
	BranchCode  string `json:"branch_code,omitempty"`
	BranchName  string `json:"branch_name,omitempty"`
	CompanyCode string `json:"company_code,omitempty"`
	CompanyName string `json:"company_name,omitempty"`
	CreatedAt   string `json:"created_at,omitempty"`
	CreatedBy   string `json:"created_by,omitempty"`
	UpdatedAt   string `json:"updated_at,omitempty"`
	UpdatedBy   string `json:"updated_by,omitempty"`
}

type GetAllUsersParams struct {
	CompanyCode string `json:"company_code,omitempty"`
	BranchCode  string `json:"branch_code,omitempty"`
	CounterId   string `json:"counter_id,omitempty"`
}

type GetUserParams struct {
	ID          string `json:"id,omitempty"`
	CompanyCode string `json:"company_code,omitempty"`
	BranchCode  string `json:"branch_code,omitempty"`
}
type AuthCounterUser struct {
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
	CounterId string `json:"counter_id"`
}

type UpdateUserEmailParams struct {
	Email     string `json:"email,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	UpdatedBy string `json:"updated_by,omitempty"`
	ID        string `json:"id,omitempty"`
}

type UpdateUserNamesParams struct {
	Firstname string `json:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	UpdatedBy string `json:"updated_by,omitempty"`
	ID        string `json:"id,omitempty"`
}

type UpdateUserPasswordParams struct {
	Password  string `json:"password,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	UpdatedBy string `json:"updated_by,omitempty"`
	ID        string `json:"id,omitempty"`
}

type UpdateUserSuspentionParams struct {
	Suspended bool   `json:"suspended,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	UpdatedBy string `json:"updated_by,omitempty"`
	ID        string `json:"id,omitempty"`
}

type UpdateUserTypeParams struct {
	UserType  string `json:"user_type,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	UpdatedBy string `json:"updated_by,omitempty"`
	ID        string `json:"id,omitempty"`
}

type AddTicketParams struct {
	ID               string `json:"id,omitempty"`
	Service          string `json:"service,omitempty"`
	CounterID        string `json:"counter_id,omitempty"`
	TicketStatus     string `json:"ticket_status,omitempty"`
	TransferedTo     string `json:"transfered_to,omitempty"`
	TransferedBy     string `json:"transfered_by,omitempty"`
	TicketNumber     string `json:"ticket_number,omitempty"`
	StartedServingAt string `json:"started_serving_at,omitempty"`
	EndServingAt     string `json:"end_serving_at,omitempty"`
	TicketName       string `json:"ticket_name,omitempty"`
	BranchCode       string `json:"branch_code,omitempty"`
	BranchName       string `json:"branch_name,omitempty"`
	CompanyCode      string `json:"company_code,omitempty"`
	CompanyName      string `json:"company_name,omitempty"`
	CreatedAt        string `json:"created_at,omitempty"`
	UpdatedAt        string `json:"updated_at,omitempty"`
	UpdatedBy        string `json:"updated_by,omitempty"`
}

type GetAllTicketsForDayParams struct {
	CompanyCode string `json:"company_code,omitempty"`
	BranchCode  string `json:"branch_code,omitempty"`
	CreatedAt   string `json:"created_at,omitempty"`
}

type GetTicketParams struct {
	ID          string `json:"id,omitempty"`
	CompanyCode string `json:"company_code,omitempty"`
	BranchCode  string `json:"branch_code,omitempty"`
	CounterId   string `json:"counter_id,omitempty"`
}

type GetTicketsByBranchParams struct {
	CompanyCode string `json:"company_code,omitempty"`
	BranchCode  string `json:"branch_code,omitempty"`
}

type GetTicketsByStatusParams struct {
	CompanyCode  string `json:"company_code,omitempty"`
	BranchCode   string `json:"branch_code,omitempty"`
	TicketStatus string `json:"ticket_status,omitempty"`
}

type GetTicketsByTransferParams struct {
	CompanyCode  string `json:"company_code,omitempty"`
	BranchCode   string `json:"branch_code,omitempty"`
	TransferedTo string `json:"transfered_to,omitempty"`
}
type GetTicketsByUserParams struct {
	CompanyCode string `json:"company_code,omitempty"`
	BranchCode  string `json:"branch_code,omitempty"`
	CounterID   string `json:"counter_id,omitempty"`
}
type UpdateTicketStartTimeParams struct {
	StartedServingAt string `json:"started_serving_at,omitempty"`
	UpdatedAt        string `json:"updated_at,omitempty"`
	UpdatedBy        string `json:"updated_by,omitempty"`
	ID               string `json:"id,omitempty"`
	CounterId        string `json:"counterId,omitempty"`
	UserId           string `json:"userId,omitempty"`
}
type UpdateTicketEndTimeParams struct {
	EndServingAt string `json:"end_serving_at,omitempty"`
	UpdatedAt    string `json:"updated_at,omitempty"`
	UpdatedBy    string `json:"updated_by,omitempty"`
	ID           string `json:"id,omitempty"`
}

type UpdateTicketStatusParams struct {
	TicketStatus string `json:"ticket_status,omitempty"`
	UpdatedAt    string `json:"updated_at,omitempty"`
	UpdatedBy    string `json:"updated_by,omitempty"`
	ID           string `json:"id,omitempty"`
}

type UpdateTicketTransferedToParams struct {
	TransferedTo string `json:"transfered_to,omitempty"`
	UpdatedAt    string `json:"updated_at,omitempty"`
	UpdatedBy    string `json:"updated_by,omitempty"`
	ID           string `json:"id,omitempty"`
}

type UpdateTicketUserParams struct {
	CounterID string `json:"counter_id,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	UpdatedBy string `json:"updated_by,omitempty"`
	ID        string `json:"id,omitempty"`
}

type UpdateCounterParams struct {
	Section          string `json:"section,omitempty"`
	SubSection       string `json:"sub_section,omitempty"`
	UserID           string `json:"user_id,omitempty"`
	CounterNumber    string `json:"counter_number,omitempty"`
	CounterName      string `json:"counter_name,omitempty"`
	OverrideFifo     bool   `json:"override_fifo,omitempty"`
	TransferQ        bool   `json:"transfer_q,omitempty"`
	AsssignUser      bool   `json:"asssign_user,omitempty"`
	AssignService    bool   `json:"assign_service,omitempty"`
	TransferPriority bool   `json:"transfer_priority,omitempty"`
	TransferFinish   bool   `json:"transfer_finish,omitempty"`
	CancelQ          bool   `json:"cancel_q,omitempty"`
	MergeSection     bool   `json:"merge_section,omitempty"`
	EndQ             bool   `json:"end_q,omitempty"`
	UpdatedAt        string `json:"updated_at,omitempty"`
	UpdatedBy        string `json:"updated_by,omitempty"`
	ID               string `json:"id,omitempty"`
}

type UpdateSectionParams struct {
	Name         string `json:"name,omitempty"`
	BenchWait    int    `json:"bench_wait,omitempty"`
	BenchProcess int    `json:"bench_process,omitempty"`
	Description  string `json:"description,omitempty"`
	UpdatedAt    string `json:"updated_at,omitempty"`
	UpdatedBy    string `json:"updated_by,omitempty"`
	ID           string `json:"id,omitempty"`
}

type UpdateServiceParams struct {
	Name         string `json:"name,omitempty"`
	Code         string `json:"code,omitempty"`
	Prefix       string `json:"prefix,omitempty"`
	NumberStarts int    `json:"number_starts,omitempty"`
	NumberEnds   int    `json:"number_ends,omitempty"`
	Hide         bool   `json:"hide,omitempty"`
	ShowDisplay  bool   `json:"show_display,omitempty"`
	StartTime    string `json:"start_time,omitempty"`
	EndTime      string `json:"end_time,omitempty"`
	DefaultTime  int    `json:"default_time,omitempty"`
	Workflow     string `json:"workflow,omitempty"`
	NumberEnds_2 int    `json:"number_ends___2,omitempty"`
	UpdatedAt    string `json:"updated_at,omitempty"`
	UpdatedBy    string `json:"updated_by,omitempty"`
	ID           string `json:"id,omitempty"`
}

type UpdateUserParams struct {
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
	Firstname string `json:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
	Suspended bool   `json:"suspended,omitempty"`
	ImageUrl  string `json:"image_url,omitempty"`
	Title     string `json:"title,omitempty"`
	UserType  string `json:"user_type,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	UpdatedBy string `json:"updated_by,omitempty"`
	ID        string `json:"id,omitempty"`
}
type UpdateSystemParams struct {
	ServerIP     string `json:"server_ip,omitempty"`
	ServerCPU    string `json:"server_cpu,omitempty"`
	ServerDiskId string `json:"server_disk_id,omitempty"`
	ClientIP     string `json:"client_ip,omitempty"`
	ClientCPU    string `json:"client_cpu,omitempty"`
	ClientDiskId string `json:"client_disk_id,omitempty"`
	IsDeleted    bool   `json:"is_deleted,omitempty"`
	IsActive     bool   `json:"is_active,omitempty"`
	CounterName  string `json:"counter_name,omitempty"`
	CounterId    string `json:"counter_id,omitempty"`
}
type CounterActivation struct {
	ClientIP     string `json:"client_ip,omitempty"`
	ClientCPU    string `json:"client_cpu,omitempty"`
	ClientDiskId string `json:"client_disk_id,omitempty"`
	Id           string `json:"id,omitempty"`
}
type ServerDetails struct {
	ServerIP     string `json:"server_ip,omitempty"`
	ServerCPU    string `json:"server_cpu,omitempty"`
	ServerDiskId string `json:"server_disk_id,omitempty"`
	Id           string `json:"id,omitempty"`
	Type         string `json:"type,omitempty"`
}
