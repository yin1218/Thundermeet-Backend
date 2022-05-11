//負責堆放需要在database所創建的table資料
package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID             int64  `gorm:"primary_key;auto_increment" json:"ID"`
	UserId         string `gorm:"size:100;not null;unique" json:"userId"`
	UserName       string `gorm:"size:100;not null" json:"userName"`
	PasswordHash   string `gorm:"size:100;not null" json:"passwordHash"`
	PasswordAnswer string `gorm:"size:100;not null" json:"passwordAnswer"`
}

// 13 column
type Event struct {
	gorm.Model
	EventId             int64     `gorm:"primary_key;auto_increment" json:"event_id"`
	EventName           string    `gorm:"size:100;not null" json:"event_name"`
	IsPriorityEnabled   bool      `gorm:"size:100;not null" json:"event_nameis_priority_enabled"`
	IsConfirmed         bool      `gorm:"size:100;not null" json:"is_confirmed"`
	StartTime           time.Time `gorm:"size:100;not null" json:"start_time"`
	EndTime             time.Time `gorm:"size:100;not null" json:"end_time"`
	DateOrDays          bool      `gorm:"size:100;not null" json:"date_or_days"`
	StartDay            string    `gorm:"size:100" json:"start_day"`
	EndDay              string    `gorm:"size:100" json:"end_day"`
	StartDate           time.Time `gorm:"size:100" json:"start_date"`
	EndDate             time.Time `gorm:"size:100" json:"end_date"`
	ConfirmedTimeblocks string    `gorm:"size:100" json:"confirmed_timeblocks"`
	AdminId             string    `gorm:"size:100;not null" json:"admin_id"`
}
