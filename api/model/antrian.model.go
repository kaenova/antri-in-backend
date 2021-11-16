package model

import (
	"antri-in-backend/db"
	"antri-in-backend/entity"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
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
	// Ambil nomor antrian
	if err := db.Where("id = ?", obj.ID).First(&obj).Error; err != nil {
		return err
	}

	tx := db.Begin()
	tx.SavePoint("sp1")
	if obj.MaxNomorAntrian >= obj.CurrNomorAntrian {
		// Hapus data pengantri
		if err := tx.Delete(&entity.Pengantri{}, "no_antrian = ? AND antrian_id = ?", obj.CurrNomorAntrian, obj.ID).Error; err != nil && !errors.Is(gorm.ErrRecordNotFound, err) {
			tx.RollbackTo("sp1")
			tx.Commit()
			return err
		}

		obj.CurrNomorAntrian = obj.CurrNomorAntrian + 1
		if err := tx.Save(&obj).Error; err != nil {
			tx.RollbackTo("sp1")
			tx.Commit()
			return err
		}
	}
	tx.Commit()
	return nil
}

func AmbilPengantribyAntrianID(id uuid.UUID) (interface{}, error) {
	var objs []entity.Pengantri
	var antri entity.Antrian

	db := db.GetDB()

	if err := db.Where("id = ?", id).Find(&antri).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	if err := db.Where("antrian_id = ?", id).Find(&objs).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return echo.Map{"antrian": antri, "pengantri": objs}, nil
}

func AntrianbyID(id uuid.UUID) (entity.Antrian, error) {
	var obj entity.Antrian
	db := db.GetDB()

	if err := db.Where("id = ?", id).First(&obj).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.Antrian{}, nil
		}
		return entity.Antrian{}, err
	}
	return obj, nil
}
