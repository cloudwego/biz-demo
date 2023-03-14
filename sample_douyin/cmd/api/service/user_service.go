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
	"math/rand"
	"strconv"

	douyinapi "github.com/cloudwego/biz-demo/sample_douyin/hertz_gen/douyinapi"

	//"mydouyin/cmd/api/biz/cache"

	"github.com/cloudwego/biz-demo/sample_douyin/cmd/api/pack"
	"github.com/cloudwego/biz-demo/sample_douyin/cmd/api/rpc"
	"github.com/cloudwego/biz-demo/sample_douyin/kitex_gen/douyinuser"
	"github.com/cloudwego/biz-demo/sample_douyin/pkg/consts"
	"github.com/cloudwego/biz-demo/sample_douyin/pkg/errno"
)

type UserService struct {
	ctx context.Context
}

func NewUserService(ctx context.Context) *UserService {
	return &UserService{
		ctx: ctx,
	}
}

func (s *UserService) RegisterUser(req douyinapi.RegisterUserRequest) (*douyinapi.RegisterUserResponse, error) {
	resp := new(douyinapi.RegisterUserResponse)
	rpc_resp, err := rpc.CreateUser(context.Background(), &douyinuser.CreateUserRequest{
		Username:        req.Username,
		Password:        req.Password,
		Avatar:          consts.AvatarList[rand.Intn(len(consts.AvatarList))],
		BackgroundImage: consts.BackgroundList[rand.Intn(len(consts.BackgroundList))],
		Signature:       "Hello World!",
	})
	if err != nil {
		return resp, err
	}
	if rpc_resp.BaseResp.StatusCode != 0 {
		return resp, errno.NewErrNo(rpc_resp.BaseResp.StatusCode, rpc_resp.BaseResp.StatusMessage)
	}
	return resp, nil
}

func (s *UserService) GetUser(req douyinapi.GetUserRequest) (*douyinapi.GetUserResponse, error) {
	resp := new(douyinapi.GetUserResponse)
	id, err := strconv.Atoi(req.UserID)
	if err != nil {
		err = errno.ParamErr
		return resp, err
	}

	rpc_resp, err := rpc.MGetUser(s.ctx, &douyinuser.MGetUserRequest{UserIds: []int64{int64(id)}})
	if err != nil {
		return nil, err
	}
	if rpc_resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(rpc_resp.BaseResp.StatusCode, rpc_resp.BaseResp.StatusMessage)
	}
	if len(rpc_resp.Users) < 1 {
		return nil, errno.QueryErr
	}
	resp.User = pack.PackUser(rpc_resp.Users[0])
	return resp, nil
}
