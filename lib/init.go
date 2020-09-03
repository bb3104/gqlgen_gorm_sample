package main

import (
	"fmt"
	"github.com/bb3104/gqlgen_gorm_sample/db_model"
)

func main() {
	db := db_model.GormConnect()
	db.AutoMigrate(&db_model.User{})
	db.AutoMigrate(&db_model.Item{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")

	user := db_model.User{}
	user.Name = "hogehoge"
	user.Password = "hogehoge"
	user.Email = "hogehoge@gmail.com"
	db.Create(&user)

	users := []*db_model.User{}
	err := db.Find(&users).First

	item := db_model.Item{}
	item.Name = "item"
	item.UserID = users[0].ID
	db.Create(&item)

	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	db.LogMode(true)
}
