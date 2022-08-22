package service

import (
	"fmt"
	"memo_gin/model"
	"memo_gin/serializer"
	"time"
)

type CreateTaskService struct {
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
	Status  int    `json:"status" form:"status"` //0 未做，1 已做
}

type ShowTaskService struct {
}

type ListTaskService struct {
	PageNum  int `json:"page_num" form:"page_num"`
	PageSize int `json:"page_size" form:"page_size"`
}
type UpdateTaskService struct {
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
	Status  int    `json:"status" form:"status"`
}
type SearchTaskService struct {
	Info     string `json:"info" form:"info"`
	PageNum  int    `json:"page_num" form:"page_num"`
	PageSize int    `json:"page_size" form:"page_size"`
}
type DeleteTaskService struct {
}

//新增一条备忘录
func (service *CreateTaskService) Create(id uint) serializer.Response {
	var user model.User
	code := 200
	model.DB.First(&user, id)
	task := model.Task{
		User:     user,
		Uid:      user.ID,
		Title:    service.Title,
		Status:   0,
		Content:  service.Content,
		StarTime: time.Now().Unix(),
		EndTime:  0,
	}
	err := model.DB.Create(&task).Error
	if err != nil {
		code = 500
		return serializer.Response{
			Status: code,
			Msg:    "备忘录创建失败",
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    "备忘录创建成功",
	}
}

//展示一条备忘录
func (service *ShowTaskService) Show(tid string) serializer.Response {
	var task model.Task
	code := 200
	err := model.DB.First(&task, tid).Error
	if err != nil {
		code = 500
		return serializer.Response{
			Status: code,
			Msg:    "查询失败",
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildTask(task),
		Msg:    "查询成功",
	}
}

//展示所有备忘录
func (service *ListTaskService) List(uid uint) serializer.Response {
	var tasks []model.Task
	count := 0
	if service.PageSize == 0 {
		service.PageSize = 15
	}
	model.DB.Model(&model.Task{}).Preload("User").Where("uid=?", uid).Count(&count).
		Limit(service.PageSize).Offset((service.PageNum - 1) * service.PageSize).Find(&tasks)
	return serializer.BuildListResponse(serializer.BuildTasks(tasks), uint(count))
}

//更新备忘录
func (service *UpdateTaskService) Update(tid string) serializer.Response {
	var task model.Task
	model.DB.First(&task, tid)
	task.Content = service.Content
	task.Title = service.Title
	task.Status = service.Status
	err := model.DB.Save(&task).Error
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	return serializer.Response{
		Status: 200,
		Data:   serializer.BuildTask(task),
		Msg:    "更新完成",
	}
}

//查询备忘录
func (service *SearchTaskService) Search(uid uint) serializer.Response {
	var tasks []model.Task
	count := 0
	if service.PageSize == 0 {
		service.PageSize = 10
	}
	model.DB.Model(&model.Task{}).Preload("User").Where("uid=?", uid).
		Where("title LIKE ? OR content LIKE ?", "%"+service.Info+"%", "%"+service.Info+"%").
		Count(&count).Limit(service.PageSize).Offset((service.PageNum - 1) * service.PageSize).Find(&tasks)
	return serializer.BuildListResponse(serializer.BuildTasks(tasks), uint(count))
}

//删除备忘录
func (service *DeleteTaskService) Delete(tid string) serializer.Response {
	var task model.Task
	err := model.DB.Delete(&task, tid).Error
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return serializer.Response{
			Status: 500,
			Msg:    "删除失败",
		}
	}
	return serializer.Response{
		Status: 200,
		Msg:    fmt.Sprintf("uid=%s的备忘录已删除", tid),
	}
}
