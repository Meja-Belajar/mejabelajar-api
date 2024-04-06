package middlewares

import (
	"net/http"
	"github.com/meja_belajar/utils"
	"github.com/gin-gonic/gin"
	"github.com/meja_belajar/models/outputs"
)

func RequiredAuth() gin.HandlerFunc {
	return func (c *gin.Context){
		//get the cookie
		tokenString, err := c.Cookie("Authorization")
		if err != nil {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				outputs.UnauthorizedOutput{
					Code:    401,
					Message: "Unauthorized",
				},
			)
			return
		}
		is_valid, err := utils.ValidateJWTToken(tokenString)
		if err != nil || !is_valid {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				outputs.UnauthorizedOutput{
					Code:    401,
					Message: "Unauthorized",
				},
			)
			return
		}
		c.Next()
	}
}
