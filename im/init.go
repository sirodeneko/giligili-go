package im

import (
	"fmt"
	"sync"

	"github.com/sirodeneko/giligili-go/model"

	"gopkg.in/olahol/melody.v1"
)

// ROOM 用map来维护聊天房间
var ROOM map[uint][]*melody.Session

//读写锁
var mutex sync.RWMutex

// CloseChan 处理断开连接
var CloseChan chan *melody.Session = make(chan *melody.Session, 100)

// RoomLoding 程序启动初始化ROOM
func RoomLoding() {
	ROOM = make(map[uint][]*melody.Session)

	var groups []model.Group

	if err := model.DB.Find(&groups).Error; err != nil {
		panic(err)
	}

	for _, group := range groups {
		if ROOM[group.ID] == nil {
			ROOM[group.ID] = make([]*melody.Session, 0, 100)
		}

	}
	//这里启动一个协程 进行房间和链接对应关心的维护
	go func() {
		for {
			select {
			case s := <-CloseChan:
				go func() {
					if !s.IsClosed() {
						s.Close()
					}
					mutex.Lock()
					for index1, room := range ROOM {
						for index2, session := range room {
							if s == session {
								if index2 == len(room)-1 {
									ROOM[index1] = ROOM[index1][:index2]
								} else {
									ROOM[index1] = append(ROOM[index1][:index2], ROOM[index1][index2+1:]...)
								}
								break
							}

						}
					}
					mutex.Unlock()
				}()
			}
		}
	}()
	fmt.Println("ROOM CREATE OJBK!!!")
}

// Join 加入
func Join(id uint, s *melody.Session) {
	mutex.Lock()
	for _, session := range ROOM[id] {
		if s == session {
			return
		}
	}
	ROOM[id] = append(ROOM[id], s)
	mutex.Unlock()
}

// Create 创建
func Create(id uint) {
	mutex.Lock()
	if ROOM[id] == nil {
		ROOM[id] = make([]*melody.Session, 0, 100)
	}
	mutex.Unlock()
}

// Exit 退出
func Exit(s *melody.Session) {
	CloseChan <- s
}
