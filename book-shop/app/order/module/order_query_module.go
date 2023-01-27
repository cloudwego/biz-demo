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

	"github.com/cloudwego/biz-demo/book-shop/app/order/dal/db"
)

type QueryModule struct {
	ctx context.Context
}

func NewQueryModule(ctx context.Context) QueryModule {
	return QueryModule{
		ctx: ctx,
	}
}

func (m QueryModule) ListOrder(userId int64, status *int64) ([]*db.Order, error) {
	filter := make(map[string]interface{})
	filter["user_id"] = userId
	if status != nil {
		filter["status"] = *status
	}
	res, err := db.ListOrders(m.ctx, filter)
	return res, err
}

func (m QueryModule) GetOrderById(orderId int64) (*db.Order, error) {
	po, err := db.GetOrderById(m.ctx, orderId)
	return po, err
}
