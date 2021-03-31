package service

import (
	"github.com/sirodeneko/giligili-go/model"
	"github.com/sirodeneko/giligili-go/serializer"
)

// UpdateVideoService 更新视频的服务
type UpdateVideoService struct { //将前端的数据绑定到结构体内
	Title string `form:"title" json:"title" binding:"required,min=2,max=30"`
	Info  string `form:"info" json:"info" binding:"max=300"`
}

// Update 更新视频
func (service *UpdateVideoService) Update(id string) serializer.Response {
	var video model.Video
	//找到视频
	err := model.DB.First(&video, id).Error
	if err != nil {
		return serializer.Response{
			Status: 404,
			Msg:    "视频不存在",
			Error:  err.Error(),
		}
	}

	video.Title = service.Title
	video.Info = service.Info
	err = model.DB.Save(&video).Error
	if err != nil {
		return serializer.Response{
			Status: 50002,
			Msg:    "视频保存失败",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Data: serializer.BuildVideo(video),
	}
}
