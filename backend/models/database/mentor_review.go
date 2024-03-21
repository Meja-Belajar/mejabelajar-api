package database

import "time"

type MentorReviews struct {
	ID          string    `gorm:"type:uuid;primaryKey;not null;default:uuid_generate_v4()"`
	MentorID    string    `gorm:"type:uuid;not null"`	
	Description string    `gorm:"type:text; not null"`
	Is_Active   bool      `gorm:"type:boolean;not null;default:true"`
	CreatedBy   string    `gorm:"type:varchar(50);not null;default:'system'"`
	UpdatedBy   string    `gorm:"type:varchar(50);not null;default:'system'"`
	CreatedAt   time.Time `gorm:"autoCreateTime;not null;default:now()"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime;not null;default:now()"`
}