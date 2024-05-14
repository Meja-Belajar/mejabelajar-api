package database

import (
	"time"

	"github.com/google/uuid"
)

type Users struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey;not null;default:uuid_generate_v4()"`
	Username       string    `gorm:"type:varchar(50);not null"`
	University     string    `gorm:"type:varchar(50)"`
	Email          string    `gorm:"type:varchar(50);not null;unique"`
	Password       string    `gorm:"type:varchar(100);not null"`
	Phone          string    `gorm:"type:varchar(20);unique"`
	Description    string    `gorm:"type:text"`
	ProfilePicture string    `gorm:"type:varchar(255)"`
	BOD            time.Time `gorm:"type:date"`

	Bookings     []Bookings   `gorm:"foreignKey:user_id;references:id"`
	Mentor       Mentors      `gorm:"foreignKey:user_id;references:id"`
	Notification Notification `gorm:"foreignKey:user_id;references:id"`
}
