package service

import (
	"giligili/serializer"
	"os"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/google/uuid"
)

// UploadTackService 获得上传oss token的服务
type UploadTackService struct {
	Filename string `form:"filename" json:"filename"`
}

// Post 创建tack
func (service *UploadTackService) Post() serializer.Response {
	client, err := oss.New(os.Getenv("OSS_END_POINT"), os.Getenv("OSS_ACCESS_KEY_ID"), os.Getenv("OSS_ACCESS_KEY_SECRET"))
	if err != nil {
		return serializer.Response{
			Status: 50002,
			Msg:    "OSS配置错误",
			Error:  err.Error(),
		}
	}

	// 获取存储空间。
	bucket, err := client.Bucket(os.Getenv("OSS_BUCKET"))
	if err != nil {
		return serializer.Response{
			Status: 50002,
			Msg:    "OSS配置错误",
			Error:  err.Error(),
		}
	}

	// 带可选参数的签名直传。
	options := []oss.Option{
		oss.ContentType("video/mp4"),
	}

	key := "upload/video/" + uuid.Must(uuid.NewRandom()).String() + ".mp4"
	// 签名直传。
	signedPutURL, err := bucket.SignURL(key, oss.HTTPPut, 600, options...)
	if err != nil {
		return serializer.Response{
			Status: 50002,
			Msg:    "OSS配置错误",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Data: map[string]string{
			"key": key,
			"put": signedPutURL,
		},
	}
}
