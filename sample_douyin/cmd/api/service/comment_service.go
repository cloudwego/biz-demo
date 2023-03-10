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

package service

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/cloudwego/biz-demo/sample_douyin/cmd/api/pack"
	"github.com/cloudwego/biz-demo/sample_douyin/cmd/api/rpc"
	douyinapi "github.com/cloudwego/biz-demo/sample_douyin/hertz_gen/douyinapi"
	"github.com/cloudwego/biz-demo/sample_douyin/kitex_gen/douyincomment"
	"github.com/cloudwego/biz-demo/sample_douyin/kitex_gen/douyinuser"
	"github.com/cloudwego/biz-demo/sample_douyin/pkg/errno"
)

type CommentService struct {
	ctx context.Context
}

func NewCommentService(ctx context.Context) *CommentService {
	return &CommentService{
		ctx: ctx,
	}
}

func (s *CommentService) CommentAction(req douyinapi.CommentActionRequest, user *douyinapi.User) (*douyinapi.CommentActionResponse, error) {
	resp := new(douyinapi.CommentActionResponse)
	actionType, err := strconv.Atoi(req.ActionType)
	if err != nil {
		return resp, err
	}
	switch actionType {
	case 1:
		// create the date
		getMonth := time.Now().Format("01")
		getDay := time.Now().Format("02")
		var build strings.Builder
		build.WriteString(getMonth)
		build.WriteString("-")
		build.WriteString(getDay)
		date := build.String()
		// create the VideoID
		videoID, err := strconv.ParseInt(req.VideoID, 10, 64)
		if err != nil {
			return resp, err
		}

		rpc_resp, err := rpc.CreateComment(s.ctx, &douyincomment.CreateCommentRequest{
			Video:      videoID,
			User:       user.ID,
			Content:    req.CommentText,
			CreateDate: date,
		})
		if err != nil {
			return resp, err
		}

		if rpc_resp.BaseResp.StatusCode != 0 {
			return resp, errno.NewErrNo(rpc_resp.BaseResp.StatusCode, rpc_resp.BaseResp.StatusMessage)
		}
		resp.Comment = &douyinapi.Comment{
			ID:         rpc_resp.CommentId,
			User:       user,
			Content:    req.CommentText,
			CreateDate: date,
		}
	case 2:
		// delete the date
		commentID, err := strconv.ParseInt(req.CommentID, 10, 64)
		if err != nil {
			return resp, err
		}
		rpc_resp, err := rpc.DeleteComment(s.ctx, &douyincomment.DeleteCommentRequest{
			CommentId: commentID,
		})
		if err != nil {
			return resp, err
		}
		if rpc_resp.BaseResp.StatusCode != 0 {
			return resp, errno.NewErrNo(rpc_resp.BaseResp.StatusCode, rpc_resp.BaseResp.StatusMessage)
		}
	default:
		return nil, errno.ParamErr
	}
	return resp, nil
}

func (s *CommentService) CommentList(req douyinapi.CommentListRequest) (*douyinapi.CommentListResponse, error) {
	resp := new(douyinapi.CommentListResponse)
	// get the VideoID
	videoID, err := strconv.ParseInt(req.VideoID, 10, 64)
	if err != nil {
		return resp, err
	}

	rpc_resp, err := rpc.GetVideoComments(s.ctx, &douyincomment.GetVideoCommentsRequest{
		Video: videoID,
	})
	if err != nil {
		return resp, err
	}
	if rpc_resp.BaseResp.StatusCode != 0 {
		return resp, errno.NewErrNo(rpc_resp.BaseResp.StatusCode, rpc_resp.BaseResp.StatusMessage)
	}
	resp.CommentList = make([]*douyinapi.Comment, 0, 50)
	for _, rpc_comment := range rpc_resp.Comments {
		r, err := rpc.MGetUser(s.ctx, &douyinuser.MGetUserRequest{UserIds: []int64{rpc_comment.User}})
		if err != nil || r.BaseResp.StatusCode != 0 || len(r.Users) < 1 {
			continue
		}
		user := pack.PackUser(r.Users[0])
		comment := pack.PackComment(rpc_comment)
		comment.User = user
		resp.CommentList = append(resp.CommentList, comment)
	}
	return resp, nil
}
