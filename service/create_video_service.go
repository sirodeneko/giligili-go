package service

import (
	"os"

	"giligili/model"
	"giligili/serializer"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

// CreateVideoService 视频投稿的服务
type CreateVideoService struct { //将前端的数据绑定到结构体内
	Title  string `form:"title" json:"title" binding:"required,min=2,max=30"`
	Info   string `form:"info" json:"info" binding:"max=300"`
	URL    string `form:"url" json:"url"`
	Avatar string `form:"avatar" json:"avatar"`
}

// Create 创建视频
func (service *CreateVideoService) Create(user *model.User) serializer.Response {
	video := model.Video{
		Title:  service.Title,
		Info:   service.Info,
		URL:    service.URL,
		Avatar: service.Avatar,
		UserID: user.ID,
	}
	err := model.DB.First(&user, user.ID).Error
	if err != nil {
		return serializer.Response{
			Status: 50002,
			Msg:    "查询用户今日投稿数失败",
			Error:  err.Error(),
		}
	}
	if user.Upcnt >= 5 {
		return serializer.Response{
			Status: 50002,
			Msg:    "今日投稿数已达上限",
			Error:  err.Error(),
		}
	}

	model.DB.Model(&user).Update("upcnt", (user.Upcnt)+1)
	//判断是否上传了封面
	if video.Avatar == "" {
		return serializer.Response{
			Status: 50002,
			Msg:    "请上传封面",
			//Error:  err.Error(),
		}
	}
	//判断是否上传视频
	if video.URL == "" {
		return serializer.Response{
			Status: 50002,
			Msg:    "请上传视频",
			//Error:  err.Error(),
		}
	}
	//将头像，视频从临时文件夹拷到文件夹
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

	key1 := "upload/avatars/" + video.Avatar[14:len(video.Avatar)]
	key2 := "upload/videos/" + video.URL[13:len(video.URL)]
	_, err = bucket.CopyObject(video.Avatar, key1)
	if err != nil {
		return serializer.Response{
			Status: 50002,
			Msg:    "OSS配置错误3",
			Error:  err.Error(),
		}
	}
	_, err = bucket.CopyObject(video.URL, key2)
	if err != nil {
		return serializer.Response{
			Status: 50002,
			Msg:    "OSS配置错误3",
			Error:  err.Error(),
		}
	}
	video.Avatar = key1
	video.URL = key2

	err = model.DB.Create(&video).Error
	if err != nil {
		return serializer.Response{
			Status: 50001,
			Msg:    "视频保存失败",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Data: serializer.BuildVideo(video),
	}
}
