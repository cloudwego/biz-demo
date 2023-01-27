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

package converter

import (
	"context"

	"github.com/cloudwego/biz-demo/book-shop/app/item/common/entity"
	"github.com/cloudwego/biz-demo/book-shop/app/item/common/po"
)

type productDO2POConverter struct{}

var ProductDO2POConverter = &productDO2POConverter{}

func (converter *productDO2POConverter) Convert2po(ctx context.Context, do *entity.ProductEntity) (*po.Product, error) {
	po := &po.Product{
		ProductId:   do.ProductId,
		Name:        do.Name,
		Pic:         do.Pic,
		Description: do.Description,
		ISBN:        "",
		SpuName:     "",
		SpuPrice:    0,
		Price:       do.Price,
		Stock:       do.Stock,
		Status:      do.Status,
	}
	if do.Property != nil {
		po.ISBN = do.Property.ISBN
		po.SpuName = do.Property.SpuName
		po.SpuPrice = do.Property.SpuPrice
	}

	return po, nil
}
