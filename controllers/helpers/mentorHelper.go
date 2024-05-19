package helpers

import (
	"context"
	"log"

	"github.com/meja_belajar/models/database"
	"github.com/meja_belajar/models/outputs"
	"github.com/meja_belajar/models/responses"
	"github.com/meja_belajar/repositories"
	"github.com/meja_belajar/utils"
	"gorm.io/gorm"
)

func GetAllMentor() (int, interface{}) {
	var mentors []database.Mentors
	var data []responses.MentorResponseDTO
	mentors, err := repositories.FindAllMentor()

	//validasi timeout
	if err == context.DeadlineExceeded {
		return utils.HandleTimeout(err)
	}
	//validasi not found
	if err == gorm.ErrRecordNotFound {
		return utils.HandleNotFound("Mentors")
	}

	for _, mentor := range mentors {
		//cari user
		user, err := repositories.FindUserByUserID(mentor.UserID.String())
		//validasi timeout
		if err == context.DeadlineExceeded {
			return utils.HandleTimeout(err)
		}
		if mentor.UserID != user.ID {
			//print mentor dengan UserID yang tidak ada di table user
			log.Printf("Mentor with ID %s has an wrong UserID", mentor.ID)
			continue
		}

		//cari reviews
		reviewResponseDTO, err := repositories.FindReviewByMentorID(mentor.ID.String())
		// validasi timeout
		if err == context.DeadlineExceeded {
			return utils.HandleTimeout(err)
		}

		//cari course
		var mentorCourses []database.MentorCourses
		mentorCourses, err = repositories.FindMentorCourseByMentorID(mentor.ID.String())

		//validasi timeout
		if err == context.DeadlineExceeded {
			return utils.HandleTimeout(err)
		}

		var courseResponseDTO []responses.CourseResponseDTO
		for _, mentorCourse := range mentorCourses {
			course, err := repositories.FindCourseByID(mentorCourse.CourseID.String())
			//validasi timeout
			if err == context.DeadlineExceeded {
				return utils.HandleTimeout(err)
			}
			courseResponseDTO = append(courseResponseDTO, responses.CourseResponseDTO{
				CourseID:        course.ID.String(),
				Name:            course.Name,
				Detail:          course.Detail,
				Rating:          mentorCourse.Rating,
				HourlyRate:      mentorCourse.HourlyRate,
				CourseStartTime: mentorCourse.CourseStartTime,
				CourseEndTime:   mentorCourse.CourseEndTime,
			})
		}

		data = append(data, responses.MentorResponseDTO{
			MentorID:           mentor.ID.String(),
			Username:           user.Username,
			University:         user.University,
			Email:              user.Email,
			Phone:              user.Phone,
			Description:        user.Description,
			ProfilePicture:     user.ProfilePicture,
			BOD:                user.BOD,
			Revenue:            mentor.Revenue,
			Rating:             mentor.Rating,
			TotalTeachingHours: mentor.TotalTeachingHours,
			TeachingFrequency:  mentor.TeachingFrequency,
			Courses:            courseResponseDTO,
			Reviews:            reviewResponseDTO,
		})
	}
	if err != nil {
		return utils.HandleInternalServerError(err)
	}
	output := outputs.GetAllMentorOutput{
		BaseOutput: outputs.BaseOutput{
			Code:    200,
			Message: "Success: Mentor found",
		},
		Data: data,
	}
	return 200, output
}

func GetMentorByMentorID(mentorID string) (int, interface{}) {
	var mentor database.Mentors
	mentor, err := repositories.FindMentorByID(mentorID)
	//validasi timeout
	if err == context.DeadlineExceeded {
		return utils.HandleTimeout(err)
	}
	//validasi not found
	if err == gorm.ErrRecordNotFound {
		return utils.HandleNotFound("Mentor")
	}

	var user database.Users
	user, err = repositories.FindUserByUserID(mentor.UserID.String())

	//validasi timeout
	if err == context.DeadlineExceeded {
		return utils.HandleTimeout(err)
	}

	//validasi jika user tidak ditemukan
	if err == gorm.ErrRecordNotFound {
		return utils.HandleNotFound("User")
	}

	var reviewResponseDTO []responses.ReviewResponseDTO
	reviewResponseDTO, err = repositories.FindReviewByMentorID(mentorID)

	//validasi timeout
	if err == context.DeadlineExceeded {
		return utils.HandleTimeout(err)
	}

	var mentorCourses []database.MentorCourses
	mentorCourses, err = repositories.FindMentorCourseByMentorID(mentorID)

	//validasi timeout
	if err == context.DeadlineExceeded {
		return utils.HandleTimeout(err)
	}

	var courseResponseDTO []responses.CourseResponseDTO
	for _, mentorCourse := range mentorCourses {
		course, err := repositories.FindCourseByID(mentorCourse.CourseID.String())

		courseResponseDTO = append(courseResponseDTO, responses.CourseResponseDTO{
			CourseID:        course.ID.String(),
			Name:            course.Name,
			Detail:          course.Detail,
			Rating:          mentorCourse.Rating,
			HourlyRate:      mentorCourse.HourlyRate,
			CourseStartTime: mentorCourse.CourseStartTime,
			CourseEndTime:   mentorCourse.CourseEndTime,
		})

		//validasi timeout
		if err == context.DeadlineExceeded {
			return utils.HandleTimeout(err)
		}
	}
	if err != nil {
		return utils.HandleInternalServerError(err)
	}
	output := outputs.GetMentorByIDOutput{
		BaseOutput: outputs.BaseOutput{
			Code:    200,
			Message: "Success: Mentor found",
		},
		Data: responses.MentorResponseDTO{
			MentorID:           mentor.ID.String(),
			Username:           user.Username,
			University:         user.University,
			Email:              user.Email,
			Phone:              user.Phone,
			Description:        user.Description,
			ProfilePicture:     user.ProfilePicture,
			BOD:                user.BOD,
			Revenue:            mentor.Revenue,
			Rating:             mentor.Rating,
			TotalTeachingHours: mentor.TotalTeachingHours,
			TeachingFrequency:  mentor.TeachingFrequency,
			Courses:            courseResponseDTO,
			Reviews:            reviewResponseDTO,
		},
	}
	return 200, output
}

func GetMentorByUserID(userID string) (int, interface{}) {
	user, err := repositories.FindUserByUserID(userID)
	//validasi timeout
	if err == context.DeadlineExceeded {
		return utils.HandleTimeout(err)
	}
	if err == gorm.ErrRecordNotFound {
		return utils.HandleNotFound("User")
	}

	mentor, err := repositories.FindMentorByUserID(userID)
	//validasi timeout
	if err == context.DeadlineExceeded {
		return utils.HandleTimeout(err)
	}
	if err == gorm.ErrRecordNotFound {
		return utils.HandleNotFound("Mentor")
	}

	var reviewResponseDTO []responses.ReviewResponseDTO
	reviewResponseDTO, err = repositories.FindReviewByMentorID(mentor.ID.String())
	//validasi timeout
	if err == context.DeadlineExceeded {
		return utils.HandleTimeout(err)
	}
	var mentorCourses []database.MentorCourses
	mentorCourses, err = repositories.FindMentorCourseByMentorID(mentor.ID.String())
	//validasi timeout
	if err == context.DeadlineExceeded {
		return utils.HandleTimeout(err)
	}
	var courseResponseDTO []responses.CourseResponseDTO
	for _, mentorCourse := range mentorCourses {
		course, err := repositories.FindCourseByID(mentorCourse.CourseID.String())
		courseResponseDTO = append(courseResponseDTO, responses.CourseResponseDTO{
			CourseID:        course.ID.String(),
			Name:            course.Name,
			Detail:          course.Detail,
			Rating:          mentorCourse.Rating,
			HourlyRate:      mentorCourse.HourlyRate,
			CourseStartTime: mentorCourse.CourseStartTime,
			CourseEndTime:   mentorCourse.CourseEndTime,
		})
		//validasi timeout
		if err == context.DeadlineExceeded {
			return utils.HandleTimeout(err)
		}
	}
	if err != nil {
		return utils.HandleInternalServerError(err)
	}
	output := outputs.GetMentorByIDOutput{
		BaseOutput: outputs.BaseOutput{
			Code:    200,
			Message: "Success: Mentor found",
		},
		Data: responses.MentorResponseDTO{
			MentorID:           mentor.ID.String(),
			Username:           user.Username,
			University:         user.University,
			Email:              user.Email,
			Phone:              user.Phone,
			Description:        user.Description,
			ProfilePicture:     user.ProfilePicture,
			BOD:                user.BOD,
			Revenue:            mentor.Revenue,
			Rating:             mentor.Rating,
			TotalTeachingHours: mentor.TotalTeachingHours,
			TeachingFrequency:  mentor.TeachingFrequency,
			Courses:            courseResponseDTO,
			Reviews:            reviewResponseDTO,
		},
	}
	return 200, output
}
