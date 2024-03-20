package database

type Discounts struct {
	ID            string  `gorm:"type:uuid;primaryKey;not null;default:uuid_generate_v4()"`
	Description   string  `gorm:"type:text";`
	TotalDiscount float64 `gorm:"type:float;not null;default:0"`
}