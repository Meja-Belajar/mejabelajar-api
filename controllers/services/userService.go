package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/meja_belajar/controllers/helpers"
	"github.com/meja_belajar/models/outputs"
	"github.com/meja_belajar/models/requests"
)

func RegisterUser(c *gin.Context) {
	var RegisterUserRequestDTO requests.RegisterUserRequestDTO
	//validas format input
	if err := c.ShouldBindJSON(&RegisterUserRequestDTO); err != nil {
		outputs := outputs.BadRequestOutput{
			Code:    400,
			Message: "Bad Request: " + err.Error(),
		}
		c.JSON(http.StatusBadRequest, outputs)
		return
	}
	code, output := helpers.RegisterUser(RegisterUserRequestDTO)
	c.JSON(code, output)
}

func LoginUser(c *gin.Context) {
	var LoginUserRequestDTO requests.LoginUserRequestDTO
	//validasi format input
	if err := c.ShouldBindJSON(&LoginUserRequestDTO); err != nil {
		outputs := outputs.BadRequestOutput{
			Code:    400,
			Message: "Bad Request: " + err.Error(),
		}
		c.JSON(http.StatusBadRequest, outputs)
		return
	}
	code, output, tokenString := helpers.LoginUser(LoginUserRequestDTO)
	c.SetCookie("Authorization", tokenString, 3600*24, "/", "localhost", true, true)
	c.SetSameSite(http.StatusOK)
	c.JSON(code, output)
}

func GetUserByID(c *gin.Context) {
	userID := c.Param("id")
	//validasi userID merupakan uuid
	if _, err := uuid.Parse(userID); err != nil {
		output := outputs.BadRequestOutput{
			Code:    400,
			Message: "Bad Request: " + err.Error(),
		}
		c.JSON(http.StatusBadRequest, output)
		return
	}
	code, output := helpers.GetUserByID(userID)
	c.JSON(code, output)
}

func UpdateUser(c *gin.Context) {
	var UpdateUserRequestDTO requests.UpdateUserRequestDTO

	if err := c.ShouldBindJSON(&UpdateUserRequestDTO); err != nil {
		outputs := outputs.BadRequestOutput{
			Code:    400,
			Message: "Bad Request: " + err.Error(),
		}
		c.JSON(http.StatusBadRequest, outputs)
		return
	}
	code, output := helpers.UpdateUser(UpdateUserRequestDTO)
	c.JSON(code, output)
}

func UserServiceBasic(router *gin.RouterGroup) {
	router.POST("/users/register", RegisterUser)
	router.POST("/users/login", LoginUser)
}

func UserServiceAuth(router *gin.RouterGroup) {
	router.POST("/users/update", UpdateUser)
	router.GET("/users/:id", GetUserByID)
}

