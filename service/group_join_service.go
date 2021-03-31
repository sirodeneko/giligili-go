package service

import (
	"github.com/sirodeneko/giligili-go/model"
	"github.com/sirodeneko/giligili-go/serializer"
)

// GroupJoinService 聊天室加入的服务
type GroupJoinService struct {
	ID uint `json:"id"`
}

// Join 加入聊天室
func (service *GroupJoinService) Join(userid uint) serializer.Response {
	group := model.Group{}

	err := model.DB.First(&group, service.ID).Error
	if err != nil {
		return serializer.Response{
			Status: 50001,
			Msg:    "秘密房间不存在",
			Error:  err.Error(),
		}
	}

	//加入房间
	groupUser := model.GroupUser{
		GroupID:   group.ID,
		UserID:    userid,
		GroupRole: model.Common,
	}
	err = model.DB.Create(&groupUser).Error
	if err != nil {
		return serializer.Response{
			Status: 50001,
			Msg:    "秘密房间加入失败",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Data: serializer.BuildGroup(group),
	}
}
