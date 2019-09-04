package service

import (
	"giligili/model"
	"giligili/serializer"

	"time"
)

// UserChangeService 用户修改信息的服务
type UserChangeService struct {
	Nickname string `form:"nickname" json:"nickname" binding:"required,min=2,max=30"`
	Birthday int64  `form:"birthday" json:"birthday" `
	Sex      string `form:"sex" json:"sex"`
	Sign     string `form:"sign" json:"sign" binding:"required,min=2,max=200"`
	Avatar   string `form:"avatar" json:"avatar"`
}

// Change 用户修改信息
func (service *UserChangeService) Change(ID uint) serializer.Response {
	var user model.User
	//找到用户
	err := model.DB.First(&user, ID).Error
	if err != nil {
		return serializer.Response{
			Status: 404,
			Msg:    "查询用户失败",
			Error:  err.Error(),
		}
	}

	user.Nickname = service.Nickname
	user.Birthday = time.Unix(service.Birthday, 0)
	user.Sex = service.Sex
	user.Sign = service.Sign
	if service.Avatar != "" {
		user.Avatar = service.Avatar
	}
	err = model.DB.Save(&user).Error
	if err != nil {
		return serializer.Response{
			Status: 50002,
			Msg:    "用户信息保存失败",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Data: serializer.BuildUser(user),
	}
}
