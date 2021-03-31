package service

import (
	"github.com/sirodeneko/giligili-go/model"
	"github.com/sirodeneko/giligili-go/serializer"
)

// GroupMsgsService 聊天室消息列表服务
type GroupMsgsService struct {
	Limit int `form:"limit"`
	Start int `form:"start"`
}

// List 聊天室消息列表
func (service *GroupMsgsService) List(ID string) serializer.Response {
	var chats []model.Chat
	total := 30
	if service.Limit == 0 {
		service.Limit = 30
	}
	if err := model.DB.Order("id desc").Where("to_group = ?", ID).Limit(service.Limit).Offset(service.Start).Find(&chats).Error; err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "数据库连接错误",
			Error:  err.Error(),
		}
	}

	return serializer.BuildListResponse(serializer.BuildChats(chats), uint(total))
}
