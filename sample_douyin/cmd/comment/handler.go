// Copyright 2023 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package main

import (
	"context"

	"github.com/cloudwego/biz-demo/sample_douyin/cmd/comment/pack"
	"github.com/cloudwego/biz-demo/sample_douyin/cmd/comment/service"
	douyincomment "github.com/cloudwego/biz-demo/sample_douyin/kitex_gen/douyincomment"
	"github.com/cloudwego/biz-demo/sample_douyin/pkg/errno"
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
