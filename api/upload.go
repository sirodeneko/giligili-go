package api

import (
	"giligili/service"

	"github.com/gin-gonic/gin"
)

// UploadToken 图片上传授权
func UploadToken(c *gin.Context) {
	service := service.UploadTokenService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Post()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
// UploadTack 视频上传授权
func UploadTack(c *gin.Context) {
	service := service.UploadTackService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Post()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
