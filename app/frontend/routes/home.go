package routes

import (
	"context"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/kitex/client"

	"github.com/baiyutang/gomall/app/frontend/kitex_gen/product"
	"github.com/baiyutang/gomall/app/frontend/kitex_gen/product/productcatalogservice"
	frontendutils "github.com/baiyutang/gomall/app/frontend/utils"
)

func RegisterHome(h *server.Hertz) {
	productClient, _ := productcatalogservice.NewClient("product", client.WithHostPorts("localhost:8881"))
	h.GET("/", func(ctx context.Context, c *app.RequestContext) {
		p, _ := productClient.ListProducts(ctx, &product.ListProductsReq{})
		c.HTML(http.StatusOK, "home", frontendutils.WarpResponse(ctx, c, utils.H{
			"title":    "Hot sale",
			"cart_num": 10,
			"items":    p.Products,
		}))
	})
}
