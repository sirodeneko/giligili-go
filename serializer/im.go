package serializer

import (
	"github.com/sirodeneko/giligili-go/model"
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
		GroupType:     item.GroupType,
		GroupName:     item.GroupName,
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
		groups = append(groups, group)
	}
	return groups
}

// Chat 返回聊天消息序列化器
type Chat struct {
	ID         uint   `json:"id"`
	UserID     uint   `json:"userID"` //创建人id
	CreatedAt  int64  `json:"created_at"`
	FromAvatar string `json:"from_avatar"`
	FromName   string `json:"from_name"`
	To         uint   `json:"to"`
	MsgType    uint   `json:"msg_type"`
	Msg        string `json:"msg"`
}

// BuildChat 序列化返回聊天消息
func BuildChat(item model.Chat) Chat {
	if item.MsgType == 2 {
		item.Msg = model.AvatarURL(item.Msg)
	}
	return Chat{
		ID:         item.ID,
		UserID:     item.UserID,
		CreatedAt:  item.CreatedAt.Unix(),
		FromAvatar: model.AvatarURL(item.FromAvatar),
		FromName:   item.FromName,
		To:         item.ToGroup,
		MsgType:    item.MsgType,
		Msg:        item.Msg,
	}
}

// BuildChats 序列化聊天室消息列表
func BuildChats(items []model.Chat) []Chat {
	var chats []Chat

	for _, item := range items {
		group := BuildChat(item)
		chats = append(chats, group)
	}
	return chats
}
