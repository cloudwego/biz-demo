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

// AddProduct godoc
// @Summary add product
// @Description add product
// @Tags product module
// @Accept json
// @Produce json
// @Param addProductRequest body model.AddProductRequest true "request param of adding product"
// @Security TokenAuth
// @Success 200 {object} model.Response
// @Router /item2b/add [post]
func AddProduct(ctx context.Context, c *app.RequestContext) {
	var addReq model.AddProductRequest
	if err := c.BindAndValidate(&addReq); err != nil {
		model.SendResponse(c, errno.ConvertErr(err), nil)
		return
	}

	req := &item.AddReq{
		Name:        addReq.Name,
		Pic:         addReq.Pic,
		Description: addReq.Description,
		Property: &item.BookProperty{
			Isbn:     addReq.ISBN,
			SpuName:  addReq.SpuName,
			SpuPrice: addReq.SpuPrice,
		},
		Price: addReq.Price,
		Stock: addReq.Stock,
	}
	pid, err := client.AddProduct(ctx, req)
	if err != nil {
		model.SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	model.SendResponse(c, errno.Success, map[string]interface{}{
		"product_id": strconv.FormatInt(pid, 10),
	})
}
