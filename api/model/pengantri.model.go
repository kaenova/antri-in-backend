package model

import (
	"antri-in-backend/db"
	"antri-in-backend/entity"
)

func AddPengantriToAntrian(obj *entity.Pengantri) error {
	var noAntrian *int
	db := db.GetDB()

	// Get max number of antrian
	err := db.Table("pengantri").Select("max(no_antrian)").Where("antrian_id = ?", obj.AntrianID).Scan(&noAntrian).Error
	if err != nil {
		return err
	}

	if noAntrian == nil {
		obj.NoAntrian = 1
	} else {
		obj.NoAntrian = *noAntrian + 1
	}
	// Add obj to Antrian
	if err := db.Create(&obj).Error; err != nil {
		return err
	}

	return nil
}
