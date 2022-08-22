package serializer

import "memo_gin/model"

type User struct {
	ID       uint   `json:"id" form:"id" example:"1"` //用户id
	UserName string `json:"user_name" form:"user_name" example:"huanghai"`
	Status   string `json:"status" form:"status"`       //用户状态
	CreateAt int64  `json:"create_at" form:"create_at"` //创建
}

//BuildUser序列化用户
func BuildUser(user model.User) User {
	return User{
		ID:       user.ID,
		UserName: user.UserName,
		CreateAt: user.CreatedAt.Unix(),
	}
}
