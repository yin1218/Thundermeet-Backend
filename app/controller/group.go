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

	// "thundermeet_backend/app/middleware/crypto"
	"thundermeet_backend/app/middleware/jwt"
	"thundermeet_backend/app/service"

	"github.com/gin-gonic/gin"
)

type GroupController struct{}

type createGroupFormat struct {
	Group_name string `json:"group_name" example:"OR-related" binding:"required"` //required
	Event_ids  []int  `json:"event_ids" example:[7,8,9] binding:"required"`       //required
} //@name createGroupFormat

// CreateGroup CreateGroup @Summary
// @Tags group
// @version 1.0
// @produce application/json
// @Param Authorization header string true "Bearer eyJNoCkSqDCEVLw0xRO8CzTg"
// @Param Body body createGroupFormat true "The body to create an group"
// @Success 201 string string successful return data
// @Failure 500 string string ErrorResponse
// @Router /v1/groups/ [post]
func CreateGroupsController() GroupController {
	return GroupController{}
}

func (u GroupController) CreateGroup(c *gin.Context) {
	//===== validate token ============//
	token := c.Request.Header.Get("Authorization")
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
	}

	//parse request body
	var form createGroupFormat
	bindErr := c.BindJSON(&form)
	if bindErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": -1,
			"msg":    "invalid input : " + bindErr.Error(),
			"data":   nil,
		})
		return
	}

	// form.Group_ids : array
	// form.Group_name

	//===== validate event_ids ============//
	var temp_event []int64
	for _, event_id := range form.Event_ids {
		fmt.Println(event_id)
		format_event_id := int64(event_id)
		// SelectOneEvent
		event, err := service.SelectOneEvent(format_event_id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": -1,
				"msg":    "invalid event : " + err.Error(),
				"data":   nil,
			})
			return
		}
		if event.AdminId != userOne.UserId {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": -1,
				"msg":    "The event doesn't belong to current user!",
				"data":   nil,
			})
			return
		}
		//add event into a list
		temp_event = append(temp_event, format_event_id)
	}

	//start adding process

	//add group: id, name, user_id
	group_id, createErr := service.CreateGroup(form.Group_name, userOne.UserId)
	if createErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": -1,
			"msg":    "fail to add group : " + createErr.Error(),
			"data":   nil,
		})
		return
	}

	// iterate slice : temp_event
	for _, event_id := range temp_event {
		createErr := service.AddEventToGroup(int(event_id), group_id)
		if createErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": -1,
				"msg":    "fail to add event into group : " + createErr.Error(),
				"data":   nil,
			})
			return
		}
	}

	//return 201 created
	c.JSON(http.StatusCreated, gin.H{
		"status": -1,
		"msg":    "Success to add group and init. event!",
		"data":   nil,
	})
	// return

}
