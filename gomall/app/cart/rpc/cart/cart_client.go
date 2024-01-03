package cart

import (
	"context"

	cart "github.com/baiyutang/gomall/app/cart/kitex_gen/cart"
	"github.com/baiyutang/gomall/app/cart/kitex_gen/cart/cartservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

type RPCClient interface {
	KitexClient() cartservice.Client
	Service() string
	AddItem(ctx context.Context, req *cart.AddItemRequest, callOptions ...callopt.Option) (resp *cart.Empty, err error)

	GetCart(ctx context.Context, req *cart.GetCartRequest, callOptions ...callopt.Option) (resp *cart.Cart, err error)

	EmptyCart(ctx context.Context, req *cart.EmptyCartRequest, callOptions ...callopt.Option) (resp *cart.Empty, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := cartservice.NewClient(dstService, opts...)
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
	kitexClient cartservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() cartservice.Client {
	return c.kitexClient
}

func (c *clientImpl) AddItem(ctx context.Context, req *cart.AddItemRequest, callOptions ...callopt.Option) (resp *cart.Empty, err error) {
	return c.kitexClient.AddItem(ctx, req, callOptions...)
}

func (c *clientImpl) GetCart(ctx context.Context, req *cart.GetCartRequest, callOptions ...callopt.Option) (resp *cart.Cart, err error) {
	return c.kitexClient.GetCart(ctx, req, callOptions...)
}

func (c *clientImpl) EmptyCart(ctx context.Context, req *cart.EmptyCartRequest, callOptions ...callopt.Option) (resp *cart.Empty, err error) {
	return c.kitexClient.EmptyCart(ctx, req, callOptions...)
}
