package api

import (
	"giligili/serializer"
	"giligili/service"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// UserRegister 用户注册接口
func UserRegister(c *gin.Context) {
	var service service.UserRegisterService
	if err := c.ShouldBind(&service); err == nil {
		if user, err := service.Register(); err != nil {
			c.JSON(200, err)
		} else {
			res := serializer.BuildUserResponse(user)
			c.JSON(200, res)
		}
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// UserLogin 用户登录接口
func UserLogin(c *gin.Context) {
	var service service.UserLoginService
	if err := c.ShouldBind(&service); err == nil {
		if user, err := service.Login(); err == nil {
			// 设置Session
			s := sessions.Default(c)
			s.Clear()
			s.Set("user_id", user.ID)
			s.Save()
			//c.JSON(200, gin.H{"user_id": s.Get("user_id")})
			res := serializer.BuildUserResponse(user)
			c.JSON(200, res)
		} else {
			c.JSON(200, err)
		}
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// UserMe 用户详情
func UserMe(c *gin.Context) {
	user := CurrentUser(c)
	res := serializer.BuildUserResponse(*user)
	c.JSON(200, res)
}

// UserLogout 用户登出
func UserLogout(c *gin.Context) {
	s := sessions.Default(c)
	s.Clear()
	s.Save()
	c.JSON(200, serializer.Response{
		Status: 0,
		Msg:    "登出成功",
	})
}

// UserChange 用户修改信息
func UserChange(c *gin.Context) {
	user := CurrentUser(c)
	ID := user.ID
	var service service.UserChangeService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Change(ID)
		c.JSON(200, res)

	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// ShowUser 查看固定id用户
func ShowUser(c *gin.Context) {
	service := service.ShowUserService{}
	//c.ShouldBind(&service)将前端的数据绑定到结构体内
	if err := c.ShouldBind(&service); err == nil {
		res := service.Show(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
