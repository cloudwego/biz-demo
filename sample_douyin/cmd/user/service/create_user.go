package service

import (
	"context"
	"crypto/md5"
	"fmt"
	"io"
	"mydouyin/cmd/user/dal/db"
	"mydouyin/kitex_gen/douyinuser"
	"mydouyin/pkg/errno"
)

type CreateUserService struct {
	ctx context.Context
}

// NewCreateUserService new CreateUserService
func NewCreateUserService(ctx context.Context) *CreateUserService {
	return &CreateUserService{ctx: ctx}
}

// CreateUser create user info.
func (s *CreateUserService) CreateUser(req *douyinuser.CreateUserRequest) error {
	users, err := db.QueryUser(s.ctx, req.Username)
	if err != nil {
		return err
	}
	if len(users) != 0 {
		return errno.UserAlreadyExistErr
	}

	h := md5.New()
	if _, err := io.WriteString(h, req.Password); err != nil {
		return err
	}
	passWord := fmt.Sprintf("%x", h.Sum(nil))
	return db.CreateUser(s.ctx, []*db.User{{
		Username: req.Username,
		Password: passWord,
		Avatar: req.Avatar,
		Signature: req.Signature,
		BackgroundImage: req.BackgroundImage,
	}})
}
