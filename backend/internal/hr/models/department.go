package models

import "time"

type Department struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Name        string `gorm:"uniqueIndex" json:"name"` // e.g. Engineering, HR, Finance
	Code        string `gorm:"uniqueIndex" json:"code"` // e.g. ENG, HR, FIN
	Description string `json:"description"`             // Optional: purpose/overview

	// Organizational structure
	ManagerID *uint `json:"manager_id"` // FK to Employee.ID
	ParentID  *uint `json:"parent_id"`  // Optional: for nested departments

	// Audit
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
