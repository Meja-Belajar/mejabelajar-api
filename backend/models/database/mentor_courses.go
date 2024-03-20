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
	Rating          float32   `gorm:"type:float;not null;default:0"`
	HourlyRate      float32   `gorm:"type:float;not null;default:0"`
	IsAvailable bool      `gorm:"type:boolean;not null;default:true"`
	CreatedBy   string    `gorm:"type:varchar(50);not null;default:'system'"`
	UpdatedBy   string    `gorm:"type:varchar(50);not null;default:'system'"`
	CreatedAt   time.Time `gorm:"autoCreateTime;not null;default:now()"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime;not null;default:now()"`
}
