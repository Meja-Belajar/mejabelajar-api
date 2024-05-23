package repositories

import (
	"context"
	"os"
	"time"

	"github.com/meja_belajar/configs"
	"github.com/meja_belajar/models/database"
	"gorm.io/gorm"
)

func InsertMentor(mentor database.Mentors) (database.Mentors, error) {
	timeout, err := time.ParseDuration(os.Getenv("TIMEOUT"))
	if err != nil {
		return mentor, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	db := configs.GetDB().WithContext(ctx)

	err = db.Create(&mentor).Error
	if ctx.Err() == context.DeadlineExceeded {
		return mentor, ctx.Err()
	}
	if err != nil {
		return mentor, err
	}
	return mentor, nil
}

func FindPopularMentor() ([]database.Mentors, error) {
	var mentors []database.Mentors
	timeout, err := time.ParseDuration(os.Getenv("TIMEOUT"))
	if err != nil {
		return mentors, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	db := configs.GetDB().WithContext(ctx)
	err = db.
		Find(&mentors).
		Order("rating desc").
		Limit(10).
		Error

	//validasi timeout
	if ctx.Err() == context.DeadlineExceeded {
		return mentors, ctx.Err()
	}

	//validasi jika mentor tidak ditemukan
	if err == gorm.ErrRecordNotFound {
		return mentors, err
	}
	return mentors, nil
}

func FindAllMentor() ([]database.Mentors, error) {
	var mentors []database.Mentors
	timeout, err := time.ParseDuration(os.Getenv("TIMEOUT"))
	if err != nil {
		return mentors, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	db := configs.GetDB().WithContext(ctx)
	err = db.Find(&mentors).Error

	//validasi timeout
	if ctx.Err() == context.DeadlineExceeded {
		return mentors, ctx.Err()
	}

	//validasi jika mentor tidak ditemukan
	if err == gorm.ErrRecordNotFound {
		return mentors, err
	}
	return mentors, nil
}

func FindMentorByUserID(userID string) (database.Mentors, error) {
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
