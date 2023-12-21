package checkout

import (
	"context"
	checkout "github.com/baiyutang/gomall/app/checkout/kitex_gen/checkout"

	"github.com/baiyutang/gomall/app/checkout/kitex_gen/checkout/checkoutservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

type RPCClient interface {
	KitexClient() checkoutservice.Client
	Service() string
	Checkout(ctx context.Context, req *checkout.CheckoutReq, callOptions ...callopt.Option) (resp *checkout.CheckoutRes, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := checkoutservice.NewClient(dstService, opts...)
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
	kitexClient checkoutservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() checkoutservice.Client {
	return c.kitexClient
}

func (c *clientImpl) Checkout(ctx context.Context, req *checkout.CheckoutReq, callOptions ...callopt.Option) (resp *checkout.CheckoutRes, err error) {
	return c.kitexClient.Checkout(ctx, req, callOptions...)
}
