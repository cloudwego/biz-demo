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
	"errors"
	"strconv"

	"github.com/cloudwego/biz-demo/book-shop/app/facade/infras/client"
	"github.com/cloudwego/biz-demo/book-shop/app/facade/model"
	"github.com/cloudwego/biz-demo/book-shop/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
)

// GetOrder godoc
// @Summary get order by order_id
// @Description get order by order_id
// @Tags order module
// @Accept json
// @Produce json
// @Param order_id query int true "order id"
// @Security TokenAuth
// @Success 200 {object} model.Response
// @Router /order/get [get]
func GetOrder(ctx context.Context, c *app.RequestContext) {
	orderIdStr := c.Query("order_id")
	if orderIdStr == "" {
		model.SendResponse(c, errno.ConvertErr(errors.New("未传入order_id")), nil)
		return
	}

	orderId, err := strconv.ParseInt(orderIdStr, 10, 64)
	if err != nil {
		model.SendResponse(c, errno.ConvertErr(err), nil)
		return
	}

	order, err := client.GetOrderById(ctx, orderId)
	if err != nil {
		model.SendResponse(c, errno.ConvertErr(err), nil)
		return
	}

	model.SendResponse(c, errno.Success, order)
}
