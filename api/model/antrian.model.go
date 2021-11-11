package model

import (
	"antri-in-backend/db"
	"antri-in-backend/entity"
)

func AntrianCreate(obj *entity.Antrian) error {
	db := db.GetDB()

	if err := db.Create(&obj).Error; err != nil {
		return err
	}

	return nil
}
