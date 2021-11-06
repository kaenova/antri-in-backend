package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Antrian struct {
	ID               uuid.UUID      `gorm:"type:uuid;primary_key;" json:"_id,omitempty"`
	Nama             string         `gorm:"not null" json:"nama"`
	Deskripsi        string         `gorm:"" json:"deskripsi"`
	CurrNomorAntrian int            `gorm:"not null" json:"curr_antrian,omitempty"`
	EstimasiAntrian  time.Time      `gorm:"not null" json:"estimasi"`
	CreatedAt        time.Time      `gorm:"type:timestamptz;not null" json:"created_at,omitempty"`
	UpdatedAt        time.Time      `gorm:"type:timestamptz;not null" json:"updated_at,omitempty"`
	DeletedAt        gorm.DeletedAt `json:"deleted_at,omitempty"`
}

func (Antrian) TableName() string {
	return "antrian"
}
