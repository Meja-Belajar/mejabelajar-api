package repositories

import (
	"context"
	"os"
	"time"

	"github.com/meja_belajar/configs"
	"github.com/meja_belajar/models/database"
	"gorm.io/gorm"
)
func FindMentorByUserID(userID string) (database.Mentors, error){
	var mentor database.Mentors
	timeout, err := time.ParseDuration(os.Getenv("TIMEOUT"))
	if err != nil {
		return mentor, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	db := configs.GetDB().WithContext(ctx)
	err = db.
		Table("mentors").
		Where("user_id = ?", userID).
		First(&mentor).
		Error

	//validasi timeout
	if ctx.Err() == context.DeadlineExceeded {
		return mentor, ctx.Err()
	}

	//validasi jika mentor tidak ditemukan
	if err == gorm.ErrRecordNotFound {
		return mentor, err
	}

	return mentor, nil
}
func FindMentorByID(mentorID string) (database.Mentors, error) {
	var mentor database.Mentors
	timeout, err := time.ParseDuration(os.Getenv("TIMEOUT"))
	if err != nil {
		return mentor, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	db := configs.GetDB().WithContext(ctx)
	err = db.
		Table("mentors").
		Where("id = ?", mentorID).
		First(&mentor).
		Error

	//validasi timeout
	if ctx.Err() == context.DeadlineExceeded {
		return mentor, ctx.Err()
	}

	//validasi jika mentor tidak ditemukan
	if err == gorm.ErrRecordNotFound {
		return mentor, err
	}

	return mentor, nil
}