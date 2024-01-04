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

package cart

import (
	"context"

	cart "github.com/cloudwego/biz-demo/gomall/app/checkout/kitex_gen/cart"

	"github.com/cloudwego/biz-demo/gomall/app/checkout/kitex_gen/cart/cartservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

type RPCClient interface {
	KitexClient() cartservice.Client
	Service() string
	AddItem(ctx context.Context, Req *cart.AddItemRequest, callOptions ...callopt.Option) (r *cart.Empty, err error)
	GetCart(ctx context.Context, Req *cart.GetCartRequest, callOptions ...callopt.Option) (r *cart.Cart, err error)
	EmptyCart(ctx context.Context, Req *cart.EmptyCartRequest, callOptions ...callopt.Option) (r *cart.Empty, err error)
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

func (c *clientImpl) AddItem(ctx context.Context, Req *cart.AddItemRequest, callOptions ...callopt.Option) (r *cart.Empty, err error) {
	return c.kitexClient.AddItem(ctx, Req, callOptions...)
}

func (c *clientImpl) GetCart(ctx context.Context, Req *cart.GetCartRequest, callOptions ...callopt.Option) (r *cart.Cart, err error) {
	return c.kitexClient.GetCart(ctx, Req, callOptions...)
}

func (c *clientImpl) EmptyCart(ctx context.Context, Req *cart.EmptyCartRequest, callOptions ...callopt.Option) (r *cart.Empty, err error) {
	return c.kitexClient.EmptyCart(ctx, Req, callOptions...)
}
