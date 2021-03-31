package service

import (
	"github.com/sirodeneko/giligili-go/model"
	"github.com/sirodeneko/giligili-go/serializer"
)

// UserRegisterService 管理用户注册服务
type UserRegisterService struct {
	Nickname        string `form:"nickname" json:"nickname" binding:"required,min=2,max=30"`
	UserName        string `form:"user_name" json:"user_name" binding:"required,min=5,max=30"`
	Password        string `form:"password" json:"password" binding:"required,min=8,max=40"`
	PasswordConfirm string `form:"password_confirm" json:"password_confirm" binding:"required,min=8,max=40"`
}

// Valid 验证表单
func (service *UserRegisterService) Valid() *serializer.Response {
	//验证注册时的前后密码
	if service.PasswordConfirm != service.Password {
		return &serializer.Response{
			Status: 40001,
			Msg:    "两次输入的密码不相同",
		}
	}
	//去同名
	count := 0
	model.DB.Model(&model.User{}).Where("nickname = ?", service.Nickname).Count(&count)
	if count > 0 {
		return &serializer.Response{
			Status: 40001,
			Msg:    "昵称被占用",
		}
	}
	//去账号相同（）
	count = 0
	model.DB.Model(&model.User{}).Where("user_name = ?", service.UserName).Count(&count)
	if count > 0 {
		return &serializer.Response{
			Status: 40001,
			Msg:    "用户账号已经注册",
		}
	}

	return nil
}

// Register 用户注册
func (service *UserRegisterService) Register() (model.User, *serializer.Response) {
	user := model.User{
		Nickname: service.Nickname,
		UserName: service.UserName,
		Status:   model.Active,
	}

	// 表单验证
	if err := service.Valid(); err != nil {
		return user, err
	}

	// 加密密码
	if err := user.SetPassword(service.Password); err != nil {
		return user, &serializer.Response{
			Status: 40002,
			Msg:    "密码加密失败",
		}
	}
	//默认头像
	user.Avatar = "img/noface.png"
	// 创建用户
	if err := model.DB.Create(&user).Error; err != nil {
		return user, &serializer.Response{
			Status: 40002,
			Msg:    "注册失败",
		}
	}

	return user, nil
}
