package handler

import (
	"gorm.io/gorm"

	"github.com/Bran00/oportuniesgo/config"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitializeHandler() {
	logger = config.GetLogger("handler")
	db = config.GetSQLite()
}
