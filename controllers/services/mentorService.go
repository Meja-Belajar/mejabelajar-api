package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/meja_belajar/controllers/helpers"
	"github.com/meja_belajar/models/outputs"
)

func GetMentorByMentorID(c *gin.Context) {
	mentorID := c.Param("id")
	//validasi mentorID merupakan uuid
	if _, err := uuid.Parse(mentorID); err != nil {
		output := outputs.BadRequestOutput{
			Code:    400,
			Message: "Bad Request: " + err.Error(),
		}
		c.JSON(http.StatusBadRequest, output)
		return
	}
	code, output := helpers.GetMentorByMentorID(mentorID)

	c.JSON(code, output)
}

func MentorServiceAuth(router *gin.RouterGroup) {
	router.GET("/mentor/:id", GetMentorByMentorID)
}
