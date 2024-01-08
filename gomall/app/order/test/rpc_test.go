// Copyright 2024 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
