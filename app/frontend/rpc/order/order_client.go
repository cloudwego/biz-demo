package order

import (
	"context"
	"context"
	order "github.com/baiyutang/gomall/app/frontend/kitex_gen/order"

	"github.com/baiyutang/gomall/app/frontend/kitex_gen/order/orderservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

type RPCClient interface {
	KitexClient() orderservice.Client
	Service() string
	PlaceOrder(ctx context.Context, req *order.PlaceOrderRequest, callOptions ...callopt.Option) (resp *order.PlaceOrderResponse, err error)

	ListOrder(ctx context.Context, req *order.ListOrderRequest, callOptions ...callopt.Option) (resp *order.ListOrderResponse, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := orderservice.NewClient(dstService, opts...)
	if err != nil {
		return nil, err
	}
	cli := &clientImpl{
		service:     dstService,
		kitexClient: kitexClient,
	}

	return cli, nil
}

type clientImpl struct {
	service     string
	kitexClient orderservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() orderservice.Client {
	return c.kitexClient
}

func (c *clientImpl) PlaceOrder(ctx context.Context, req *order.PlaceOrderRequest, callOptions ...callopt.Option) (resp *order.PlaceOrderResponse, err error) {
	return c.kitexClient.PlaceOrder(ctx, req, callOptions...)
}

func (c *clientImpl) ListOrder(ctx context.Context, req *order.ListOrderRequest, callOptions ...callopt.Option) (resp *order.ListOrderResponse, err error) {
	return c.kitexClient.ListOrder(ctx, req, callOptions...)
}
