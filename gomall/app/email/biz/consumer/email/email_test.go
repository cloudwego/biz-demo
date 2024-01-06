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
