package database

import (
	"github.com/google/uuid"
)

type Courses struct {
	ID     uuid.UUID `gorm:"type:uuid;primaryKey;not null;default:uuid_generate_v4()"`
	Name   string    `gorm:"type:varchar(50);not null"`
	Detail string    `gorm:"type:varchar(255);not null"`

	// Relationship
	MentorCourses []MentorCourses `gorm:"foreignKey:course_id;references:id"`
	// Mentors []Mentors `gorm:"many2many:mentor_courses;"`
	Booking Bookings  `gorm:"foreignKey:course_id;references:id"`
}
