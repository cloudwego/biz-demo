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

package main

import (
	"context"

	"github.com/cloudwego/biz-demo/book-shop/app/order/common"
	"github.com/cloudwego/biz-demo/book-shop/app/order/module"
	order "github.com/cloudwego/biz-demo/book-shop/kitex_gen/cwg/bookshop/order"
	"github.com/cloudwego/biz-demo/book-shop/pkg/errno"
)

// OrderServiceImpl implements the last service interface defined in the IDL.
type OrderServiceImpl struct{}

// CreateOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) CreateOrder(ctx context.Context, req *order.CreateOrderReq) (resp *order.CreateOrderResp, err error) {
	resp = order.NewCreateOrderResp()
	updateModule := module.NewUpdateModule(ctx)
	err = updateModule.CreateOrder(req)
	if err != nil {
		resp.BaseResp = errno.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = errno.BuildBaseResp(errno.Success)
	return resp, nil
}

// CancelOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) CancelOrder(ctx context.Context, req *order.CancelOrderReq) (resp *order.CancelOrderResp, err error) {
	resp = order.NewCancelOrderResp()
	updateModule := module.NewUpdateModule(ctx)
	err = updateModule.CancelOrder(req)
	if err != nil {
		resp.BaseResp = errno.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = errno.BuildBaseResp(errno.Success)
	return resp, nil
}

// ListOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) ListOrder(ctx context.Context, req *order.ListOrderReq) (resp *order.ListOrderResp, err error) {
	resp = order.NewListOrderResp()
	queryModule := module.NewQueryModule(ctx)
	pos, err := queryModule.ListOrder(req.UserId, (*int64)(req.Status))
	if err != nil {
		resp.BaseResp = errno.BuildBaseResp(err)
		return resp, nil
	}
	orders := make([]*order.OrderItem, 0)
	for _, v := range pos {
		orders = append(orders, common.ConvertPO2DTO(ctx, v))
	}
	resp.Orders = orders
	resp.BaseResp = errno.BuildBaseResp(errno.Success)
	return resp, nil
}

// GetOrderById implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) GetOrderById(ctx context.Context, req *order.GetOrderByIdReq) (resp *order.GetOrderByIdResp, err error) {
	resp = order.NewGetOrderByIdResp()
	queryModule := module.NewQueryModule(ctx)
	po, err := queryModule.GetOrderById(req.OrderId)
	if err != nil {
		resp.BaseResp = errno.BuildBaseResp(err)
		return resp, nil
	}
	resp.Order = common.ConvertPO2DTO(ctx, po)
	resp.BaseResp = errno.BuildBaseResp(errno.Success)
	return resp, nil
}
