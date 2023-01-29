// Copyright 2022 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package client

import (
	"context"
	"time"

	"github.com/cloudwego/biz-demo/book-shop/kitex_gen/cwg/bookshop/order"
	"github.com/cloudwego/biz-demo/book-shop/kitex_gen/cwg/bookshop/order/orderservice"
	"github.com/cloudwego/biz-demo/book-shop/pkg/conf"
	"github.com/cloudwego/biz-demo/book-shop/pkg/errno"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var orderClient orderservice.Client

func initOrderRpc() {
	r, err := etcd.NewEtcdResolver([]string{conf.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := orderservice.NewClient(
		conf.OrderRpcServiceName,
		client.WithRPCTimeout(3*time.Second),              // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond),    // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		client.WithResolver(r),                            // resolver
	)
	if err != nil {
		panic(err)
	}
	orderClient = c
}

func CreateOrder(ctx context.Context, req *order.CreateOrderReq) error {
	resp, err := orderClient.CreateOrder(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(int64(resp.BaseResp.StatusCode), resp.BaseResp.StatusMessage)
	}
	return nil
}

func CancelOrder(ctx context.Context, orderId int64) error {
	resp, err := orderClient.CancelOrder(ctx, &order.CancelOrderReq{OrderId: orderId})
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(int64(resp.BaseResp.StatusCode), resp.BaseResp.StatusMessage)
	}
	return nil
}

func GetOrderById(ctx context.Context, orderId int64) (*order.OrderItem, error) {
	resp, err := orderClient.GetOrderById(ctx, &order.GetOrderByIdReq{
		OrderId: orderId,
	})
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(int64(resp.BaseResp.StatusCode), resp.BaseResp.StatusMessage)
	}
	return resp.Order, nil
}

func ListOrder(ctx context.Context, userId int64, status *int64) ([]*order.OrderItem, error) {
	resp, err := orderClient.ListOrder(ctx, &order.ListOrderReq{
		UserId: userId,
		Status: (*order.Status)(status),
	})
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(int64(resp.BaseResp.StatusCode), resp.BaseResp.StatusMessage)
	}
	return resp.Orders, nil
}
