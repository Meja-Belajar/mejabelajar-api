package database

import (
	"github.com/google/uuid"
)

type MentorReviews struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey;not null;default:uuid_generate_v4()"`
	MentorID    string    `gorm:"type:uuid;not null"`
	Description string    `gorm:"type:text; not null"`
}
