package service

import (
	"context"
	"mydouyin/cmd/user/dal/db"
	"mydouyin/cmd/user/pack"
	"mydouyin/kitex_gen/douyinuser"
)

type MGetUserService struct {
	ctx context.Context
}

// NewMGetUserService new CheckUserService
func NewMGetUserService(ctx context.Context) *MGetUserService {
	return &MGetUserService{
		ctx: ctx,
	}
}

// CreateUser create user info.
func (s *MGetUserService) GetUsers(req *douyinuser.MGetUserRequest) ([]*douyinuser.User, error) {
	users, err := db.MGetUSers(s.ctx, req.UserIds)
	if err != nil {
		return nil, err
	}
	return pack.Users(users), nil
}
