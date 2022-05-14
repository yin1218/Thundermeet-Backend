//負責堆放需要在database所創建的table資料
package model

import (
	"time"
)

// 4 column
type Timeblock struct {
	// gorm.Model
	// ID          int64     `gorm:"primary_key;auto_increment" json:"ID"`
	TimeBlockId string    `gorm:"size:100;not null;unique" json:"time_block_id"`
	EventId     int64     `gorm:"size:100;not null" json:"event_name"`
	BlockTime   time.Time `gorm:"size:100;not null" json:"block_time"`
}

// 4 column
type TimeblockParticipants struct {
	// gorm.Model
	// ID          int64  `gorm:"primary_key;auto_increment" json:"ID"`
	UserId      string `gorm:"primary_key;size:100;not null;unique" json:"userId"`
	TimeBlockId string `gorm:"primary_key;size:100;not null;unique" json:"timeblock_id"`
	Priority    bool   `gorm:"primary_key;size:100;" json:"priority"`
}
