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

package converter

import (
	"github.com/cloudwego/biz-demo/book-shop/app/item/common/entity"
	"github.com/cloudwego/biz-demo/book-shop/kitex_gen/cwg/bookshop/item"
	"github.com/cloudwego/biz-demo/book-shop/pkg/utils"
)

func ConvertDTO2Entity(dto *item.Product) *entity.ProductEntity {
	ret := &entity.ProductEntity{
		ProductId:   dto.ProductId,
		Name:        dto.Name,
		Pic:         dto.Pic,
		Description: dto.Description,
		Price:       dto.Price,
		Stock:       dto.Stock,
		Status:      dto.Stock,
	}
	if dto.Property != nil {
		ret.Property = &entity.PropertyEntity{
			ISBN:     dto.Property.Isbn,
			SpuName:  dto.Property.SpuName,
			SpuPrice: dto.Property.SpuPrice,
		}
	}

	return ret
}

func ConvertAddReq2Entity(req *item.AddReq) (*entity.ProductEntity, error) {
	pid, err := utils.GenerateID()
	if err != nil {
		return nil, err
	}

	ret := &entity.ProductEntity{
		ProductId:   pid,
		Name:        req.Name,
		Pic:         req.Pic,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
		Status:      int64(item.Status_Online),
	}

	if req.Property != nil {
		ret.Property = &entity.PropertyEntity{
			ISBN:     req.Property.Isbn,
			SpuName:  req.Property.SpuName,
			SpuPrice: req.Property.SpuPrice,
		}
	}

	return ret, err
}

func ConvertEditReq2Entity(originEntity *entity.ProductEntity, req *item.EditReq) (*entity.ProductEntity, error) {
	targetEntity, err := originEntity.Clone()
	if err != nil {
		return nil, err
	}
	if req.Name != nil {
		targetEntity.Name = *req.Name
	}
	if req.Pic != nil {
		targetEntity.Pic = *req.Pic
	}
	if req.Description != nil {
		targetEntity.Description = *req.Description
	}
	if req.Price != nil {
		targetEntity.Price = *req.Price
	}
	if req.Stock != nil {
		targetEntity.Stock = *req.Stock
	}
	if req.Property != nil {
		targetEntity.Property = &entity.PropertyEntity{
			ISBN:     req.Property.Isbn,
			SpuName:  req.Property.SpuName,
			SpuPrice: req.Property.SpuPrice,
		}
	}
	return targetEntity, nil
}
