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

package repository

import (
	"github.com/cloudwego/biz-demo/book-shop/app/item/domain/repository"
	"github.com/cloudwego/biz-demo/book-shop/pkg/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func register() {
	productRepository := ProductRepositoryImpl{}
	stockRepository := StockRepositoryImpl{}
	product2CRepository := Product2CRepositoryImpl{}
	repository.GetRegistry().SetProductRepository(productRepository)
	repository.GetRegistry().SetStockRepository(stockRepository)
	repository.GetRegistry().SetProduct2CRepository(product2CRepository)
}

func initDB() {
	var err error
	DB, err = gorm.Open(mysql.Open(conf.MySQLDefaultDSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}
}

func Init() {
	register()
	initDB()
}
