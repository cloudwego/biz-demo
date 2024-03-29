// Code generated by Kitex v0.4.3. DO NOT EDIT.
package paymentsvc

import (
	payment "github.com/cloudwego/biz-demo/open-payment-platform/kitex_gen/payment"
	server "github.com/cloudwego/kitex/server"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler payment.PaymentSvc, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}
