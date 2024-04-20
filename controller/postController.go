package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"www.blog.com/common/handle"
	"www.blog.com/common/helper"
	"www.blog.com/common/service"
	"www.blog.com/config"
	"www.blog.com/dto"
	"www.blog.com/model"
)

type PostController struct{}

/*
After login Create a post
first map dto objects and bind them with JSON and check for errors
trimmer func to trim the spaces in front
get userid from the jwt token to match the post and user
connect database and check for errors while creatinng post
return response : post created

*/

func (con PostController) CreatePost(c *gin.Context) {
	var InputDTO dto.CreatePost

	if errDTO := c.ShouldBindJSON(&InputDTO); errDTO != nil {
		msg := handle.Error(errDTO)
		c.AbortWithStatusJSON(http.StatusBadRequest, msg)
		return
	}
	helper.Trimmer(&InputDTO)

	user_id := service.GetUserId(c.GetHeader("Token"))
	var post model.Post
	post.Title = InputDTO.Title
	post.Post = InputDTO.Post
	post.UserID = user_id

	db := config.GoConnect()

	if result := db.Create(&post); result.Error != nil {
		helper.ELog.Error(result.Error.Error())
		response := helper.Error("SQL error", result.Error.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.Success(true, "ok", "Post Created Successfully")
	c.JSON(http.StatusOK, response)

}
