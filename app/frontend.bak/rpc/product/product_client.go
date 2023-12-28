package product

import (
	"context"

	product "github.com/baiyutang/gomall/app/frontend/kitex_gen/product"
	"github.com/baiyutang/gomall/app/frontend/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

type RPCClient interface {
	KitexClient() productcatalogservice.Client
	Service() string
	ListProducts(ctx context.Context, req *product.ListProductsReq, callOptions ...callopt.Option) (resp *product.ListProductsResponse, err error)

	GetProduct(ctx context.Context, req *product.GetProductRequest, callOptions ...callopt.Option) (resp *product.Product, err error)

	SearchProducts(ctx context.Context, req *product.SearchProductsRequest, callOptions ...callopt.Option) (resp *product.SearchProductsResponse, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := productcatalogservice.NewClient(dstService, opts...)
	if err != nil {
		return nil, err
	}
	cli := &clientImpl{
		service:     dstService,
		kitexClient: kitexClient,
	}

	return cli, nil
}

type clientImpl struct {
	service     string
	kitexClient productcatalogservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() productcatalogservice.Client {
	return c.kitexClient
}

func (c *clientImpl) ListProducts(ctx context.Context, req *product.ListProductsReq, callOptions ...callopt.Option) (resp *product.ListProductsResponse, err error) {
	return c.kitexClient.ListProducts(ctx, req, callOptions...)
}

func (c *clientImpl) GetProduct(ctx context.Context, req *product.GetProductRequest, callOptions ...callopt.Option) (resp *product.Product, err error) {
	return c.kitexClient.GetProduct(ctx, req, callOptions...)
}

func (c *clientImpl) SearchProducts(ctx context.Context, req *product.SearchProductsRequest, callOptions ...callopt.Option) (resp *product.SearchProductsResponse, err error) {
	return c.kitexClient.SearchProducts(ctx, req, callOptions...)
}
