// Code generated by Kitex v0.4.2. DO NOT EDIT.

package productpageservice

import (
	"context"
	product "github.com/cloudwego/biz-demo/bookinfo/kitex_gen/cwg/bookinfo/product"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return productPageServiceServiceInfo
}

var productPageServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "ProductPageService"
	handlerType := (*product.ProductPageService)(nil)
	methods := map[string]kitex.MethodInfo{
		"GetProduct":   kitex.NewMethodInfo(getProductHandler, newProductPageServiceGetProductArgs, newProductPageServiceGetProductResult, false),
		"ListProducts": kitex.NewMethodInfo(listProductsHandler, newProductPageServiceListProductsArgs, newProductPageServiceListProductsResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "product",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.4.2",
		Extra:           extra,
	}
	return svcInfo
}

func getProductHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*product.ProductPageServiceGetProductArgs)
	realResult := result.(*product.ProductPageServiceGetProductResult)
	success, err := handler.(product.ProductPageService).GetProduct(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newProductPageServiceGetProductArgs() interface{} {
	return product.NewProductPageServiceGetProductArgs()
}

func newProductPageServiceGetProductResult() interface{} {
	return product.NewProductPageServiceGetProductResult()
}

func listProductsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*product.ProductPageServiceListProductsArgs)
	realResult := result.(*product.ProductPageServiceListProductsResult)
	success, err := handler.(product.ProductPageService).ListProducts(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newProductPageServiceListProductsArgs() interface{} {
	return product.NewProductPageServiceListProductsArgs()
}

func newProductPageServiceListProductsResult() interface{} {
	return product.NewProductPageServiceListProductsResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) GetProduct(ctx context.Context, req *product.GetProductReq) (r *product.GetProductResp, err error) {
	var _args product.ProductPageServiceGetProductArgs
	_args.Req = req
	var _result product.ProductPageServiceGetProductResult
	if err = p.c.Call(ctx, "GetProduct", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ListProducts(ctx context.Context, req *product.ListProductsReq) (r *product.ListProductsResp, err error) {
	var _args product.ProductPageServiceListProductsArgs
	_args.Req = req
	var _result product.ProductPageServiceListProductsResult
	if err = p.c.Call(ctx, "ListProducts", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
