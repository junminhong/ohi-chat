package db

import (
	"github.com/junminhong/ohi-chat/app/models"
	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(&models.Member{})
}
