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

package handler_order

import (
	"context"
	"strconv"

	"github.com/cloudwego/biz-demo/book-shop/app/facade/infras/client"
	"github.com/cloudwego/biz-demo/book-shop/app/facade/model"
	"github.com/cloudwego/biz-demo/book-shop/kitex_gen/cwg/bookshop/order"
	"github.com/cloudwego/biz-demo/book-shop/pkg/conf"
	"github.com/cloudwego/biz-demo/book-shop/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/jwt"
)

// CreateOrder godoc
// @Summary consumer creates order
// @Description consumer creates order
// @Tags order module
// @Accept json
// @Produce json
// @Param createOrderReq body model.CreateOrderReq true "request param to create one order"
// @Security TokenAuth
// @Success 200 {object} model.Response
// @Router /order/create [post]
func CreateOrder(ctx context.Context, c *app.RequestContext) {
	var createReq model.CreateOrderReq
	if err := c.BindAndValidate(&createReq); err != nil {
		model.SendResponse(c, errno.ConvertErr(err), nil)
		return
	}

	claims := jwt.ExtractClaims(ctx, c)
	userID := int64(claims[conf.IdentityKey].(float64))

	pid, err := strconv.ParseInt(createReq.ProductId, 10, 64)
	if err != nil {
		model.SendResponse(c, errno.ConvertErr(err), nil)
		return
	}

	req := &order.CreateOrderReq{
		UserId:    userID,
		Address:   createReq.Address,
		ProductId: pid,
		StockNum:  createReq.StockNum,
	}
	err = client.CreateOrder(ctx, req)
	if err != nil {
		model.SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	model.SendResponse(c, errno.Success, nil)
}
