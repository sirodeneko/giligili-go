package model

import (
	"github.com/jinzhu/gorm"
)

const (
	// Master 群主
	Master string = "master"
	// Administrator 管理员
	Administrator string = "administrator"
)

// Group 聊天室
type Group struct {
	gorm.Model
	UserID        uint   //创建人id
	GroupName     string //聊天室名称
	GroupAvatar   string `gorm:"size:1000"` //聊天室头像
	GroupType     uint   //聊天室类型  1：普通群聊 2：一对一私人聊天
	GroupIntrduce string `gorm:"size:1000"` //群组介绍
	LastID        uint   //最后发送信息id
}

// GroupUser 聊天室成员
type GroupUser struct {
	ID        uint   `gorm:"primary_key"` //聊天室id
	UserID    uint   //成员id
	GroupRole string //角色
}

// Chat 聊天消息
type Chat struct {
	gorm.Model
	UserID     uint   //发送人id
	IsMe       uint   //是否是自己发送的
	FromAvatar string `gorm:"size:1000"` //发送人头像
	FromName   string //发送者名称
	To         uint   //接受者
	MsgType    uint   // 消息类型 1: 文本 2：图片
	Msg        string `gorm:"size:1000"` //内容
}

// Contact 群列表
type Contact struct {
	gorm.Model
	UserID     uint   //所属人id
	GroupID uint //群id
}