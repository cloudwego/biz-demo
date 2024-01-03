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
	"errors"
	"time"

	"github.com/bytedance/sonic"
	"github.com/cloudwego/biz-demo/book-shop/kitex_gen/cwg/bookshop/item"
	"github.com/cloudwego/biz-demo/book-shop/kitex_gen/cwg/bookshop/item/itemservice"
	"github.com/cloudwego/biz-demo/book-shop/pkg/conf"
	"github.com/cloudwego/biz-demo/book-shop/pkg/errno"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var itemClient itemservice.Client

func initItemRpc() {
	r, err := etcd.NewEtcdResolver([]string{conf.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := itemservice.NewClient(
		conf.ItemRpcServiceName,
		client.WithRPCTimeout(3*time.Second),              // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond),    // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		client.WithResolver(r),                            // resolver
	)
	if err != nil {
		panic(err)
	}
	itemClient = c
}

func DecreaseStock(ctx context.Context, productId, stockNum int64) error {
	req := &item.DecrStockReq{
		ProductId: productId,
		StockNum:  stockNum,
	}
	resp, err := itemClient.DecrStock(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(int64(resp.BaseResp.StatusCode), resp.BaseResp.StatusMessage)
	}
	return nil
}

func DecreaseStockRevert(ctx context.Context, productId, stockNum int64) error {
	req := &item.DecrStockReq{
		ProductId: productId,
		StockNum:  stockNum,
	}
	resp, err := itemClient.DecrStockRevert(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(int64(resp.BaseResp.StatusCode), resp.BaseResp.StatusMessage)
	}
	return nil
}

func GetProductSnapshot(ctx context.Context, productId int64) (string, error) {
	req := &item.MGet2CReq{ProductIds: []int64{productId}}
	resp, err := itemClient.MGet2C(ctx, req)
	if err != nil {
		return "", err
	}
	if resp.BaseResp.StatusCode != 0 {
		return "", errno.NewErrNo(int64(resp.BaseResp.StatusCode), resp.BaseResp.StatusMessage)
	}
	if _, ok := resp.ProductMap[productId]; !ok {
		return "", errors.New("该商品不存在")
	}
	productStr, err := sonic.MarshalString(resp.ProductMap[productId])
	if err != nil {
		return "", err
	}
	return productStr, nil
}
