package controller

import (
	"context"

	"github.com/google/wire"
	"github.com/xh-polaris/gopkg/util"
	"github.com/xh-polaris/psych-idl/kitex_gen/profile"
	"github.com/xh-polaris/psych-profile/biz/application/service"
	"github.com/xh-polaris/psych-profile/pkg/errorx"
	"github.com/xh-polaris/psych-profile/pkg/logs"
)

var _ IDashboardController = (*DashboardController)(nil)

type IDashboardController interface {
	// 数据看板
	DashboardGetDataOverview(ctx context.Context, req *profile.DashboardGetDataOverviewReq) (resp *profile.DashboardGetDataOverviewResp, err error)
	DashboardGetDataTrend(ctx context.Context, req *profile.DashboardGetDataTrendReq) (resp *profile.DashboardGetDataTrendResp, err error)
	DashboardGetPsycheTrend(ctx context.Context, req *profile.DashboardGetPsycheTrendReq) (resp *profile.DashboardGetPsycheTrendResp, err error)
	// 预警管理
	DashboardGetAlertOverview(ctx context.Context, req *profile.DashboardGetAlertOverviewReq) (resp *profile.DashboardGetAlertOverviewResp, err error)
	DashboardListAlertUsers(ctx context.Context, req *profile.DashboardListAlertUsersReq) (resp *profile.DashboardListAlertUsersResp, err error)
	// 用户管理
	DashboardListClasses(ctx context.Context, req *profile.DashboardListClassesReq) (resp *profile.DashboardListClassesResp, err error)
	DashboardListUsers(ctx context.Context, req *profile.DashboardListUsersReq) (resp *profile.DashboardListUsersResp, err error)
}

type DashboardController struct {
	DashboardService *service.DashboardService
}

var DashboardControllerSet = wire.NewSet(
	wire.Struct(new(DashboardController), "*"),
	wire.Bind(new(IDashboardController), new(*DashboardController)),
)

func (c *DashboardController) DashboardGetDataOverview(ctx context.Context, req *profile.DashboardGetDataOverviewReq) (resp *profile.DashboardGetDataOverviewResp, err error) {
	resp, err := c.DashboardService.DashboardGetDataOverview(ctx, req)
	logs.CtxInfof(ctx, "[%s] req=%s, resp=%s, err=%s", "DashboardGetDataOverview", util.JSONF(req), util.JSONF(resp), errorx.ErrorWithoutStack(err))
	return resp, err
}

func (c *DashboardController) DashboardGetDataTrend(ctx context.Context, req *profile.DashboardGetDataTrendReq) (resp *profile.DashboardGetDataTrendResp, err error) {
	resp, err := c.DashboardService.DashboardGetDataTrend(ctx, req)
	logs.CtxInfof(ctx, "[%s] req=%s, resp=%s, err=%s", "DashboardGetDataTrend", util.JSONF(req), util.JSONF(resp), errorx.ErrorWithoutStack(err))
	return resp, err
}

func (c *DashboardController) DashboardGetPsycheTrend(ctx context.Context, req *profile.DashboardGetPsycheTrendReq) (resp *profile.DashboardGetPsycheTrendResp, err error) {
	resp, err := c.DashboardService.DashboardGetPsycheTrend(ctx, req)
	logs.CtxInfof(ctx, "[%s] req=%s, resp=%s, err=%s", "DashboardGetPsycheTrend", util.JSONF(req), util.JSONF(resp), errorx.ErrorWithoutStack(err))
	return resp, err
}

func (c *DashboardController) DashboardGetAlertOverview(ctx context.Context, req *profile.DashboardGetAlertOverviewReq) (resp *profile.DashboardGetAlertOverviewResp, err error) {
	resp, err := c.DashboardService.DashboardGetAlertOverview(ctx, req)
	logs.CtxInfof(ctx, "[%s] req=%s, resp=%s, err=%s", "DashboardGetAlertOverview", util.JSONF(req), util.JSONF(resp), errorx.ErrorWithoutStack(err))
	return resp, err
}

func (c *DashboardController) DashboardListAlertUsers(ctx context.Context, req *profile.DashboardListAlertUsersReq) (resp *profile.DashboardListAlertUsersResp, err error) {
	resp, err := c.DashboardService.DashboardListAlertUsers(ctx, req)
	logs.CtxInfof(ctx, "[%s] req=%s, resp=%s, err=%s", "DashboardListAlertUsers", util.JSONF(req), util.JSONF(resp), errorx.ErrorWithoutStack(err))
	return resp, err
}

func (c *DashboardController) DashboardListClasses(ctx context.Context, req *profile.DashboardListClassesReq) (resp *profile.DashboardListClassesResp, err error) {
	resp, err := c.DashboardService.DashboardListClasses(ctx, req)
	logs.CtxInfof(ctx, "[%s] req=%s, resp=%s, err=%s", "DashboardListClasses", util.JSONF(req), util.JSONF(resp), errorx.ErrorWithoutStack(err))
	return resp, err
}

func (c *DashboardController) DashboardListUsers(ctx context.Context, req *profile.DashboardListUsersReq) (resp *profile.DashboardListUsersResp, err error) {
	resp, err := c.DashboardService.DashboardListUsers(ctx, req)
	logs.CtxInfof(ctx, "[%s] req=%s, resp=%s, err=%s", "DashboardListUsers", util.JSONF(req), util.JSONF(resp), errorx.ErrorWithoutStack(err))
	return resp, err
}
