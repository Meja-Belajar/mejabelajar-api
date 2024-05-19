package repositories

import (
	"context"
	"os"
	"time"

	"github.com/meja_belajar/configs"
	"github.com/meja_belajar/models/database"
)

func FindMentorCourseByMentorID(mentorID string) ([]database.MentorCourses, error) {
	var mentorCourses []database.MentorCourses
	timeout, err := time.ParseDuration(os.Getenv("TIMEOUT"))
	if err != nil {
		return mentorCourses, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	db := configs.GetDB().WithContext(ctx)
	err = db.Table("mentor_courses").
		Where("mentor_courses.mentor_id", mentorID).
		Find(&mentorCourses).
		Error

	if ctx.Err() == context.DeadlineExceeded {
		return mentorCourses, ctx.Err()
	}
	if err != nil {
		return mentorCourses, err
	}
	return mentorCourses, nil
}
