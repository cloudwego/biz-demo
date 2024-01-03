package service

import (
	"context"

	"github.com/baiyutang/gomall/app/user/biz/dal/mysql"
	"github.com/baiyutang/gomall/app/user/biz/model"
	user "github.com/baiyutang/gomall/app/user/kitex_gen/user"
	"github.com/cloudwego/kitex/pkg/klog"
	"golang.org/x/crypto/bcrypt"
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
	klog.Infof("LoginReq:%+v", req)
	userRow, err := model.GetByEmail(mysql.DB, s.ctx, req.Email)
	if err != nil {
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(userRow.PasswordHashed), []byte(req.Password))
	if err != nil {
		return
	}
	return &user.LoginRes{Userid: int32(userRow.ID)}, nil
}
