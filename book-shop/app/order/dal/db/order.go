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

package db

import (
	"context"
	"errors"

	"github.com/cloudwego/biz-demo/book-shop/pkg/conf"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	OrderId         int64  `json:"order_id"`
	UserId          int64  `json:"user_id"`
	Address         string `json:"address"`
	ProductId       int64  `json:"product_id"`
	StockNum        int64  `json:"stock_num"`
	ProductSnapshot string `json:"product_snapshot"`
	Status          int64  `json:"status"`
}

func (o *Order) TableName() string {
	return conf.OrderTableName
}

func CreateOrder(ctx context.Context, orders []*Order) error {
	return DB.WithContext(ctx).Create(orders).Error
}

func UpdateOrder(ctx context.Context, orderId int64, updateMap map[string]interface{}) error {
	return DB.WithContext(ctx).Model(&Order{}).Where("order_id = ?", orderId).
		Updates(updateMap).Error
}

func ListOrders(ctx context.Context, filterMap map[string]interface{}) ([]*Order, error) {
	res := make([]*Order, 0)
	db := DB.WithContext(ctx)
	for k, v := range filterMap {
		db = db.Where(k+" = ?", v)
	}
	err := db.Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func GetOrderById(ctx context.Context, orderId int64) (*Order, error) {
	res := make([]*Order, 0)
	err := DB.WithContext(ctx).Where("order_id = ?", orderId).Find(&res).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	if len(res) == 0 {
		return nil, errors.New("不存在该订单")
	}
	return res[0], nil
}
