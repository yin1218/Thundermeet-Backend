//關於orm相關的變數以及function，這邊也是主要對資料庫進行資料CRUD的地方。
package service

import (
	"fmt"
	"thundermeet_backend/app/dao"
	"thundermeet_backend/app/model"
	"time"
)

func CreateEvent(eventName string, isPriorityEnabled bool, startTime time.Time, endTime time.Time, dateOrDays bool, startDay string, endDay string, startDate time.Time, endDate time.Time, adminId string) error {

	fmt.Println("Here")
	event := model.Event{
		EventName:         eventName,
		IsPriorityEnabled: isPriorityEnabled,
		IsConfirmed:       false,
		StartTime:         startTime,
		EndTime:           endTime,
		DateOrDays:        dateOrDays,
		StartDay:          startDay,
		EndDay:            endDay,
		StartDate:         startDate,
		EndDate:           endDate,
		AdminId:           adminId,
	}

	insertErr := dao.SqlSession.Model(&model.Event{}).Create(&event).Error
	return insertErr
}
