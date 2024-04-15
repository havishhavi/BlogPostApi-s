package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"www.blog.com/common/handle"
	"www.blog.com/common/helper"
	"www.blog.com/config"
	"www.blog.com/dto"
	"www.blog.com/model"
)

type UserController struct{}

func (con UserController) Register(c *gin.Context) {

	var InputDTO dto.Register

	if errDTO := c.ShouldBindJSON(&InputDTO); errDTO != nil {
		msg := handle.Error(errDTO)
		c.AbortWithStatusJSON(http.StatusBadRequest, msg)
		return
	}
	// this Trimmer method removes the white spacxe from the given details from the front
	helper.Trimmer(&InputDTO)

	//Finduserbyemail check if the email already exists
	result, err := model.FindUserByEmail(InputDTO.Email)
	if err != nil {
		response := helper.Error("SQL error", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if result != 0 {
		//helper.Elog.Error("user already exist")
		response := helper.Error("User error", "user Already exist", helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
		return
	}
	// encrypting the password using Pwdencryprion method
	password, err := helper.PwdEncryption(InputDTO.Password)
	if err != nil {
		response := helper.Error("encryption error", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadGateway, response)
		return
	}

	// convert string to int64
	mobile_no, err := helper.ConvertStoI(InputDTO.Mobile)
	if err != nil {
		response := helper.Error("invalid mobile no", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// create ba variable to store the inputdto verified variable to the database

	var user model.User
	//for every siingle variable that we checked the error
	user.Name = InputDTO.Name
	user.Email = InputDTO.Email
	user.Mobile = mobile_no
	user.Password = password

	//database conn
	db := config.GoConnect()
	if result := db.Create(&user); result.Error != nil {
		//helper.Elog.Error(result.Error.Error())
		response := helper.Error("Sql Error", result.Error.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.Success(true, "ok", user)

	c.JSON(http.StatusCreated, response)
}
