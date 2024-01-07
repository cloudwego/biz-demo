package main

import (
	"context"

	"github.com/cloudwego/biz-demo/gomall/demo/demo_proto/biz/service"
	pbapi "github.com/cloudwego/biz-demo/gomall/demo/demo_proto/kitex_gen/pbapi"
)

// EchoImpl implements the last service interface defined in the IDL.
type EchoImpl struct{}

// Echo implements the EchoImpl interface.
func (s *EchoImpl) Echo(ctx context.Context, req *pbapi.Request) (resp *pbapi.Response, err error) {
	resp, err = service.NewEchoService(ctx).Run(req)

	return resp, err
}
