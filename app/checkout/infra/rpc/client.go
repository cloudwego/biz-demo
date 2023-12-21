package rpc

import (
	"github.com/baiyutang/gomall/app/checkout/kitex_gen/order/orderservice"
	"github.com/baiyutang/gomall/app/common/suite"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"os"
	"sync"

	"github.com/baiyutang/gomall/app/checkout/kitex_gen/cart/cartservice"
	"github.com/baiyutang/gomall/app/checkout/kitex_gen/payment/paymentservice"
	"github.com/baiyutang/gomall/app/checkout/kitex_gen/product/productcatalogservice"
	checkoututils "github.com/baiyutang/gomall/app/checkout/utils"
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
)

var (
	CartClient    cartservice.Client
	ProductClient productcatalogservice.Client
	PaymentClient paymentservice.Client
	OrderClient   orderservice.Client
	once          sync.Once
	err           error
)

var (
	commonOpts = []client.Option{}
)

func InitClient() {
	once.Do(func() {
		initCartClient()
		initProductClient()
		initPaymentClient()
		initOrderClient()
	})
}

func initProductClient() {
	ProductClient, err = productcatalogservice.NewClient(
		"product",
		client.WithSuite(suite.CommonGrpcClientSuite{
			DestServiceAddr:    "localhost:8881",
			CurrentServiceName: "checkout",
		}),
	)
	checkoututils.MustHandleError(err)
}

func initCartClient() {
	var opts []client.Option
	if os.Getenv("REGISTRY_ENABLE") == "true" {
		r, err := consul.NewConsulResolver(os.Getenv("REGISTRY_ADDR"))
		checkoututils.MustHandleError(err)
		opts = append(opts, client.WithResolver(r))
	} else {
		opts = append(opts, client.WithHostPorts("localhost:8881"))
	}
	opts = append(opts, client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "checkout-cart-client"}))
	opts = append(opts, commonOpts...)
	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	checkoututils.MustHandleError(err)
}

func initPaymentClient() {
	var opts []client.Option
	if os.Getenv("REGISTRY_ENABLE") == "true" {
		r, err := consul.NewConsulResolver(os.Getenv("REGISTRY_ADDR"))
		checkoututils.MustHandleError(err)
		opts = append(opts, client.WithResolver(r))
	} else {
		opts = append(opts, client.WithHostPorts("localhost:8881"))
	}
	opts = append(opts, client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "checkout-payment-client"}))
	opts = append(opts, commonOpts...)
	PaymentClient, err = paymentservice.NewClient("payment", opts...)
	checkoututils.MustHandleError(err)
}

func initOrderClient() {
	var opts []client.Option
	if os.Getenv("REGISTRY_ENABLE") == "true" {
		r, err := consul.NewConsulResolver(os.Getenv("REGISTRY_ADDR"))
		checkoututils.MustHandleError(err)
		opts = append(opts, client.WithResolver(r))
	} else {
		opts = append(opts, client.WithHostPorts("localhost:8881"))
	}
	opts = append(opts, client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "checkout-order-client"}))
	opts = append(opts, commonOpts...)
	OrderClient, err = orderservice.NewClient("order", opts...)
	checkoututils.MustHandleError(err)
}
