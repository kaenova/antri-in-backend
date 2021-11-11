package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Pengantri struct {
	ID        uuid.UUID `gorm:"primary_key; index;" json:"_id,omitempty"`
	Nama      string    `gorm:"not null" json:"nama"`
	NoTelp    string    `gorm:"not null" json:"no_telp"`
	NoAntrian int       `gorm:"not null" json:"no_antrian"`
	IsCancel  bool      `gorm:"not null" json:"is_cancel"`

	AntrianID uuid.UUID `gorm:"not null;index" json:"antrian_id"`
	Antrian   *Antrian  `gorm:"foreignKey:AntrianID" json:"antrian,omitempty"`

	CreatedAt time.Time      `gorm:"type:timestamptz;not null" json:"created_at,omitempty"`
	UpdatedAt time.Time      `gorm:"type:timestamptz;not null" json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`
}

func (Pengantri) TableName() string {
	return "pengantri"
}

func (u *Pengantri) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	u.IsCancel = false
	return
}
