package model

import (
	"os"
	"strings"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/jinzhu/gorm"
)

const (
	// Master 群主
	Master string = "master"
	// Administrator 管理员
	Administrator string = "administrator"
	// Common 普通群员
	Common = "common"
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
	gorm.Model
	GroupID   uint   `gorm:"index"` //聊天室id
	UserID    uint   `gorm:"index"` //成员id
	GroupRole string //角色
}

// Chat 聊天消息
type Chat struct {
	gorm.Model
	UserID     uint   //发送人id
	FromAvatar string `gorm:"size:1000"` //发送人头像
	FromName   string //发送者名称
	ToGroup    uint   `gorm:"index"` //接受聊天室
	MsgType    uint   // 消息类型 1: 文本 2：图片
	Msg        string `gorm:"size:1000"` //内容
}

// AvatarURL 封面地址
func (group *Group) AvatarURL() string {
	client, _ := oss.New(os.Getenv("OSS_END_POINT"), os.Getenv("OSS_ACCESS_KEY_ID"), os.Getenv("OSS_ACCESS_KEY_SECRET"))
	bucket, _ := client.Bucket(os.Getenv("OSS_BUCKET"))
	signedGetURL, _ := bucket.SignURL(group.GroupAvatar, oss.HTTPGet, 60)
	if strings.Contains(signedGetURL, "http://giligili-img-av.oss-cn-hangzhou.aliyuncs.com/?Exp") {
		signedGetURL = "https://giligili-img-av.oss-cn-hangzhou.aliyuncs.com/img/noface.png"
	}
	return signedGetURL
}

// AvatarURL 封面地址
func AvatarURL(url string) string {
	client, _ := oss.New(os.Getenv("OSS_END_POINT"), os.Getenv("OSS_ACCESS_KEY_ID"), os.Getenv("OSS_ACCESS_KEY_SECRET"))
	bucket, _ := client.Bucket(os.Getenv("OSS_BUCKET"))
	signedGetURL, _ := bucket.SignURL(url, oss.HTTPGet, 60)
	if strings.Contains(signedGetURL, "http://giligili-img-av.oss-cn-hangzhou.aliyuncs.com/?Exp") {
		signedGetURL = "https://giligili-img-av.oss-cn-hangzhou.aliyuncs.com/img/noface.png"
	}
	return signedGetURL
}
