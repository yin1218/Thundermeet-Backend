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
	helper "thundermeet_backend/app/helpers"
	"thundermeet_backend/app/middleware/jwt"
	"thundermeet_backend/app/service"
	"time"

	"github.com/gin-gonic/gin"
)

type TimeblockController struct{}

type createTimeblockFormat struct {
	Event_id int64    `json:"eventId" example:"1" binding:"required"`       //required
	Normal   []string `json:"normal" example:"2021-01-01T11:00:00+08:00"`   //optional
	Priority []string `json:"priority" example:"2021-01-02T12:00:00+08:00"` //optional
} //@name Timeblock

func CreateTimeblocksController() TimeblockController {
	return TimeblockController{}
}

func UpdateTimeblocksController() TimeblockController {
	return TimeblockController{}
}

func GetTimeblocksController() TimeblockController {
	return TimeblockController{}
}

func StartNewTime(curHour int, curMin int, endHour int, endMin int) bool {
	if curHour > endHour {
		return true
	}
	if curHour == endHour && curMin == endMin {
		return true
	}
	return false
}

func CreateManyTimeblocks(dateOrDays bool, startTime string, endTime string, startDate time.Time, endDate time.Time, eventId int64) error {
	s_hh, _ := strconv.Atoi(startTime[0:2])
	s_min, _ := strconv.Atoi(startTime[2:4])
	e_hh, _ := strconv.Atoi(endTime[0:2])
	e_min, _ := strconv.Atoi(endTime[2:4])
	fmt.Print("strateim = ", startTime, " ", startTime, " ", endTime, " ", endTime)

	if dateOrDays { // date
		s_yyyy, s_mm, s_dd := startDate.Date()
		e_yyyy, e_mm, e_dd := endDate.Date()

		start_t := time.Date(s_yyyy, s_mm, s_dd, s_hh, s_min, 0, 0, time.Local)
		end_t := time.Date(e_yyyy, e_mm, e_dd, e_hh, e_min, 0, 0, time.Local)

		time_string := start_t.Format(time.RFC3339)
		timeblock_id := time_string + "A" + strconv.Itoa(int(eventId))

		err := service.CreateOneTimeblock(timeblock_id, eventId, start_t)
		if err != nil {
			return fmt.Errorf(err.Error())
		}

		for start_t.Add(time.Minute * time.Duration(30)).Before(end_t) {
			add30Min := start_t.Add(time.Minute * time.Duration(30))
			if StartNewTime(add30Min.Hour(), add30Min.Minute(), e_hh, e_min) {
				start_t = start_t.AddDate(0, 0, 1)
				start_t = time.Date(start_t.Year(), start_t.Month(), start_t.Day(), s_hh, s_min, 0, start_t.Nanosecond(), start_t.Location())
			} else {
				start_t = add30Min
			}
			fmt.Print("time = ", start_t)
			time_string := start_t.Format(time.RFC3339)
			timeblock_id := time_string + "A" + strconv.Itoa(int(eventId))
			err := service.CreateOneTimeblock(timeblock_id, eventId, start_t)
			if err != nil {
				return fmt.Errorf(err.Error())
			}
		}

		return nil
	} else { // days

	}
	return nil
}

func CreateManyTimeblocksParticipants(eventId int64, userId string, normal []string, priority []string) error {
	timeset := map[string]bool{}
	for _, timeblock := range priority {
		fmt.Print(timeblock)
		timeset[timeblock] = true
	}

	for _, timeblock := range normal {
		fmt.Print(timeblock)
		if timeset[timeblock] == true {
			return fmt.Errorf("normal and priority have times that overlap")
		}
		timeset[timeblock] = false
	}

	for timeblock, priority := range timeset {
		timeblock_id := helper.ConvertToTimeblockId(timeblock, eventId)
		fmt.Print("timeid = ", timeblock_id, "-------")
		err := service.CreateOneTimeblockParticipant(userId, timeblock_id, priority)
		if err != nil {
			return err
		}
	}
	return nil
}

func UpdateManyTimeblocksParticipants(eventId int64, userId string, normal []string, priority []string) error {
	timeset := map[string]bool{}
	for _, timeblock := range priority {
		fmt.Print(timeblock)
		timeset[timeblock] = true
	}

	for _, timeblock := range normal {
		fmt.Print(timeblock)
		if timeset[timeblock] == true {
			return fmt.Errorf("normal and priority have times that overlap")
		}
		timeset[timeblock] = false
	}

	err := service.DeletePreviousTimeblockParticipant(userId, eventId)
	if err != nil {
		return err
	}

	for timeblock, priority := range timeset {
		timeblock_id := helper.ConvertToTimeblockId(timeblock, eventId)
		fmt.Print("timeid = ", timeblock_id, "-------")

		err := service.CreateOneTimeblockParticipant(userId, timeblock_id, priority)
		if err != nil {
			return err
		}
	}
	return nil
}

// CreateTimeblock CreateTimeblock @Summary
// @Tags timeblock
// @version 1.0
// @produce application/json
// @Param Authorization header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param Body body createTimeblockFormat true "The body to create an event"
// @Success 200 string string successful return data
// @Failure 500 string string ErrorResponse
// @Router /v1/timeblocks/ [post]
func (u TimeblockController) CreateTimeblock(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	//validate token
	userId, err := jwt.ValidateToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	var form createTimeblockFormat
	bindErr := c.BindJSON(&form)

	if bindErr == nil {
		err := CreateManyTimeblocksParticipants(form.Event_id, userId, form.Normal, form.Priority)
		err = service.UpdateEventParticipants(form.Event_id, userId)
		if err == nil {
			c.JSON(http.StatusCreated, gin.H{
				"status": 1,
				"msg":    "timeblocks saved successfully!",
				"data":   nil,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": -1,
				"msg":    "Cannot create timeblocks!" + err.Error(),
				"data":   nil,
			})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": -1,
			"msg":    "Failed to parse timeblocks data : " + bindErr.Error(),
			"data":   nil,
		})
	}

}

// UpdateTimeblock UpdateTimeblock @Summary
// @Tags timeblock
// @version 1.0
// @produce application/json
// @Param Authorization header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Param Body body createTimeblockFormat true "The body to update an event"
// @Success 200 string string successful return data
// @Failure 500 string string ErrorResponse
// @Router /v1/timeblocks/ [put]
func (u TimeblockController) UpdateTimeblock(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	//validate token
	userId, err := jwt.ValidateToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	var form createTimeblockFormat
	bindErr := c.BindJSON(&form)
	if bindErr == nil {
		err := UpdateManyTimeblocksParticipants(form.Event_id, userId, form.Normal, form.Priority)
		if err == nil {
			c.JSON(http.StatusAccepted, gin.H{
				"status": 1,
				"msg":    "timeblocks updated successfully!",
				"data":   nil,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": -1,
				"msg":    "Cannot create timeblocks!" + err.Error(),
				"data":   nil,
			})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": -1,
			"msg":    "Failed to parse timeblocks data : " + bindErr.Error(),
			"data":   nil,
		})
	}

}

type GetTimeblockFormat struct {
	Time         string   `json:"time" example:"2021-01-01T11:00:00+08:00" binding:"required"` //required
	Normal       []string `json:"normal" example:"小葉"`                                         //optional
	Priority     []string `json:"priority" example:"小巫"`                                       //optional
	NotAvailable []string `json:"not_available" example:"小陳"`                                  //optional
} //@name GetTimeblockResponse

// GetTimeblock GetTimeblock @Summary
// @Tags timeblock
// @version 1.0
// @produce application/json
// @Param Authorization header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Success 200 string string successful return data
// @Failure 500 string string ErrorResponse
// @param event_id path int64 true "event id"
// @Router /v1/timeblocks/{event_id} [get]
func (u TimeblockController) GetTimeblock(c *gin.Context) {
	event_id, err := strconv.ParseInt(c.Param("event_id"), 10, 64)
	fmt.Print(event_id)

	timeblocks, err := service.GetTimeblocksForEvent(event_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": -1,
			"msg":    "Cannot get timeblocks!" + err.Error(),
			"data":   nil,
		})
		return
	}
	var fullTimeblockRes []GetTimeblockFormat
	for _, timeblock := range timeblocks {
		timeblockId := timeblock.TimeBlockId

		var timeblockRes GetTimeblockFormat

		timeblockRes.Time = timeblock.BlockTime.Format(time.RFC3339)
		participants, err := service.GetEventParticipants(timeblock.EventId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": -1,
				"msg":    "Cannot get timeblocks!" + err.Error(),
				"data":   nil,
			})
			return
		}

		fmt.Print("participants = ", participants)

		normal, priority, notAvailable, err := service.GetMembersStatusPerTimeBlock(timeblockId, participants)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": -1,
				"msg":    "Cannot get timeblocks!" + err.Error(),
				"data":   nil,
			})
			return
		}

		timeblockRes.Normal = normal
		timeblockRes.Priority = priority
		timeblockRes.NotAvailable = notAvailable

		fullTimeblockRes = append(fullTimeblockRes, timeblockRes)
	}

	c.JSON(http.StatusAccepted, fullTimeblockRes)

}

type GetTimeblockPreviewFormat struct {
	EventId  int64    `json:"event_id" example:"26" binding:"required"`     //required
	Normal   []string `json:"normal" example:"2021-01-01T11:00:00+08:00"`   //optional
	Priority []string `json:"priority" example:"2021-01-01T11:00:00+08:00"` //optional
} //@name GetTimeblockPreviewResponse

// GetTimeblockPreview GetTimeblockPreview @Summary
// @Tags timeblock
// @version 1.0
// @produce application/json
// @Param Authorization header string true "Bearer 31a165baebe6dec616b1f8f3207b4273"
// @Success 200 string string successful return data
// @Failure 500 string string ErrorResponse
// @param event_id path int64 true "event id"
// @Router /v1/timeblocks/{event_id}/preview [get]
func (u TimeblockController) GetTimeblockPreview(c *gin.Context) {
	eventId, err := strconv.ParseInt(c.Param("event_id"), 10, 64)
	fmt.Print(eventId)

	token := c.Request.Header.Get("Authorization")
	//validate token
	userId, err := jwt.ValidateToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	var timeblockPreviewRes GetTimeblockPreviewFormat

	timeblockPreviewRes.EventId = eventId

	normal, priority, err := service.GetStatusForTimeblock(userId, eventId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": -1,
			"msg":    "Cannot get timeblocks!" + err.Error(),
			"data":   nil,
		})
		return
	}
	timeblockPreviewRes.Normal = normal
	timeblockPreviewRes.Priority = priority

	c.JSON(http.StatusAccepted, timeblockPreviewRes)

}
