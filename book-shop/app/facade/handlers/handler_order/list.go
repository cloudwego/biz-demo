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

	"github.com/cloudwego/biz-demo/book-shop/app/facade/infras/client"
	"github.com/cloudwego/biz-demo/book-shop/app/facade/model"
	"github.com/cloudwego/biz-demo/book-shop/pkg/conf"
	"github.com/cloudwego/biz-demo/book-shop/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/jwt"
)

// ListOrder godoc
// @Summary get order list of a consumer
// @Description get order list of a consumer
// @Tags order module
// @Accept json
// @Produce json
// @Param listOrderReq body model.ListOrderReq true "request param to get order list"
// @Security TokenAuth
// @Success 200 {object} model.Response
// @Router /order/list [post]
func ListOrder(ctx context.Context, c *app.RequestContext) {
	var listReq model.ListOrderReq
	if err := c.BindAndValidate(&listReq); err != nil {
		model.SendResponse(c, errno.ConvertErr(err), nil)
		return
	}

	claims := jwt.ExtractClaims(ctx, c)
	userID := int64(claims[conf.IdentityKey].(float64))

	orders, err := client.ListOrder(ctx, userID, listReq.Status)
	if err != nil {
		model.SendResponse(c, errno.ConvertErr(err), nil)
		return
	}

	model.SendResponse(c, errno.Success, orders)
}
