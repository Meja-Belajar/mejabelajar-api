package database

import (
	"github.com/google/uuid"
	"time"
)

type Bookings struct {
	ID       		  uuid.UUID `gorm:"type:uuid;primaryKey;not null;default:uuid_generate_v4()"`
	//	FK
	UserID   		  uuid.UUID `gorm:"type:uuid;not null"`
	MentorID 		  uuid.UUID `gorm:"type:uuid;not null"`
	CourseID	      uuid.UUID `gorm:"type:uuid;not null"`
	InvoiceID 		  uuid.UUID `gorm:"type:uuid;not null"`
	//	FK
	Date		 	  time.Time `gorm:"type:date;not null"`
	Location 		  string    `gorm:"type:varchar(255);not null"`
	IsActive	      bool      `gorm:"type:boolean;not null;default:true"`
	CreatedBy         string    `gorm:"type:varchar(50);not null;default:'system'"`
	UpdatedBy         string    `gorm:"type:varchar(50);not null;default:'system'"`
	CreatedAt         time.Time `gorm:"autoCreateTime;not null;default:now()"`
	UpdatedAt         time.Time `gorm:"autoUpdateTime;not null;default:now()"`

}
