package service

import (
	"giligili/model"
	"giligili/serializer"
)

// GetMyGroupsService 聊天室列表服务
type GetMyGroupsService struct {
}

// List 聊天室列表
func (service *GetMyGroupsService) List(userID uint) serializer.Response {
	var GroupUsers []model.GroupUser

	if err := model.DB.Where("user_id=?", userID).Find(&GroupUsers).Error; err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "数据库连接错误",
			Error:  err.Error(),
		}
	}
	groups := ListGroups(GroupUsers)
	return serializer.BuildListResponse(serializer.BuildGroups(groups), uint(len(groups)))
}

// ListGroups 根据id渲染列表数据
func ListGroups(GroupUsers []model.GroupUser) []model.Group {
	var groups []model.Group
	for _, item := range GroupUsers {
		var group model.Group
		model.DB.Where("id=?", item.GroupID).Find(&group)
		groups = append(groups, group)
	}
	return groups
}
