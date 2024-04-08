package database

import (
	"github.com/google/uuid"
	"time"
);

//  CEK LAGI NOT NULLNYA

type MentorCourses struct {
	MentorID        uuid.UUID `gorm:"type:uuid;primaryKey;not null"`
	CourseID        uuid.UUID `gorm:"type:uuid;primaryKey;not null"`
	CourseStartTime time.Time `gorm:"not null"`
	CourseEndTime   time.Time `gorm:"not null"`
	Rating          float64   `gorm:"type:float;not null;default:0"`
	HourlyRate      float64   `gorm:"type:float;not null;default:0"`
	Log      Log       	`gorm:"embedded"`
}
