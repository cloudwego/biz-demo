package main

import (
	order "github.com/cloudwego/biz-demo/book-shop/kitex_gen/cwg/bookshop/order/orderservice"
	"log"
)

func main() {
	svr := order.NewServer(new(OrderServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
