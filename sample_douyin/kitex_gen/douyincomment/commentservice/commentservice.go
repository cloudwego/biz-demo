// Code generated by Kitex v0.4.4. DO NOT EDIT.

package commentservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	douyincomment "mydouyin/kitex_gen/douyincomment"
)

func serviceInfo() *kitex.ServiceInfo {
	return commentServiceServiceInfo
}

var commentServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "CommentService"
	handlerType := (*douyincomment.CommentService)(nil)
	methods := map[string]kitex.MethodInfo{
		"CreateComment":    kitex.NewMethodInfo(createCommentHandler, newCommentServiceCreateCommentArgs, newCommentServiceCreateCommentResult, false),
		"DeleteComment":    kitex.NewMethodInfo(deleteCommentHandler, newCommentServiceDeleteCommentArgs, newCommentServiceDeleteCommentResult, false),
		"GetVideoComments": kitex.NewMethodInfo(getVideoCommentsHandler, newCommentServiceGetVideoCommentsArgs, newCommentServiceGetVideoCommentsResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "douyincomment",
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

func createCommentHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*douyincomment.CommentServiceCreateCommentArgs)
	realResult := result.(*douyincomment.CommentServiceCreateCommentResult)
	success, err := handler.(douyincomment.CommentService).CreateComment(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newCommentServiceCreateCommentArgs() interface{} {
	return douyincomment.NewCommentServiceCreateCommentArgs()
}

func newCommentServiceCreateCommentResult() interface{} {
	return douyincomment.NewCommentServiceCreateCommentResult()
}

func deleteCommentHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*douyincomment.CommentServiceDeleteCommentArgs)
	realResult := result.(*douyincomment.CommentServiceDeleteCommentResult)
	success, err := handler.(douyincomment.CommentService).DeleteComment(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newCommentServiceDeleteCommentArgs() interface{} {
	return douyincomment.NewCommentServiceDeleteCommentArgs()
}

func newCommentServiceDeleteCommentResult() interface{} {
	return douyincomment.NewCommentServiceDeleteCommentResult()
}

func getVideoCommentsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*douyincomment.CommentServiceGetVideoCommentsArgs)
	realResult := result.(*douyincomment.CommentServiceGetVideoCommentsResult)
	success, err := handler.(douyincomment.CommentService).GetVideoComments(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newCommentServiceGetVideoCommentsArgs() interface{} {
	return douyincomment.NewCommentServiceGetVideoCommentsArgs()
}

func newCommentServiceGetVideoCommentsResult() interface{} {
	return douyincomment.NewCommentServiceGetVideoCommentsResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) CreateComment(ctx context.Context, req *douyincomment.CreateCommentRequest) (r *douyincomment.CreateCommentResponse, err error) {
	var _args douyincomment.CommentServiceCreateCommentArgs
	_args.Req = req
	var _result douyincomment.CommentServiceCreateCommentResult
	if err = p.c.Call(ctx, "CreateComment", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DeleteComment(ctx context.Context, req *douyincomment.DeleteCommentRequest) (r *douyincomment.DeleteCommentResponse, err error) {
	var _args douyincomment.CommentServiceDeleteCommentArgs
	_args.Req = req
	var _result douyincomment.CommentServiceDeleteCommentResult
	if err = p.c.Call(ctx, "DeleteComment", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetVideoComments(ctx context.Context, req *douyincomment.GetVideoCommentsRequest) (r *douyincomment.GetVideoCommentsResponse, err error) {
	var _args douyincomment.CommentServiceGetVideoCommentsArgs
	_args.Req = req
	var _result douyincomment.CommentServiceGetVideoCommentsResult
	if err = p.c.Call(ctx, "GetVideoComments", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}