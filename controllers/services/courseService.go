package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/meja_belajar/controllers/helpers"
	"github.com/meja_belajar/models/outputs"
	"github.com/meja_belajar/models/requests"
)

// Get Course by ID
func GetCourse(c *gin.Context) {
	var GetCourseRequestDTO requests.GetCourseRequestDTO
	
	if err := c.ShouldBindJSON(&GetCourseRequestDTO); err != nil {
		outputs := outputs.BadRequestOutput{
			Code:    400,
			Message: "Bad Request: " + err.Error(),
		}
		c.JSON(http.StatusBadRequest, outputs)
		return
	}
	code, output := helpers.GetCourse(GetCourseRequestDTO)
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
	router.POST("/course/:id", GetCourse)
	router.POST("/course/create", AddCourse)
	router.POST("/course/update", UpdateCourse)
}