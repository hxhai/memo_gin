package routes

import (
	"memo_gin/api"
	"memo_gin/middleware"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

//路由
func NewRouter() *gin.Engine { //返回gin引擎
	e := gin.Default() //创建实例
	store := cookie.NewStore([]byte("something-very-secret"))
	e.Use(sessions.Sessions("mysesion", store))
	v1 := e.Group("api/v1")
	{
		//用户注册
		v1.POST("user/register", api.UserRegister)
		//用户登录
		v1.POST("user/login", api.UserLogin)
		//身份认证
		authed := v1.Group("/")
		//使用将中间件添加到组
		authed.Use(middleware.JWT())
		{
			authed.POST("task", api.CreateTask)
			//展示备忘录
			authed.GET("task/:id", api.ShowTask)
			//展示所有备忘录
			authed.GET("tasks", api.ListTask)
			//更新备忘录
			authed.PUT("task/:id", api.UpdateTask)
			//查询备忘录
			authed.POST("search", api.SearchTask)
			//删除备忘录
			authed.DELETE("task/:id", api.DeleteTask)
		}
	}
	return e
}
