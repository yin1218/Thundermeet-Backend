//關於orm相關的變數以及function，這邊也是主要對資料庫進行資料CRUD的地方。
package service

import (
	"thundermeet_backend/app/dao"
	"thundermeet_backend/app/model"
)

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

func AddEventToGroup(eventId int, groupId int) error {
	group_event := model.GroupEvent{
		GroupId: groupId,
		EventId: eventId,
	}
	insertErr := dao.SqlSession.Model(&model.GroupEvent{}).Create(&group_event).Error
	return insertErr

}
