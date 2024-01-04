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

package model

import (
	"context"

	"gorm.io/gorm"
)

type Consignee struct {
	Email string

	StreetAddress string
	City          string
	State         string
	Country       string
	ZipCode       int32
}

type Order struct {
	Base
	OrderId      string `gorm:"uniqueIndex;size:256"`
	UserId       uint32
	UserCurrency string
	Consignee    Consignee   `gorm:"embedded"`
	OrderItems   []OrderItem `gorm:"foreignKey:OrderIdRefer;references:OrderId"`
}

func (o Order) TableName() string {
	return "order"
}

func ListOrder(db *gorm.DB, ctx context.Context, userId uint32) (orders []Order, err error) {
	err = db.Debug().Model(&Order{}).Where(&Order{UserId: userId}).Preload("OrderItems").Find(&orders).Error
	return
}
