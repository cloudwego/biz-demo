package test

import (
	"context"
	"testing"

	api "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/order"
	"github.com/cloudwego/biz-demo/gomall/rpc_gen/rpc/order"
	"github.com/cloudwego/kitex/client"
)

func TestMarkPaid(t *testing.T) {
	var opts []client.Option
	opts = append(opts, client.WithHostPorts("localhost:8885"))
	order.InitClient("order", opts...)
	resp, err := order.MarkOrderPaid(context.TODO(), &api.MarkOrderPaidReq{
		UserId:  1,
		OrderId: "42423444234",
	})
	if err != nil {
		t.Errorf("MarkOrderPaid err:%s", err)
		return
	}
	t.Log(resp)
}
