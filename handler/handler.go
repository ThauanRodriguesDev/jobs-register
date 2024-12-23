package handler

import (
	"github.com/ThauanRodriguesDev/jobs-register/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitHandler() {
	logger = config.GetLogger("handler")
	db = config.GetSqlite()
}
