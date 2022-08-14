package main

import (
	product "github.com/cloudwego/biz-demo/mall/cmd/cmp_ecom_product/kitex_gen/cmp/ecom/product/productservice"
	"log"
)

func main() {
	svr := product.NewServer(new(ProductServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
