package routes

import (
	"context"
	"github.com/baiyutang/gomall/app/frontend/infra/rpc"
	"github.com/baiyutang/gomall/app/frontend/kitex_gen/product"
	frontendutils "github.com/baiyutang/gomall/app/frontend/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func RegisterCategory(h *server.Hertz) {

	productClient := rpc.ProductClient
	h.GET("/category/:category", func(ctx context.Context, c *app.RequestContext) {
		category := c.Param("category")
		p, _ := productClient.ListProducts(ctx, &product.ListProductsReq{CategoryName: category})
		c.HTML(consts.StatusOK, "category", frontendutils.WarpResponse(ctx, c, utils.H{
			"title":    "Category",
			"items":    p.Products,
			"cart_num": 10,
		}))
	})

	h.GET("/category", func(ctx context.Context, c *app.RequestContext) {
		p, _ := productClient.ListProducts(ctx, &product.ListProductsReq{})
		c.HTML(consts.StatusOK, "category", frontendutils.WarpResponse(ctx, c, utils.H{
			"title":    "Category",
			"items":    p.Products,
			"cart_num": 10,
		}))
	})

	h.GET("/search", func(ctx context.Context, c *app.RequestContext) {
		p, _ := productClient.SearchProducts(ctx, &product.SearchProductsRequest{Query: c.Query("q")})
		c.HTML(consts.StatusOK, "search", frontendutils.WarpResponse(ctx, c, utils.H{
			"q":        c.Query("q"),
			"items":    p.Results,
			"cart_num": 10,
		}))
	})
}
