package api

import (
	"memo_gin/pkg/utils"
	"memo_gin/service"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus" //推荐日志依赖
)

//创建备忘录
func CreateTask(c *gin.Context) {
	var createTask service.CreateTaskService
	//验证token，GetHeader从请求标头返回token值
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&createTask); err == nil {
		res := createTask.Create(claim.Id)
		c.JSON(200, res)
	} else {
		logrus.Error(err) //打印错误日志
		c.JSON(400, ErrorResponse(err))
	}
}

//展示一条备忘录
func ShowTask(c *gin.Context) {
	var showTask service.ShowTaskService
	if err := c.ShouldBind(&showTask); err == nil {
		res := showTask.Show(c.Param("id"))
		c.JSON(200, res)
	} else {
		logrus.Error(err) //打印错误日志
		c.JSON(400, ErrorResponse(err))
	}
}

//展示所有备忘录
func ListTask(c *gin.Context) {
	var listTask service.ListTaskService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&listTask); err == nil {
		res := listTask.List(claim.Id)
		c.JSON(200, res)
	} else {
		logrus.Error(err) //打印错误日志
		c.JSON(400, ErrorResponse(err))
	}
}

//更新备忘录
func UpdateTask(c *gin.Context) {
	var updateTask service.UpdateTaskService
	if err := c.ShouldBind(&updateTask); err == nil {
		res := updateTask.Update(c.Param("id"))
		c.JSON(200, res)
	} else {
		logrus.Error(err) //打印错误日志
		c.JSON(400, ErrorResponse(err))
	}
}

//查询备忘录
func SearchTask(c *gin.Context) {
	var searchTask service.SearchTaskService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&searchTask); err == nil {
		res := searchTask.Search(claim.Id)
		c.JSON(200, res)
	} else {
		logrus.Error(err) //打印错误日志
		c.JSON(400, ErrorResponse(err))
	}
}

//删除备忘录
func DeleteTask(c *gin.Context) {
	var deleteTask service.DeleteTaskService
	if err := c.ShouldBind(&deleteTask); err == nil {
		res := deleteTask.Delete(c.Param("id"))
		c.JSON(200, res)
	} else {
		logrus.Error(err) //打印错误日志
		c.JSON(400, ErrorResponse(err))
	}
}
