package service

import (
	"os"

	"giligili/model"
	"giligili/serializer"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

// DeleteVideoService 删除视频的服务
type DeleteVideoService struct { //将前端的数据绑定到结构体内
}

// Delete 删除视频
func (service *DeleteVideoService) Delete(id string, userid uint) serializer.Response {
	var video model.Video
	err := model.DB.First(&video, id).Error
	if err != nil {
		return serializer.Response{
			Status: 404,
			Msg:    "视频不存在",
			Error:  err.Error(),
		}
	}
	if userid != video.UserID {
		if userid != 1 {
			return serializer.Response{
				Status: 404,
				Msg:    "没有权限删除",
			}
		}

	}
	// DeleteVideo 删除阿里云oss里视频
	err = DeleteVideo(video)
	if err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "源视频删除失败",
			Error:  err.Error(),
		}
	}
	//删除数据库内容
	err = model.DB.Delete(&video).Error
	if err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "视频删除失败",
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Msg: "视频删除成功",
	}
}

// DeleteVideo 删除阿里云oss里视频
func DeleteVideo(video model.Video) error {
	client, err := oss.New(os.Getenv("OSS_END_POINT"), os.Getenv("OSS_ACCESS_KEY_ID"), os.Getenv("OSS_ACCESS_KEY_SECRET"))
	if err != nil {
		return err
	}

	// 获取存储空间。
	bucket, err := client.Bucket(os.Getenv("OSS_BUCKET"))
	if err != nil {
		return err
	}
	key1 := "delet/avatars" + video.Avatar[14:len(video.Avatar)]
	key2 := "delet/videos" + video.URL[13:len(video.URL)]

	_, err = bucket.CopyObject(video.Avatar, key1)
	if err != nil {
		return err
	}
	err = bucket.DeleteObject(video.Avatar)
	if err != nil {
		return err
	}

	_, err = bucket.CopyObject(video.URL, key2)
	if err != nil {
		return err
	}
	err = bucket.DeleteObject(video.URL)
	if err != nil {
		return err
	}

	//处理视频被删除后的问题
	video.DeleteVideo()
	return nil
}
