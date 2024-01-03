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

package checkout

import (
	"context"

	checkout "github.com/cloudwego/biz-demo/gomall/app/frontend/kitex_gen/checkout"

	"github.com/cloudwego/biz-demo/gomall/app/frontend/kitex_gen/checkout/checkoutservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

type RPCClient interface {
	KitexClient() checkoutservice.Client
	Service() string
	Checkout(ctx context.Context, Req *checkout.CheckoutReq, callOptions ...callopt.Option) (r *checkout.CheckoutRes, err error)
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

func (c *clientImpl) Checkout(ctx context.Context, Req *checkout.CheckoutReq, callOptions ...callopt.Option) (r *checkout.CheckoutRes, err error) {
	return c.kitexClient.Checkout(ctx, Req, callOptions...)
}
