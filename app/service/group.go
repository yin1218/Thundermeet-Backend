//關於orm相關的變數以及function，這邊也是主要對資料庫進行資料CRUD的地方。
package service

import (
	"thundermeet_backend/app/dao"
	"thundermeet_backend/app/model"
)

var GroupFields = []string{"group_id", "group_name", "user_id"}

func CreateGroup(groupName string, userId string) (int, error) {

	// fmt.Println("Here")
	group := model.Group{
		GroupName: groupName,
		UserId:    userId,
	}

	insertErr := dao.SqlSession.Model(&model.Group{}).Create(&group).Error
	if insertErr != nil {
		return -1, insertErr
	}
	return group.GroupId, nil
}

func SelectOneGroup(groupId int) (*model.Group, error) {
	groupOne := &model.Group{}
	err := dao.SqlSession.Select(GroupFields).Where("Group_id=?", groupId).First(&groupOne).Error
	if err != nil {
		return nil, err
	} else {
		return groupOne, nil
	}

}

func GroupNameNotExist(groupName string, userId string) bool {
	groupOne := &model.Group{}
	err := dao.SqlSession.Select(GroupFields).Where("Group_name=? AND User_id=?", groupName, userId).First(&groupOne).Error
	return err == nil
}

type GroupEventInfo struct {
	Event_id     int
	Event_name   string
	Is_confirmed bool
}

func SelectGroupEvents(groupId int) ([]GroupEventInfo, error) {

	var results []GroupEventInfo
	var events_id []int
	db := dao.SqlSession.Model(&model.GroupEvent{}).Where("Group_id=?", groupId).Pluck("event_id", &events_id)
	if db.Error != nil {
		return nil, db.Error
	}

	//get info of the group
	for _, event_id := range events_id {
		var temp GroupEventInfo
		db := dao.SqlSession.Model(&model.Event{}).Where("Event_id=?", event_id).First(&temp)
		if db.Error != nil {
			return nil, db.Error
		}
		results = append(results, temp)
	}

	return results, nil

}

type GroupInfoItem struct {
	Group_id   int    `json:"group_id"`
	Group_name string `json:"group_name"`
}

type Groups []GroupInfoItem

func SelectGroups(userId string) (Groups, error) {
	var results Groups

	db := dao.SqlSession.Model(&model.Group{}).Select("group_id", "group_name").Where("User_id=?", userId).Scan(&results)
	if db.Error != nil {
		return nil, db.Error
	}
	return results, nil
}

func AddEventToGroup(eventId int, groupId int) error {
	group_event := model.GroupEvent{
		GroupId: groupId,
		EventId: eventId,
	}

	insertErr := dao.SqlSession.Model(&model.GroupEvent{}).Create(&group_event).Error
	return insertErr

}

func DeleteGroup(userId string, groupId int) error {
	delErr := dao.SqlSession.Where("user_id = ? AND group_id = ?", userId, groupId).Delete(&model.Group{}).Error
	return delErr
}

func DeleteGroupEvent(eventId int, groupId int) error {
	delErr := dao.SqlSession.Where("event_id = ? AND group_id = ?", eventId, groupId).Delete(&model.GroupEvent{}).Error
	return delErr
}

func DeleteEvents(groupId int) error {
	delErr := dao.SqlSession.Where("group_id = ?", groupId).Delete(&model.GroupEvent{}).Error
	return delErr
}

func ChangeGroupName(groupId int, groupName string) error {
	changeErr := dao.SqlSession.Model(&model.Group{}).Where("group_id = ?", groupId).Update("group_name", groupName).Error
	return changeErr

}
