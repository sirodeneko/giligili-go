package service

import (
	"giligili/model"
	"giligili/serializer"
)

// ShowVideoService 视频详情的服务
type ShowVideoService struct { //将前端的数据绑定到结构体内
}

// Show 视频详情
func (service *ShowVideoService) Show(id string) serializer.Response {
	var video model.Video
	err := model.DB.First(&video, id).Error
	if err != nil {
		return serializer.Response{
			Status: 404,
			Msg:    "查询不成功",
			Error:  err.Error(),
		}
	}
	//处理视频被观看后的问题
	video.AddView()

	return serializer.Response{
		Data: serializer.BuildVideo(video),
	}
}
