package api

import (
	"giligili/service"

	"github.com/gin-gonic/gin"
)

// VideoComment 用户视频评论接口
func VideoComment(c *gin.Context) {
	service := service.VideoCommentService{}
	user := CurrentUser(c)
	if err := c.ShouldBind(&service); err == nil {
		res := service.AddComment(c.Param("id"), user.ID)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// ListComment 视频列表详情接口
func ListComment(c *gin.Context) {
	service := service.ListCommentService{}
	var ID uint
	if CurrentUser(c)!=nil{
		ID=CurrentUser(c).ID
	}else{
		ID=0
	}
	//c.ShouldBind(&service)将前端的数据绑定到结构体内
	if err := c.ShouldBind(&service); err == nil {
		res := service.List(c.Param("id"),ID)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
// DeleteComment 删除视频评论的接口
func DeleteComment(c *gin.Context) {
	user := CurrentUser(c)
	userid := user.ID
	service := service.DeleteCommentService{}
	res := service.Delete(c.Param("id"), userid)
	c.JSON(200, res)
}