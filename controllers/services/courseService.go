package services

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/meja_belajar/controllers/helpers"
	"github.com/meja_belajar/models/outputs"
	"github.com/meja_belajar/models/requests"
)

// Get Course by ID
func GetCourse(c *gin.Context) {
	courseID := c.Param("id")

	if _, err := uuid.Parse(courseID); err != nil {
		output := outputs.BadRequestOutput{
			Code:    400,
			Message: "Bad Request: " + err.Error(),
		}
		c.JSON(http.StatusBadRequest, output)
		return
	}
	requestDTO := requests.GetCourseRequestDTO{CourseID: courseID}
	// Call the helper function to get the mentor review
	code, output := helpers.GetCourse(requestDTO)

	c.JSON(code, output)
}

func AddCourse(c *gin.Context) {
	var AddCourseRequestDTO requests.AddCourseRequestDTO

	if err := c.ShouldBindJSON(&AddCourseRequestDTO); err != nil {
		outputs := outputs.BadRequestOutput{
			Code:    400,
			Message: "Bad Request: " + err.Error(),
		}
		c.JSON(http.StatusBadRequest, outputs)
		return
	}
	code, output := helpers.AddCourse(AddCourseRequestDTO)
	c.JSON(code, output)
}

func UpdateCourse(c *gin.Context) {
	var UpdateCourseRequestDTO requests.UpdateCourseRequestDTO

	if err := c.ShouldBindJSON(&UpdateCourseRequestDTO); err != nil {
		outputs := outputs.BadRequestOutput{
			Code:    400,
			Message: "Bad Request: " + err.Error(),
		}
		c.JSON(http.StatusBadRequest, outputs)
		return
	}
	code, output := helpers.UpdateCourse(UpdateCourseRequestDTO)
	c.JSON(code, output)
}

func CourseService(router *gin.RouterGroup){
	router.GET("/course/:id", GetCourse)
	router.POST("/course/create", AddCourse)
	router.POST("/course/update", UpdateCourse)
}