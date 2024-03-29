// Code generated by Kitex v0.4.3. DO NOT EDIT.

package paymentsvc

import (
	"context"

	payment "github.com/cloudwego/biz-demo/open-payment-platform/kitex_gen/payment"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	UnifyPay(ctx context.Context, req *payment.UnifyPayReq, callOptions ...callopt.Option) (r *payment.UnifyPayResp, err error)
	QRPay(ctx context.Context, req *payment.QRPayReq, callOptions ...callopt.Option) (r *payment.QRPayResp, err error)
	QueryOrder(ctx context.Context, req *payment.QueryOrderReq, callOptions ...callopt.Option) (r *payment.QueryOrderResp, err error)
	CloseOrder(ctx context.Context, req *payment.CloseOrderReq, callOptions ...callopt.Option) (r *payment.CloseOrderResp, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kPaymentSvcClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kPaymentSvcClient struct {
	*kClient
}

func (p *kPaymentSvcClient) UnifyPay(ctx context.Context, req *payment.UnifyPayReq, callOptions ...callopt.Option) (r *payment.UnifyPayResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UnifyPay(ctx, req)
}

func (p *kPaymentSvcClient) QRPay(ctx context.Context, req *payment.QRPayReq, callOptions ...callopt.Option) (r *payment.QRPayResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.QRPay(ctx, req)
}

func (p *kPaymentSvcClient) QueryOrder(ctx context.Context, req *payment.QueryOrderReq, callOptions ...callopt.Option) (r *payment.QueryOrderResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.QueryOrder(ctx, req)
}

func (p *kPaymentSvcClient) CloseOrder(ctx context.Context, req *payment.CloseOrderReq, callOptions ...callopt.Option) (r *payment.CloseOrderResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CloseOrder(ctx, req)
}
