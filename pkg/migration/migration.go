package migration

import (
	"fmt"

	"www.blog.com/config"
	"www.blog.com/model"
)

func Migrate() {
	// migrate does migration of models to database

	db := config.GoConnect()
	//
	db.AutoMigrate(&model.User{}, &model.Post{})
	fmt.Println("Migration Successfull")
}
