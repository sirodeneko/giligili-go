package service

import (
	"github.com/sirodeneko/giligili-go/im"
	"github.com/sirodeneko/giligili-go/model"
	"github.com/sirodeneko/giligili-go/serializer"
)

// GroupCreateService 聊天室创建的服务
type GroupCreateService struct { //将前端的数据绑定到结构体内
	GroupName string `form:"name" json:"name" binding:"required,min=2,max=30"`
	GroupType uint   `form:"type" json:"type"`
}

// Create 创建聊天室
func (service *GroupCreateService) Create(userid uint) serializer.Response {
	group := model.Group{
		GroupName:     service.GroupName,
		GroupType:     service.GroupType,
		GroupAvatar:   "图片库/1.jpg",
		GroupIntrduce: "",
		UserID:        userid,
		LastID:        0,
	}

	err := model.DB.Create(&group).Error
	if err != nil {
		return serializer.Response{
			Status: 50001,
			Msg:    "秘密房间创建失败",
			Error:  err.Error(),
		}
	}

	//将房间加入ROOM
	im.Create(group.ID)
	//将群主加入房间
	groupUser := model.GroupUser{
		GroupID:   group.ID,
		UserID:    userid,
		GroupRole: model.Master,
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
