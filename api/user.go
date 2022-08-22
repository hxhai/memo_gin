package api

import (
	"memo_gin/service"

	"github.com/gin-gonic/gin"
)

//用户注册
func UserRegister(c *gin.Context) {
	var userRegister service.UserService
	//进行绑定
	if err := c.ShouldBind(&userRegister); err == nil {
		res := userRegister.Register()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}

//用户登录
func UserLogin(c *gin.Context) {
	var userLogin service.UserService
	//绑定
	if err := c.ShouldBind(&userLogin); err == nil {
		res := userLogin.Login()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}
