package database

import (
	"time"

	"github.com/google/uuid"
)

type Mentors struct {
	ID                 uuid.UUID `gorm:"type:uuid;primaryKey;not null;default:uuid_generate_v4()"`
	Revenue 		   float64   `gorm:"type:float;default:0"`
	Rating             float64   `gorm:"type:float;default:0"`
	TotalTeachingHours int       `gorm:"type:int;default:0"`
	TeachingFrequency  int       `gorm:"type:int;default:0"`
	IsActive	       bool      `gorm:"type:boolean;not null;default:true"`
	CreatedBy          string    `gorm:"type:varchar(50);not null;default:'system'"`
	UpdatedBy          string    `gorm:"type:varchar(50);not null;default:'system'"`
	CreatedAt          time.Time `gorm:"autoCreateTime;not null;default:now()"`
	UpdatedAt          time.Time `gorm:"autoUpdateTime;not null;default:now()"`

	Courses            []Courses  `gorm:"many2many:mentor_courses;"`
 	MentorReviews      []MentorReviews `gorm:"foreignKey:MentorID;references:ID"`
	Bookings 		   []Bookings`gorm:"foreignKey:MentorID;references:ID"`
	UserID             uuid.UUID  `gorm:"type:uuid"`

}
