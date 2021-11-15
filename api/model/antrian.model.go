package model

import (
	"antri-in-backend/db"
	"antri-in-backend/entity"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func AntrianCreate(obj *entity.Antrian) error {
	db := db.GetDB()

	if err := db.Create(&obj).Error; err != nil {
		return err
	}

	return nil
}

func AntrianGetAll() ([]entity.Antrian, error) {
	var objs []entity.Antrian

	db := db.GetDB()

	if tx := db.Find(&objs); tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return []entity.Antrian{}, nil
		}
		return []entity.Antrian{}, tx.Error
	}
	return objs, nil
}

func AntrianSearchName(cat string) ([]entity.Antrian, error) {
	var (
		objs []entity.Antrian
	)

	db := db.GetDB()

	tx := db.Where("LOWER(nama) LIKE LOWER(?)", fmt.Sprintf("%%%s%%", cat)).Find(&objs)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return []entity.Antrian{}, nil
		}
		return []entity.Antrian{}, tx.Error
	}

	return objs, nil
}

func AntrianDelete(id uuid.UUID) (interface{}, error) {
	db := db.GetDB()
	var pengantri []entity.Pengantri
	var antrian entity.Antrian
	// Hapus user yang di antrian itu
	tx := db.Begin()
	tx.SavePoint("sp1")
	err := tx.Clauses(clause.Returning{}).Where("antrian_id = ?", id).Delete(&pengantri).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) && err != nil {
		tx.RollbackTo("sp1")
		tx.Commit()
		return entity.Antrian{}, err
	}

	// Hapus Antriannya
	err = tx.Clauses(clause.Returning{}).Where("id = ?", id).Delete(&antrian).Error
	if err != nil {
		tx.RollbackTo("sp1")
		tx.Commit()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	tx.Commit()

	return map[string]interface{}{"antrian": antrian, "pengantri": pengantri}, nil
}

func AntrianUbah(obj *entity.Antrian) error {
	db := db.GetDB()
	var tempAntrian entity.Antrian

	if err := db.Where("id = ?", obj.ID).First(&tempAntrian).Error; err != nil {
		return err
	}

	obj.CurrNomorAntrian = tempAntrian.CurrNomorAntrian
	obj.EstimasiAntrian = tempAntrian.EstimasiAntrian

	if err := db.Save(&obj).Error; err != nil {
		return err
	}

	return nil
}

func TambahNomorAntrian(obj *entity.Antrian) error {
	db := db.GetDB()
	if err := db.Where("id = ?", obj.ID).First(&obj).Error; err != nil {
		return err
	}
	obj.CurrNomorAntrian = obj.CurrNomorAntrian + 1
	if err := db.Save(&obj).Error; err != nil {
		return err
	}
	return nil
}
