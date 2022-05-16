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

type addGroupEventFormat struct {
	Event_ids []int `json:"event_ids" example:[9,10,11] binding:"required"` //required
} //@name addGroupEventFormat

type patchGroupFormat struct {
	Group_name string `json:"group_name" example:"SAD-related" binding:"required"` //required
} //@name patchGroupFormat

// DeleteEventFromGroup DeleteEventFromGroup @Summary
// @Tags group
// @version 1.0
// @produce application/json
// @Param Authorization header string true "Bearer eyJhbGcikDCEVLw0xRO8CzTg"
// @Success 200 string string successful return data
// @Failure 500 string string ErrorResponse
// @param group_id path int64 true "5"
// @param event_id path int64 true "20"
// @Router /v1/groups/{group_id}/{event_id} [delete]
func DeleteEventFromGroupController() GroupController {
	return GroupController{}
}

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

// GetGroup GetGroup @Summary
// @Tags group
// @version 1.0
// @produce application/json
// @Param Authorization header string true "Bearer eyJhbGcikDCEVLw0xRO8CzTg"
// @Success 200 string string successful return data
// @Failure 500 string string ErrorResponse
// @param group_id path int64 true "7"
// @Router /v1/groups/{group_id} [get]
func GetGroupController() GroupController {
	return GroupController{}
}

// GetGroup GetGroup @Summary
// @Tags group
// @version 1.0
// @produce application/json
// @Param Authorization header string true "Bearer eyJhbGcikDCEVLw0xRO8CzTg"
// @Success 200 string string successful return data
// @Failure 500 string string ErrorResponse
// @Router /v1/groups/ [get]
func GetGroupListController() GroupController {
	return GroupController{}
}

// DeleteGroup DeleteGroup @Summary
// @Tags group
// @version 1.0
// @produce application/json
// @Param Authorization header string true "Bearer eyJhbGcikDCEVLw0xRO8CzTg"
// @Success 200 string string successful return data
// @Failure 500 string string ErrorResponse
// @param group_id path int64 true "7"
// @Router /v1/groups/{group_id} [delete]
func DeleteGroupController() GroupController {
	return GroupController{}
}

// AddEventsToGroup AddEventsToGroup @Summary
// @Tags group
// @version 1.0
// @produce application/json
// @Param Authorization header string true "Bearer eyJhbGcikDCEVLw0xRO8CzTg"
// @Param Body body addGroupEventFormat true "The body to change the group's name"
// @Success 200 string string successful return data
// @Failure 500 string string ErrorResponse
// @param group_id path int64 true "5"
// @Router /v1/groups/{group_id} [post]
func AddEventsToGroupController() GroupController {
	return GroupController{}
}

// ChangeGroupName ChangeGroupName @Summary
// @Tags group
// @version 1.0
// @produce application/json
// @Param Authorization header string true "Bearer eyJhbGcikDCEVLw0xRO8CzTg"
// @Param Body body patchGroupFormat true "The body to change the group's name"
// @Success 200 string string successful return data
// @Failure 500 string string ErrorResponse
// @param group_id path int64 true "5"
// @Router /v1/groups/{group_id} [patch]
func ReviseGroupController() GroupController {
	return GroupController{}
}

func (u GroupController) DeleteEventsFromGroup(c *gin.Context) {
	//validate token and group owner
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
		return
	}

	group_id, err := strconv.ParseInt(c.Param("group_id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": -1,
			"msg":    "Fail to parse group id : " + err.Error(),
			"data":   nil,
		})
		return
	}

	//Check Group Admin
	groupOne, err := service.SelectOneGroup(int(group_id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": -1,
			"msg":    "Group not found : " + err.Error(),
			"data":   nil,
		})
		return
	}

	if groupOne.UserId != userOne.UserId {
		c.JSON(http.StatusNotFound, gin.H{
			"status": -1,
			"msg":    "This group doesn't belong to the user",
			"data":   nil,
		})
		return
	}

	event_id, err := strconv.ParseInt(c.Param("event_id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": -1,
			"msg":    "Fail to parse event id : " + err.Error(),
			"data":   nil,
		})
		return
	}

	//delete
	deleteErr := service.DeleteGroupEvent(int(event_id), int(group_id))
	if deleteErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": -1,
			"msg":    "Can't delete event from group : " + deleteErr.Error(),
			"data":   nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"msg":    "event deleted successfully from group",
		"data":   nil,
	})
	return

}

func (u GroupController) AddEventsToGroup(c *gin.Context) {
	//validate token and group owner
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
		return
	}

	group_id, err := strconv.ParseInt(c.Param("group_id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": -1,
			"msg":    "Fail to parse group id : " + err.Error(),
			"data":   nil,
		})
		return
	}

	//Check Group Admin
	groupOne, err := service.SelectOneGroup(int(group_id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": -1,
			"msg":    "Group not found : " + err.Error(),
			"data":   nil,
		})
		return
	}

	if groupOne.UserId != userOne.UserId {
		c.JSON(http.StatusNotFound, gin.H{
			"status": -1,
			"msg":    "This group doesn't belong to the user : ",
			"data":   nil,
		})
		return
	}

	//get datas
	var form addGroupEventFormat
	bindErr := c.BindJSON(&form)
	if bindErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": -1,
			"msg":    "invalid input : " + bindErr.Error(),
			"data":   nil,
		})
		return
	}

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

	// iterate slice : temp_event
	for _, event_id := range temp_event {
		createErr := service.AddEventToGroup(int(event_id), int(group_id))
		if createErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": -1,
				"msg":    "fail to add event into group : " + createErr.Error(),
				"data":   nil,
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"msg":    "events added to group!",
		"data":   nil,
	})

}

func (u GroupController) ChangeGroupName(c *gin.Context) {
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
		return
	}

	group_id, err := strconv.ParseInt(c.Param("group_id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": -1,
			"msg":    "Fail to parse group id : " + err.Error(),
			"data":   nil,
		})
		return
	}

	//Check Group Admin
	groupOne, err := service.SelectOneGroup(int(group_id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": -1,
			"msg":    "Group not found : " + err.Error(),
			"data":   nil,
		})
		return
	}

	if groupOne.UserId != userOne.UserId {
		c.JSON(http.StatusNotFound, gin.H{
			"status": -1,
			"msg":    "This group doesn't belong to the user : ",
			"data":   nil,
		})
		return
	}

	//get group new name
	//parse request body
	var form patchGroupFormat
	bindErr := c.BindJSON(&form)
	if bindErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": -1,
			"msg":    "invalid input : " + bindErr.Error(),
			"data":   nil,
		})
		return
	}

	// form.Group_name
	changeErr := service.ChangeGroupName(int(group_id), form.Group_name)
	if changeErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": -1,
			"msg":    "Fail to change group name : " + changeErr.Error(),
			"data":   nil,
		})
		return
	}

	//return 200
	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"msg":    "Success to change group name!",
		"data":   nil,
	})

}

func (u GroupController) GetGroupList(c *gin.Context) {
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
		return
	}

	groupList, err := service.SelectGroups(userOne.UserId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": -1,
			"msg":    "fail to get groups : " + err.Error(),
			"data":   nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"data":   groupList,
	})

}

func (u GroupController) GetGroup(c *gin.Context) {
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
		return
	}

	group_id, err := strconv.ParseInt(c.Param("group_id"), 10, 64)
	fmt.Println(c.Param("event_id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": -1,
			"msg":    "Group id error : " + err.Error(),
			"data":   nil,
		})
		return
	}

	groupOne, err := service.SelectOneGroup(int(group_id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": -1,
			"msg":    "Group not found : " + err.Error(),
			"data":   nil,
		})
		return
	}

	if groupOne.UserId != userOne.UserId {
		c.JSON(http.StatusNotFound, gin.H{
			"status": -1,
			"msg":    "This group doesn't belong to the user : ",
			"data":   nil,
		})
		return
	}

	// var eventList []int
	eventList, err := service.SelectGroupEvents(groupOne.GroupId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": -1,
			"msg":    "Get event fail : " + err.Error(),
			"data":   nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":     0,
		"group_name": groupOne.GroupName,
		"event_ids":  eventList,
	})
	return

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

	//check group name
	GroupNamrErr := service.GroupNameNotExist(form.Group_name)
	if GroupNamrErr {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": -1,
			"msg":    "Group Name Already Exist!",
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
		"status":   0,
		"msg":      "Success to add group and init. event!",
		"group_id": group_id,
		"data":     nil,
	})
}

func (u GroupController) DeleteGroup(c *gin.Context) {
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
		return
	}

	group_id, err := strconv.ParseInt(c.Param("group_id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": -1,
			"msg":    "Fail to parse group id : " + err.Error(),
			"data":   nil,
		})
		return
	}

	//delete event group relationship
	err = service.DeleteEvents(int(group_id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": -1,
			"msg":    "Fail to delete event-group relationship : " + err.Error(),
			"data":   nil,
		})
		return

	}

	//delete group
	err = service.DeleteGroup(userOne.UserId, int(group_id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": -1,
			"msg":    "Fail to delete group : " + err.Error(),
			"data":   nil,
		})
		return
	}

	//return value
	c.JSON(http.StatusOK, gin.H{
		"status": 1,
		"msg":    "group delete successfully!",
		"data":   nil,
	})

}
