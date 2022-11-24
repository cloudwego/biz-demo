package usecase

import (
	"context"

	"github.com/cloudwego/biz-demo/open-payment-platform/internal/payment/entity"
)

type Repository interface {
	GetByOutOrderNo(ctx context.Context, outOrderNo string) (*entity.Order, error)
	UpdateOrderStatus(ctx context.Context, outOrderNo string, orderStatus int8) error
	Create(ctx context.Context, order *entity.Order) error
}
