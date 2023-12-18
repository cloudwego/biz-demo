package service

import (
	"context"
	payment "github.com/baiyutang/gomall/app/payment/kitex_gen/payment"
)

type ChargeService struct {
	ctx context.Context
} // NewChargeService new ChargeService
func NewChargeService(ctx context.Context) *ChargeService {
	return &ChargeService{ctx: ctx}
}

// Run create note info
func (s *ChargeService) Run(req *payment.ChargeRequest) (resp *payment.ChargeResponse, err error) {
	// Finish your business logic.

	return
}
