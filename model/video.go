package model

import (
	"giligili/cache"
	"os"
	"strconv"
	"strings"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/jinzhu/gorm"
)

// Video 视频模型
type Video struct {
	gorm.Model
	Title  string
	Info   string
	URL    string
	Avatar string
	UserID uint
}

// AvatarURL 封面地址
func (video *Video) AvatarURL() string {
	client, _ := oss.New(os.Getenv("OSS_END_POINT"), os.Getenv("OSS_ACCESS_KEY_ID"), os.Getenv("OSS_ACCESS_KEY_SECRET"))
	bucket, _ := client.Bucket(os.Getenv("OSS_BUCKET"))
	signedGetURL, _ := bucket.SignURL(video.Avatar, oss.HTTPGet, 60)
	if strings.Contains(signedGetURL, "http://giligili-img-av.oss-cn-hangzhou.aliyuncs.com/?Exp") {
		signedGetURL = "https://giligili-img-av.oss-cn-hangzhou.aliyuncs.com/img/noface.png"
	}
	return signedGetURL
}

// VideoURL 视频地址
func (video *Video) VideoURL() string {
	client, _ := oss.New(os.Getenv("OSS_END_POINT"), os.Getenv("OSS_ACCESS_KEY_ID"), os.Getenv("OSS_ACCESS_KEY_SECRET"))
	bucket, _ := client.Bucket(os.Getenv("OSS_BUCKET"))
	signedGetURL, _ := bucket.SignURL(video.URL, oss.HTTPGet, 3600)
	// if strings.Contains(signedGetURL, "http://giligili-img-av.oss-cn-hangzhou.aliyuncs.com/?Exp") {
	// 	signedGetURL = "https://giligili-img-av.oss-cn-hangzhou.aliyuncs.com/img/noface.png"
	// }
	return signedGetURL
}

// View 取 点击数
func (video *Video) View() uint64 {
	//返回字符串
	countStr, _ := cache.RedisClient.Get(cache.VideoViewKey(video.ID)).Result()
	//转回uint
	count, _ := strconv.ParseUint(countStr, 10, 64)
	return count
}

// AddView 视频浏览
func (video *Video) AddView() {
	// 增加视频点击数 incr()将键值++
	cache.RedisClient.Incr(cache.VideoViewKey(video.ID))
	// 增加排行点击数
	// 键名 加多少 成员(视频id)
	cache.RedisClient.ZIncrBy(cache.DailyRankKey, 1, strconv.Itoa(int(video.ID)))
}

// DeleteVideo 删除排行榜视频
func (video *Video) DeleteVideo() {
	//删除排行榜视频
	cache.RedisClient.ZRem(cache.DailyRankKey, strconv.Itoa(int(video.ID)))
}
