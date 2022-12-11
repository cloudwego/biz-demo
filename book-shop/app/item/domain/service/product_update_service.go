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
)

// ProductUpdateService 商品更新服务
type ProductUpdateService struct {
}

var productUpdateService ProductUpdateService

// GetProductUpdateServiceInstance 单例
func GetProductUpdateServiceInstance() *ProductUpdateService {
	return &productUpdateService
}

func (s *ProductUpdateService) AddProduct(ctx context.Context, entity *entity.ProductEntity) error {
	return nil
}

func (s *ProductUpdateService) EditProduct(ctx context.Context, origin *entity.ProductEntity, target *entity.ProductEntity) error {
	return nil
}
