package service

import (
	"giligili/model"
	"giligili/serializer"
	"strconv"
)

// VideoCommentService 视频评论的服务
type VideoCommentService struct {
	Txet string `form:"text" json:"text" binding:"required,max=10000"`
}

// AddComment 创建视频评论
func (service *VideoCommentService) AddComment(id string, userid uint) serializer.Response {
	var a uint64
	var b uint
	a, _ = strconv.ParseUint(id, 10, 0)
	b = uint(a)
	comment := model.Comment{
		UserID:  userid,
		VideoID: b,
		Txet:    service.Txet,
	}

	var video model.Video
	var user model.User
	err := model.DB.First(&video, id).Error
	if err != nil {
		return serializer.Response{
			Status: 404,
			Msg:    "查询视频不存在",
			Error:  err.Error(),
		}
	}
	err = model.DB.First(&user, userid).Error
	if err != nil {
		return serializer.Response{
			Status: 50002,
			Msg:    "查询评论用户信息失败",
			Error:  err.Error(),
		}
	}

	err = model.DB.Create(&comment).Error
	if err != nil {
		return serializer.Response{
			Status: 50001,
			Msg:    "评论保存失败",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Data: serializer.BuildComment(comment, userid),
	}
}
