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

func AddProduct(ctx context.Context, req *item.AddReq) (int64, error) {
	resp, err := itemClient.Add(ctx, req)
	if err != nil {
		return 0, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return 0, errno.NewErrNo(int64(resp.BaseResp.StatusCode), resp.BaseResp.StatusMessage)
	}
	return resp.ProductId, nil
}

func EditProduct(ctx context.Context, req *item.EditReq) error {
	resp, err := itemClient.Edit(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(int64(resp.BaseResp.StatusCode), resp.BaseResp.StatusMessage)
	}
	return nil
}

func OperateProduct(ctx context.Context, productId int64, operate string) error {
	if operate == "del" {
		resp, err := itemClient.Delete(ctx, &item.DeleteReq{ProductId: productId})
		if err != nil {
			return err
		}
		if resp.BaseResp.StatusCode != 0 {
			return errno.NewErrNo(int64(resp.BaseResp.StatusCode), resp.BaseResp.StatusMessage)
		}
	} else if operate == "offline" {
		resp, err := itemClient.Offline(ctx, &item.OfflineReq{ProductId: productId})
		if err != nil {
			return err
		}
		if resp.BaseResp.StatusCode != 0 {
			return errno.NewErrNo(int64(resp.BaseResp.StatusCode), resp.BaseResp.StatusMessage)
		}
	} else if operate == "online" {
		resp, err := itemClient.Online(ctx, &item.OnlineReq{ProductId: productId})
		if err != nil {
			return err
		}
		if resp.BaseResp.StatusCode != 0 {
			return errno.NewErrNo(int64(resp.BaseResp.StatusCode), resp.BaseResp.StatusMessage)
		}
	}

	return nil
}

func GetProduct(ctx context.Context, productId int64) (*item.Product, error) {
	resp, err := itemClient.Get(ctx, &item.GetReq{ProductId: productId})
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(int64(resp.BaseResp.StatusCode), resp.BaseResp.StatusMessage)
	}
	return resp.Product, nil
}

func MGetProducts2C(ctx context.Context, productIds []int64) (map[int64]*item.Product, error) {
	resp, err := itemClient.MGet2C(ctx, &item.MGet2CReq{
		ProductIds: productIds,
	})
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(int64(resp.BaseResp.StatusCode), resp.BaseResp.StatusMessage)
	}
	return resp.ProductMap, nil
}

func SearchProduct(ctx context.Context, req *item.SearchReq) ([]*item.Product, error) {
	resp, err := itemClient.Search(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(int64(resp.BaseResp.StatusCode), resp.BaseResp.StatusMessage)
	}
	return resp.Products, nil
}

func ListProduct(ctx context.Context, req *item.ListReq) ([]*item.Product, error) {
	resp, err := itemClient.List(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(int64(resp.BaseResp.StatusCode), resp.BaseResp.StatusMessage)
	}
	return resp.Products, nil
}
