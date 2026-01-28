package handler

import (
	"github.com/IgorVianadF/goLang/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func Init() {
	logger = config.GetLogger("Handler")
	db = config.GetDB()
}
