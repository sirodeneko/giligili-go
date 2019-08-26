package service

import (
	"giligili/model"
	"giligili/serializer"
)

// DeleteVideoService 删除视频的服务
type DeleteVideoService struct { //将前端的数据绑定到结构体内
}

// Delete 删除视频
func (service *DeleteVideoService) Delete(id string) serializer.Response {
	var video model.Video
	err := model.DB.First(&video, id).Error
	if err != nil {
		return serializer.Response{
			Status: 404,
			Msg:    "视频不存在",
			Error:  err.Error(),
		}
	}
	err = model.DB.Delete(&video).Error
	if err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "视频删除失败",
			Error:  err.Error(),
		}
	}
	return serializer.Response{}
}
