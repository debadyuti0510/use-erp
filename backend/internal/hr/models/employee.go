package models

import (
	"time"
)

const (
	NationalIDTypePassport      = "passport"
	NationalIDTypeNationalID    = "national_id"
	NationalIDTypeSSN           = "ssn"
	NationalIDTypeAadhaar       = "aadhaar"
	NationalIDTypeNIN           = "nin"
	NationalIDTypeSIN           = "sin"
	NationalIDTypeDNI           = "dni"
	NationalIDTypeRCN           = "rcn"
	NationalIDTypeDriverLicense = "driver_license"
	NationalIDTypeWorkPermit    = "work_permit"
	NationalIDTypeTaxID         = "tax_id"
	NationalIDTypeOther         = "other"
)

var ValidNationalIDTypes = map[string]struct{}{
	NationalIDTypePassport:      {},
	NationalIDTypeNationalID:    {},
	NationalIDTypeSSN:           {},
	NationalIDTypeAadhaar:       {},
	NationalIDTypeNIN:           {},
	NationalIDTypeSIN:           {},
	NationalIDTypeDNI:           {},
	NationalIDTypeRCN:           {},
	NationalIDTypeDriverLicense: {},
	NationalIDTypeWorkPermit:    {},
	NationalIDTypeTaxID:         {},
	NationalIDTypeOther:         {},
}

type Employee struct {
	ID                uint      `gorm:"primaryKey" json:"id"`
	FirstName         string    `json:"first_name"`
	LastName          string    `json:"last_name"`
	Gender            string    `json:"gender"`
	DOB               time.Time `json:"dob"`
	Nationality       string    `json:"nationality"`
	NationalID        string    `json:"national_id"`
	NationalIDType    string    `json:"national_id_type"`
	NationalIDCountry string    `json:"national_id_country"`

	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`

	JobTitleID     uint   `json:"job_title_id"`
	DepartmentID   uint   `json:"department_id"`
	ManagerID      *uint  `json:"manager_id"` // Self-reference
	WorkingCountry string `json:"working_country"`
	EmployeeCode   string `gorm:"uniqueIndex" json:"employee_code"`

	DateOfJoining  time.Time  `json:"date_of_joining"`
	DateOfExit     *time.Time `json:"date_of_exit,omitempty"`
	EmploymentType string     `json:"employment_type"`
	IsActive       bool       `json:"is_active"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// Relations
	Manager    *Employee  `gorm:"foreignKey:ManagerID" json:"manager,omitempty"`
	Department Department `gorm:"foreignKey:DepartmentID" json:"department,omitempty"`
	JobTitle   JobTitle   `gorm:"foreignKey:JobTitleID" json:"job_title,omitempty"`
}
