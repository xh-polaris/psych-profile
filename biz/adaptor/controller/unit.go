package controller

import (
	"context"

	"github.com/google/wire"
	"github.com/xh-polaris/psych-idl/kitex_gen/basic"
	u "github.com/xh-polaris/psych-idl/kitex_gen/user"
	"github.com/xh-polaris/psych-profile/biz/application/service"
	"github.com/xh-polaris/psych-profile/pkg/logs"
)

var _ IUnitController = (*UnitController)(nil)

type IUnitController interface {
	UnitSignUp(ctx context.Context, req *u.UnitSignUpReq) (res *u.UnitSignUpResp, err error)
	UnitGetInfo(ctx context.Context, req *u.UnitGetInfoReq) (res *u.UnitGetInfoResp, err error)
	UnitUpdateInfo(ctx context.Context, req *u.UnitUpdateInfoReq) (res *basic.Response, err error)
	UnitUpdatePassword(ctx context.Context, req *u.UnitUpdatePasswordReq) (res *basic.Response, err error)
	UnitCreateAndLinkUser(ctx context.Context, req *u.UnitCreateAndLinkUserReq) (res *basic.Response, err error)
	UnitCreateAndLinkView(ctx context.Context, req *u.UnitCreateAndLinkViewReq) (res *basic.Response, err error)
	UnitSignIn(ctx context.Context, req *u.UnitSignInReq) (res *u.UnitSignInResp, err error)
	UnitCreateVerify(ctx context.Context, req *u.UnitCreateVerifyReq) (res *u.UnitCreateVerifyResp, err error)
	UnitUpdateVerifyPassword(ctx context.Context, req *u.UnitUpdateVerifyReq) (res *basic.Response, err error)
	UnitLinkUser(ctx context.Context, req *u.UnitLinkUserReq) (res *basic.Response, err error)
	UnitLinkView(ctx context.Context, req *u.UnitLinkViewReq) (res *basic.Response, err error)
	UnitPageQueryUser(ctx context.Context, req *u.UnitPageQueryUserReq) (res *u.UnitPageQueryUserResp, err error)
	UnitPageQueryView(ctx context.Context, req *u.UnitPageQueryViewReq) (res *u.UnitPageQueryViewResp, err error)
	UnitGetAppInfo(ctx context.Context, req *u.UnitGetAppInfoReq) (res *u.UnitGetAppInfoResp, err error)
	UnitModelGetInfo(ctx context.Context, req *u.UnitModelGetInfoReq) (res *u.UnitModelGetInfoResp, err error)
	UnitModelUpdateInfo(ctx context.Context, req *u.UnitModelUpdateInfoReq) (res *basic.Response, err error)
}

type UnitController struct {
	UnitService *service.UnitService
}

var UnitControllerSet = wire.NewSet(
	wire.Struct(new(UnitController), "*"),
	wire.Bind(new(IUnitController), new(*UnitController)),
)

func (u *UnitController) UnitSignUp(ctx context.Context, req *u.UnitSignUpReq) (res *u.UnitSignUpResp, err error) {
	logs.Info("UnitSignUp", req)
	return u.UnitService.UnitSignUp(ctx, req)
}

func (u *UnitController) UnitGetInfo(ctx context.Context, req *u.UnitGetInfoReq) (res *u.UnitGetInfoResp, err error) {
	logs.Info("UnitGetInfo", req)
	return u.UnitService.UnitGetInfo(ctx, req)
}
func (u *UnitController) UnitUpdateInfo(ctx context.Context, req *u.UnitUpdateInfoReq) (res *basic.Response, err error) {
	logs.Info("UnitUpdateInfo", req)
	return u.UnitService.UnitUpdateInfo(ctx, req)
}
func (u *UnitController) UnitUpdatePassword(ctx context.Context, req *u.UnitUpdatePasswordReq) (res *basic.Response, err error) {
	logs.Info("UnitUpdatePassword", req)
	return u.UnitService.UnitUpdatePassword(ctx, req)
}
func (u *UnitController) UnitCreateAndLinkUser(ctx context.Context, req *u.UnitCreateAndLinkUserReq) (res *basic.Response, err error) {
	logs.Info("UnitCreateAndLinkUser", req)
	return u.UnitService.UnitCreateAndLinkUser(ctx, req)
}
func (u *UnitController) UnitCreateAndLinkView(ctx context.Context, req *u.UnitCreateAndLinkViewReq) (res *basic.Response, err error) {
	logs.Info("UnitCreateAndLinkView", req)
	return u.UnitService.UnitCreateAndLinkView(ctx, req)
}
func (u *UnitController) UnitSignIn(ctx context.Context, req *u.UnitSignInReq) (res *u.UnitSignInResp, err error) {
	logs.Info("UnitSignIn", req)
	return u.UnitService.UnitSignIn(ctx, req)
}
func (u *UnitController) UnitCreateVerify(ctx context.Context, req *u.UnitCreateVerifyReq) (res *u.UnitCreateVerifyResp, err error) {
	logs.Info("UnitCreateVerify", req)
	return u.UnitService.UnitCreateVerify(ctx, req)
}
func (u *UnitController) UnitUpdateVerifyPassword(ctx context.Context, req *u.UnitUpdateVerifyReq) (res *basic.Response, err error) {
	logs.Info("UnitUpdateVerifyPassword", req)
	return u.UnitService.UnitUpdateVerifyPassword(ctx, req)
}
func (u *UnitController) UnitLinkUser(ctx context.Context, req *u.UnitLinkUserReq) (res *basic.Response, err error) {
	logs.Info("UnitLinkUser", req)
	return u.UnitService.UnitLinkUser(ctx, req)
}
func (u *UnitController) UnitLinkView(ctx context.Context, req *u.UnitLinkViewReq) (res *basic.Response, err error) {
	logs.Info("UnitLinkView", req)
	return u.UnitService.UnitLinkView(ctx, req)
}
func (u *UnitController) UnitPageQueryUser(ctx context.Context, req *u.UnitPageQueryUserReq) (res *u.UnitPageQueryUserResp, err error) {
	logs.Info("UnitPageQueryUser", req)
	return u.UnitService.UnitPageQueryUser(ctx, req)
}
func (u *UnitController) UnitPageQueryView(ctx context.Context, req *u.UnitPageQueryViewReq) (res *u.UnitPageQueryViewResp, err error) {
	logs.Info("UnitPageQueryView", req)
	return u.UnitService.UnitPageQueryView(ctx, req)
}
func (u *UnitController) UnitGetAppInfo(ctx context.Context, req *u.UnitGetAppInfoReq) (res *u.UnitGetAppInfoResp, err error) {
	logs.Info("UnitGetAppInfo", req)
	return u.UnitService.UnitGetAppInfo(ctx, req)
}
func (u *UnitController) UnitModelGetInfo(ctx context.Context, req *u.UnitModelGetInfoReq) (res *u.UnitModelGetInfoResp, err error) {
	logs.Info("UnitModelGetInfo", req)
	return u.UnitService.UnitModelGetInfo(ctx, req)
}
func (u *UnitController) UnitModelUpdateInfo(ctx context.Context, req *u.UnitModelUpdateInfoReq) (res *basic.Response, err error) {
	logs.Info("UnitModelUpdateInfo", req)
	return u.UnitService.UnitModelUpdateInfo(ctx, req)
}
