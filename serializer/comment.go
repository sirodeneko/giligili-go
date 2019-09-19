package serializer

import (
	"giligili/model"
)

// Comment 视频评论序列化器
type Comment struct {
	ID        uint   `json:"id"`
	UserID    uint   `json:"uid"`
	UserURL   string `json:"user_url"`
	UserName  string `json:"user_name"`
	Txet      string `json:"txet"`
	Like      uint32 `json:"like"`
	CreatedAt int64  `json:"created_at"`
	Me        uint   `json:"me"`
}

// BuildComment 序列化评论
func BuildComment(item model.Comment, uid uint) Comment {
	var flog uint
	flog = 0
	if uid == item.UserID || uid == 1 {
		flog = 1
	}
	UserURL, UserName := item.GetUserURL()
	return Comment{
		ID:        item.ID,
		UserID:    item.UserID,
		UserURL:   UserURL,
		UserName:  UserName,
		Txet:      item.Txet,
		Like:      item.Like,
		Me:        flog,
		CreatedAt: item.CreatedAt.Unix(), //time.Time转换为int64（时间戳）
	}
}

// BuildComments 序列化视频评论列表
func BuildComments(items []model.Comment, uid uint) []Comment {
	var comments []Comment

	for _, item := range items {
		comment := BuildComment(item, uid)
		comments = append(comments, comment)
	}
	return comments
}
