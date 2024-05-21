package database

import (
	"time"

	"github.com/google/uuid"
)

type Notification struct {
	ID          uuid.UUID `gorm:"primaryKey;default:uuid_generate_v4()"`
	Title       string    `gorm:"not null;type:VARCHAR(100)"`
	UserId      uuid.UUID `gorm:"not null"`
	Description string    `gorm:"type:TEXT"`
	CreatedAt   time.Time `gorm:"autoCreateTime;not null;default:now()"`
}
