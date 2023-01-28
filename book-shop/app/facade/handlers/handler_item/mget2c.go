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

package handler_item

import (
	"context"
	"errors"
	"strconv"
	"strings"

	"github.com/cloudwego/biz-demo/book-shop/app/facade/infras/client"
	"github.com/cloudwego/biz-demo/book-shop/app/facade/model"
	"github.com/cloudwego/biz-demo/book-shop/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
)

// MGetProduct2C godoc
// @Summary batch get products by product_id (2C interface)
// @Description batch get products by product_id (2C interface)
// @Tags product module(2C)
// @Accept json
// @Produce json
// @Param product_ids query string true "product-ids separated by commas"
// @Security TokenAuth
// @Success 200 {object} model.Response
// @Router /item2c/mget [get]
func MGetProduct2C(ctx context.Context, c *app.RequestContext) {
	productIdsStr := c.Query("product_ids")
	if productIdsStr == "" {
		model.SendResponse(c, errno.ConvertErr(errors.New("未传入product_id")), nil)
		return
	}
	productIdStrArr := strings.Split(productIdsStr, ",")
	productIds := make([]int64, 0)
	for _, v := range productIdStrArr {
		cur, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			model.SendResponse(c, errno.ConvertErr(errors.New("非法参数")), nil)
			return
		}
		productIds = append(productIds, cur)
	}

	products, err := client.MGetProducts2C(ctx, productIds)
	if err != nil {
		model.SendResponse(c, errno.ConvertErr(err), nil)
		return
	}

	model.SendResponse(c, errno.Success, products)
}
