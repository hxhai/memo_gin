package model

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	User     User   `gorm:"ForeignKey:Uid"` //备忘录所属用户
	Uid      uint   `gorm:"not null"`       //外键关联
	Title    string `gorm:"index;not null"` //备忘录标题
	Status   int    `gorm:"default:'0'"`    //备忘录状态 0 未完成，1 已完成
	Content  string `gorm:"type:longtext"`  //备忘录内容
	StarTime int64  //备忘录开始时间
	EndTime  int64  //备忘录完成时间
}
