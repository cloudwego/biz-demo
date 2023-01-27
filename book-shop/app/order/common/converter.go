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

package common

import (
	"context"

	"github.com/cloudwego/biz-demo/book-shop/app/order/dal/client"
	"github.com/cloudwego/biz-demo/book-shop/app/order/dal/db"
	"github.com/cloudwego/biz-demo/book-shop/kitex_gen/cwg/bookshop/order"
	"github.com/cloudwego/biz-demo/book-shop/pkg/utils"
)

func ConvertCreateReq2PO(ctx context.Context, req *order.CreateOrderReq) (*db.Order, error) {
	orderId, err := utils.GenerateID()
	if err != nil {
		return nil, err
	}
	snapshot, err := client.GetProductSnapshot(ctx, req.GetProductId())
	if err != nil {
		return nil, err
	}

	ret := &db.Order{
		OrderId:         orderId,
		UserId:          req.UserId,
		Address:         req.Address,
		ProductId:       req.ProductId,
		StockNum:        req.StockNum,
		ProductSnapshot: snapshot,
		Status:          int64(order.Status_Finish),
	}
	return ret, nil
}

func ConvertPO2DTO(ctx context.Context, po *db.Order) *order.OrderItem {
	ret := &order.OrderItem{
		OrderId:         po.OrderId,
		UserId:          po.UserId,
		UserName:        "",
		Address:         po.Address,
		ProductId:       po.ProductId,
		StockNum:        po.StockNum,
		ProductSnapshot: po.ProductSnapshot,
		Status:          order.Status(po.Status),
		CreateTime:      po.CreatedAt.Unix(),
		UpdateTime:      po.UpdatedAt.Unix(),
	}
	userName, err := client.GetUserName(ctx, po.UserId)
	if err == nil {
		ret.UserName = userName
	}

	return ret
}
