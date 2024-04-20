package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"www.blog.com/common/helper"
	"www.blog.com/common/service"
	"www.blog.com/model"
)

// AuthorizeJWT validates the token user given, return 401 if not valid
func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Token")
		if authHeader == "" {
			response := helper.Error("Failed to process request", "No token found", nil)
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		_, err := service.ValidateToken(authHeader)
		if err != nil {
			response := helper.Error("Token is not valid", err.Error(), nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}
		/*Handeling user logout and token hacking*/

		user_id := service.GetUserId(authHeader)

		user, err := model.FindUserByID(uint(user_id))
		if err != nil {
			helper.ELog.Error(err.Error())
			response := helper.Error("Sql error", err.Error(), helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}
		if user == nil {
			helper.WLog.Warn("Middleware:User not exists")
			response := helper.Error("User not exists", "Invalid user", helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}
		if user.JwtToken != authHeader {
			helper.ELog.Error("PANIC:attack on website")
			response := helper.Error("Invalid user", "You have to login again", helper.EmptyObj{})
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}
	}

}
