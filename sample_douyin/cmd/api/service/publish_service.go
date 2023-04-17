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

	"github.com/cloudwego/biz-demo/sample_douyin/cmd/api/pack"
	"github.com/cloudwego/biz-demo/sample_douyin/cmd/api/rpc"
	"github.com/cloudwego/biz-demo/sample_douyin/cmd/api/videohandler"
	douyinapi "github.com/cloudwego/biz-demo/sample_douyin/hertz_gen/douyinapi"
	"github.com/cloudwego/biz-demo/sample_douyin/kitex_gen/douyinuser"
	"github.com/cloudwego/biz-demo/sample_douyin/kitex_gen/douyinvideo"
	"github.com/cloudwego/biz-demo/sample_douyin/pkg/errno"
)

type PublishService struct {
	ctx context.Context
}

func NewPublishService(ctx context.Context) *PublishService {
	return &PublishService{
		ctx: ctx,
	}
}

func (s *PublishService) PublishVideo(req douyinapi.PublishVideoRequest, user *douyinapi.User) (*douyinapi.PublishVideoResponse, error) {
	resp := new(douyinapi.PublishVideoResponse)
	// err := videohandler.VH.UpLoadVideoV0(req.Data, user.UserID, req.Title)
	videoName, err := videohandler.VH.UpLoadVideo(req.Data)
	if err != nil {
		return resp, err
	}

	go videohandler.VH.CommitCommand(videoName, user.ID, req.Title)

	return resp, err
}

func (s *PublishService) GetPublishList(req douyinapi.GetPublishListRequest, user *douyinapi.User) (*douyinapi.GetPublishListResponse, error) {
	resp := new(douyinapi.GetPublishListResponse)
	userId, err := strconv.Atoi(req.UserID)
	if err != nil {
		return resp, err
	}
	rpc_resp, err := rpc.GetList(s.ctx, &douyinvideo.GetListRequest{
		UserId: int64(userId),
	})
	if err != nil {
		return nil, err
	}
	if rpc_resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(rpc_resp.BaseResp.StatusCode, rpc_resp.BaseResp.StatusMessage)
	}
	resp.VideoList = make([]*douyinapi.Video, 0, 50)
	for _, rpc_video := range rpc_resp.VideoList {
		r, err := rpc.MGetUser(s.ctx, &douyinuser.MGetUserRequest{UserIds: []int64{rpc_video.Author}})
		if err != nil || r.BaseResp.StatusCode != 0 || len(r.Users) < 1 {
			continue
		}
		author := pack.PackUser(r.Users[0])
		video := pack.PackVideo(rpc_video)
		video.Author = author
		resp.VideoList = append(resp.VideoList, video)
	}
	return resp, nil
}
