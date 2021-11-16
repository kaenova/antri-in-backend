package model

import (
	"antri-in-backend/db"
	"antri-in-backend/entity"
	"errors"

	"github.com/google/uuid"
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

func Login(email string) (bool, entity.Admin) {
	var obj entity.Admin

	db := db.GetDB()

	err := db.Where("email = ? AND is_active = true", email).First(&obj)
	if errors.Is(err.Error, gorm.ErrRecordNotFound) {
		return false, entity.Admin{}
	}
	return true, obj
}

func AdminRequestAll() ([]entity.Admin, error) {
	var objs []entity.Admin

	db := db.GetDB()

	if err := db.Select("id, nama, email, role, created_at").Where("role = ? AND is_active = false", entity.ROLES_ADMIN).Find(&objs).Error; err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			return []entity.Admin{}, nil
		}
		return []entity.Admin{}, err
	}
	return objs, nil
}

func AcceptAdmin(adminID uuid.UUID) error {
	var obj entity.Admin

	db := db.GetDB()
	obj.ID = adminID

	if err := db.Model(&obj).Update("is_active", true).Error; err != nil {
		return err
	}
	return nil
}

func DeleteAdmin(adminID uuid.UUID) error {
	var obj entity.Admin

	db := db.GetDB()
	obj.ID = adminID

	if err := db.Delete(&obj).Error; err != nil {
		return err
	}
	return nil
}
