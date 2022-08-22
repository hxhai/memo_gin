package model

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "gorm.io/driver/mysql"
)

var DB *gorm.DB

func Database(connstring string) {
	fmt.Printf("connstring: %v\n", connstring) //root:root@tcp(127.0.0.1:3306)/memo_list_db?charset=utf8&parseTime=True&loc=Local
	db, err := gorm.Open("mysql", connstring)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		panic("mysql数据库连接错误！")
	}
	fmt.Println("mysql连接成功！")
	// db.LogMode(true) //打印日志
	if gin.Mode() == "release" {
		db.LogMode(false) //如果gin.Mode是发行版，不用输出日志
	}
	db.SingularTable(true)                       //表名不加s，user->user
	db.DB().SetMaxIdleConns(20)                  //最大连接池
	db.DB().SetMaxOpenConns(100)                 //最大连接数
	db.DB().SetConnMaxLifetime(time.Second * 30) //连接时间
	DB = db
	migration() //建表
}
