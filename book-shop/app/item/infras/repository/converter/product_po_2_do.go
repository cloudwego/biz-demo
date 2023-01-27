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

type productPO2DOConverter struct{}

var ProductPO2DOConverter = productPO2DOConverter{}

func (converter *productPO2DOConverter) Convert2do(ctx context.Context, po *po.Product) (*entity.ProductEntity, error) {
	do := &entity.ProductEntity{
		ProductId:   po.ProductId,
		Name:        po.Name,
		Pic:         po.Pic,
		Description: po.Description,
		Property: &entity.PropertyEntity{
			ISBN:     po.ISBN,
			SpuName:  po.SpuName,
			SpuPrice: po.SpuPrice,
		},
		Price:  po.Price,
		Stock:  po.Stock,
		Status: po.Status,
	}

	return do, nil
}
