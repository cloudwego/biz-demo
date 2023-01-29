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
	"context"
	"errors"

	"github.com/cloudwego/kitex/pkg/klog"

	"github.com/cloudwego/biz-demo/book-shop/app/item/common/entity"
	"github.com/cloudwego/biz-demo/book-shop/app/item/common/po"
	"github.com/cloudwego/biz-demo/book-shop/app/item/infras/es"
	"github.com/cloudwego/biz-demo/book-shop/app/item/infras/repository/converter"
	"github.com/cloudwego/biz-demo/book-shop/app/item/infras/repository/differ"
)

type ProductRepositoryImpl struct{}

func (i ProductRepositoryImpl) AddProduct(ctx context.Context, product *entity.ProductEntity) error {
	if product == nil {
		return errors.New("插入数据不可为空")
	}
	po, err := converter.ProductDO2POConverter.Convert2po(ctx, product)
	if err != nil {
		return err
	}
	// update es async
	go func() {
		err := es.UpsertProductES(ctx, po.ProductId, product)
		if err != nil {
			klog.CtxErrorf(ctx, "UpsertProductES err: %v", err)
		}
	}()
	return DB.WithContext(ctx).Create(po).Error
}

func (i ProductRepositoryImpl) UpdateProduct(ctx context.Context, origin, target *entity.ProductEntity) error {
	productId := target.ProductId
	originPO, err := converter.ProductDO2POConverter.Convert2po(ctx, origin)
	if err != nil {
		return err
	}
	targetPO, err := converter.ProductDO2POConverter.Convert2po(ctx, target)
	if err != nil {
		return err
	}
	// update es async
	go func() {
		err := es.UpsertProductES(ctx, productId, target)
		if err != nil {
			klog.CtxErrorf(ctx, "UpsertProductES err: %v", err)
		}
	}()
	changeMap := differ.ProductPODiffer.GetChangedMap(originPO, targetPO)
	return DB.WithContext(ctx).Model(&po.Product{}).Where("product_id = ?", productId).
		Updates(changeMap).Error
}

func (i ProductRepositoryImpl) GetProductById(ctx context.Context, productId int64) (*entity.ProductEntity, error) {
	products := make([]*po.Product, 0)
	err := DB.WithContext(ctx).Where("product_id = ?", productId).Find(&products).Error
	if err != nil {
		return nil, err
	}
	if len(products) == 0 {
		return nil, errors.New("该商品不存在")
	}
	do, err := converter.ProductPO2DOConverter.Convert2do(ctx, products[0])
	if err != nil {
		return nil, err
	}
	return do, nil
}

func (i ProductRepositoryImpl) ListProducts(ctx context.Context, filterParam map[string]interface{}) ([]*entity.ProductEntity, error) {
	products := make([]*po.Product, 0)
	productEntities := make([]*entity.ProductEntity, 0)
	DB = DB.Debug().WithContext(ctx)
	for k, v := range filterParam {
		DB = DB.Where(k+" = ?", v)
	}
	if err := DB.Find(&products).Error; err != nil {
		return nil, err
	}
	for _, v := range products {
		entity, err := converter.ProductPO2DOConverter.Convert2do(ctx, v)
		if err != nil {
			return nil, err
		}
		productEntities = append(productEntities, entity)
	}
	return productEntities, nil
}
