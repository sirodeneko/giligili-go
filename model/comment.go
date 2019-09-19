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

// GetUserURL 用ID获取用户头像
func (comment *Comment) GetUserURL() (string, string) {
	var user User
	DB.First(&user, comment.UserID)

	return user.AvatarURL(), user.Nickname
}
