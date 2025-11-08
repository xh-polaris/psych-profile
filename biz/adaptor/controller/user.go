package controller

import (
	"context"

	"github.com/google/wire"
	"github.com/xh-polaris/psych-idl/kitex_gen/basic"
	"github.com/xh-polaris/psych-idl/kitex_gen/profile"
	"github.com/xh-polaris/psych-profile/biz/application/service"
	"github.com/xh-polaris/psych-profile/pkg/logs"
)

var _ IUserController = (*UserController)(nil)

type IUserController interface {
	UserSignUp(ctx context.Context, req *profile.UserSignUpReq) (resp *profile.UserSignUpResp, err error)
	UserGetInfo(ctx context.Context, req *profile.UserGetInfoReq) (resp *profile.UserGetInfoResp, err error)
	UserUpdateInfo(ctx context.Context, req *profile.UserUpdateInfoReq) (resp *basic.Response, err error)
	UserUpdatePassword(ctx context.Context, req *profile.UserUpdatePasswordReq) (resp *basic.Response, err error)
	UserSignIn(ctx context.Context, req *profile.UserSignInReq) (resp *profile.UserSignInResp, err error)
}

type UserController struct {
	UserService *service.UserService
}

var UserControllerSet = wire.NewSet(
	wire.Struct(new(UserController), "*"),
	wire.Bind(new(IUserController), new(*UserController)),
)

func (u *UserController) UserSignUp(ctx context.Context, req *profile.UserSignUpReq) (resp *profile.UserSignUpResp, err error) {
	logs.Info("UserSignUp", req)
	return u.UserService.UserSignUp(ctx, req)
}

func (u *UserController) UserGetInfo(ctx context.Context, req *profile.UserGetInfoReq) (resp *profile.UserGetInfoResp, err error) {
	logs.Info("UserGetInfo", req)
	return u.UserService.UserGetInfo(ctx, req)
}

func (u *UserController) UserUpdateInfo(ctx context.Context, req *profile.UserUpdateInfoReq) (resp *basic.Response, err error) {
	logs.Info("UserUpdateInfo", req)
	return u.UserService.UserUpdateInfo(ctx, req)
}
func (u *UserController) UserUpdatePassword(ctx context.Context, req *profile.UserUpdatePasswordReq) (resp *basic.Response, err error) {
	logs.Info("UserUpdatePassword", req)
	return u.UserService.UserUpdatePassword(ctx, req)
}

func (u *UserController) UserSignIn(ctx context.Context, req *profile.UserSignInReq) (resp *profile.UserSignInResp, err error) {
	logs.Info("UserSignIn", req)
	return u.UserService.UserSignIn(ctx, req)
}
