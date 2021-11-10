package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const (
	ROLES_SUPER_USER = "super"
	ROLES_ADMIN      = "admin"
)

type Admin struct {
	ID        uuid.UUID      `gorm:"type:uuid;primary_key;" json:"_id,omitempty"`
	Nama      string         `gorm:"not null" json:"nama"`
	Email     string         `gorm:"not null" json:"email"`
	Password  string         `gorm:"not null" json:"password,omitempty"`
	Role      string         `gorm:"not null" json:"role,omitempty"`
	IsActive  bool           `gorm:"not null" json:"is_active,omitempty"`
	CreatedAt time.Time      `gorm:"type:timestamptz;not null" json:"created_at,omitempty"`
	UpdatedAt time.Time      `gorm:"type:timestamptz;not null" json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`
}

func (Admin) TableName() string {
	return "admin"
}

func (b *Admin) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uuid.New()
	return
}
