/*
最後則是這次新增的controller directory
這邊主要是用來撰寫API邏輯的地方
舉例來說一次的GET DATA API
我們會將其切割為routing, business logic, orm query三大部分
controller這層就是用來實現business logic的地方
*/
package controller

import (
	// "backend/app/service"

	"fmt"
	"net/http"
	"strconv"
	"thundermeet_backend/app/middleware/crypto"
	"thundermeet_backend/app/middleware/jwt"
	"thundermeet_backend/app/service"

	"github.com/gin-gonic/gin"
)

type UsersController struct{}

func NewUsersController() UsersController {
	return UsersController{}
}

func QueryUsersController() UsersController {
	return UsersController{}
}

type Login struct {
	User_name string `json:"userName" binding:"required" example:"christineWang"`
	Password  string `json:"password" binding:"required" example:"password"`
} // @name Login

// LoginUser LoginUser @Summary
// @Tags user
// @version 1.0
// @produce application/json
// @Param Body body Login true "The body to login a user"
// @Success 200 string string successful return data
// @Failure 500 string string ErrorResponse
// @Router /v1/users/login [post]
func (u UsersController) Login(c *gin.Context) {
	var form Login
	bindErr := c.BindJSON(&form)
	if bindErr == nil {

	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": -1,
			"msg":    "Failed to parse login data" + bindErr.Error(),
			"data":   nil,
		})
	}

	userId, passwordHash, err := service.GetOneUserUsernamePasswordHash(form.User_name)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("user %s not found", form.User_name),
		})
		return
	}
	// validate password
	passErr := crypto.Compare(passwordHash, form.Password)
	if passErr != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "incorrect password",
		})
		return
	}

	// generate jwt token
	token, err := jwt.GenToken(userId)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

type Register struct {
	User_name       string `json:"userName" binding:"required" example:"christineWang"`
	Password        string `json:"password" binding:"required" example:"password"`
	Password_answer string `json:"passwordAnswer" binding:"required" example:"NTU"`
} // @name Register

// CreateUser CreateUser @Summary
// @Tags user
// @version 1.0
// @produce application/json
// @Param Body body Register true "The body to create a user"
// @Success 200 string string successful return data
// @Failure 500 string string ErrorResponse
// @Router /v1/users [post]
func (u UsersController) CreateUser(c *gin.Context) {
	var form Register
	bindErr := c.BindJSON(&form)
	if bindErr == nil {

		err := service.RegisterOneUser(form.User_name, form.Password, form.Password_answer)
		if err == nil {
			fmt.Println("Good register!")
			c.JSON(http.StatusOK, gin.H{
				"status": 1,
				"msg":    "success Register",
				"data":   nil,
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": -1,
				"msg":    "Register Failed" + err.Error(),
				"data":   nil,
			})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": -1,
			"msg":    "Failed to parse register data" + bindErr.Error(),
			"data":   nil,
		})
	}
}

func (u UsersController) GetUser(c *gin.Context) {
	id := c.Params.ByName("ID")

	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": -1,
			"msg":    "Failed to parse params" + err.Error(),
			"data":   nil,
		})
	}

	userOne, err := service.SelectOneUser(userId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": -1,
			"msg":    "User not found" + err.Error(),
			"data":   nil,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": 0,
			"msg":    "Successfully get user data",
			"user":   &userOne,
		})
	}
}
