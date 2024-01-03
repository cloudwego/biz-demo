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
	"errors"

	"github.com/cloudwego/biz-demo/book-shop/app/item/common/constant"
	"github.com/cloudwego/biz-demo/book-shop/app/item/common/entity"
	"github.com/cloudwego/biz-demo/book-shop/app/item/domain/repository"
	"github.com/cloudwego/kitex/pkg/klog"
)

// ProductStateService product state machine service
type ProductStateService struct{}

var productStateService ProductStateService

func GetProductStateService() *ProductStateService {
	return &productStateService
}

type ProductStateInfo struct {
	Status constant.ProductStatus
}

type CanTransferFunc func(originalInfo *ProductStateInfo) error

type ConstructTargetInfoFunc func(originalInfo *ProductStateInfo) *ProductStateInfo

var canTransferFuncMap = map[constant.StateOperationType]CanTransferFunc{
	constant.StateOperationTypeAdd: func(originalInfo *ProductStateInfo) error {
		return nil
	},
	constant.StateOperationTypeSave: func(originalInfo *ProductStateInfo) error {
		if originalInfo.Status == constant.ProductStatusDelete {
			return errors.New("商品已删除")
		}
		return nil
	},
	constant.StateOperationTypeDel: func(originalInfo *ProductStateInfo) error {
		return nil
	},
	constant.StateOperationTypeOnline: func(originalInfo *ProductStateInfo) error {
		if originalInfo.Status != constant.ProductStatusOffline {
			return errors.New("商品非下架状态")
		}
		return nil
	},
	constant.StateOperationTypeOffline: func(originalInfo *ProductStateInfo) error {
		if originalInfo.Status != constant.ProductStatusOnline {
			return errors.New("商品非上架状态")
		}
		return nil
	},
}

var constructTargetInfoFuncMap = map[constant.StateOperationType]ConstructTargetInfoFunc{
	constant.StateOperationTypeAdd: func(originalInfo *ProductStateInfo) (ret *ProductStateInfo) {
		ret = &ProductStateInfo{}
		ret.Status = constant.ProductStatusOnline
		return
	},
	constant.StateOperationTypeSave: func(originalInfo *ProductStateInfo) (ret *ProductStateInfo) {
		ret = &ProductStateInfo{}
		ret.Status = constant.ProductStatusOnline
		return
	},
	constant.StateOperationTypeDel: func(originalInfo *ProductStateInfo) (ret *ProductStateInfo) {
		ret = &ProductStateInfo{}
		ret.Status = constant.ProductStatusDelete
		return
	},
	constant.StateOperationTypeOnline: func(originalInfo *ProductStateInfo) (ret *ProductStateInfo) {
		ret = &ProductStateInfo{}
		ret.Status = constant.ProductStatusOnline
		return
	},
	constant.StateOperationTypeOffline: func(originalInfo *ProductStateInfo) (ret *ProductStateInfo) {
		ret = &ProductStateInfo{}
		ret.Status = constant.ProductStatusOffline
		return
	},
}

// GetCanTransferFunc get the validating func
func (s *ProductStateService) GetCanTransferFunc(operationType constant.StateOperationType) (CanTransferFunc, error) {
	if canTransferFunc, ok := canTransferFuncMap[operationType]; ok {
		return canTransferFunc, nil
	}

	return nil, errors.New("GetCanTransferFunc not found")
}

// GetConstructTargetInfoFunc get func to change product status
func (s *ProductStateService) getConstructTargetInfoFunc(operationType constant.StateOperationType) (ConstructTargetInfoFunc, error) {
	if constructTargetInfoFunc, ok := constructTargetInfoFuncMap[operationType]; ok {
		return constructTargetInfoFunc, nil
	}

	return nil, errors.New("GetConstructTargetInfoFunc not found")
}

// ConstructTargetInfo change product status
func (s *ProductStateService) ConstructTargetInfo(originProduct *entity.ProductEntity,
	operation constant.StateOperationType,
) (*entity.ProductEntity, error) {
	targetProduct, err := originProduct.Clone()
	if err != nil {
		return nil, err
	}
	originStateInfo := &ProductStateInfo{
		Status: originProduct.Status,
	}
	constructFunc, err := s.getConstructTargetInfoFunc(operation)
	if err != nil {
		return nil, err
	}
	targetState := constructFunc(originStateInfo)
	targetProduct.Status = targetState.Status
	return targetProduct, nil
}

// OperateProduct update product
func (s *ProductStateService) OperateProduct(ctx context.Context, origin, target *entity.ProductEntity) error {
	repo := repository.GetRegistry().GetProductRepository()
	// update status
	err := repo.UpdateProduct(ctx, origin, target)
	if err != nil {
		klog.CtxErrorf(ctx, "OperateProduct err: %v", err)
		return err
	}
	return nil
}
