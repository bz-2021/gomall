package main

import (
	"github.com/bz-2021/gomall/app/user/biz/dal"
	"github.com/bz-2021/gomall/app/user/biz/dal/mysql"
	"github.com/bz-2021/gomall/app/user/biz/model"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	dal.Init()

	mysql.DB.Create(&model.User{
		Email: "bz2021@mail.sdu.edu.cn",
		Password: "string",
	})
}