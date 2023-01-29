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

// ProductUpdateService product update service
type ProductUpdateService struct{}

var productUpdateService ProductUpdateService

func GetProductUpdateServiceInstance() *ProductUpdateService {
	return &productUpdateService
}

func (s *ProductUpdateService) AddProduct(ctx context.Context, entity *entity.ProductEntity) error {
	err := repository.GetRegistry().GetProductRepository().AddProduct(ctx, entity)
	if err != nil {
		return err
	}
	return nil
}

func (s *ProductUpdateService) EditProduct(ctx context.Context, origin, target *entity.ProductEntity) error {
	err := repository.GetRegistry().GetProductRepository().UpdateProduct(ctx, origin, target)
	if err != nil {
		return err
	}
	return nil
}
