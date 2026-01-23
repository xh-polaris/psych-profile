package service

import (
	"context"

	"github.com/google/wire"
	"github.com/xh-polaris/psych-idl/kitex_gen/profile"
)

var _ IDashboardService = (*DashboardService)(nil)

type IDashboardService interface {
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

type DashboardService struct {
}

var DashboardServiceSet = wire.NewSet(
	wire.Struct(new(DashboardService), "*"),
	wire.Bind(new(IDashboardService), new(*DashboardService)),
)

func (s *DashboardService) DashboardGetDataOverview(ctx context.Context, req *profile.DashboardGetDataOverviewReq) (resp *profile.DashboardGetDataOverviewResp, err error) {
	panic("implement me")
}

func (s *DashboardService) DashboardGetDataTrend(ctx context.Context, req *profile.DashboardGetDataTrendReq) (resp *profile.DashboardGetDataTrendResp, err error) {
	panic("implement me")
}

func (s *DashboardService) DashboardGetPsycheTrend(ctx context.Context, req *profile.DashboardGetPsycheTrendReq) (resp *profile.DashboardGetPsycheTrendResp, err error) {
	panic("implement me")
}

func (s *DashboardService) DashboardGetAlertOverview(ctx context.Context, req *profile.DashboardGetAlertOverviewReq) (resp *profile.DashboardGetAlertOverviewResp, err error) {
	panic("implement me")
}

func (s *DashboardService) DashboardListAlertUsers(ctx context.Context, req *profile.DashboardListAlertUsersReq) (resp *profile.DashboardListAlertUsersResp, err error) {
	panic("implement me")
}

func (s *DashboardService) DashboardListClasses(ctx context.Context, req *profile.DashboardListClassesReq) (resp *profile.DashboardListClassesResp, err error) {
	panic("implement me")
}

func (s *DashboardService) DashboardListUsers(ctx context.Context, req *profile.DashboardListUsersReq) (resp *profile.DashboardListUsersResp, err error) {
	panic("implement me")
}
