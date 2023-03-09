package service

import (
	"context"
	"math/rand"
	"mydouyin/cmd/api/biz/apimodel"

	//"mydouyin/cmd/api/biz/cache"
	"mydouyin/cmd/api/biz/rpc"
	"mydouyin/kitex_gen/douyinuser"
	"mydouyin/pkg/consts"
	"mydouyin/pkg/errno"
	"strconv"
)

type UserService struct {
	ctx context.Context
}

func NewUserService(ctx context.Context) *UserService {
	return &UserService{
		ctx: ctx,
	}
}

func (s *UserService) RegistUser(req apimodel.RegistUserRequest) (*apimodel.RegistUserResponse, error) {
	resp := new(apimodel.RegistUserResponse)
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

func (s *UserService) GetUser(req apimodel.GetUserRequest) (*apimodel.GetUserResponse, error) {
	resp := new(apimodel.GetUserResponse)
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
	resp.User = *apimodel.PackUser(rpc_resp.Users[0])
	return resp, nil
}
