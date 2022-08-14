package main

import (
	"context"
	"github.com/cloudwego/biz-demo/mall/cmd/cmp_ecom_user/app"
	"github.com/cloudwego/biz-demo/mall/cmd/cmp_ecom_user/common"
	"github.com/cloudwego/biz-demo/mall/cmd/cmp_ecom_user/kitex_gen/cmp/ecom/user"
	"github.com/cloudwego/biz-demo/mall/pkg/errno"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// CreateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CreateUser(ctx context.Context, req *user.CreateUserReq) (resp *user.CreateUserResp, err error) {
	resp = user.NewCreateUserResp()

	if len(req.GetUserName()) == 0 || len(req.GetPassword()) == 0 {
		resp.BaseResp = common.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	err = app.NewUserService(ctx).CreateUser(req)
	if err != nil {
		resp.BaseResp = common.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = common.BuildBaseResp(errno.Success)
	return resp, nil
}

// MGetUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) MGetUser(ctx context.Context, req *user.MGetUserReq) (resp *user.MGetUserResp, err error) {
	resp = user.NewMGetUserResp()

	if len(req.Ids) == 0 {
		resp.BaseResp = common.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	users, err := app.NewUserService(ctx).MGetUser(req)
	if err != nil {
		resp.BaseResp = common.BuildBaseResp(err)
		return resp, nil
	}
	resp.Users = users
	resp.BaseResp = common.BuildBaseResp(errno.Success)
	return resp, nil
}

// CheckUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CheckUser(ctx context.Context, req *user.CheckUserReq) (resp *user.CheckUserResp, err error) {
	resp = user.NewCheckUserResp()

	if len(req.Password) == 0 || len(req.UserName) == 0 {
		resp.BaseResp = common.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	userId, err := app.NewUserService(ctx).CheckUser(req)
	if err != nil {
		resp.BaseResp = common.BuildBaseResp(err)
		return resp, nil
	}
	resp.UserId = userId
	resp.BaseResp = common.BuildBaseResp(errno.Success)
	return resp, nil
}

// AddUserRole implements the UserServiceImpl interface.
func (s *UserServiceImpl) AddUserRole(ctx context.Context, req *user.AddUserRoleReq) (resp *user.AddUserRoleResp, err error) {
	resp = user.NewAddUserRoleResp()

	err = app.NewUserRoleService(ctx).AddUserRole(req)
	if err != nil {
		resp.BaseResp = common.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = common.BuildBaseResp(errno.Success)
	return resp, nil
}

// DelUserRole implements the UserServiceImpl interface.
func (s *UserServiceImpl) DelUserRole(ctx context.Context, req *user.DelUserRoleReq) (resp *user.DelUserRoleResp, err error) {
	resp = user.NewDelUserRoleResp()
	err = app.NewUserRoleService(ctx).DelUserRole(req)
	if err != nil {
		resp.BaseResp = common.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = common.BuildBaseResp(errno.Success)
	return resp, nil
}

// ValidateUserRole implements the UserServiceImpl interface.
func (s *UserServiceImpl) ValidateUserRole(ctx context.Context, req *user.ValidateUserRolesReq) (resp *user.ValidateUserRoleResp, err error) {
	resp = user.NewValidateUserRoleResp()
	pass, err := app.NewUserRoleService(ctx).ValidateUserRole(req)
	if err != nil {
		resp.BaseResp = common.BuildBaseResp(err)
		return resp, nil
	}
	resp.IsPass = pass
	resp.BaseResp = common.BuildBaseResp(errno.Success)
	return resp, nil
}
