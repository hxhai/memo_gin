package model

func migration() {
	//自动迁移到mysql
	DB.Set("gorm:table_options", "charset=utf8").AutoMigrate(&User{}).AutoMigrate(&Task{}) //迁移User表和Task表
	//AddForeignKey 将外键添加到给定的作用域
	DB.Model(&Task{}).AddForeignKey("uid", "User(id)", "CASCADE", "CASCADE") //外键关联uid关联User中的id
}
