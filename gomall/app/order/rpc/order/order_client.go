// Copyright 2024 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package order

import (
	"context"

	order "github.com/cloudwego/biz-demo/gomall/app/order/kitex_gen/order"

	"github.com/cloudwego/biz-demo/gomall/app/order/kitex_gen/order/orderservice"
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
