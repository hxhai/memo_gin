package model

func migration() {
	//自动迁移
	DB.Set("gorm:table_options", "charset=utf8").AutoMigrate(&User{}).AutoMigrate(&Task{}) //迁移Task表
	DB.Model(&Task{}).AddForeignKey("uid", "User(id)", "CASCADE", "CASCADE")
}
