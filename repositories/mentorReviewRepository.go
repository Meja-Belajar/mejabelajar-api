package repositories

import (
	"context"
	"os"
	"time"

	"github.com/meja_belajar/configs"
	"github.com/meja_belajar/models/database"
	"gorm.io/gorm"
)

func GetAllMentorReviewByMentorID(MentorID string) ([]database.MentorReviews, error){
	var mentorReview []database.MentorReviews
	timeout, err := time.ParseDuration(os.Getenv("TIMEOUT"))
	if err != nil {
		return mentorReview, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	db := configs.GetDB().WithContext(ctx)

	err = db.
		Table("mentor_reviews").
		Where("mentor_id = ?", MentorID).
		Find(&mentorReview).
		Error

	if ctx.Err() == context.DeadlineExceeded {
		return mentorReview, ctx.Err()
	}

	if err == gorm.ErrRecordNotFound {
		return mentorReview, err
	}

	return mentorReview, nil
}