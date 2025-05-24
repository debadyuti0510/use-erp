package models

import "time"

type LeaveType struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Name         string    `gorm:"uniqueIndex" json:"name"`
	Code         string    `gorm:"uniqueIndex" json:"code"`
	Description  string    `json:"description"`
	ExpiryInDays *int      `json:"expiry_in_days"` // Nullable
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
