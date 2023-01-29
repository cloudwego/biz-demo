// Copyright 2023 CloudWeGo Authors
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
	"github.com/cloudwego/biz-demo/book-shop/app/item/common/entity"
	"github.com/cloudwego/biz-demo/book-shop/kitex_gen/cwg/bookshop/item"
)

func ConvertEntity2DTO(e *entity.ProductEntity) *item.Product {
	ret := &item.Product{
		ProductId:   e.ProductId,
		Name:        e.Name,
		Pic:         e.Pic,
		Description: e.Description,
		Price:       e.Price,
		Stock:       e.Stock,
		Status:      item.Status(e.Status),
	}
	if e.Property != nil {
		ret.Property = &item.BookProperty{
			Isbn:     e.Property.ISBN,
			SpuName:  e.Property.SpuName,
			SpuPrice: e.Property.SpuPrice,
		}
	}
	return ret
}
