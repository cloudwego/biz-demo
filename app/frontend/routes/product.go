package routes

import (
	"context"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/kitex/client"

	"github.com/baiyutang/gomall/app/frontend/kitex_gen/product"
	"github.com/baiyutang/gomall/app/frontend/kitex_gen/product/productcatalogservice"
)

func RegisterProductRoute(h *server.Hertz) {
	productClient, _ := productcatalogservice.NewClient("product", client.WithHostPorts("localhost:8881"))
	h.GET("/product", func(ctx context.Context, c *app.RequestContext) {
		p, _ := productClient.ListProducts(context.Background(), &product.ListProductsReq{CategoryNames: []string{"T-Shirt"}})
		fmt.Println(p)
		c.HTML(consts.StatusOK, "product", utils.H{
			"title":   "Product",
			"product": p,
		})
	})
}
