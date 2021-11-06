package db

import (
	"antri-in-backend/config"
	"antri-in-backend/entity"
	"antri-in-backend/utils"
	"antri-in-backend/utils/errlogger"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func initData(db *gorm.DB) {
	config := config.GetConfig()

	// Create SuperUser account
	password, err := utils.HashPassword(config.SuperUser.Password)
	if err != nil {
		errlogger.ErrFatalPanic(err)
	}
	su := entity.Admin{
		ID:       uuid.New(),
		Nama:     config.SuperUser.Email,
		Email:    config.SuperUser.Email,
		Password: password,
		Role:     entity.ROLES_SUPER_USER,
		IsActive: true,
	}
	if err := db.Create(&su).Error; err != nil {
		errlogger.ErrFatalPanic(err)
	}

}
