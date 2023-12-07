package routes

import (
	"context"
	"fmt"
	"github.com/baiyutang/gomall/app/frontend/kitex_gen/product"
	"github.com/baiyutang/gomall/app/frontend/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/kitex/client"
)

func RegisterCategory(h *server.Hertz) {

	productClient, _ := productcatalogservice.NewClient("product", client.WithHostPorts("localhost:8881"))
	h.GET("/category/:category", func(ctx context.Context, c *app.RequestContext) {
		category := c.Param("category")
		fmt.Println(category)
		p, _ := productClient.ListProducts(ctx, &product.ListProductsReq{CategoryName: category})
		c.HTML(consts.StatusOK, "category", utils.H{
			"title":    "Category",
			"items":    p.Products,
			"cart_num": 10,
		})
	})

	h.GET("/category", func(ctx context.Context, c *app.RequestContext) {
		p, _ := productClient.ListProducts(ctx, &product.ListProductsReq{})
		c.HTML(consts.StatusOK, "category", utils.H{
			"title":    "Category",
			"items":    p.Products,
			"cart_num": 10,
		})
	})

	h.GET("/search", func(ctx context.Context, c *app.RequestContext) {
		p, _ := productClient.SearchProducts(ctx, &product.SearchProductsRequest{Query: c.Query("q")})
		c.HTML(consts.StatusOK, "search", utils.H{
			"q":        c.Query("q"),
			"items":    p.Results,
			"cart_num": 10,
		})
	})
}
