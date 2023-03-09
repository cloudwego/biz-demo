// Code generated by Kitex v0.4.4. DO NOT EDIT.

package favoriteservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	douyinfavorite "mydouyin/kitex_gen/douyinfavorite"
)

func serviceInfo() *kitex.ServiceInfo {
	return favoriteServiceServiceInfo
}

var favoriteServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "FavoriteService"
	handlerType := (*douyinfavorite.FavoriteService)(nil)
	methods := map[string]kitex.MethodInfo{
		"FavoriteAction": kitex.NewMethodInfo(favoriteActionHandler, newFavoriteServiceFavoriteActionArgs, newFavoriteServiceFavoriteActionResult, false),
		"GetList":        kitex.NewMethodInfo(getListHandler, newFavoriteServiceGetListArgs, newFavoriteServiceGetListResult, false),
		"GetIsFavorite":  kitex.NewMethodInfo(getIsFavoriteHandler, newFavoriteServiceGetIsFavoriteArgs, newFavoriteServiceGetIsFavoriteResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "douyinfavorite",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.4.4",
		Extra:           extra,
	}
	return svcInfo
}

func favoriteActionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*douyinfavorite.FavoriteServiceFavoriteActionArgs)
	realResult := result.(*douyinfavorite.FavoriteServiceFavoriteActionResult)
	success, err := handler.(douyinfavorite.FavoriteService).FavoriteAction(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newFavoriteServiceFavoriteActionArgs() interface{} {
	return douyinfavorite.NewFavoriteServiceFavoriteActionArgs()
}

func newFavoriteServiceFavoriteActionResult() interface{} {
	return douyinfavorite.NewFavoriteServiceFavoriteActionResult()
}

func getListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*douyinfavorite.FavoriteServiceGetListArgs)
	realResult := result.(*douyinfavorite.FavoriteServiceGetListResult)
	success, err := handler.(douyinfavorite.FavoriteService).GetList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newFavoriteServiceGetListArgs() interface{} {
	return douyinfavorite.NewFavoriteServiceGetListArgs()
}

func newFavoriteServiceGetListResult() interface{} {
	return douyinfavorite.NewFavoriteServiceGetListResult()
}

func getIsFavoriteHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*douyinfavorite.FavoriteServiceGetIsFavoriteArgs)
	realResult := result.(*douyinfavorite.FavoriteServiceGetIsFavoriteResult)
	success, err := handler.(douyinfavorite.FavoriteService).GetIsFavorite(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newFavoriteServiceGetIsFavoriteArgs() interface{} {
	return douyinfavorite.NewFavoriteServiceGetIsFavoriteArgs()
}

func newFavoriteServiceGetIsFavoriteResult() interface{} {
	return douyinfavorite.NewFavoriteServiceGetIsFavoriteResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) FavoriteAction(ctx context.Context, req *douyinfavorite.FavoriteActionRequest) (r *douyinfavorite.FavoriteActionResponse, err error) {
	var _args douyinfavorite.FavoriteServiceFavoriteActionArgs
	_args.Req = req
	var _result douyinfavorite.FavoriteServiceFavoriteActionResult
	if err = p.c.Call(ctx, "FavoriteAction", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetList(ctx context.Context, req *douyinfavorite.GetListRequest) (r *douyinfavorite.GetListResponse, err error) {
	var _args douyinfavorite.FavoriteServiceGetListArgs
	_args.Req = req
	var _result douyinfavorite.FavoriteServiceGetListResult
	if err = p.c.Call(ctx, "GetList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetIsFavorite(ctx context.Context, req *douyinfavorite.GetIsFavoriteRequest) (r *douyinfavorite.GetIsFavoriteResponse, err error) {
	var _args douyinfavorite.FavoriteServiceGetIsFavoriteArgs
	_args.Req = req
	var _result douyinfavorite.FavoriteServiceGetIsFavoriteResult
	if err = p.c.Call(ctx, "GetIsFavorite", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}