package main

import (
	"context"
	"mydouyin/cmd/comment/pack"
	"mydouyin/cmd/comment/service"
	douyincomment "mydouyin/kitex_gen/douyincomment"
	"mydouyin/pkg/errno"
)

// CommentServiceImpl implements the last service interface defined in the IDL.
type CommentServiceImpl struct{}

// CreateComment implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CreateComment(ctx context.Context, req *douyincomment.CreateCommentRequest) (resp *douyincomment.CreateCommentResponse, err error) {
	resp = new(douyincomment.CreateCommentResponse)
	if err = req.IsValid(); err != nil {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}
	id, err := service.NewCreateCommentService(ctx).CreateComment(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.CommentId = id
	return resp, nil
}

// DeleteComment implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) DeleteComment(ctx context.Context, req *douyincomment.DeleteCommentRequest) (resp *douyincomment.DeleteCommentResponse, err error) {
	resp = new(douyincomment.DeleteCommentResponse)
	err = service.NewDeleteCommentService(ctx).DeleteComment(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// GetVideoComments implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) GetVideoComments(ctx context.Context, req *douyincomment.GetVideoCommentsRequest) (resp *douyincomment.GetVideoCommentsResponse, err error) {
	resp = new(douyincomment.GetVideoCommentsResponse)
	var comments []*douyincomment.Comment
	comments, err = service.NewGetVideoCommentsService(ctx).GetVideoComments(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.Comments = comments
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}
