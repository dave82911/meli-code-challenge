package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"meli-code-challenge/helper"
	"os"
)

func DatabaseConnection() *gorm.DB {
	// Load environment variables
	databaseConnection := os.Getenv("DATABASE_CONNECTION")

	db, err := gorm.Open(postgres.Open(databaseConnection), &gorm.Config{})
	helper.ErrorPanic(err)

	return db
}
