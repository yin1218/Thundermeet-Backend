//負責堆放需要在database所創建的table資料
package model

// 13 column
type Group struct {
	// gorm.Model
	GroupId   int    `gorm:"primary_key;auto_increment;constraint:OnDelete:CASCADE;" json:"group_id"`
	GroupName string `gorm:"size:100;not null" json:"group_name"`
	UserId    string `gorm:"size:100;not null" json:"user_id"`
}

type GroupEvent struct {
	// gorm.Model
	GroupEventsId int `gorm:"primary_key;auto_increment" json:"group_events_id"`
	GroupId       int `gorm:"size:100;not null;" json:"group_id"`
	EventId       int `gorm:"size:100;not null;" json:"event_id"`
}
