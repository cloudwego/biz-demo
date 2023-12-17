package service

import (
	"context"

	user "github.com/baiyutang/gomall/app/user/kitex_gen/user"
)

type LoginService struct {
	ctx context.Context
} // NewLoginService new LoginService
func NewLoginService(ctx context.Context) *LoginService {
	return &LoginService{ctx: ctx}
}

// Run create note info
func (s *LoginService) Run(req *user.LoginReq) (resp *user.LoginRes, err error) {
	// Finish your business logic.

	return &user.LoginRes{Userid: 1}, nil
}
