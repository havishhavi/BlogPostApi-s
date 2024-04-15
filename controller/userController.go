package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"www.blog.com/dto"
)

type UserController struct{}

func (con UserController) Register(c *gin.Context) {

	var InputDTO dto.Register

	if errDTO := c.ShouldBindJSON(&InputDTO); errDTO != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errDTO.Error())
		return
	}

	c.JSON(200, gin.H{
		"message ": InputDTO,
	})
}
