package service

import (
	"context"

	user "github.com/baiyutang/gomall/app/user/kitex_gen/user"
)

type RegisterService struct {
	ctx context.Context
} // NewRegisterService new RegisterService
func NewRegisterService(ctx context.Context) *RegisterService {
	return &RegisterService{ctx: ctx}
}

// Run create note info
func (s *RegisterService) Run(req *user.RegisterReq) (resp *user.RegisterRes, err error) {
	// Finish your business logic.

	return &user.RegisterRes{Userid: 1}, nil
}
