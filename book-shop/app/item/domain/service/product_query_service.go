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

package service

import (
	"context"

	"github.com/cloudwego/biz-demo/book-shop/app/item/common/entity"
	"github.com/cloudwego/biz-demo/book-shop/app/item/domain/repository"
)

// ProductQueryService product query service
type ProductQueryService struct{}

var productQueryService ProductQueryService

// GetProductQueryServiceInstance single instance
func GetProductQueryServiceInstance() *ProductQueryService {
	return &productQueryService
}

func (s *ProductQueryService) GetProduct(ctx context.Context, productId int64) (*entity.ProductEntity, error) {
	do, err := repository.GetRegistry().GetProductRepository().GetProductById(ctx, productId)
	if err != nil {
		return nil, err
	}
	return do, nil
}

func (s *ProductQueryService) ListProducts(ctx context.Context, name, spuName *string, status *int64) ([]*entity.ProductEntity, error) {
	filterParam := make(map[string]interface{})
	if name != nil {
		filterParam["name"] = *name
	}
	if spuName != nil {
		filterParam["spu_name"] = *spuName
	}
	if status != nil {
		filterParam["status"] = *status
	}
	entities, err := repository.GetRegistry().GetProductRepository().ListProducts(ctx, filterParam)
	if err != nil {
		return nil, err
	}
	return entities, nil
}

func (s *ProductQueryService) MGet2C(ctx context.Context, productIds []int64) ([]*entity.ProductEntity, error) {
	do, err := repository.GetRegistry().GetProduct2CRepository().MGetProducts2C(ctx, productIds)
	if err != nil {
		return nil, err
	}
	return do, nil
}

func (s *ProductQueryService) Search(ctx context.Context, name, description, spuName *string) ([]*entity.ProductEntity, error) {
	do, err := repository.GetRegistry().GetProduct2CRepository().SearchProducts(ctx, name, description, spuName)
	if err != nil {
		return nil, err
	}
	return do, nil
}
