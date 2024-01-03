package service

import (
	"context"
	"testing"

	order "github.com/cloudwego/biz-demo/gomall/app/order/kitex_gen/order"
)

func TestListOrder_Run(t *testing.T) {
	ctx := context.Background()
	s := NewListOrderService(ctx)
	// init req and assert value

	req := &order.ListOrderRequest{}
	resp, err := s.Run(req)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if resp == nil {
		t.Errorf("unexpected nil response")
	}
	// todo: edit your unit test

}
