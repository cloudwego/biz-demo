package service

import (
	"context"
	"strconv"

	common "github.com/cloudwego/biz-demo/gomall/app/frontend/hertz_gen/frontend/common"
	"github.com/cloudwego/biz-demo/gomall/app/frontend/infra/rpc"
	rpccart "github.com/cloudwego/biz-demo/gomall/app/frontend/kitex_gen/cart"
	rpcproduct "github.com/cloudwego/biz-demo/gomall/app/frontend/kitex_gen/product"
	frontendutils "github.com/cloudwego/biz-demo/gomall/app/frontend/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type GetCartService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetCartService(Context context.Context, RequestContext *app.RequestContext) *GetCartService {
	return &GetCartService{RequestContext: RequestContext, Context: Context}
}

func (h *GetCartService) Run(req *common.Empty) (resp map[string]any, err error) {
	var items []map[string]string
	carts, err := rpc.CartClient.GetCart(h.Context, &rpccart.GetCartRequest{
		UserId: uint32(h.Context.Value(frontendutils.UserIdKey).(float64)),
	})
	if err != nil {
		return nil, err
	}
	var total float32
	for _, v := range carts.Items {
		p, err := rpc.ProductClient.GetProduct(h.Context, &rpcproduct.GetProductRequest{Id: v.GetProductId()})
		if err != nil {
			continue
		}
		items = append(items, map[string]string{"Name": p.Name, "Description": p.Description, "Picture": p.Picture, "Price": strconv.FormatFloat(float64(p.Price), 'f', 2, 64), "Qty": strconv.Itoa(int(v.Quantity))})
		total += float32(v.Quantity) * p.Price
	}

	return utils.H{
		"title": "Cart",
		"items": items,
		"total": strconv.FormatFloat(float64(total), 'f', 2, 64),
	}, nil
}
