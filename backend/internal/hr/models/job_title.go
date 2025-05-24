package models

import "time"

type JobTitle struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Title       string    `gorm:"uniqueIndex" json:"title"` // e.g., "Software Engineer"
	Code        string    `gorm:"uniqueIndex" json:"code"`  // e.g., "SE-II"
	Description string    `json:"description"`
	Level       string    `json:"level"` // e.g., "Junior", "Senior"
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
