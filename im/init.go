package im

import (
	"fmt"
	"giligili/model"

	"gopkg.in/olahol/melody.v1"
)

// ROOM 用map来维护聊天房间
var ROOM map[uint][]*melody.Melody

// RoomLoding 晨旭启动初始化ROOM
func RoomLoding() {
	ROOM = make(map[uint][]*melody.Melody)

	var groups []model.Group

	if err := model.DB.Find(&groups).Error; err != nil {
		panic(err)
	}

	for _, group := range groups {
		if ROOM[group.ID] == nil {
			ROOM[group.ID] = make([]*melody.Melody, 0, 100)
		}

	}
	fmt.Println(ROOM)
}
