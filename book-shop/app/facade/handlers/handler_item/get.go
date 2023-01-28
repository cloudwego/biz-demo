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

	"github.com/cloudwego/biz-demo/book-shop/app/facade/infras/client"
	"github.com/cloudwego/biz-demo/book-shop/app/facade/model"
	"github.com/cloudwego/biz-demo/book-shop/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
)

// GetProduct godoc
// @Summary get product by product_id
// @Description get product by product_id
// @Tags product module
// @Accept json
// @Produce json
// @Param product_id query int true "product_id"
// @Security TokenAuth
// @Success 200 {object} model.Response
// @Router /item2b/get [get]
func GetProduct(ctx context.Context, c *app.RequestContext) {
	productIdStr := c.Query("product_id")
	if productIdStr == "" {
		model.SendResponse(c, errno.ConvertErr(errors.New("未传入product_id")), nil)
		return
	}

	productId, err := strconv.ParseInt(productIdStr, 10, 64)
	if err != nil {
		model.SendResponse(c, errno.ConvertErr(err), nil)
		return
	}

	product, err := client.GetProduct(ctx, productId)
	if err != nil {
		model.SendResponse(c, errno.ConvertErr(err), nil)
		return
	}

	model.SendResponse(c, errno.Success, product)
}
