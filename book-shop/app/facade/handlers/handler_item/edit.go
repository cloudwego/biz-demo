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
	"strconv"

	"github.com/cloudwego/biz-demo/book-shop/app/facade/infras/client"
	"github.com/cloudwego/biz-demo/book-shop/app/facade/model"
	"github.com/cloudwego/biz-demo/book-shop/kitex_gen/cwg/bookshop/item"
	"github.com/cloudwego/biz-demo/book-shop/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
)

// EditProduct godoc
// @Summary edit product
// @Description edit product
// @Tags product module
// @Accept json
// @Produce json
// @Param editProductRequest body model.EditProductRequest true "request param of editing product"
// @Security TokenAuth
// @Success 200 {object} model.Response
// @Router /item2b/edit [post]
func EditProduct(ctx context.Context, c *app.RequestContext) {
	var editReq model.EditProductRequest
	if err := c.BindAndValidate(&editReq); err != nil {
		model.SendResponse(c, errno.ConvertErr(err), nil)
		return
	}

	pid, err := strconv.ParseInt(editReq.ProductId, 10, 64)
	if err != nil {
		model.SendResponse(c, errno.ConvertErr(err), nil)
		return
	}

	req := &item.EditReq{
		ProductId:   pid,
		Name:        editReq.Name,
		Pic:         editReq.Pic,
		Description: editReq.Description,
		Price:       editReq.Price,
		Stock:       editReq.Stock,
	}
	property := &item.BookProperty{}
	if editReq.SpuPrice != nil {
		property.SpuPrice = *editReq.SpuPrice
	}
	if editReq.SpuName != nil {
		property.SpuName = *editReq.SpuName
	}
	if editReq.ISBN != nil {
		property.Isbn = *editReq.ISBN
	}
	if editReq.SpuPrice != nil || editReq.SpuName != nil || editReq.ISBN != nil {
		req.Property = property
	}
	err = client.EditProduct(ctx, req)
	if err != nil {
		model.SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	model.SendResponse(c, errno.Success, nil)
}
