package main

import (
	item "github.com/cloudwego/biz-demo/book-shop/kitex_gen/cwg/bookshop/item/itemservice"
	"log"
)

func main() {
	svr := item.NewServer(new(ItemServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
