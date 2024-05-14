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
	Log      Log       	`gorm:"embedded"`
}
