package model

import (
	"antri-in-backend/db"
	"antri-in-backend/entity"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func AddPengantriToAntrian(obj *entity.Pengantri) error {
	var noAntrian int
	var antrian entity.Antrian
	db := db.GetDB()

	tx := db.Begin()
	tx.SavePoint("sp1")

	// Get max number of antrian
	err := tx.Table("antrian").Select("max_nomor_antrian").Where("id = ?", obj.AntrianID).Scan(&noAntrian).Error
	if err != nil {
		tx.RollbackTo("sp1")
		tx.Commit()
		return err
	}

	if noAntrian == 0 {
		obj.NoAntrian = 1
	} else {
		obj.NoAntrian = noAntrian + 1
	}

	// Update max nomor antrian
	err = db.Where("id = ?", obj.AntrianID).First(&antrian).Error
	if err != nil {
		tx.RollbackTo("sp1")
		tx.Commit()
		return err
	}

	antrian.MaxNomorAntrian = obj.NoAntrian
	err = tx.Save(&antrian).Error
	if err != nil {
		tx.RollbackTo("sp1")
		tx.Commit()
		return err
	}

	// Add obj to Antrian
	if err := tx.Create(&obj).Error; err != nil {
		tx.RollbackTo("sp1")
		tx.Commit()
		return err
	}
	tx.Commit()
	return nil
}

func NoHPIsUsed(no string) (entity.Pengantri, bool) {
	var obj entity.Pengantri
	db := db.GetDB()

	if err := db.Where("no_telp = ?", no).First(&obj).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return obj, false
		}
		return obj, true
	}

	return obj, true
}

func DeletePengantri(idPengantri uuid.UUID) error {
	var obj entity.Pengantri
	obj.ID = idPengantri
	db := db.GetDB()
	if err := db.Clauses(clause.Returning{}).Delete(&obj).Error; err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	return nil
}
