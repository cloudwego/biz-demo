package service

import (
	"context"
	"fmt"
	common "github.com/baiyutang/gomall/app/frontend/hertz_gen/frontend/common"
	"github.com/baiyutang/gomall/app/frontend/infra/rpc"
	rpcorder "github.com/baiyutang/gomall/app/frontend/kitex_gen/order"
	"github.com/baiyutang/gomall/app/frontend/types"
	frontendutils "github.com/baiyutang/gomall/app/frontend/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type OrderListService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewOrderListService(Context context.Context, RequestContext *app.RequestContext) *OrderListService {
	return &OrderListService{RequestContext: RequestContext, Context: Context}
}

func (h *OrderListService) Run(req *common.Empty) (resp map[string]any, err error) {
	userId := frontendutils.GetUserIdFromCtx(h.Context)
	var orders []*types.Order
	listOrderResp, err := rpc.OrderClient.ListOrder(h.Context, &rpcorder.ListOrderRequest{UserId: userId})
	if err != nil {
		return nil, err
	}
	if listOrderResp != nil && len(listOrderResp.Orders) > 0 {
		for _, v := range listOrderResp.Orders {
			var items []types.OrderItem
			var total float32
			if len(v.OrderItems) > 0 {
				for _, vv := range v.OrderItems {
					total += vv.Cost
					i := vv.Item
					items = append(items, types.OrderItem{
						ProductId:   i.ProductId,
						Qty:         uint32(i.Quantity),
						ProductName: fmt.Sprintf("product - name %d", i.ProductId),
						Cost:        vv.Cost,
					})
				}
			}
			orders = append(orders, &types.Order{
				Cost:      total,
				Items:     items,
				CreatedAt: v.CreatedAt,
				OrderId:   v.OrderId,
				Consignee: types.Consignee{Email: v.Email},
			})
		}
	}

	return utils.H{
		"title":  "Order",
		"orders": orders,
	}, nil
}
