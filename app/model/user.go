//負責堆放需要在database所創建的table資料
package model

import (
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
