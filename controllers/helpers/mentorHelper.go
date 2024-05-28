package helpers

import (
	"context"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/meja_belajar/models/database"
	"github.com/meja_belajar/models/outputs"
	"github.com/meja_belajar/models/requests"
	"github.com/meja_belajar/models/responses"
	"github.com/meja_belajar/repositories"
	"github.com/meja_belajar/utils"
	"gorm.io/gorm"
)

func RegisterMentor(registerMentor requests.RegisterMentorRequestDTO) (int, interface{}) {
	//validasi user exist
	user, err := repositories.FindUserByUserID(registerMentor.UserID)
	if err == context.DeadlineExceeded {
		return utils.HandleTimeout(err)
	}
	if err == gorm.ErrRecordNotFound {
		return utils.HandleNotFound("User")
	}

	//validasi belum pernah jadi mentor
	_, err = repositories.FindMentorByUserID(registerMentor.UserID)
	if err == context.DeadlineExceeded {
		return utils.HandleTimeout(err)
	}
	if err == nil {
		return utils.HandleBadRequest("User already registered as mentor")
	}

	//change string to uuid
	userID, err := uuid.Parse(registerMentor.UserID)
	if err != nil {
		return utils.HandleInternalServerError(err)
	}
	bod, err := time.Parse("2006-01-02T15:04:05Z", registerMentor.BOD)
	if err != nil {
		return utils.HandleBadRequest("Invalid BOD format")
	}
	mentor := database.Mentors{
		UserID:             userID,
		Revenue:            0,
		Rating:             0,
		TotalTeachingHours: 0,
		TeachingFrequency:  0,
		IsActive:           false,
	}
	mentor, err = repositories.InsertMentor(mentor)
	if err == context.DeadlineExceeded {
		return utils.HandleTimeout(err)
	}
	if err != nil {
		return utils.HandleInternalServerError(err)
	}
	var courseResponseDTO []responses.CourseResponseDTO
	var mentorCourse database.MentorCourses
	var course database.Courses
	//insert course
	for _, courseName := range registerMentor.Courses {
		course, err = repositories.FindCourseByName(courseName)
		//course belum terdaftar
		if err != nil {
			course = database.Courses{
				Name:   courseName,
				Detail: "",
			}
			course, err = repositories.InsertCourse(course)
			if err == context.DeadlineExceeded {
				return utils.HandleTimeout(err)
			}
			if err != nil {
				return utils.HandleInternalServerError(err)
			}
		}
		mentorCourse = database.MentorCourses{
			MentorID:        mentor.ID,
			CourseID:        course.ID,
			CourseStartTime: time.Time{},
			CourseEndTime:   time.Time{},
			Rating:          0,
			HourlyRate:      0,
		}
		mentorCourse, err = repositories.InsertMentorCourse(mentorCourse)
		if err == context.DeadlineExceeded {
			return utils.HandleTimeout(err)
		}
		if err != nil {
			return utils.HandleInternalServerError(err)
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

	//insert mentorCourse
	output := outputs.RegisterMentor{
		BaseOutput: outputs.BaseOutput{
			Code:    201,
			Message: "Mentor Registered",
		},
		Data: responses.MentorResponseDTO{
			MentorID:           mentor.ID.String(),
			Username:           user.Username,
			University:         user.University,
			Email:              user.Email,
			Phone:              user.Phone,
			Description:        user.Description,
			ProfilePicture:     user.ProfilePicture,
			BOD:                bod,
			Revenue:            mentor.Revenue,
			Rating:             mentor.Rating,
			TotalTeachingHours: mentor.TotalTeachingHours,
			TeachingFrequency:  mentor.TeachingFrequency,
			Courses:            courseResponseDTO,
			Reviews:            []responses.ReviewResponseDTO{},
		},
	}
	return 201, output
}

func GetPopularMentor() (int, interface{}) {
	var data []responses.MentorResponseDTO
	mentors, err := repositories.FindPopularMentor()

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
		//validasi jika user tidak ditemukan
		if mentor.UserID != user.ID {
			//print mentor dengan UserID yang tidak ada di table user
			continue
		}

		//cari review
		reviewResponseDTO, err := repositories.FindReviewByMentorID(mentor.ID.String())
		//validasi timeout
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
			if mentorCourse.CourseID != course.ID {
				//print mentorCourse dengan CourseID yang tidak ada di table course
				continue
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
	output := outputs.GetPopularMentorOutput{
		BaseOutput: outputs.BaseOutput{
			Code:    200,
			Message: "Success: Mentor found",
		},
		Data: data,
	}
	return 200, output
}

func GetAllMentor() (int, interface{}) {
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
			if mentorCourse.CourseID != course.ID {
				//print mentorCourse dengan CourseID yang tidak ada di table course
				continue
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
	log.Println("in");

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
