package email

import (
	"github.com/cloudwego/biz-demo/gomall/app/email/infra/mq"
	"github.com/cloudwego/biz-demo/gomall/app/email/infra/notify"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"

	"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/email"
)

func ConsumerInit() {
	// Connect to a server

	sub, err := mq.Nc.Subscribe("email", func(m *nats.Msg) {
		var req email.EmailReq
		err := proto.Unmarshal(m.Data, &req)
		if err != nil {
			klog.Error(err)
		}
		noopEmail := notify.NewNoopEmail()
		_ = noopEmail.Send(&req)
	})
	if err != nil {
		panic(err)
	}

	server.RegisterShutdownHook(func() {
		sub.Unsubscribe() //nolint:errcheck
		mq.Nc.Close()
	})
}
