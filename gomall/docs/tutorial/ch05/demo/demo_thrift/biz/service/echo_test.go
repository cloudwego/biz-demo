package service

import (
	"context"
	api "github.com/cloudwego/biz-demo/gomall/demo/demo_thrift/kitex_gen/api"
	"testing"
)

func TestEcho_Run(t *testing.T) {
	ctx := context.Background()
	s := NewEchoService(ctx)
	// init req and assert value

	req := &api.Request{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
