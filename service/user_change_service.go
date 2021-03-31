package service

import (
	"os"
	"strings"

	"github.com/sirodeneko/giligili-go/model"
	"github.com/sirodeneko/giligili-go/serializer"

	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

// UserChangeService 用户修改信息的服务
type UserChangeService struct {
	Nickname string `form:"nickname" json:"nickname" binding:"required,min=2,max=30"`
	Birthday int64  `form:"birthday" json:"birthday" `
	Sex      string `form:"sex" json:"sex"`
	Sign     string `form:"sign" json:"sign" binding:"required,min=2,max=200"`
	Avatar   string `form:"avatar" json:"avatar"`
}

// Change 用户修改信息
func (service *UserChangeService) Change(ID uint) serializer.Response {
	var user model.User
	//找到用户
	err := model.DB.First(&user, ID).Error
	if err != nil {
		return serializer.Response{
			Status: 404,
			Msg:    "查询用户失败",
			Error:  err.Error(),
		}
	}

	user.Nickname = service.Nickname
	user.Birthday = time.Unix(service.Birthday, 0)
	user.Sex = service.Sex
	user.Sign = service.Sign
	//如果更新了图片
	if !strings.Contains(service.Avatar, "upload/avatars") {
		//将头像，文件夹拷到文件夹
		client, err := oss.New(os.Getenv("OSS_END_POINT"), os.Getenv("OSS_ACCESS_KEY_ID"), os.Getenv("OSS_ACCESS_KEY_SECRET"))
		if err != nil {
			return serializer.Response{
				Status: 50002,
				Msg:    "OSS配置错误1",
				Error:  err.Error(),
			}
		}

		// 获取存储空间。
		bucket, err := client.Bucket(os.Getenv("OSS_BUCKET"))
		if err != nil {
			return serializer.Response{
				Status: 50002,
				Msg:    "OSS配置错误2",
				Error:  err.Error(),
			}
		}

		key1 := "upload/avatars/" + service.Avatar[14:len(service.Avatar)]
		_, err = bucket.CopyObject(service.Avatar, key1)
		if err != nil {
			return serializer.Response{
				Status: 50002,
				Msg:    "OSS配置错误3",
				Error:  err.Error(),
			}
		}
		service.Avatar = key1
		user.Avatar = service.Avatar
	}
	err = model.DB.Save(&user).Error
	if err != nil {
		return serializer.Response{
			Status: 50002,
			Msg:    "用户信息保存失败",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Data: serializer.BuildUser(user),
	}
}
