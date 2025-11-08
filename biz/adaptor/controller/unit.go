package controller

import (
	"context"

	"github.com/google/wire"
	"github.com/xh-polaris/psych-idl/kitex_gen/basic"
	"github.com/xh-polaris/psych-idl/kitex_gen/profile"
	"github.com/xh-polaris/psych-profile/biz/application/service"
	"github.com/xh-polaris/psych-profile/pkg/logs"
)

var _ IUnitController = (*UnitController)(nil)

type IUnitController interface {
	UnitSignUp(ctx context.Context, req *profile.UnitSignUpReq) (resp *profile.UnitSignUpResp, err error)
	UnitGetInfo(ctx context.Context, req *profile.UnitGetInfoReq) (resp *profile.UnitGetInfoResp, err error)
	UnitUpdateInfo(ctx context.Context, req *profile.UnitUpdateInfoReq) (resp *basic.Response, err error)
	UnitUpdatePassword(ctx context.Context, req *profile.UnitUpdatePasswordReq) (resp *basic.Response, err error)
	UnitCreateAndLinkUser(ctx context.Context, req *profile.UnitCreateAndLinkUserReq) (resp *basic.Response, err error)
	UnitSignIn(ctx context.Context, req *profile.UnitSignInReq) (resp *profile.UnitSignInResp, err error)
	UnitLinkUser(ctx context.Context, req *profile.UnitLinkUserReq) (resp *basic.Response, err error)
}

type UnitController struct {
	UnitService *service.UnitService
}

var UnitControllerSet = wire.NewSet(
	wire.Struct(new(UnitController), "*"),
	wire.Bind(new(IUnitController), new(*UnitController)),
)

func (u *UnitController) UnitSignUp(ctx context.Context, req *profile.UnitSignUpReq) (resp *profile.UnitSignUpResp, err error) {
	logs.Info("UnitSignUp", req)
	return u.UnitService.UnitSignUp(ctx, req)
}

func (u *UnitController) UnitGetInfo(ctx context.Context, req *profile.UnitGetInfoReq) (resp *profile.UnitGetInfoResp, err error) {
	logs.Info("UnitGetInfo", req)
	return u.UnitService.UnitGetInfo(ctx, req)
}
func (u *UnitController) UnitUpdateInfo(ctx context.Context, req *profile.UnitUpdateInfoReq) (resp *basic.Response, err error) {
	logs.Info("UnitUpdateInfo", req)
	return u.UnitService.UnitUpdateInfo(ctx, req)
}
func (u *UnitController) UnitUpdatePassword(ctx context.Context, req *profile.UnitUpdatePasswordReq) (resp *basic.Response, err error) {
	logs.Info("UnitUpdatePassword", req)
	return u.UnitService.UnitUpdatePassword(ctx, req)
}
func (u *UnitController) UnitCreateAndLinkUser(ctx context.Context, req *profile.UnitCreateAndLinkUserReq) (resp *basic.Response, err error) {
	logs.Info("UnitCreateAndLinkUser", req)
	return u.UnitService.UnitCreateAndLinkUser(ctx, req)
}

func (u *UnitController) UnitSignIn(ctx context.Context, req *profile.UnitSignInReq) (resp *profile.UnitSignInResp, err error) {
	logs.Info("UnitSignIn", req)
	return u.UnitService.UnitSignIn(ctx, req)
}

func (u *UnitController) UnitLinkUser(ctx context.Context, req *profile.UnitLinkUserReq) (resp *basic.Response, err error) {
	logs.Info("UnitLinkUser", req)
	return u.UnitService.UnitLinkUser(ctx, req)
}
