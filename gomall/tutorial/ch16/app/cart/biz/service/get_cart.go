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

package service

import (
	"context"

	"github.com/cloudwego/biz-demo/gomall/app/cart/biz/dal/mysql"
	"github.com/cloudwego/biz-demo/gomall/app/cart/biz/model"
	cart "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/cart"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

type GetCartService struct {
	ctx context.Context
} // NewGetCartService new GetCartService
func NewGetCartService(ctx context.Context) *GetCartService {
	return &GetCartService{ctx: ctx}
}

// Run create note info
func (s *GetCartService) Run(req *cart.GetCartReq) (resp *cart.GetCartResp, err error) {
	// resp = &cart.Cart{}
	// Finish your business logic.
	carts, err := model.GetCartByUserId(mysql.DB, s.ctx, req.GetUserId())
	if err != nil {
		return nil, kerrors.NewBizStatusError(50000, err.Error())
	}
	var items []*cart.CartItem
	for _, v := range carts {
		items = append(items, &cart.CartItem{ProductId: v.ProductId, Quantity: int32(v.Qty)})
	}

	return &cart.GetCartResp{Cart: &cart.Cart{UserId: req.GetUserId(), Items: items}}, nil
}
