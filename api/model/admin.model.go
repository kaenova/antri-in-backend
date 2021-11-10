package model

import (
	"antri-in-backend/db"
	"antri-in-backend/entity"
	"errors"

	"gorm.io/gorm"
)

func AddAdmin(data entity.Admin) (entity.Admin, error) {
	db := db.GetDB()

	if err := db.Create(&data).Error; err != nil {
		return entity.Admin{}, err
	}

	return data, nil
}

func AdminEmailIsUsed(email string) (bool, entity.Admin) {
	var obj entity.Admin

	db := db.GetDB()

	err := db.Where("email = ?", email).First(&obj)
	if errors.Is(err.Error, gorm.ErrRecordNotFound) {
		return false, entity.Admin{}
	}
	return true, obj
}
