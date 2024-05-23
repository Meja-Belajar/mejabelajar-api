package database

import (
	"time"

	"github.com/google/uuid"
)

type Courses struct {
	ID     			uuid.UUID 	`gorm:"type:uuid;primaryKey;not null;default:uuid_generate_v4()"`
	Name  	 		string    	`gorm:"type:varchar(50);not null"`
	Detail 			string    	`gorm:"type:varchar(50);not null"`
	Rating 			float64 	`gorm:"type:float;not null;default:0"`
	HourlyRate		float64 	`gorm:"type:float;not null;default:0"`
	CourseStartTime time.Time 	`gorm:"type:timestamp;not null"`
	CourseEndTime 	time.Time 	`gorm:"type:timestamp;not null"`
	IsActive 		bool 		`gorm:"type:boolean;not null;default:true"`
	CreatedBy 		string 		`gorm:"type:varchar(50);not null"`
	UpdatedBy 		string 		`gorm:"type:varchar(50);not null"`
	CreatedAt 		time.Time 	`gorm:"type:timestamp;not null;default:now()"`
	UpdatedAt 		time.Time 	`gorm:"type:timestamp;not null;default:now()"`

	// Relationship
	MentorCourses []MentorCourses `gorm:"foreignKey:course_id;references:id"`
	// Mentors []Mentors `gorm:"many2many:mentor_courses;"`
	Booking Bookings  `gorm:"foreignKey:course_id;references:id"`
}
