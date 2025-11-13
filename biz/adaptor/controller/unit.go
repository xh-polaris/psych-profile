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

var _ IUnitController = (*UnitController)(nil)

type IUnitController interface {
	UnitSignUp(ctx context.Context, req *profile.UnitSignUpReq) (resp *profile.UnitSignUpResp, err error)
	UnitGetInfo(ctx context.Context, req *profile.UnitGetInfoReq) (resp *profile.UnitGetInfoResp, err error)
	UnitUpdateInfo(ctx context.Context, req *profile.UnitUpdateInfoReq) (resp *basic.Response, err error)
	UnitUpdatePassword(ctx context.Context, req *profile.UnitUpdatePasswordReq) (resp *basic.Response, err error)
	UnitCreateAndLinkUser(ctx context.Context, req *profile.UnitCreateAndLinkUserReq) (resp *profile.UnitCreateAndLinkUserResp, err error)
	UnitSignIn(ctx context.Context, req *profile.UnitSignInReq) (resp *profile.UnitSignInResp, err error)
	UnitLinkUser(ctx context.Context, req *profile.UnitLinkUserReq) (resp *basic.Response, err error) // Deprecated
}

type UnitController struct {
	UnitService *service.UnitService
}

var UnitControllerSet = wire.NewSet(
	wire.Struct(new(UnitController), "*"),
	wire.Bind(new(IUnitController), new(*UnitController)),
)

func (u *UnitController) UnitSignUp(ctx context.Context, req *profile.UnitSignUpReq) (resp *profile.UnitSignUpResp, err error) {
	resp, err = u.UnitService.UnitSignUp(ctx, req)
	logs.CtxInfof(ctx, "[%s] req=%s, resp=%s, err=%s", "UnitSignUp", util.JSONF(req), util.JSONF(resp), errorx.ErrorWithoutStack(err))
	return
}

func (u *UnitController) UnitGetInfo(ctx context.Context, req *profile.UnitGetInfoReq) (resp *profile.UnitGetInfoResp, err error) {
	resp, err = u.UnitService.UnitGetInfo(ctx, req)
	logs.CtxInfof(ctx, "[%s] req=%s, resp=%s, err=%s", "UnitGetInfo", util.JSONF(req), util.JSONF(resp), errorx.ErrorWithoutStack(err))
	return
}
func (u *UnitController) UnitUpdateInfo(ctx context.Context, req *profile.UnitUpdateInfoReq) (resp *basic.Response, err error) {
	resp, err = u.UnitService.UnitUpdateInfo(ctx, req)
	logs.CtxInfof(ctx, "[%s] req=%s, resp=%s, err=%s", "UnitUpdateInfo", util.JSONF(req), util.JSONF(resp), errorx.ErrorWithoutStack(err))
	return
}
func (u *UnitController) UnitUpdatePassword(ctx context.Context, req *profile.UnitUpdatePasswordReq) (resp *basic.Response, err error) {
	resp, err = u.UnitService.UnitUpdatePassword(ctx, req)
	logs.CtxInfof(ctx, "[%s] req=%s, resp=%s, err=%s", "UnitUpdatePassword", util.JSONF(req), util.JSONF(resp), errorx.ErrorWithoutStack(err))
	return
}
func (u *UnitController) UnitCreateAndLinkUser(ctx context.Context, req *profile.UnitCreateAndLinkUserReq) (resp *profile.UnitCreateAndLinkUserResp, err error) {
	resp, err = u.UnitService.UnitCreateAndLinkUser(ctx, req)
	logs.CtxInfof(ctx, "[%s] req=%s, resp=%s, err=%s", "UnitCreateAndLinkUser", util.JSONF(req), util.JSONF(resp), errorx.ErrorWithoutStack(err))
	return
}

func (u *UnitController) UnitSignIn(ctx context.Context, req *profile.UnitSignInReq) (resp *profile.UnitSignInResp, err error) {
	resp, err = u.UnitService.UnitSignIn(ctx, req)
	logs.CtxInfof(ctx, "[%s] req=%s, resp=%s, err=%s", "UnitSignIn", util.JSONF(req), util.JSONF(resp), errorx.ErrorWithoutStack(err))
	return
}

func (u *UnitController) UnitLinkUser(ctx context.Context, req *profile.UnitLinkUserReq) (resp *basic.Response, err error) {
	resp, err = u.UnitService.UnitLinkUser(ctx, req)
	logs.CtxInfof(ctx, "[%s] req=%s, resp=%s, err=%s", "UnitLinkUser", util.JSONF(req), util.JSONF(resp), errorx.ErrorWithoutStack(err))
	return
}
