package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/meja_belajar/controllers/helpers"
	"github.com/meja_belajar/models/outputs"
	"github.com/meja_belajar/models/requests"
)

// Get Mentor Review by ID
func GetMentorReview(c *gin.Context) {
	// var GetMentorReviewsRequestDTO requests.GetMentorReviewsRequestDTO
	// //validasi format input
	// if err := c.ShouldBindJSON(&GetMentorReviewsRequestDTO); err != nil {
	// 	outputs := outputs.BadRequestOutput{
	// 		Code:    400,
	// 		Message: "Bad Request: " + err.Error(),
	// 	}
	
	reviewID := c.Param("id")

	if _, err := uuid.Parse(reviewID); err != nil {
		output := outputs.BadRequestOutput{
			Code:    400,
			Message: "Bad Request: " + err.Error(),
		}
		c.JSON(http.StatusBadRequest, output)
		return
	}
	// code, output := helpers.GetMentorReview(requests.GetMentorReviewsRequestDTO{ID: reviewID})
	requestDTO := requests.GetMentorReviewsRequestDTO{ID: reviewID}
	// Call the helper function to get the mentor review
	code, output := helpers.GetMentorReview(requestDTO)

	c.JSON(code, output)
}

func CreateMentorReview(c *gin.Context) {
	var CreateMentorReviewRequestDTO requests.CreateMentorReviewRequestDTO
	//validasi format input
	if err := c.ShouldBindJSON(&CreateMentorReviewRequestDTO); err != nil {
		outputs := outputs.BadRequestOutput{
			Code:    400,
			Message: "Bad Request: " + err.Error(),
		}
		c.JSON(http.StatusBadRequest, outputs)
		return
	}
	code, output := helpers.CreateMentorReview(CreateMentorReviewRequestDTO)
	c.JSON(code, output)
}

func UpdateMentorReview(c *gin.Context) {
	var UpdateMentorReviewRequestDTO requests.UpdateMentorReviewRequestDTO
	//validasi format input
	if err := c.ShouldBindJSON(&UpdateMentorReviewRequestDTO); err != nil {
		outputs := outputs.BadRequestOutput{
			Code:    400,
			Message: "Bad Request: " + err.Error(),
		}
		c.JSON(http.StatusBadRequest, outputs)
		return
	}
	code, output := helpers.UpdateMentorReview(UpdateMentorReviewRequestDTO)
	c.JSON(code, output)
}

func MentorReviewService(router *gin.RouterGroup){
	router.GET("/mentor-review/:id", GetMentorReview)
	router.POST("/mentor-review/create", CreateMentorReview)
	router.POST("/mentor-review/update", UpdateMentorReview)
}