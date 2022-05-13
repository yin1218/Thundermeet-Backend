//關於orm相關的變數以及function，這邊也是主要對資料庫進行資料CRUD的地方。
package service

import (
	"fmt"
	"thundermeet_backend/app/dao"
	"thundermeet_backend/app/model"
	"time"
)

func CreateOneTimeblock(timeblockId string, eventId int64, blockTime time.Time) error {
	if !CheckOneTimeblock(timeblockId) {
		return fmt.Errorf("Timeblock exists")
	}

	timeblock := model.Timeblock{
		TimeBlockId: timeblockId,
		EventId:     eventId,
		BlockTime:   blockTime,
	}

	insertErr := dao.SqlSession.Model(&model.Timeblock{}).Create(&timeblock).Error
	return insertErr
}

func CheckOneTimeblock(timeblockId string) bool {
	result := false
	var timeblock model.Timeblock
	fmt.Print(result)
	dbResult := dao.SqlSession.Where("time_block_id = ?", timeblockId).Find(&timeblock)
	fmt.Print(dbResult)
	if dbResult.Error != nil {
		fmt.Printf("Get timeblock info Failed:%v\n", dbResult.Error)
	} else {
		result = true
	}
	fmt.Print(result)
	return result
}

func CreateOneTimeblockParticipant(userId string, timeblockId string, priority bool) error {
	if !CheckOneTimeblockParticipant(userId, timeblockId) {
		return fmt.Errorf("Timeblock exists")
	}

	timeblockparticipant := model.TimeblockParticipants{
		TimeBlockId: timeblockId,
		UserId:      userId,
		Priority:    priority,
	}

	insertErr := dao.SqlSession.Model(&model.TimeblockParticipants{}).Create(&timeblockparticipant).Error
	return insertErr
}

func CheckOneTimeblockParticipant(userId string, timeblockId string) bool {
	result := false
	var timeblockparticipant model.TimeblockParticipants
	fmt.Print(result)
	dbResult := dao.SqlSession.Where("user_id = ? AND time_block_id = ?", userId, timeblockId).Find(&timeblockparticipant)
	fmt.Print(dbResult)
	if dbResult.Error != nil {
		fmt.Printf("Get timeblockparticipant info Failed:%v\n", dbResult.Error)
	} else {
		result = true
	}
	fmt.Print(result)
	return result
}
