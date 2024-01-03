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

	"github.com/cloudwego/biz-demo/book-shop/app/facade/infras/client"
	"github.com/cloudwego/biz-demo/book-shop/app/facade/model"
	"github.com/cloudwego/biz-demo/book-shop/kitex_gen/cwg/bookshop/item"
	"github.com/cloudwego/biz-demo/book-shop/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
)

// SearchProduct godoc
// @Summary search products (2C interface)
// @Description search products (2C interface)
// @Tags product module(2C)
// @Accept json
// @Produce json
// @Param searchProductReq body model.SearchProductReq true "request param of searching products"
// @Security TokenAuth
// @Success 200 {object} model.Response
// @Router /item2c/search [post]
func SearchProduct(ctx context.Context, c *app.RequestContext) {
	var searchReq model.SearchProductReq
	if err := c.BindAndValidate(&searchReq); err != nil {
		model.SendResponse(c, errno.ConvertErr(err), nil)
		return
	}

	req := &item.SearchReq{
		Name:        searchReq.Name,
		Description: searchReq.Description,
		SpuName:     searchReq.SpuName,
	}
	products, err := client.SearchProduct(ctx, req)
	if err != nil {
		model.SendResponse(c, errno.ConvertErr(err), nil)
		return
	}

	model.SendResponse(c, errno.Success, products)
}
