package api

import (
	"giligili/service"

	"github.com/gin-gonic/gin"
)

// GetMyGroups 获取用户的聊天室
func GetMyGroups(c *gin.Context) {
	user := CurrentUser(c)
	ID := user.ID
	service := service.GetMyGroupsService{}
	res := service.List(ID)
	c.JSON(200, res)
}
