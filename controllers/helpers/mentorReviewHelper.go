package helpers

import (
	"context"

	"github.com/google/uuid"
	"github.com/meja_belajar/configs"
	"github.com/meja_belajar/models/database"
	"github.com/meja_belajar/models/outputs"
	"github.com/meja_belajar/models/requests"
	"github.com/meja_belajar/models/responses"
	"github.com/meja_belajar/utils"
)

// Get Mentor Review by ID
func GetMentorReview(GetMentorReviewsRequestDTO requests.GetMentorReviewsRequestDTO) (int, interface{}){	
	db := configs.GetDB()
	var mentorReview database.MentorReviews

	err := db.First(&mentorReview, "id = ?", utils.StringToUUID(GetMentorReviewsRequestDTO.ID)).Error

	if err != nil {
		return 500, err.Error()
	}

	if mentorReview.ID == uuid.Nil {
		return 404, "Mentor Review not found"
	}
	
	if err == context.DeadlineExceeded {
		output := outputs.RequestTimeoutOutput{
			Code:    408,
			Message: "Request Timeout",
		}
		return 408, output
	}

	output := outputs.GetMentorReviewsOutput{}
	output.Code = 200
	output.Message = "Success"
	output.Data = responses.MentorReviewsResponseDTO{
		ID:          mentorReview.ID,
		MentorID:    mentorReview.MentorID,
		Description: mentorReview.Description,
		IsActive:    mentorReview.IsActive,
		CreatedBy:   mentorReview.CreatedBy,	
		UpdatedBy:   mentorReview.UpdatedBy,
		CreatedAt:   mentorReview.CreatedAt,
		UpdatedAt:   mentorReview.UpdatedAt,
	}
	return 200, output		
}

// Create Mentor Reviews
func CreateMentorReview(CreateMentorReviewRequestDTO requests.CreateMentorReviewRequestDTO) (int, interface{}) {
	db := configs.GetDB()
	mentorReview := database.MentorReviews{
		MentorID:    CreateMentorReviewRequestDTO.MentorID,
		Description: CreateMentorReviewRequestDTO.Description,
		IsActive:    CreateMentorReviewRequestDTO.IsActive,
		CreatedBy:   CreateMentorReviewRequestDTO.CreatedBy,
	}

	err := db.Create(&mentorReview).Error

	if err != nil {
		return 500, err.Error()
	}

	output := outputs.CreateMentorReviewRequestDTO{}
	output.Code = 200
	output.Message = "Success"
	output.Data = responses.MentorReviewsResponseDTO{
		ID:          mentorReview.ID,
		MentorID:    mentorReview.MentorID,
		Description: mentorReview.Description,
		IsActive:    mentorReview.IsActive,
		CreatedBy:   mentorReview.CreatedBy,
		UpdatedBy:   mentorReview.UpdatedBy,
		CreatedAt:   mentorReview.CreatedAt,
		UpdatedAt:   mentorReview.UpdatedAt,
	}
	return 201, output
}

// Update Mentor Reviews
func UpdateMentorReview(UpdateMentorReviewRequestDTO requests.UpdateMentorReviewRequestDTO) (int, interface{}) {
	db := configs.GetDB()
	var mentorReview database.MentorReviews

	err := db.First(&mentorReview, "id = ?", utils.StringToUUID(UpdateMentorReviewRequestDTO.ID)).Error

	if err != nil {
		return 500, err.Error()
	}

	if mentorReview.ID == uuid.Nil {
		return 404, "Mentor Review not found"
	}

	mentorReview.MentorID = utils.StringToUUID(UpdateMentorReviewRequestDTO.MentorID)
	mentorReview.Description = UpdateMentorReviewRequestDTO.Description
	mentorReview.IsActive = UpdateMentorReviewRequestDTO.IsActive
	mentorReview.UpdatedBy = UpdateMentorReviewRequestDTO.UpdatedBy

	err = db.Save(&mentorReview).Error

	if err != nil {
		return 500, err.Error()
	}

	output := outputs.UpdateMentorReviewRequestDTO{}
	output.Code = 200
	output.Message = "Success"
	output.Data = responses.MentorReviewsResponseDTO{
		ID:          mentorReview.ID,
		MentorID:    mentorReview.MentorID,
		Description: mentorReview.Description,
		IsActive:    mentorReview.IsActive,
		CreatedBy:   mentorReview.CreatedBy,
		UpdatedBy:   mentorReview.UpdatedBy,
		CreatedAt:   mentorReview.CreatedAt,
		UpdatedAt:   mentorReview.UpdatedAt,
	}
	return 200, output
}