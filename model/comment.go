package model

import (
	"github.com/jinzhu/gorm"
)

// Comment 用户模型
type Comment struct {
	gorm.Model
	UserID  uint
	VideoID uint
	Txet    string `gorm:"size:10000"`
	Like    uint32 
}
