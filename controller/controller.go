package controller

import (
	"github.com/danielmesquitta/go-rate-api/config"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() {
	db = config.GetDatabase()
}
