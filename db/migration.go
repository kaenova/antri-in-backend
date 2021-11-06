package db

import (
	"antri-in-backend/entity"

	"gorm.io/gorm"
)

func Migration(db *gorm.DB) {
	db.AutoMigrate(&entity.Admin{}, &entity.Antrian{}, &entity.Pengantri{})
}
