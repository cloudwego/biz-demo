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

	"github.com/cloudwego/biz-demo/book-shop/app/item/domain/repository"
)

type ProductStockService struct{}

var productStockService ProductStockService

func GetProductStockServiceInstance() *ProductStockService {
	return &productStockService
}

func (s *ProductStockService) IncreaseStockNum(ctx context.Context, productId, incrNum int64) error {
	return repository.GetRegistry().GetStockRepository().IncrStock(ctx, productId, incrNum)
}

func (s *ProductStockService) DecreaseStockNum(ctx context.Context, productId, decrNum int64) error {
	return repository.GetRegistry().GetStockRepository().DecrStock(ctx, productId, decrNum)
}
