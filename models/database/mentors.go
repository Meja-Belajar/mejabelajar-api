package database

import (
	"github.com/google/uuid"
)

type Mentors struct {
	ID                 uuid.UUID `gorm:"type:uuid;primaryKey;not null;default:uuid_generate_v4()"`
	Revenue            float64   `gorm:"type:float;default:0"`
	Rating             float64   `gorm:"type:float;default:0"`
	TotalTeachingHours int       `gorm:"type:int;default:0"`
	TeachingFrequency  int       `gorm:"type:int;default:0"`
	Log                Log       `gorm:"embedded"`

	Courses       []Courses       `gorm:"many2many:mentor_courses;"`
	MentorReviews []MentorReviews `gorm:"foreignKey:mentor_id;references:id"`
	Bookings      []Bookings      `gorm:"foreignKey:mentor_id;references:id"`
	UserID        uuid.UUID       `gorm:"type:uuid"`
}
