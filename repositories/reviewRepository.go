package repositories

import (
	"context"
	"os"
	"time"

	"github.com/meja_belajar/configs"
	"github.com/meja_belajar/models/responses"
	"gorm.io/gorm"
)

func FindReviewByMentorID(mentorID string) ([]responses.ReviewResponseDTO, error) {
	var reviewResponseDTO []responses.ReviewResponseDTO
	timeout, err := time.ParseDuration(os.Getenv("TIMEOUT"))
	if err != nil {
		return reviewResponseDTO, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	db := configs.GetDB().WithContext(ctx)

	err = db.
		Table("mentor_reviews").
		Where("mentor_id = (?)", mentorID).
		Find(&reviewResponseDTO).
		Error

	//validasi timeout
	if ctx.Err() == context.DeadlineExceeded {
		return reviewResponseDTO, ctx.Err()
	}

	//validasi not found
	if err == gorm.ErrRecordNotFound {
		return reviewResponseDTO, err
	}

	return reviewResponseDTO, err
}
