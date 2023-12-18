package service

import (
	"context"
	"testing"
	payment "github.com/baiyutang/gomall/app/payment/kitex_gen/payment"
)

func TestCharge_Run(t *testing.T) {
	ctx := context.Background()
	s := NewChargeService(ctx)
	// init req and assert value

	req := &payment.ChargeRequest{}
	resp, err := s.Run(req)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if resp == nil {
		t.Errorf("unexpected nil response")
	}
	// todo: edit your unit test

}
