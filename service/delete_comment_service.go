package service

import (
	"giligili/model"
	"giligili/serializer"
)

// DeleteCommentService 删除视频评论的服务
type DeleteCommentService struct { //将前端的数据绑定到结构体内
}

// Delete 删除视频评论
func (service *DeleteCommentService) Delete(id string, userid uint) serializer.Response {
	var comment model.Comment
	err := model.DB.First(&comment, id).Error
	if err != nil {
		return serializer.Response{
			Status: 404,
			Msg:    "视频评论不存在",
			Error:  err.Error(),
		}
	}
	if userid != comment.UserID {
		if userid != 1 {
			return serializer.Response{
				Status: 404,
				Msg:    "没有权限删除",
			}
		}

	}
	//删除数据库内容
	err = model.DB.Delete(&comment).Error
	if err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "评论删除失败",
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Msg: "评论删除成功",
	}
}
