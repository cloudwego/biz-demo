package main

import (
	"context"

	"github.com/cloudwego/biz-demo/gomall/app/email/biz/service"
	email "github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/email"
)

// EmailServiceImpl implements the last service interface defined in the IDL.
type EmailServiceImpl struct{}

// Send implements the EmailServiceImpl interface.
func (s *EmailServiceImpl) Send(ctx context.Context, req *email.EmailReq) (resp *email.EmailResp, err error) {
	resp, err = service.NewSendService(ctx).Run(req)

	return resp, err
}
