package controller

import (
	"context"

	"github.com/google/wire"
	"github.com/xh-polaris/gopkg/util"
	"github.com/xh-polaris/psych-idl/kitex_gen/basic"
	"github.com/xh-polaris/psych-idl/kitex_gen/profile"
	"github.com/xh-polaris/psych-profile/biz/application/service"
	"github.com/xh-polaris/psych-profile/pkg/errorx"
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
	resp, err = u.UserService.UserSignUp(ctx, req)
	logs.CtxInfof(ctx, "[%s] req=%s, resp=%s, err=%s", "UserSignUp", util.JSONF(req), util.JSONF(resp), errorx.ErrorWithoutStack(err))
	return
}

func (u *UserController) UserGetInfo(ctx context.Context, req *profile.UserGetInfoReq) (resp *profile.UserGetInfoResp, err error) {
	resp, err = u.UserService.UserGetInfo(ctx, req)
	logs.CtxInfof(ctx, "[%s] req=%s, resp=%s, err=%s", "UserGetInfo", util.JSONF(req), util.JSONF(resp), errorx.ErrorWithoutStack(err))
	return
}

func (u *UserController) UserUpdateInfo(ctx context.Context, req *profile.UserUpdateInfoReq) (resp *basic.Response, err error) {
	resp, err = u.UserService.UserUpdateInfo(ctx, req)
	logs.CtxInfof(ctx, "[%s] req=%s, resp=%s, err=%s", "UserUpdateInfo", util.JSONF(req), util.JSONF(resp), errorx.ErrorWithoutStack(err))
	return
}
func (u *UserController) UserUpdatePassword(ctx context.Context, req *profile.UserUpdatePasswordReq) (resp *basic.Response, err error) {
	resp, err = u.UserService.UserUpdatePassword(ctx, req)
	logs.CtxInfof(ctx, "[%s] req=%s, resp=%s, err=%s", "UserUpdatePassword", util.JSONF(req), util.JSONF(resp), errorx.ErrorWithoutStack(err))
	return
}

func (u *UserController) UserSignIn(ctx context.Context, req *profile.UserSignInReq) (resp *profile.UserSignInResp, err error) {
	resp, err = u.UserService.UserSignIn(ctx, req)
	logs.CtxInfof(ctx, "[%s] req=%s, resp=%s, err=%s", "UserSignIn", util.JSONF(req), util.JSONF(resp), errorx.ErrorWithoutStack(err))
	return
}
