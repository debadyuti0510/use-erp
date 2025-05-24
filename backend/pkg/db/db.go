package db

import (
	"fmt"
	"log/slog"

	"github.com/debadyuti0510/use-erp/internal/hr/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		"localhost",
		"postgres",
		"postgres",
		"oserp_db",
		"5432")
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		slog.Error("Failed to connect to DB: %v", "err", err)
	}

	// Auto migrate models
	err = DB.AutoMigrate(&models.Employee{})
	if err != nil {
		slog.Error("AutoMigrate failed: %v", "err", err)
	} else {
		slog.Info("Succesfully migrated models!")
	}
}
