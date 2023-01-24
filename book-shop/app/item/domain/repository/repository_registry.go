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

type RepositoryRegistry struct {
	productRepository   ProductRepository
	stockRepository     StockRepository
	product2CRepository Product2CRepository
}

var inst = &RepositoryRegistry{}

func GetRegistry() *RepositoryRegistry {
	return inst
}

func (r *RepositoryRegistry) GetProductRepository() ProductRepository {
	return r.productRepository
}

func (r *RepositoryRegistry) SetProductRepository(productRepositoryIns ProductRepository) {
	r.productRepository = productRepositoryIns
}

func (r *RepositoryRegistry) GetProduct2CRepository() Product2CRepository {
	return r.product2CRepository
}

func (r *RepositoryRegistry) SetProduct2CRepository(product2CRepositoryIns Product2CRepository) {
	r.product2CRepository = product2CRepositoryIns
}

func (r *RepositoryRegistry) GetStockRepository() StockRepository {
	return r.stockRepository
}

func (r *RepositoryRegistry) SetStockRepository(stockRepositoryIns StockRepository) {
	r.stockRepository = stockRepositoryIns
}
