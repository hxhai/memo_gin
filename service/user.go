package service

import (
	"fmt"
	"memo_gin/model"
	"memo_gin/pkg/utils"
	"memo_gin/serializer"

	"github.com/jinzhu/gorm"
)

type UserService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=3,max=15"`
	Password string `form:"password" json:"password" binding:"required,min=5,max=16"`
}

func (service *UserService) Register() serializer.Response {
	var user model.User
	var count int
	model.DB.Model(user).Where("user_name=?", service.UserName).
		First(&user).Count(&count) //First:查找与给定条件匹配的第一条记录,Count:计数
	if count == 1 {
		return serializer.Response{
			Status: 400,
			Msg:    "用户名重复",
		}
	}
	user.UserName = service.UserName
	//密码加密
	err := user.SetPassword(service.Password)
	if err != nil {
		return serializer.Response{
			Status: 400,
			Msg:    "密码加密失败",
			Error:  err.Error(),
		}
	}
	//创建用户
	err = model.DB.Create(&user).Error
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "用户创建错误",
			Error:  err.Error(), //返回错误信息
		}
	}
	return serializer.Response{
		Status: 200,
		Msg:    "注册成功",
	}
}

func (service *UserService) Login() serializer.Response {
	var user model.User
	//数据库查找是否有这条用户信息
	if err := model.DB.Where("user_name=?", service.UserName).First(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) { //包含 RecordNotFound 错误，则返回 true
			return serializer.Response{
				Status: 400,
				Msg:    "用户不存在，请先登录",
			}
		}
		//如果是其他不可抗拒原因
		return serializer.Response{
			Status: 500,
			Msg:    "数据库错误！",
		}
	}
	//验证密码
	if !(user.CheckPassword(service.Password)) {
		return serializer.Response{
			Status: 400,
			Msg:    "密码错误！",
		}
	}
	//发一个token，为了其他功能需要身份验证所给前端存储的。
	//创建一个备忘录，就要token，不然不知道是谁创建的
	token, err := utils.GenerateToken(user.ID, service.UserName, service.Password)
	fmt.Printf("token: %v\n", token)
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "Token签发错误！",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: 200,
		Data:   serializer.TokenData{User: serializer.BuildUser(user), Token: token},
		Msg:    "登录成功",
	}
}
