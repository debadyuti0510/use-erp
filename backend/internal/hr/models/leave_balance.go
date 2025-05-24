package models

import "time"

type LeaveBalance struct {
	ID          uint       `gorm:"primaryKey" json:"id"`
	EmployeeID  uint       `json:"employee_id"`
	LeaveTypeID uint       `json:"leave_type_id"`
	Year        int        `json:"year"`
	TotalDays   float64    `json:"total_days"`
	UsedDays    float64    `json:"used_days"`
	GrantedAt   time.Time  `json:"granted_at"`           // NEW: when this balance was issued
	ExpiresAt   *time.Time `json:"expires_at,omitempty"` // NEW: optional precomputed expiry
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`

	// Relationships
	Employee  Employee  `gorm:"foreignKey:EmployeeID" json:"employee"`
	LeaveType LeaveType `gorm:"foreignKey:LeaveTypeID" json:"leave_type"`
}
