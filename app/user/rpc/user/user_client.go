package user

import (
	"context"
	"context"
	user "github.com/baiyutang/gomall/app/user/kitex_gen/user"

	"github.com/baiyutang/gomall/app/user/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

type RPCClient interface {
	KitexClient() userservice.Client
	Service() string
	Register(ctx context.Context, req *user.RegisterReq, callOptions ...callopt.Option) (resp *user.RegisterRes, err error)

	Login(ctx context.Context, req *user.LoginReq, callOptions ...callopt.Option) (resp *user.LoginRes, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := userservice.NewClient(dstService, opts...)
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
	kitexClient userservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() userservice.Client {
	return c.kitexClient
}

func (c *clientImpl) Register(ctx context.Context, req *user.RegisterReq, callOptions ...callopt.Option) (resp *user.RegisterRes, err error) {
	return c.kitexClient.Register(ctx, req, callOptions...)
}

func (c *clientImpl) Login(ctx context.Context, req *user.LoginReq, callOptions ...callopt.Option) (resp *user.LoginRes, err error) {
	return c.kitexClient.Login(ctx, req, callOptions...)
}
