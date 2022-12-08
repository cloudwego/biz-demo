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

package repository

import (
	"context"
	"fmt"

	"github.com/cloudwego/biz-demo/open-payment-platform/internal/payment/entity"
	"github.com/cloudwego/biz-demo/open-payment-platform/internal/payment/infrastructure/ent"
	"github.com/cloudwego/biz-demo/open-payment-platform/internal/payment/infrastructure/ent/order"
	"github.com/cloudwego/biz-demo/open-payment-platform/internal/payment/usecase"
	"github.com/cloudwego/kitex/pkg/klog"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"github.com/pkg/errors"
)

const (
	_dataBase = "payment"
)

type OrderRepository struct {
	db *ent.Client
}

var SQLProviderSet = wire.NewSet(NewEntClient, NewOrderSQL)

var _ usecase.Repository = (*OrderRepository)(nil)

// NewEntClient create an ent client with default config
func NewEntClient() *ent.Client {
	entClient, err := ent.Open(
		"mysql",
		fmt.Sprintf("root:root@tcp(127.0.0.1:3306)/%s?parseTime=true", _dataBase),
	)
	if err != nil {
		klog.Fatal(err)
	}
	if err = entClient.Schema.Create(context.TODO()); err != nil {
		klog.Fatalf("failed creating schema resources: %v", err)
	}
	return entClient
}

// GetByOutOrderNo implements usecase.Repository.GetByOutOrderNo
func (o *OrderRepository) GetByOutOrderNo(ctx context.Context, outOrderNo string) (*entity.Order, error) {
	ret, err := o.db.Order.Query().
		Where(order.OutOrderNo(outOrderNo)).
		Limit(1).
		All(ctx)
	if err != nil {
		klog.Error(err)
		return nil, err
	}
	if len(ret) == 0 {
		return nil, errors.New("OrdersNoFound")
	}
	row := ret[0]
	return &entity.Order{
		ID:              row.ID,
		MerchantID:      row.MerchantID,
		Channel:         row.Channel,
		PayWay:          row.PayWay,
		OutOrderNo:      row.OutOrderNo,
		OrderStatus:     row.OrderStatus,
		TotalAmount:     row.TotalAmount,
		Body:            row.Body,
		AuthCode:        row.AuthCode,
		WxAppid:         row.WxAppid,
		SubOpenid:       row.SubOpenid,
		JumpURL:         row.JumpURL,
		NotifyURL:       row.NotifyURL,
		ClientIP:        row.ClientIP,
		Attach:          row.Attach,
		OrderExpiration: row.OrderExpiration,
		ExtendParams:    row.ExtendParams,
	}, nil
}

// Create implements usecase.Repository.Create
func (o *OrderRepository) Create(ctx context.Context, order *entity.Order) error {
	ret, err := o.db.Order.Create().
		SetMerchantID(order.MerchantID).
		SetPayWay(order.PayWay).
		SetTotalAmount(order.TotalAmount).
		SetBody(order.Body).
		SetAttach(order.Attach).
		SetChannel(order.Channel).
		SetClientIP(order.ClientIP).
		SetAuthCode(order.AuthCode).
		SetJumpURL(order.JumpURL).
		SetNotifyURL(order.NotifyURL).
		SetOrderExpiration(order.OrderExpiration).
		SetSubOpenid(order.SubOpenid).
		SetOutOrderNo(order.OutOrderNo).
		SetWxAppid(order.WxAppid).
		SetOrderStatus(order.OrderStatus).
		SetExtendParams(order.ExtendParams).
		Save(ctx)
	if err != nil {
		return err
	}
	order.ID = ret.ID
	return nil
}

// UpdateOrderStatus implements usecase.Repository.UpdateOrderStatus
func (o *OrderRepository) UpdateOrderStatus(ctx context.Context, outOrderNo string, orderStatus int8) error {
	return o.db.Order.Update().Where(order.OutOrderNo(outOrderNo)).SetOrderStatus(orderStatus).Exec(ctx)
}

// NewOrderSQL creates a new OrderRepository.
// This is the concrete implementation of Repository with SQL.
func NewOrderSQL(dbClient *ent.Client) usecase.Repository {
	return &OrderRepository{
		db: dbClient,
	}
}
