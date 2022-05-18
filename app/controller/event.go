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
	"strconv"
	"time"

	// "thundermeet_backend/app/middleware/crypto"
	"thundermeet_backend/app/middleware/jwt"
	"thundermeet_backend/app/service"

	"github.com/gin-gonic/gin"
)

type EventController struct{}

type createEventFormat struct {
	Event_name          string `json:"eventName" example:"OR first meet" binding:"required"` //required
	Is_priority_enabled *bool  `json:"isPriorityEnabled" example:"true" binding:"required"`  //required
	Start_time          string `json:"startTime" example:"0900" binding:"required"`          //required
	End_time            string `json:"endTime" example:"2100" binding:"required"`            //required
	Date_or_days        *bool  `json:"dateOrDays" example:"true" binding:"required"`         //required
	Start_day           string `json:"startDay" example:"1" `                                //optional
	End_day             string `json:"endDay" example:"7"`                                   //optional
	Start_date          string `json:"startDate" example:"2021-01-01T12:00:00.000Z"`         //optional
	End_date            string `json:"endDate" example:"2021-01-02T12:00:00.000Z"`           //optional
	Event_description   string `json:"eventDescription" example:"description of event"`      //optional
} //@name EventFormat

func CreateEventsController() EventController {
	return EventController{}
}

func GetEventsController() EventController {
	return EventController{}
}

func UpdateEventsController() EventController {
	return EventController{}
}

func isValidTime(startTime string, endTIme string) bool {
	s, _ := strconv.Atoi(startTime[0:2])
	t, _ := strconv.Atoi(endTIme[0:2])
	if s < t {
		return true
	} else {
		smin, _ := strconv.Atoi(startTime[2:4])
		tmin, _ := strconv.Atoi(endTIme[2:4])
		if smin < tmin {
			return true
		}
		return false
	}
}

// CreateEvent CreateEvent @Summary
// @Tags event
// @version 1.0
// @produce application/json
// @Param Authorization header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param Body body createEventFormat true "The body to create an event"
// @Success 200 string string successful return data
// @Failure 500 string string ErrorResponse
// @Router /v1/events/ [post]
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
	// fmt.Println(userOne)

	// 利用userOne抓userId -> 作為admin id
	var adminId string = userOne.UserId
	// fmt.Println(adminId)

	//get request body
	var form createEventFormat
	bindErr := c.BindJSON(&form)

	if bindErr == nil {
		//================== check the correctness of input value and format ==================
		// change time format
		// start_time, err := time.Parse(time.RFC3339, form.Start_time)
		// if err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{
		// 		"status": -1,
		// 		"msg":    "invalid start time",
		// 		"data":   nil,
		// 	})
		// 	return
		// }
		// end_time, err := time.Parse(time.RFC3339, form.End_time)
		// if err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{
		// 		"status": -1,
		// 		"msg":    "invalid end time",
		// 		"data":   nil,
		// 	})
		// 	return
		// }

		//check if time period is valid
		if !isValidTime(form.Start_time, form.End_time) {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": -1,
				"msg":    "invalid time period",
				"data":   nil,
			})
			return
		}

		var start_date time.Time = time.Now()
		var end_date time.Time = time.Now()
		var start_day string = ""
		var end_day string = ""

		// fmt.Println("start date = ", start_date)
		// fmt.Println("start date = ", end_date)

		// check if needed information is in the json: start/end day || start/end date
		if *form.Date_or_days {
			if form.Start_date == "" || form.End_date == "" {
				c.JSON(http.StatusBadRequest, gin.H{
					"status": -1,
					"msg":    "request should contain both start_date and end_date",
					"data":   nil,
				})
				return
			}
			// change time format
			temp_start_date, err := time.Parse(time.RFC3339, form.Start_date)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status": -1,
					"msg":    "invalid start date format" + err.Error(),
					"data":   nil,
				})
				return
			}
			temp_end_date, err := time.Parse(time.RFC3339, form.End_date)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status": -1,
					"msg":    "invalid end date format : " + err.Error(),
					"data":   nil,
				})
				return
			}
			//check if period is correct
			if temp_end_date.Before(temp_start_date) {
				c.JSON(http.StatusBadRequest, gin.H{
					"status": -1,
					"msg":    "Invalid date period",
					"data":   nil,
				})
				return
			}
			start_date = temp_start_date
			end_date = temp_end_date
		} else {
			if form.Start_day == "" || form.End_day == "" {
				c.JSON(http.StatusBadRequest, gin.H{
					"status": -1,
					"msg":    "request should contain both start_day and end_day",
					"data":   nil,
				})
				return
			}

			if form.End_day < form.Start_day {
				c.JSON(http.StatusBadRequest, gin.H{
					"status": -1,
					"msg":    "invalid weekday period",
					"data":   nil,
				})
				return
			}

			start_day = form.Start_day
			end_day = form.End_day

		}

		// fmt.Println("Successfully reset password")
		// c.JSON(http.StatusOK, gin.H{
		// 	"status": 1,
		// 	"msg":    "success Create Event",
		// 	"data":   nil,
		// })

		// return

		// ========== 到這邊檢查都沒有問題 ================

		// ================ complete checking process, start to add things to db ================

		fmt.Println("Before go into loop")
		event_id, createErr := service.CreateEvent(form.Event_name, form.Is_priority_enabled, form.Start_time, form.End_time, form.Date_or_days, start_day, end_day, start_date, end_date, adminId, form.Event_description)
		if createErr == nil {

			err := CreateManyTimeblocks(*form.Date_or_days, form.Start_time, form.End_time, start_date, end_date, event_id)
			if err == nil {
				c.JSON(http.StatusOK, gin.H{
					"status":   1,
					"msg":      "success Create Event",
					"event_id": event_id,
					"data":     nil,
				})
			} else {
				c.JSON(http.StatusBadRequest, gin.H{
					"status": -1,
					"msg":    "Event Create Failed : " + err.Error(),
					"data":   nil,
				})
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": -1,
				"msg":    "Event Create Failed : " + createErr.Error(),
				"data":   nil,
			})
			return
		}

	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": -1,
			"msg":    "Failed to parse event data : " + bindErr.Error(),
			"data":   nil,
		})
	}

}

// GetEvent GetEvent @Summary
// @Tags event
// @version 1.0
// @produce application/json
// @Param Authorization header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Success 200 {object} string
// @Failure 404 string string ErrorResponse
// @Failure 500 string string ErrorResponse
// @param event_id path int64 true "event id"
// @Router /v1/events/{event_id} [get]
func (u EventController) GetEvent(c *gin.Context) {

	//validate user
	token := c.Request.Header.Get("Authorization")
	//validate token
	user_id, err := jwt.ValidateToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	fmt.Println(user_id)

	//validate user of event

	event_id, err := strconv.ParseInt(c.Param("event_id"), 10, 64)
	fmt.Print(event_id)
	event, err := service.SelectOneEvent(event_id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": -1,
			"msg":    "The event id does not exist.",
			"data":   nil,
		})
		return
	} else if event.AdminId != user_id {
		c.JSON(http.StatusNotFound, gin.H{
			"status": -1,
			"msg":    "the event's admin is not the current user.",
			"data":   nil,
		})
		return

	} else {

		// get groups of the event

		groupList, err := service.SelectEventGroups(int(event.EventId))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"status": -1,
				"msg":    "Get group fail : " + err.Error(),
				"data":   nil,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":               0,
			"event_id":             event.EventId,
			"event_name":           event.EventName,
			"event_description":    event.EventDescription,
			"is_priority_enabled":  event.IsPriorityEnabled,
			"is_confirmed":         event.IsConfirmed,
			"confirmed_timeblocks": event.ConfirmedTimeblocks,
			"start_time":           event.StartTime,
			"end_time":             event.EndTime,
			"date_or_days":         event.DateOrDays,
			"start_day":            event.StartDay,
			"end_day":              event.EndDay,
			"start_date":           event.StartDate,
			"end_date":             event.EndDate,
			"admin_id":             event.AdminId,
			"groups":               groupList,
		})
	}
}

// GetEvent GetEvent @Summary
// @Tags event
// @version 1.0
// @produce application/json
// @Param Authorization header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Success 200 {object} string
// @Failure 404 string string ErrorResponse
// @Failure 500 string string ErrorResponse
// @Router /v1/events [get]
func (u EventController) GetEvents(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	//validate token
	user_id, err := jwt.ValidateToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	//check id and return needed data
	events, err := service.GetEventsByUser(user_id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": -1,
			"msg":    "Events not found : " + err.Error(),
			"data":   nil,
		})
	} else {
		c.JSON(http.StatusOK, events)
	}

}

type UpdateEventFormat struct {
	EventId             string   `json:"event_id" example:"26"`
	EventName           string   `json:"event_name" example:"Sad 2nd meeting"`
	ConfirmedTimeblocks []string `json:"confirmed_timeblocks" example:"2021-01-01T11:00:00.000Z"`
} //@name UpdateEventFormat

// UpdateEvent UpdateEvent @Summary
// @Tags event
// @version 1.0
// @produce application/json
// @Param Authorization header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param Body body UpdateEventFormat true "The body to update a event"
// @Success 200 string string successful return data
// @Failure 401 string string ErrorResponse
// @Failure 403 string string ErrorResponse
// @Failure 500 string string ErrorResponse
// @Router /v1/events [patch]
func (u EventController) UpdateEvent(c *gin.Context) {
	var form UpdateEventFormat
	bindErr := c.BindJSON(&form)
	if bindErr == nil {
		event_id, _ := strconv.ParseInt(form.EventId, 10, 64)
		fmt.Print("event-id = ", event_id)

		//check the existance of event_name
		_, SelectEventErr := service.SelectOneEvent(event_id)

		if SelectEventErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": -1,
				"msg":    "The event id does not exist : " + SelectEventErr.Error(),
				"data":   nil,
			})
			return
		}

		err := service.UpdateOneEvent(event_id, form.EventName, form.ConfirmedTimeblocks)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": -1,
				"msg":    "Fail to Update Event Info" + bindErr.Error(),
				"data":   nil,
			})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": 1,
				"msg":    "success Update",
				"data":   nil,
			})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": -1,
			"msg":    "Fail to Update Event Info" + bindErr.Error(),
			"data":   nil,
		})
	}
}
