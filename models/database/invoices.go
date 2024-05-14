package database

import (
	"github.com/google/uuid"
)

type Invoices struct {
	ID            uuid.UUID `gorm:"type:uuid;primaryKey;not null;default:uuid_generate_v4()"`
	PaymentMethod string    `gorm:"type:varchar(50);not null"`
	PaymentName   string    `gorm:"type:varchar(50);not null"`
	PaymentStatus string    `gorm:"type:varchar(50);not null"`
	PaymentAmount float64   `gorm:"type:float;not null;default:0"`
	PaymentFee    float64   `gorm:"type:float;not null;default:0"`
	PaymentTotal  float64   `gorm:"type:float;not null;default:0"`

	// Relationship
	Bookings Bookings `gorm:"foreignKey:invoice_id;references:id"`
}
