package models

import (
	"github.com/google/uuid"
)

type User struct {
	UserID       uuid.UUID `json:"id"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Email        string    `json:"email"`
	PhoneNumber  string    `json:"phone_number"`
	HireDate     string    `json:"hire_date"`
	JobId        int       `json:"job_id"`
	Salary       int       `json:"salary"`
	ManagerId    int       `json:"manager_id"`
	DepartmentId int       `json:"department_id"`
}

type UserAccount struct {
	UserID      uuid.UUID `json:"id"`
	Email       string    `json:"email"`
	Password    string    `json:"password" `
	FirstName   string    `json:"firstname"`
	LastName    string    `json:"lastname"`
	Company     string    `json:"company"`
	CompanyCode string    `json:"company_code"`
	Country     string    `json:"country"`
	Timezone    string    `json:"timezone"`
	UpdatedBy   string    `json:"updated_by"`
	CreatedAt   string    `json:"created_at"`
	UpdatedAt   string    `json:"updated_at"`
}
type BranchAdmin struct {
	Email       string `json:"email"`
	Password    string `json:"password" `
	BranchName  string `json:"branch_name"`
	BranchCode  string `json:"branch_code"`
	Company     string `json:"company_name"`
	CompanyCode string `json:"company_code"`
	UpdatedBy   string `json:"updated_by"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	Isdeleted   bool   `json:"isdeleted"`
}

type LicensePayload struct {
	Firm                     string          `json:"firm,omitempty" bson:"firm,omitempty"`
	ExpirationInDays         string          `json:"expiration_in_days"`
	AllowedBranches          int             `json:"allowedBranches,omitempty" bson:"allowedBranches,omitempty"`
	AllowedCountersPerBranch int             `json:"allowedCountersPerBranch,omitempty" bson:"allowedCountersPerFirm,omitempty"`
	ActivatedBranches        []ActiveBranch  `json:"activatedBranches,omitempty" bson:"activatedBranches,omitempty"`
	ActiveCounterPerBranch   []ActiveCounter `json:"activeCounterPerBranch,omitempty" bson:"activeCounterPerBranch,omitempty"`
}
type ActiveBranch struct {
	BranchId   string `json:"branchId,omitempty" bson:"branchId,omitempty"`
	BranchName string `json:"branchName,omitempty" bson:"branchName,omitempty"`
}
type ActiveCounter struct {
	BranchId    string `json:"branchId,omitempty" bson:"branchId,omitempty"`
	CounterId   string `json:"counterId,omitempty" bson:"counterId,omitempty"`
	CounterName string `json:"counterName,omitempty" bson:"counterName,omitempty"`
	CpuId       string `json:"cpuId,omitempty" bson:"cpuId,omitempty"`
	HardDiskId  string `json:"hardDiskId,omitempty" bson:"hardDiskId,omitempty"`
	BranchName  string `json:"branchName,omitempty" bson:"branchName,omitempty"`
	Id          string `json:"id,omitempty"`
}
