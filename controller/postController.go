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

// View Post by post id
/*
create a Inputdto variable and bind uri to that variable to get the data
check for errors
extract the post data by id and display
*/

func (con PostController) ViewPost(c *gin.Context) {
	var inputDTO dto.ViewPost
	//ShouldBindUri binds the passed struct pointer

	errDTO := c.ShouldBindUri(&inputDTO)
	if errDTO != nil {
		//helper.ELog.Error(errDTO.Error())
		msg := handle.Error(errDTO)
		c.AbortWithStatusJSON(http.StatusBadRequest, msg)
		return
	}
	helper.Trimmer(&inputDTO)

	post_data, err := model.FindPostById(inputDTO.Postid)
	if err != nil {
		helper.ELog.Error(err.Error())
		response := helper.Error("SQL Error", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
		return

	}
	response := helper.Success(true, "ok", post_data)

	c.JSON(http.StatusOK, response)

}

//View All User Posts

/*
create a Inputdto variable and bind uri to that variable to get the data
check for errors
extract the post data by id and display
*/
func (con PostController) UserPosts(c *gin.Context) {
	user_id := service.GetUserId(c.GetHeader("Token"))
	allPosts, err := model.FindAllPostsByUserId(user_id)
	if err != nil {
		helper.ELog.Error(err.Error())
		response := helper.Error("SQL Error", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.Success(true, "ok", allPosts)
	c.JSON(http.StatusOK, response)

}

// viewall the users posts from loggedin user
func (con PostController) AllUserPosts(c *gin.Context) {
	allPosts, err := model.FindallPostData()
	if err != nil {
		helper.ELog.Error(err.Error())
		response := helper.Error("SQL Error", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.Success(true, "ok", allPosts)
	c.JSON(http.StatusOK, response)

}
