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

package repository

import (
	"context"

	"github.com/cloudwego/biz-demo/book-shop/app/item/common/entity"
	"github.com/cloudwego/biz-demo/book-shop/app/item/infras/es"
)

type Product2CRepositoryImpl struct{}

func (i Product2CRepositoryImpl) MGetProducts2C(ctx context.Context, productIds []int64) ([]*entity.ProductEntity, error) {
	entities, err := es.BatchGetProductById(ctx, productIds)
	return entities, err
}

func (i Product2CRepositoryImpl) SearchProducts(ctx context.Context, name, description, spuName *string) ([]*entity.ProductEntity, error) {
	filterMap := make(map[string]interface{})
	if name != nil {
		filterMap["name"] = *name
	}
	if description != nil {
		filterMap["description"] = *description
	}
	if spuName != nil {
		filterMap["spu_name"] = *spuName
	}
	entities, err := es.SearchProduct(ctx, filterMap)
	return entities, err
}
