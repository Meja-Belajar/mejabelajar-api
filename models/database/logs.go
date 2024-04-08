package database

import "time"

type Log struct {
	IsActive  bool      `gorm:"type:boolean;not null;default:true"`
	CreatedBy string    `gorm:"type:varchar(50);not null;default:'system'"`
	UpdatedBy string    `gorm:"type:varchar(50);not null;default:'system'"`
	CreatedAt time.Time `gorm:"autoCreateTime;not null;default:now()"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;not null;default:now()"`
}