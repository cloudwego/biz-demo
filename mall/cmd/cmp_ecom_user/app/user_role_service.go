package app

import (
	"context"
	"github.com/cloudwego/biz-demo/mall/cmd/cmp_ecom_user/infras/persistence/dal/db"
	"github.com/cloudwego/biz-demo/mall/cmd/cmp_ecom_user/kitex_gen/cmp/ecom/user"
)

type UserRoleService struct {
	ctx context.Context
}

func NewUserRoleService(ctx context.Context) *UserRoleService {
	return &UserRoleService{
		ctx: ctx,
	}
}

func (s *UserRoleService) AddUserRole(req *user.AddUserRoleReq) error {
	return db.AddUserRole(s.ctx, req.UserName, req.Role)
}

func (s *UserRoleService) DelUserRole(req *user.DelUserRoleReq) error {
	return db.DelUserRole(s.ctx, req.GetUserName(), req.GetRole())
}

func (s *UserRoleService) ValidateUserRole(req *user.ValidateUserRolesReq) (bool, error) {
	return db.ValidateUserRole(s.ctx, req.GetUserName(), req.GetRoles())
}
