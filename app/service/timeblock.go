//關於orm相關的變數以及function，這邊也是主要對資料庫進行資料CRUD的地方。
package service

import (
	"fmt"
	"strconv"
	"strings"
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

func DeletePreviousTimeblockParticipant(userId string, eventId int64) error {
	eventMatchString := "%" + strconv.Itoa(int(eventId)) + "%"
	delErr := dao.SqlSession.Where("user_id = ? AND time_block_id LIKE ?", userId, eventMatchString).Delete(&model.TimeblockParticipants{}).Error
	return delErr
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

func GetTimeblocksForEvent(eventId int64) ([]model.Timeblock, error) {
	var timeblocks []model.Timeblock
	dbResult := dao.SqlSession.Where("event_id = ? ", eventId).Find(&timeblocks)

	if dbResult.Error != nil {
		return nil, dbResult.Error
	} else {
		return timeblocks, nil
	}
}

func remove(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

func GetMembersStatusPerTimeBlock(timeblockId string, participants []string) ([]string, []string, []string, error) {
	var TimeblockParticipants []model.TimeblockParticipants
	dbResult := dao.SqlSession.Where("time_block_id = ? ", timeblockId).Find(&TimeblockParticipants)
	if dbResult.Error != nil {
		return nil, nil, nil, dbResult.Error
	} else {
		var priority []string
		var normal []string
		var notAvailable []string = participants
		for _, timeblockparticipant := range TimeblockParticipants {
			if timeblockparticipant.Priority {
				priority = append(priority, timeblockparticipant.UserId)
				notAvailable = remove(notAvailable, timeblockparticipant.UserId)
			} else {
				normal = append(normal, timeblockparticipant.UserId)
				notAvailable = remove(notAvailable, timeblockparticipant.UserId)
			}
		}

		return normal, priority, notAvailable, nil
	}
}

func GetStatusForTimeblock(userId string, eventId int64) ([]string, []string, error) {
	timeblocks, err := GetTimeblocksForEvent(eventId)
	if err != nil {
		return nil, nil, err
	}

	var priority []string
	var normal []string

	for _, timeblock := range timeblocks {
		var TimeblockParticipants []model.TimeblockParticipants
		dbResult := dao.SqlSession.Where("time_block_id = ? AND user_id = ?", timeblock.TimeBlockId, userId).Find(&TimeblockParticipants)
		if dbResult.Error != nil {
			return nil, nil, dbResult.Error
		} else {

			for _, timeblockparticipant := range TimeblockParticipants {
				blocktime := strings.Split(timeblockparticipant.TimeBlockId, "A")[0]
				if timeblockparticipant.Priority {
					priority = append(priority, blocktime)
				} else {
					normal = append(normal, blocktime)
				}
			}
		}
	}
	return normal, priority, nil
}
