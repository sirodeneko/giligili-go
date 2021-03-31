package api

import (
	"encoding/json"

	"github.com/sirodeneko/giligili-go/im"
	"github.com/sirodeneko/giligili-go/serializer"
	"github.com/sirodeneko/giligili-go/service"

	"gopkg.in/olahol/melody.v1"

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

// GroupCreate 创建聊天室
func GroupCreate(c *gin.Context) {
	user := CurrentUser(c)
	ID := user.ID
	service := service.GroupCreateService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create(ID)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}

}

// GroupMsgs 拉取聊天消息
func GroupMsgs(c *gin.Context) {
	service := service.GroupMsgsService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.List(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// GroupJoin 加入聊天室
func GroupJoin(c *gin.Context) {
	user := CurrentUser(c)
	ID := user.ID
	service := service.GroupJoinService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Join(ID)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

type WsHeader struct {
	WsHeader string `json:"ws_header"`
}

// Select 判断websocket的消息
func Select(msg []byte, s *melody.Session) {
	var WH WsHeader
	json.Unmarshal(msg, &WH)
	switch {
	case WH.WsHeader == "link":
		link(s, msg)
	case WH.WsHeader == "msg":
		putMsg(s, msg)
	case WH.WsHeader == "ping":
		pong(s)
	}
}

type LinkGroup struct {
	Id uint `json:"id"`
}

// link 连接
func link(s *melody.Session, msg []byte) {
	var LG LinkGroup
	json.Unmarshal(msg, &LG)
	im.Join(LG.Id, s)
}

// putMsg 发送消息
func putMsg(s *melody.Session, msg []byte) {
	service := service.GroupPutMsgService{}
	err := json.Unmarshal(msg, &service)
	if err == nil {
		res := service.PutMsg()
		msg, err := json.Marshal(res)
		if err == nil {
			if res.Status == 0 {
				//广播消息
				//强行for干他妈的
				for _, item := range im.ROOM[service.To] {
					item.Write(msg)
				}
			} else {
				//处理错误信息
				s.Write(msg)
			}

		}
	} else {
		s.Write([]byte(err.Error()))
	}

}

// ping PingPong
func pong(s *melody.Session) {
	res := serializer.Response{
		Status: 1,
		Msg:    "pong",
	}
	msg, err := json.Marshal(res)
	if err == nil {
		s.Write(msg)
	}
}
