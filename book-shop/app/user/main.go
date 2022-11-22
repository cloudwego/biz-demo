package main

import (
	user "github.com/cloudwego/biz-demo/book-shop/kitex_gen/cwg/bookshop/user/userservice"
	"log"
)

func main() {
	svr := user.NewServer(new(UserServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
