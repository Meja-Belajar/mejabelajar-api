package database

import (
	"time"

	"github.com/google/uuid"
)

type invoices struct {
	ID             uuid.UUID  `gorm:"type:uuid;primaryKey;not null;default:uuid_generate_v4()"`
	PaymentMethod  string     `gorm:"type:varchar(50);not null"`
	PaymentName    string     `gorm:"type:varchar(50);not null"`
	PaymentStatus  string     `gorm:"type:varchar(50);not null"`
	PaymentAmount  float64    `gorm:"type:float;not null;default:0"`
	PaymentFee     float64    `gorm:"type:float;not null;default:0"`
	PaymentTotal   float64    `gorm:"type:float;not null;default:0"`
	IsActive       bool       `gorm:"type:boolean;not null;default:true"`
	CreatedBy      string     `gorm:"type:varchar(50);not null;default:'system'"`
	UpdatedBy      string     `gorm:"type:varchar(50);not null;default:'system'"`
	CreatedAt      time.Time  `gorm:"autoCreateTime;not null;default:now()"`
	UpdatedAt      time.Time  `gorm:"autoUpdateTime;not null;default:now()"`
	
	Bookings       Bookings   `gorm:"foreignKey:InvoiceID;references:ID"`
	
}