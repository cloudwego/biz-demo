// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package ratings

import (
	"context"
	"github.com/cloudwego/biz-demo/bookinfo/internal/service/ratings"
	"github.com/cloudwego/biz-demo/bookinfo/pkg/configparser"
)

// Injectors from wire.go:

func NewServer(ctx context.Context) (*Server, error) {
	provider := configparser.Default()
	options, err := Configure(provider)
	if err != nil {
		return nil, err
	}
	serverOptions := options.Server
	ratingService := ratings.New()
	server := &Server{
		opts: serverOptions,
		svc:  ratingService,
	}
	return server, nil
}
