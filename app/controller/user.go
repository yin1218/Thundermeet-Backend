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

	// "strings"
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

func UpdateUsersController() UsersController {
	return UsersController{}
}

type Login struct {
	User_id  string `json:"userId" binding:"required" example:"christine891225"`
	Password string `json:"password" binding:"required" example:"password"`
} // @name Login

// LoginUser LoginUser @Summary
// @Tags user
// @version 1.0
// @produce application/json
// @Param Body body Login true "The body to login a user"
// @Success 200 string string successful return data
// @Failure 500 string string ErrorResponse
// @Router /v1/users/login/ [post]
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

	userId, passwordHash, err := service.GetOneUserUsernamePasswordHash(form.User_id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("user %s not found", form.User_id),
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
	User_id         string `json:"userId" binding:"required" example:"christine891225"`
	User_name       string `json:"userName" example:"Christine Wang"`
	Password        string `json:"password" binding:"required" example:"password"`
	Password_answer string `json:"passwordAnswer" binding:"required" example:"NTU"`
} // @name Register

type Update struct {
	User_name       string `json:"userName" example:"Christine Wang"`
	Password        string `json:"password" example:"password"`
	Password_answer string `json:"passwordAnswer" example:"NTU"`
} //@name Update

type ForgotInfo struct {
	User_id         string `json:"userId" binding:"required" example:"christine891225"`
	Password        string `json:"password" binding:"required" example:"password"`
	Password_answer string `json:"passwordAnswer" binding:"required" example:"NTU"`
}

// CreateUser CreateUser @Summary
// @Tags user
// @version 1.0
// @produce application/json
// @Param Body body Register true "The body to create a user"
// @Success 200 string string successful return data
// @Failure 500 string string ErrorResponse
// @Router /v1/users/ [post]
func (u UsersController) CreateUser(c *gin.Context) {
	var form Register
	bindErr := c.BindJSON(&form)
	if bindErr == nil {

		err := service.RegisterOneUser(form.User_id, form.User_name, form.Password, form.Password_answer)
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

// CheckUser CheckUser @Summary
// @Tags user
// @version 1.0
// @produce application/json
// @Param Authorization header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Success 200 string string successful return data
// @Failure 500 string string ErrorResponse
// @Router /v1/users/ [get]
func (u UsersController) CheckUser(c *gin.Context) {

	//get and check token format
	token := c.Request.Header.Get("Authorization")

	//validate token
	id, err := jwt.ValidateToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	//check id and return needed data
	userOne, err := service.SelectOneUser(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": -1,
			"msg":    "User not found : " + err.Error(),
			"data":   nil,
		})
	} else {

		c.JSON(http.StatusOK, gin.H{
			"status":          0,
			"user_id":         userOne.UserId,
			"username":        userOne.UserName,
			"password_answer": userOne.PasswordAnswer,
		})
	}
}

// UpdateUserInfo UpdateUserInfo @Summary
// @Tags user
// @version 1.0
// @produce application/json
// @Param Authorization header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param Body body Update true "The body to create a user"
// @Success 200 string string successful return data
// @Failure 401 string string ErrorResponse
// @Failure 500 string string ErrorResponse
// @Router /v1/users [patch]
func (u UsersController) UpdateUserInfo(c *gin.Context) {
	//validate jwt token
	token := c.Request.Header.Get("Authorization")
	userId, err := jwt.ValidateToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	//update user info through user id
	var form Update
	bindErr := c.BindJSON(&form)
	if bindErr == nil {
		err := service.UpdateOneUser(userId, form.User_name, form.Password, form.Password_answer)
		if err == nil {
			fmt.Println("Successfully update info")
			c.JSON(http.StatusOK, gin.H{
				"status": 1,
				"msg":    "success Update",
				"data":   nil,
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": -1,
				"msg":    "Update Failed : " + err.Error(),
				"data":   nil,
			})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": -1,
			"msg":    "Fail to Update User Info" + bindErr.Error(),
			"data":   nil,
		})
	}
	//Question: Should we add "Else" statement? I can't imagine which situation would lead to this result

}

// ResetPassword ResetPassword @Summary
// @Tags user
// @version 1.0
// @produce application/json
// @Param Body body ForgotInfo true "The body to create a user"
// @Success 200 string string successful return data
// @Failure 500 string string ErrorResponse
// @Router /v1/users/resetPassword [patch]
func (u UsersController) ResetPassword(c *gin.Context) {
	var form ForgotInfo
	bindErr := c.BindJSON(&form)
	if bindErr == nil {
		err := service.ResetUserPassword(form.User_id, form.Password, form.Password_answer)
		if err == nil {
			fmt.Println("Successfully reset password")
			c.JSON(http.StatusOK, gin.H{
				"status": 1,
				"msg":    "success Update",
				"data":   nil,
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": -1,
				"msg":    "Password Reset Failed : " + err.Error(),
				"data":   nil,
			})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": -1,
			"msg":    "Failed to reset password : " + bindErr.Error(),
			"data":   nil,
		})
	}
}
