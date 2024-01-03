package order

import (
	"context"

	order "github.com/cloudwego/biz-demo/gomall/app/checkout/kitex_gen/order"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func PlaceOrder(ctx context.Context, req *order.PlaceOrderRequest, callOptions ...callopt.Option) (resp *order.PlaceOrderResponse, err error) {
	resp, err = defaultClient.PlaceOrder(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "PlaceOrder call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
