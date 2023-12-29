package service

import (
	"context"
	"strconv"

	"github.com/baiyutang/gomall/app/frontend/hertz_gen/frontend/checkout"
	"github.com/baiyutang/gomall/app/frontend/infra/rpc"
	rpccart "github.com/baiyutang/gomall/app/frontend/kitex_gen/cart"
	rpcproduct "github.com/baiyutang/gomall/app/frontend/kitex_gen/product"
	frontendutils "github.com/baiyutang/gomall/app/frontend/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type CheckoutService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCheckoutService(Context context.Context, RequestContext *app.RequestContext) *CheckoutService {
	return &CheckoutService{RequestContext: RequestContext, Context: Context}
}

func (h *CheckoutService) Run(req *checkout.CheckoutReq) (resp map[string]any, err error) {
	var items []map[string]string
	userId := frontendutils.GetUserIdFromCtx(h.Context)

	carts, err := rpc.CartClient.GetCart(h.Context, &rpccart.GetCartRequest{UserId: userId})
	if err != nil {
	}
	var total float32
	for _, v := range carts.Items {
		p, err := rpc.ProductClient.GetProduct(h.Context, &rpcproduct.GetProductRequest{Id: v.ProductId})
		if err != nil {
			continue
		}
		items = append(items, map[string]string{
			"Name":    p.Name,
			"Price":   strconv.FormatFloat(float64(p.Price), 'f', 2, 64),
			"Picture": p.Picture,
			"Qty":     strconv.Itoa(int(v.Quantity)),
		})
		total += float32(v.Quantity) * p.Price
	}

	return utils.H{

		"title":    "Checkout",
		"items":    items,
		"cart_num": len(items),
		"total":    strconv.FormatFloat(float64(total), 'f', 2, 64),
	}, nil
}
