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

package po

import (
	"github.com/cloudwego/biz-demo/book-shop/pkg/conf"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ProductId   int64  `json:"product_id"`
	Name        string `json:"name"`
	Pic         string `json:"pic"`
	Description string `json:"description"`
	ISBN        string `json:"isbn"`
	SpuName     string `json:"spu_name"`
	SpuPrice    int64  `json:"spu_price"`
	Price       int64  `json:"price"`
	Stock       int64  `json:"stock"`
	Status      int64  `json:"status"`
}

func (p *Product) TableName() string {
	return conf.ProductTableName
}
