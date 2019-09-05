package service

import (
	"giligili/model"
	"giligili/serializer"
)

// ShowUserService 用户详情的服务
type ShowUserService struct { //将前端的数据绑定到结构体内
}

// Show 视频详情
func (service *ShowUserService) Show(id string) serializer.Response {
	var user model.User
	err := model.DB.First(&user, id).Error
	if err != nil {
		return serializer.Response{
			Status: 404,
			Msg:    "查询不成功",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Data: serializer.BuildUser(user),
	}
}
