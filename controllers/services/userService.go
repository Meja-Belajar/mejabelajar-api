package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
	c.SetCookie("Authorization", tokenString, 3600*24, "", "", true, true)
	c.JSON(code, output)
}

func UserService(router *gin.RouterGroup) {
	router.POST("/user/register", RegisterUser)
	router.POST("/user/login", LoginUser)
}
