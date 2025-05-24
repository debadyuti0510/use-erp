package models

import "time"

type LeaveRequest struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	EmployeeID  uint      `json:"employee_id"`
	LeaveTypeID uint      `json:"leave_type_id"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	Reason      string    `json:"reason"`
	Status      string    `json:"status"` // pending, approved, rejected
	ApproverID  *uint     `json:"approver_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	// Relationships
	Employee  Employee  `gorm:"foreignKey:EmployeeID" json:"employee"`
	LeaveType LeaveType `gorm:"foreignKey:LeaveTypeID" json:"leave_type"`
	Approver  *Employee `gorm:"foreignKey:ApproverID" json:"approver,omitempty"`
}
