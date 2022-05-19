//關於orm相關的變數以及function，這邊也是主要對資料庫進行資料CRUD的地方。
package service

import (
	"fmt"
	"thundermeet_backend/app/dao"
	"thundermeet_backend/app/model"
	"time"
)

func CreateEvent(eventName string, isPriorityEnabled *bool, startTime string, endTime string, dateOrDays *bool, startDay string, endDay string, startDate time.Time, endDate time.Time, adminId string, eventDescription string) (int64, error) {

	participants := []string{adminId}
	event := model.Event{
		EventName:         eventName,
		IsPriorityEnabled: *isPriorityEnabled,
		IsConfirmed:       false,
		StartTime:         startTime,
		EndTime:           endTime,
		DateOrDays:        *dateOrDays,
		StartDay:          startDay,
		EndDay:            endDay,
		StartDate:         startDate,
		EndDate:           endDate,
		AdminId:           adminId,
		EventDescription:  eventDescription,
		Participants:      participants,
	}

	insertErr := dao.SqlSession.Model(&model.Event{}).Create(&event).Error
	return event.EventId, insertErr
}

func SelectOneEvent(event_id int64) (*model.Event, error) {
	eventOne := &model.Event{}
	err := dao.SqlSession.Where("event_id = ?", event_id).First(&eventOne).Error
	if err != nil {
		return nil, err
	} else {
		return eventOne, nil
	}
}

func UpdateOneEvent(eventId int64, eventName string, confirmedTimeblocks []string, eventDescription string) error {
	var event model.Event
	event = model.Event{
		EventName:           eventName,
		EventDescription:    eventDescription,
		ConfirmedTimeblocks: confirmedTimeblocks,
	}
	updateErr := dao.SqlSession.Model(&model.Event{}).Where("event_id = ?", eventId).Updates(event).Error

	return updateErr
}

func ConfirmOneEvent(eventId int64, confirmedTimeblocks []string) error {
	var event model.Event
	event = model.Event{
		IsConfirmed:         true,
		ConfirmedTimeblocks: confirmedTimeblocks,
	}
	updateErr := dao.SqlSession.Model(&model.Event{}).Where("event_id = ?", eventId).Updates(event).Error

	return updateErr
}

func GetEventParticipants(eventId int64) ([]string, error) {
	eventOne := &model.Event{}
	err := dao.SqlSession.Where("event_id = ?", eventId).First(&eventOne).Error
	if err != nil {
		return nil, err
	}
	participants := eventOne.Participants
	return participants, nil
}

func UpdateEventParticipants(eventId int64, userId string) error {
	eventOne := &model.Event{}
	err := dao.SqlSession.Where("event_id = ?", eventId).First(&eventOne).Error
	if err != nil {
		return err
	}

	participants := eventOne.Participants
	containsParticipant := false

	for _, x := range participants {
		if x == userId {
			containsParticipant = true
			break
		}
	}

	if containsParticipant {
		return nil
	} else {
		participants = append(participants, userId)
		var event model.Event
		event = model.Event{
			Participants: participants,
		}
		updateErr := dao.SqlSession.Model(&model.Event{}).Where("event_id = ?", eventId).Updates(event).Error

		return updateErr
	}
}

func GetEventsByUser(userId string) ([]model.Event, error) {
	var events []model.Event
	dbResult := dao.SqlSession.Where("? = ANY (participants)", userId).Find(&events)
	fmt.Print(dbResult)
	if dbResult.Error != nil {
		return nil, fmt.Errorf("get Event Info Failed:%v\n", dbResult.Error)
	} else {
		return events, nil
	}
}

func DeleteEvent(eventId int64) error {
	delErr := dao.SqlSession.Where("event_id = ?", eventId).Delete(&model.Event{}).Error
	return delErr
}

type Group struct {
	// gorm.Model
	GroupId   int
	GroupName string
}

func SelectEventGroups(eventId int, userId string) ([]Group, error) {

	// var groups_notfilter []int
	var groups_notfilter []model.GroupEvent
	// db := dao.SqlSession.Model(&model.GroupEvent{}).Where("Event_id=?", eventId).Pluck("group_id", &groups)
	db := dao.SqlSession.Preload("Group").Where("Event_id=?", eventId).Find(&groups_notfilter)

	if db.Error != nil {
		return nil, db.Error
	}

	var groups []int

	for _, group_id := range groups_notfilter {
		if group_id.Group.UserId == userId {
			groups = append(groups, group_id.GroupId)
		}
	}

	var results []Group
	//get name of the results
	for _, group_id := range groups {
		// fmt.Println(group)
		var temp Group
		err := dao.SqlSession.Model(&model.Group{}).Where("Group_id=?", group_id).First(&temp).Error
		if err != nil {
			return nil, err
		}
		results = append(results, temp)
	}

	fmt.Print(results)
	return results, nil

}
