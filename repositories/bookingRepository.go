package repositories

import (
	"context"

	"github.com/meja_belajar/configs"
	"github.com/meja_belajar/models/database"
)

func CreateBooking(ctx context.Context, booking database.Bookings) (*database.Bookings, error) {
	db := configs.GetDB()
	err := db.Model(&database.Bookings{}).
		WithContext(ctx).
		Create(&booking).
		Error
	return &booking, err
}
