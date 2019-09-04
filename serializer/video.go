package serializer

import (
	"giligili/model"
)

// Video 视频序列化器
type Video struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Info      string `json:"Info"`
	URL       string `json:"url"`
	CreatedAt int64  `json:"created_at"`
	Avatar    string `json:"avatar"`
	View      uint64 `json:"view"`
	User      uint   `json:"user"`
	//User      User   `json:"user"`
}

// BuildVideo 序列化视频
func BuildVideo(item model.Video) Video {
	//user, _ := model.GetUser(item.UserID)
	return Video{
		ID:     item.ID,
		Title:  item.Title,
		Info:   item.Info,
		URL:    item.VideoURL(),
		Avatar: item.AvatarURL(),
		View:   item.View(),
		//User:      BuildUser(user),
		User:      item.UserID,
		CreatedAt: item.CreatedAt.Unix(), //time.Time转换为int64（时间戳）

	}
}

// BuildVideos 序列化视频列表
func BuildVideos(items []model.Video) []Video {
	var videos []Video

	for _, item := range items {
		video := BuildVideo(item)
		//把video给videos然后赋给videos 因为不是传指针（ps:有点傻逼）
		videos = append(videos, video)
	}
	return videos
}
