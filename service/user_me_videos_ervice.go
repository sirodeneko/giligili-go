package service

import (
	"giligili/model"
	"giligili/serializer"
)

// UserMeVideosService 用户视频列表服务
type UserMeVideosService struct {
	Limit int `form:"limit"`
	Start int `form:"start"`
}

// List 用户视频列表
func (service *UserMeVideosService) List(ID uint) serializer.Response {
	videos := []model.Video{}
	total := 0

	if service.Limit == 0 {
		service.Limit = 12
	}

	if err := model.DB.Model(model.Video{}).Where("user_id = ?", ID).Count(&total).Error; err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "数据库连接错误",
			Error:  err.Error(),
		}
	}

	if err := model.DB.Where("user_id = ?", ID).Limit(service.Limit).Offset(service.Start).Find(&videos).Error; err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "数据库连接错误",
			Error:  err.Error(),
		}
	}

	return serializer.BuildListResponse(serializer.BuildVideos(videos), uint(total))
}
