package server

import (
	"os"

	"github.com/sirodeneko/giligili-go/api"
	"github.com/sirodeneko/giligili-go/im"
	"github.com/sirodeneko/giligili-go/middleware"

	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()
	// websocket连接
	m := melody.New()
	// 中间件, 顺序不能改
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.Cors()) //跨域
	// 获取用户身份
	r.Use(middleware.CurrentUser())
	// 路由
	v1 := r.Group("/api/v1")
	{
		v1.POST("ping", api.Ping)

		//视频详情
		v1.GET("video/:id", api.ShowVideo)
		//用户详情
		v1.GET("users/:id", api.ShowUser)
		//查询视频列表
		v1.GET("videos", api.ListVideo)
		// 排行榜
		v1.GET("rank/daily", api.DailyRank)
		// 用户注册
		v1.POST("user/register", api.UserRegister)

		//头像上传
		v1.POST("upload/token", api.UploadToken)
		//视频上传
		v1.POST("upload/tack", api.UploadTack)
		// 用户登录
		v1.POST("user/login", api.UserLogin)
		//查询视频评论列表
		v1.GET("videos/comments/:id", api.ListComment)
		// 需要登录保护的
		v1.Use(middleware.AuthRequired())
		{
			// User Routing
			v1.GET("user/me", api.UserMe)
			//用户退出
			v1.DELETE("user/logout", api.UserLogout)
			//用户投稿
			v1.POST("videos", api.CreateVideo)
			//更新视频
			v1.PUT("video/:id", api.UpdateVideo)
			//删除视频
			v1.DELETE("video/:id", api.DeleteVideo)
			//查询用户视频列表
			v1.GET("user/videos", api.UserMeVideos)
			//更新用户信息
			v1.PUT("user/account", api.UserChange)
			//发表评论(视频id))
			v1.POST("video/comment/:id", api.VideoComment)
			//删除评论(评论id)
			v1.DELETE("comment/:id", api.DeleteComment)
			//获取聊天室
			v1.GET("user/groups", api.GetMyGroups)
			//创建聊天室
			v1.POST("group/create", api.GroupCreate)
			//加入聊天室
			v1.POST("group/join", api.GroupJoin)
			//拉取聊天信息
			v1.GET("group/msgs/:id", api.GroupMsgs)
		}
		ws := v1.Group("")
		{
			//http升级为websocket
			ws.GET("/ws", func(c *gin.Context) {
				err := m.HandleRequest(c.Writer, c.Request)
				if err != nil {
					c.JSON(200, err)
				}
			})
			//当接收到消息
			m.HandleMessage(func(s *melody.Session, msg []byte) {
				api.Select(msg, s)
			})
			//当断开连接
			m.HandleDisconnect(func(s *melody.Session) {
				im.Exit(s)
			})

		}

	}
	return r
}
