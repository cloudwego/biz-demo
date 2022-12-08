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

package usecase

import (
	"context"

	"github.com/cloudwego/biz-demo/open-payment-platform/internal/payment/entity"
)

// Repository is the interface of usecase dependent on.
// the interface is the part of usecase logic,so we put it here.
type Repository interface {
	GetByOutOrderNo(ctx context.Context, outOrderNo string) (*entity.Order, error)
	UpdateOrderStatus(ctx context.Context, outOrderNo string, orderStatus int8) error
	Create(ctx context.Context, order *entity.Order) error
}
