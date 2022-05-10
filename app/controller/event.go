/*
最後則是這次新增的controller directory
這邊主要是用來撰寫API邏輯的地方
舉例來說一次的GET DATA API
我們會將其切割為routing, business logic, orm query三大部分
controller這層就是用來實現business logic的地方
*/
package controller

import (
	"fmt"
	"net/http"
	"time"

	// "thundermeet_backend/app/middleware/crypto"
	"thundermeet_backend/app/middleware/jwt"
	"thundermeet_backend/app/service"

	"github.com/gin-gonic/gin"
)

type EventController struct{}

type createEventFormat struct {
	event_name          string `json:"eventName" example:"OR first meet" binding:"required"`            //required
	is_priority_enabled bool   `json:"isPriorityEnabled" example:true binding:"required"`               //required
	start_time          string `json:"startTime" example:"1975-08-19T11:00:00.000Z" binding:"required"` //required
	end_time            string `json:"endTime" example:"1975-08-19T23:00:00.000Z" binding:"required"`   //required
	date_or_days        bool   `json:"dateOrDays" example:true binding:"required"`                      //required
	start_day           string `json:"startDay" example:1 `                                             //optional
	end_day             string `json:"endDay" example:7`                                                //optional
	start_date          string `json:"startDate" example:"2021-01-01T11:00:00.000Z"`                    //optional
	end_date            string `json:"endDate" example:"2021-01-10T11:00:00.000Z"`                      //optional
} //@name EventFormat

func CreateEventsController() EventController {
	return EventController{}
}

func (u EventController) CreateEvent(c *gin.Context) {

	token := c.Request.Header.Get("Authorization")
	//validate token
	id, err := jwt.ValidateToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	//check id and return needed data -> change to finduser
	userOne, err := service.SelectOneUser(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": -1,
			"msg":    "User not found : " + err.Error(),
			"data":   nil,
		})
		return
	}

	// 利用userOne抓userId -> 作為admin id
	var adminId string = userOne.UserId

	//get request body
	var form createEventFormat
	bindErr := c.BindJSON(&form)

	if bindErr == nil {
		//================== check the correctness of input value and format ==================
		//change time format
		start_time, err := time.Parse(time.RFC3339, form.start_time)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": -1,
				"msg":    "invalid start time",
				"data":   nil,
			})
			return
		}
		end_time, err := time.Parse(time.RFC3339, form.end_time)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": -1,
				"msg":    "invalid end time",
				"data":   nil,
			})
			return
		}

		//check if time period is valid
		if end_time.Before(start_time) {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": -1,
				"msg":    "invalid time period",
				"data":   nil,
			})
			return
		}

		// check if needed information is in the json: start/end day || start/end date
		if form.date_or_days == true {
			if form.start_date == "" || form.end_date == "" {
				c.JSON(http.StatusBadRequest, gin.H{
					"status": -1,
					"msg":    "request should contain both start_date and end_date",
					"data":   nil,
				})
				return
			}
			// change time format
			start_date, err := time.Parse(time.RFC3339, form.start_date)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status": -1,
					"msg":    "invalid start date format" + err.Error(),
					"data":   nil,
				})
				return
			}
			end_date, err := time.Parse(time.RFC3339, form.end_date)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status": -1,
					"msg":    "invalid end date format : " + err.Error(),
					"data":   nil,
				})
				return
			}
			//check if period is correct
			if end_date.Before(start_date) {
				c.JSON(http.StatusBadRequest, gin.H{
					"status": -1,
					"msg":    "Invalid date period",
					"data":   nil,
				})
				return
			}
		} else {
			if form.start_day == "" || form.end_day == "" {
				c.JSON(http.StatusBadRequest, gin.H{
					"status": -1,
					"msg":    "request should contain both start_day and end_day",
					"data":   nil,
				})
				return
			}

		}

		//================ complete checking process, start to add things to db ================
		err := service.createEvent()
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
			"msg":    "Failed to parse event data : " + bindErr.Error(),
			"data":   nil,
		})
	}

	//check if the user is the same as the given userId
	// if userOne.UserId !=
}
