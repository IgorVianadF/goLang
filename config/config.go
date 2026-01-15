package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error

	db, err = InitializePostgres()
	if err != nil {
		return fmt.Errorf("Error at the initialization of the database %v", err)
	}

	return nil
}

func GetDB() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}
