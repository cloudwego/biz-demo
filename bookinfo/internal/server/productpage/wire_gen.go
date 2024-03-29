// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package productpage

import (
	"context"
	"github.com/cloudwego/biz-demo/bookinfo/internal/handler/productpage"
	"github.com/cloudwego/biz-demo/bookinfo/pkg/configparser"
	"github.com/cloudwego/biz-demo/bookinfo/pkg/injectors"
)

// Injectors from wire.go:

func NewServer(ctx context.Context) (*Server, error) {
	provider := configparser.Default()
	options, err := Configure(provider)
	if err != nil {
		return nil, err
	}
	serverOptions := options.Server
	reviewClientOptions := options.Reviews
	client, err := injectors.ProvideReviewClient(reviewClientOptions)
	if err != nil {
		return nil, err
	}
	detailsClientOptions := options.Details
	detailsserviceClient, err := injectors.ProvideDetailsClient(detailsClientOptions)
	if err != nil {
		return nil, err
	}
	handler := productpage.New(client, detailsserviceClient)
	server := &Server{
		opts:    serverOptions,
		handler: handler,
	}
	return server, nil
}
