package config

import (
	"database/sql"
	"fmt"
)

var (
	db     *sql.DB
	logger *Logger
	err    error
)

func InitConfig() error {
	db, err = InitDB()
	if err != nil {
		return fmt.Errorf("Error initialize database: %v", err)
	}

	return nil
}

func GetDB() *sql.DB {
	return db
}

func GetLogger(p string) *Logger {
	logger := NewLogger(p)

	return logger
}
