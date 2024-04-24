package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"www.blog.com/common/handle"
	"www.blog.com/common/helper"
	"www.blog.com/common/service"
	"www.blog.com/config"
	"www.blog.com/dto"
	"www.blog.com/model"
)

type UserController struct{}

/*
Register user

function: Register
*/
func (con UserController) Register(c *gin.Context) {

	var InputDTO dto.Register

	if errDTO := c.ShouldBindJSON(&InputDTO); errDTO != nil {
		msg := handle.Error(errDTO)
		c.AbortWithStatusJSON(http.StatusBadRequest, msg)
		return
	}
	// this Trimmer method removes the white spaces from the given details from the front

	helper.Trimmer(&InputDTO)

	//Finduserbyemail check if the email already exists
	result, err := model.FindUserByEmail(InputDTO.Email)
	if err != nil {
		helper.WLog.Warn(err.Error())
		response := helper.Error("SQL error", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if result != 0 {
		helper.ELog.Error("user already exist")
		response := helper.Error("User error", "user Already exist", helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// encrypting the password using Pwdencryprion method
	password, err := helper.PwdEncryption(InputDTO.Password)
	helper.ELog.Error("encripton error")
	if err != nil {
		response := helper.Error("encryption error", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadGateway, response)
		return
	}

	// convert string to int64
	mobile_no, err := helper.ConvertStoI(InputDTO.Mobile)
	if err != nil {
		helper.ELog.Error(err.Error())
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
		helper.ELog.Error(result.Error.Error())
		response := helper.Error("Sql Error", result.Error.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.Success(true, "ok", user)

	c.JSON(http.StatusCreated, response)
}

/*
Login user

function: login
*/

func (con UserController) Login(c *gin.Context) {
	/*
	   Login user
	   Operations: Declare a variable for data Transfer Object
	   bind json data to variable
	   get user data using email and handle errors
	   compare hash and password (database password and received password)
	   generate JWT token
	   return json response
	   function: login
	*/

	var InputDTO dto.Login

	//binding data from json to inputdto
	if errDTO := c.ShouldBindJSON(&InputDTO); errDTO != nil {
		msg := handle.Error(errDTO)
		c.AbortWithStatusJSON(http.StatusBadRequest, msg)
		return
	}
	//remove extra spaces from the front of given values
	helper.Trimmer(&InputDTO)

	data, err := model.FindUserDataByEmail(InputDTO.Email)
	if err != nil {
		helper.ELog.Error(err.Error())
		response := helper.Error("SQL Error", err.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
		return
	}
	if data == nil {
		response := helper.Error("msg", "User not found", helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(InputDTO.Password)); err != nil {
		response := helper.Error("Invalid Password", "password is incorrect", helper.EmptyObj{})
		helper.ILog.Info(err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	//Generating JWT token

	JwtToken := service.GenerateToken(strconv.Itoa(int(data.ID)), 1)

	if _, err := model.UpdateToken(data.ID, JwtToken); err != nil {
		msg := helper.Error("SQL Error", err.Error(), helper.EmptyObj{})
		helper.ELog.Error(err.Error())
		c.JSON(http.StatusBadRequest, msg)
		return
	}

	data.Password = ""
	data.JwtToken = JwtToken

	response := helper.Success(true, "ok", data)
	c.JSON(http.StatusOK, response)

}

func (con UserController) Logout(c *gin.Context) {
	userid := service.GetUserId(c.GetHeader("Token"))
	db := config.GoConnect()
	if result := db.Model(&model.User{}).Where("id = ?", userid).Update("JwtToken", nil); result.Error != nil {
		helper.ELog.Error(result.Error.Error())
		response := helper.Error("Sql Error", result.Error.Error(), helper.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.Success(true, "ok", "Logout Successfull")
	c.JSON(http.StatusOK, response)

}
