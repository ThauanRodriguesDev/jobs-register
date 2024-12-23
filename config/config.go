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

	//Initialize SQLite
	db, err = InitializeSQLite()

	if err != nil {
		return fmt.Errorf("error initializing sqlit: %v \n", err)
	}

	return nil
}

func GetSqlite() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	//Initialize Logger
	logger = NewLogger(p)

	return logger
}
