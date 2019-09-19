package model

import (
	"github.com/jinzhu/gorm"
)

// VideoLike 视频点赞关系
type VideoLike struct {
	gorm.Model
	VideoID  uint
	UserID   uint
}

