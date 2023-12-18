package routes

import (
	"context"
	"fmt"
	"github.com/baiyutang/gomall/app/frontend/infra/rpc"
	"github.com/cloudwego/kitex/pkg/klog"
	"net/http"

	"github.com/baiyutang/gomall/app/frontend/kitex_gen/product"
	frontendutils "github.com/baiyutang/gomall/app/frontend/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

func RegisterHome(h *server.Hertz) {
	productClient := rpc.ProductClient
	fmt.Printf("%#v", productClient)
	h.GET("/", func(ctx context.Context, c *app.RequestContext) {
		p, err := productClient.ListProducts(ctx, &product.ListProductsReq{})
		if err != nil {
			klog.Error(err)
		}
		c.HTML(http.StatusOK, "home", frontendutils.WarpResponse(ctx, c, utils.H{
			"title":    "Hot sale",
			"cart_num": 10,
			"items":    p.Products,
		}))
	})
}
