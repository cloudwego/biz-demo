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

package module

import (
	"context"

	"github.com/cloudwego/biz-demo/book-shop/app/order/common"
	"github.com/cloudwego/biz-demo/book-shop/app/order/dal/client"
	"github.com/cloudwego/biz-demo/book-shop/app/order/dal/db"
	"github.com/cloudwego/biz-demo/book-shop/kitex_gen/cwg/bookshop/order"
)

type UpdateModule struct {
	ctx context.Context
}

func NewUpdateModule(ctx context.Context) UpdateModule {
	return UpdateModule{
		ctx: ctx,
	}
}

func (m UpdateModule) CreateOrder(req *order.CreateOrderReq) error {
	po, err := common.ConvertCreateReq2PO(m.ctx, req)
	if err != nil {
		return err
	}
	// 扣减库存
	err = client.DecreaseStock(m.ctx, req.ProductId, req.StockNum)
	if err != nil {
		return err
	}
	poList := make([]*db.Order, 0)
	poList = append(poList, po)
	// 插入数据
	err = db.CreateOrder(m.ctx, poList)
	if err != nil {
		// 回滚
		m.createRollback(req)
		return err
	}
	return nil
}

func (m UpdateModule) createRollback(req *order.CreateOrderReq) {
	_ = client.DecreaseStockRevert(m.ctx, req.ProductId, req.StockNum)
}

func (m UpdateModule) CancelOrder(req *order.CancelOrderReq) error {
	updateMap := map[string]interface{}{
		"status": int64(order.Status_Cancel),
	}
	orderPO, err := db.GetOrderById(m.ctx, req.OrderId)
	if err != nil {
		return err
	}
	// 库存返还
	err = client.DecreaseStockRevert(m.ctx, orderPO.ProductId, orderPO.StockNum)
	if err != nil {
		return err
	}
	// 修改状态
	err = db.UpdateOrder(m.ctx, req.OrderId, updateMap)
	if err != nil {
		m.cancelRollback(orderPO.ProductId, orderPO.StockNum)
		return err
	}
	return nil
}

func (m UpdateModule) cancelRollback(productId, stockNum int64) {
	_ = client.DecreaseStock(m.ctx, productId, stockNum)
}
