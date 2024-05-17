package database

import (
	"time"

	"github.com/google/uuid"
)

type Notification struct {
	Id               uuid.UUID `gorm:"primaryKey;default:uuid_generate_v4()"`
	UserId           uuid.UUID `gorm:"not null"`
	Description      string    `gorm:"type:TEXT"`
	NotificationType string    `gorm:"type:VARCHAR(10)"`
	IsAvailable      bool      `gorm:"not null"`
	CreatedBy        string    `gorm:"type:VARCHAR(50); not null"`
	CreatedAt        time.Time `gorm:"autoCreateTime;not null;default:now()"`
}
