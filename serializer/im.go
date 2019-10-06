package serializer

import (
	"giligili/model"
)

// Group 聊天室序列化器
type Group struct {
	ID            uint   `json:"id"`
	UserID        uint   `json:"userID"` //创建人id
	CreatedAt     int64  `json:"created_at"`
	GroupAvatar   string `json:"avatar"`        //聊天室头像
	GroupName     string `json:"groupName"`     //聊天室名称
	GroupType     uint   `json:"groupType"`     //聊天室类型  1：普通群聊 2：一对一私人聊天
	GroupIntrduce string `json:"groupIntrduce"` //群组介绍
	LastID        uint   `json:"lastID"`        //最后发送信息id
}

// BuildGroup 序列化聊天室
func BuildGroup(item model.Group) Group {
	//user, _ := model.GetUser(item.UserID)
	return Group{
		ID:            item.ID,
		UserID:        item.UserID,
		GroupName:     item.GroupName,
		GroupType:     item.GroupType,
		GroupAvatar:   item.AvatarURL(),
		GroupIntrduce: item.GroupIntrduce,
		LastID:        item.LastID,
		CreatedAt:     item.CreatedAt.Unix(), //time.Time转换为int64（时间戳）

	}
}

// BuildGroups 序列化聊天室列表
func BuildGroups(items []model.Group) []Group {
	var groups []Group

	for _, item := range items {
		group := BuildGroup(item)
		//把video给videos然后赋给videos 因为不是传指针（ps:有点傻逼）
		groups = append(groups, group)
	}
	return groups
}
