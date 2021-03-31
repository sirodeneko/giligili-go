package service

import (
	"github.com/sirodeneko/giligili-go/model"
	"github.com/sirodeneko/giligili-go/serializer"
)

// GroupPutMsgService 聊天室发送消息的服务
type GroupPutMsgService struct {
	UserID  uint   `json:"user_id"`
	To      uint   `json:"to"`
	MsgType uint   `json:"type"`
	Msg     string `json:"msg"`
}

// PutMsg 发送消息
func (service *GroupPutMsgService) PutMsg() serializer.Response {
	//fmt.Println("开始1")
	groupUser := model.GroupUser{}
	var count uint
	err := model.DB.Where("group_id = ? AND user_id= ?", service.To, service.UserID).First(&groupUser).Count(&count).Error
	if err != nil {
		return serializer.Response{
			Status: 50001,
			Msg:    "进错房间了吧你",
			Error:  err.Error(),
		}
	}
	//fmt.Println("开始2")
	if count == 0 {
		return serializer.Response{
			Status: 50001,
			Msg:    "进错房间了吧你",
			Error:  "",
		}
	}
	//fmt.Println("开始3")
	var user model.User
	err = model.DB.First(&user, service.UserID).Error
	if err != nil {
		return serializer.Response{
			Status: 50001,
			Msg:    "你是异世界穿越过来的嘛，没见过你呢",
			Error:  err.Error(),
		}
	}
	//fmt.Println("开始4")
	if service.MsgType == 2 {
		service.Msg = model.AvatarURL(service.Msg)
	}
	chat := model.Chat{
		UserID:     service.UserID,
		FromAvatar: user.Avatar,
		FromName:   user.Nickname,
		ToGroup:    service.To,
		MsgType:    service.MsgType,
		Msg:        service.Msg,
	}

	err = model.DB.Create(&chat).Error
	if err != nil {
		return serializer.Response{
			Status: 50001,
			Msg:    "消息被黑洞吃掉了啊啊啊",
			Error:  err.Error(),
		}
	}
	//fmt.Println("开始5")
	//更新聊天室最后数据
	var group model.Group
	err = model.DB.First(&group, service.To).Error
	if err != nil {
		return serializer.Response{
			Status: 50001,
			Msg:    "房间不存在",
			Error:  err.Error(),
		}
	}
	//fmt.Println("开始6")
	model.DB.Model(&group).Update("last_id", chat.ID)
	//fmt.Println("开始7")
	return serializer.Response{
		Data: serializer.BuildChat(chat),
	}
}
