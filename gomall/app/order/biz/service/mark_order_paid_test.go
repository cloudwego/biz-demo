package service

import (
	"context"
	"testing"
	order "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/order"
)

func TestMarkOrderPaid_Run(t *testing.T) {
	ctx := context.Background()
	s := NewMarkOrderPaidService(ctx)
	// init req and assert value

	req := &order.MarkOrderPaidReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
