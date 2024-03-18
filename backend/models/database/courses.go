package database

import (
	"github.com/google/uuid"
	"time"
)

type Courses struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey;not null;default:uuid_generate_v4()"`
	Name        string    `gorm:"type:varchar(50);not null"`
	Detail      string    `gorm:"type:varchar(50);not null"`
	IsAvailable bool      `gorm:"type:boolean;not null;default:true"`
	CreatedBy   string    `gorm:"type:varchar(20);not null;default:'system'"`
	UpdatedBy   string    `gorm:"type:varchar(20);not null;default:'system'"`
	CreatedAt   time.Time `gorm:"autoCreateTime;not null;default:now()"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime;not null;default:now()"`
	Mentors     []Mentors `gorm:"many2many:mentor_courses;"`

	Booking	    Bookings `gorm:"foreignKey:CourseID;references:ID"`
}
