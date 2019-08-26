package service

import (
	"giligili/model"
	"giligili/serializer"
)

// ListVideoService 视频列表
type ListVideoService struct { //将前端的数据绑定到结构体内
}

// List 视频列表
func (service *ListVideoService) List() serializer.Response {
	var videos []model.Video
	err := model.DB.Find(&videos).Error
	if err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "数据库连接错误",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Data: serializer.BuildVideos(videos),
	}
}
