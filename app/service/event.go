//關於orm相關的變數以及function，這邊也是主要對資料庫進行資料CRUD的地方。
package service

import (
	"fmt"
	"thundermeet_backend/app/dao"
	"thundermeet_backend/app/model"
	"time"
)

func CreateEvent(eventName string, isPriorityEnabled *bool, startTime string, endTime string, dateOrDays *bool, startDay string, endDay string, startDate time.Time, endDate time.Time, adminId string) (int64, error) {

	fmt.Println("Here")
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

func UpdateOneEvent(eventId int64, eventName string, confirmedTimeblocks []string) error {
	var event model.Event
	event = model.Event{
		EventName:           eventName,
		ConfirmedTimeblocks: confirmedTimeblocks,
	}
	updateErr := dao.SqlSession.Model(&model.Event{}).Where("event_id = ?", eventId).Updates(event).Error

	return updateErr
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
		return nil, fmt.Errorf("Get Event Info Failed:%v\n", dbResult.Error)
	} else {
		return events, nil
	}
}
