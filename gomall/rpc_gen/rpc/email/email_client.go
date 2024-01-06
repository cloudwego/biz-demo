package email

import (
	"context"
	email "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/email"

	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/email/emailservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

type RPCClient interface {
	KitexClient() emailservice.Client
	Service() string
	Send(ctx context.Context, Req *email.EmailReq, callOptions ...callopt.Option) (r *email.EmailResp, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := emailservice.NewClient(dstService, opts...)
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
	kitexClient emailservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() emailservice.Client {
	return c.kitexClient
}

func (c *clientImpl) Send(ctx context.Context, Req *email.EmailReq, callOptions ...callopt.Option) (r *email.EmailResp, err error) {
	return c.kitexClient.Send(ctx, Req, callOptions...)
}
