package database

import (
	"github.com/google/uuid"
	"time"
)

type Bookings struct {
	ID       		  uuid.UUID `gorm:"type:uuid;primaryKey;not null;default:uuid_generate_v4()"`
	UserID   		  uuid.UUID `gorm:"type:uuid;not null"`
	MentorID 		  uuid.UUID `gorm:"type:uuid;not null"`
	CourseID	      uuid.UUID `gorm:"type:uuid;not null"`
	TeachingStartDate time.Time `gorm:"not null"`
	TeachingEndDate   time.Time `gorm:"not null"`
	IsAvailable       bool      `gorm:"type:boolean;not null;default:true"`
	CreatedBy         string    `gorm:"type:varchar(20);not null;default:'system'"`
	UpdatedBy         string    `gorm:"type:varchar(20);not null;default:'system'"`
	CreatedAt         time.Time `gorm:"autoCreateTime;not null;default:now()"`
	UpdatedAt         time.Time `gorm:"autoUpdateTime;not null;default:now()"`
}
