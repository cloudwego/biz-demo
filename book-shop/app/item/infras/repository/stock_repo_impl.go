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
	"errors"

	"github.com/cloudwego/biz-demo/book-shop/app/item/common/po"
	"gorm.io/gorm/clause"
)

type StockRepositoryImpl struct{}

func (i StockRepositoryImpl) IncrStock(ctx context.Context, productId, stockNum int64) error {
	return i.updateStock(ctx, productId, stockNum, "incr")
}

func (i StockRepositoryImpl) DecrStock(ctx context.Context, productId, stockNum int64) error {
	return i.updateStock(ctx, productId, stockNum, "decr")
}

func (i StockRepositoryImpl) updateStock(ctx context.Context, productId, stockNum int64, updateType string) error {
	productPOArr := make([]*po.Product, 0)

	tx := DB.Begin().WithContext(ctx)
	if tx.Error != nil {
		return tx.Error
	}
	// select for update
	if err := tx.Clauses(clause.Locking{Strength: "Update"}).Where("product_id = ?", productId).Find(&productPOArr).Error; err != nil {
		tx.Rollback()
		return err
	}
	if len(productPOArr) == 0 {
		tx.Rollback()
		return errors.New("item not found")
	}

	productPO := productPOArr[0]
	curStockNum := productPO.Stock
	if updateType == "incr" {
		curStockNum += stockNum
	} else if updateType == "decr" {
		curStockNum -= stockNum
	}
	if curStockNum < 0 {
		tx.Rollback()
		return errors.New("库存不足")
	}
	if err := tx.Model(&po.Product{}).Where("product_id = ?", productId).
		Updates(map[string]interface{}{
			"stock": curStockNum,
		}).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
