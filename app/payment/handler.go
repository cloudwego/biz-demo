package main

import (
	"context"
	payment "github.com/baiyutang/gomall/app/payment/kitex_gen/payment"
	"github.com/baiyutang/gomall/app/payment/biz/service"
)

// PaymentServiceImpl implements the last service interface defined in the IDL.
type PaymentServiceImpl struct{}

// Charge implements the PaymentServiceImpl interface.
func (s *PaymentServiceImpl) Charge(ctx context.Context, req *payment.ChargeRequest) (resp *payment.ChargeResponse, err error) {
	resp, err = service.NewChargeService(ctx).Run(req)

	return resp, err
}
