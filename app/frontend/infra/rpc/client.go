package rpc

import (
	"github.com/baiyutang/gomall/app/frontend/kitex_gen/product/productcatalogservice"
	"github.com/baiyutang/gomall/app/frontend/kitex_gen/user/userservice"
	frontendutils "github.com/baiyutang/gomall/app/frontend/utils"
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
	"os"
	"sync"
)

var (
	ProductClient productcatalogservice.Client
	UserClient    userservice.Client
	once          sync.Once
	err           error
)

func InitClient() {
	once.Do(func() {
		initProductClient()
		initUserClient()
	})
}

func initProductClient() {
	var opts []client.Option
	if os.Getenv("REGISTRY_ENABLE") == "true" {
		r, err := consul.NewConsulResolver(os.Getenv("REGISTRY_ADDR"))
		frontendutils.MustHandleError(err)
		opts = append(opts, client.WithResolver(r))
	} else {
		opts = append(opts, client.WithHostPorts("localhost:8881"))
	}

	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	frontendutils.MustHandleError(err)
}

func initUserClient() {
	var opts []client.Option
	if os.Getenv("REGISTRY_ENABLE") == "true" {
		r, err := consul.NewConsulResolver(os.Getenv("REGISTRY_ADDR"))
		frontendutils.MustHandleError(err)
		opts = append(opts, client.WithResolver(r))
	} else {
		opts = append(opts, client.WithHostPorts("localhost:8882"))
	}

	UserClient, err = userservice.NewClient("user", opts...)
	frontendutils.MustHandleError(err)
}