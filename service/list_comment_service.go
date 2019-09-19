package service

import (
	"giligili/model"
	"giligili/serializer"
)

// ListCommentService 视频评论列表服务
type ListCommentService struct {
	Limit int `form:"limit"`
	Start int `form:"start"`
}

// List 视频评论列表
func (service *ListCommentService) List(id string, uid uint) serializer.Response {
	comments := []model.Comment{}
	total := 0

	if service.Limit == 0 {
		service.Limit = 20
	}

	if err := model.DB.Where("video_id = ?", id).Model(model.Comment{}).Count(&total).Error; err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "数据库连接错误",
			Error:  err.Error(),
		}
	}

	if err := model.DB.Where("video_id = ?", id).Order("id desc").Limit(service.Limit).Offset(service.Start).Find(&comments).Error; err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "数据库连接错误",
			Error:  err.Error(),
		}
	}

	return serializer.BuildListResponse(serializer.BuildComments(comments, uid), uint(total))
}
