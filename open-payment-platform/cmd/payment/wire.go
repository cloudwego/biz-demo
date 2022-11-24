//go:build wireinject
// +build wireinject

package main

import (
	"github.com/cloudwego/biz-demo/open-payment-platform/internal/payment/infrastructure/repository"
	"github.com/cloudwego/biz-demo/open-payment-platform/internal/payment/usecase"
	"github.com/cloudwego/biz-demo/open-payment-platform/kitex_gen/payment"
	"github.com/google/wire"
)

func initHandler() payment.PaymentSvc {
	panic(wire.Build(
		repository.SQLProviderSet,
		usecase.ProviderSet,
	))
}
