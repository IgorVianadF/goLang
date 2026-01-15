package config

import (
	"fmt"
	"os"

	"github.com/IgorVianadF/goLang/schemas"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializePostgres() (*gorm.DB, error) {
	logger := GetLogger("Postgres")
	logger.Info("Initializating DB Connection")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Errorf("Error connecting to database: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Drawing{})
	if err != nil {
		logger.Errorf("Automigration error: %v", err)
		return nil, err
	}

	logger.Info("Successfully initialized DB connection")
	return db, nil
}
