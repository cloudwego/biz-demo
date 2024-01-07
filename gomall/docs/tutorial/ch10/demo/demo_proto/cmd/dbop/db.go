package main

import (
	"fmt"

	"github.com/cloudwego/biz-demo/gomall/demo/demo_proto/biz/dal"
	"github.com/cloudwego/biz-demo/gomall/demo/demo_proto/biz/dal/mysql"
	"github.com/cloudwego/biz-demo/gomall/demo/demo_proto/biz/model"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	dal.Init()
	mysql.DB.Create(&model.User{Email: "demo@example.com", Password: "a958af3702caf245d205da6164afebe0"})
	var u model.User
	mysql.DB.First(&u)
	fmt.Printf("%#v", u)
}
