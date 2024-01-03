package main

import (
	"context"
	"github.com/baiyutang/gomall/app/payment/biz/service"
	payment "github.com/baiyutang/gomall/app/payment/kitex_gen/payment"
)

// PaymentServiceImpl implements the last service interface defined in the IDL.
type PaymentServiceImpl struct{}

// Charge implements the PaymentServiceImpl interface.
func (s *PaymentServiceImpl) Charge(ctx context.Context, req *payment.ChargeRequest) (resp *payment.ChargeResponse, err error) {
	resp, err = service.NewChargeService(ctx).Run(req)

	return resp, err
}
