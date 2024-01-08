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

package email

import (
	"testing"

	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/email"

	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

func TestEmailConsumer(t *testing.T) {
	// Connect to a server
	nc, _ := nats.Connect(nats.DefaultURL)
	defer nc.Close()

	data, err := proto.Marshal(&email.EmailReq{
		From:        "hello@example.com",
		To:          "to@example.com",
		ContentType: "text/plain",
		Subject:     "hello world",
		Content:     "hello world",
	})
	if err != nil {
		t.Error(err)
	}

	err = nc.PublishMsg(&nats.Msg{Subject: "email", Data: data})
	if err != nil {
		panic(err)
	}
}
