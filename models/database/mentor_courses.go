package database

import (
	"time"

	"github.com/google/uuid"
)

type MentorCourses struct {
	MentorID        uuid.UUID `gorm:"type:uuid;primaryKey;not null"`
	CourseID        uuid.UUID `gorm:"type:uuid;primaryKey;not null"`
	CourseStartTime time.Time `gorm:"not null"`
	CourseEndTime   time.Time `gorm:"not null"`
	Rating          float64   `gorm:"type:float;not null;default:0"`
	HourlyRate      float64   `gorm:"type:float;not null;default:0"`
	IsActive        bool      `gorm:"type:boolean;not null;default:true"`
	CreatedBy       string    `gorm:"type:varchar(50);not null"`
	UpdatedBy       string    `gorm:"type:varchar(50);not null"`
	CreatedAt       time.Time `gorm:"type:timestamp;not null;default:now()"`
	UpdatedAt       time.Time `gorm:"type:timestamp;not null;default:now()"`
}
