package controller

import (
	"context"

	"github.com/google/wire"
	"github.com/xh-polaris/psych-idl/kitex_gen/basic"
	u "github.com/xh-polaris/psych-idl/kitex_gen/user"
	"github.com/xh-polaris/psych-profile/biz/application/service"
	"github.com/xh-polaris/psych-profile/pkg/logs"
)

var _ IUserController = (*UserController)(nil)

type IUserController interface {
	UserSignUp(ctx context.Context, req *u.UserSignUpReq) (res *u.UserSignUpResp, err error)
	UserGetInfo(ctx context.Context, req *u.UserGetInfoReq) (res *u.UserGetInfoResp, err error)
	UserUpdateInfo(ctx context.Context, req *u.UserUpdateInfoReq) (res *basic.Response, err error)
	UserUpdatePassword(ctx context.Context, req *u.UserUpdatePasswordReq) (res *basic.Response, err error)
	UserBelongUnit(ctx context.Context, req *u.UserBelongUnitReq) (res *u.UserBelongUnitResp, err error)
	UserSignIn(ctx context.Context, req *u.UserSignInReq) (res *u.UserSignInResp, err error)
}

type UserController struct {
	UserService *service.UserService
}

var UserControllerSet = wire.NewSet(
	wire.Struct(new(UserController), "*"),
	wire.Bind(new(IUserController), new(*UserController)),
)

func (u *UserController) UserSignUp(ctx context.Context, req *u.UserSignUpReq) (res *u.UserSignUpResp, err error) {
	logs.Info("UserSignUp", req)
	return u.UserService.UserSignUp(ctx, req)
}

func (u *UserController) UserGetInfo(ctx context.Context, req *u.UserGetInfoReq) (res *u.UserGetInfoResp, err error) {
	logs.Info("UserGetInfo", req)
	return u.UserService.UserGetInfo(ctx, req)
}

func (u *UserController) UserUpdateInfo(ctx context.Context, req *u.UserUpdateInfoReq) (res *basic.Response, err error) {
	logs.Info("UserUpdateInfo", req)
	return u.UserService.UserUpdateInfo(ctx, req)
}
func (u *UserController) UserUpdatePassword(ctx context.Context, req *u.UserUpdatePasswordReq) (res *basic.Response, err error) {
	logs.Info("UserUpdatePassword", req)
	return u.UserService.UserUpdatePassword(ctx, req)
}
func (u *UserController) UserBelongUnit(ctx context.Context, req *u.UserBelongUnitReq) (res *u.UserBelongUnitResp, err error) {
	logs.Info("UserBelongUnit", req)
	return u.UserService.UserBelongUnit(ctx, req)
}
func (u *UserController) UserSignIn(ctx context.Context, req *u.UserSignInReq) (res *u.UserSignInResp, err error) {
	logs.Info("UserSignIn", req)
	return u.UserService.UserSignIn(ctx, req)
}
