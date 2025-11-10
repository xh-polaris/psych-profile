package controller

import (
	"context"

	"github.com/google/wire"
	"github.com/xh-polaris/psych-idl/kitex_gen/basic"
	"github.com/xh-polaris/psych-idl/kitex_gen/profile"
	"github.com/xh-polaris/psych-profile/biz/application/service"
	"github.com/xh-polaris/psych-profile/pkg/logs"
)

var _ IConfigController = (*ConfigController)(nil)

type IConfigController interface {
	ConfigCreate(ctx context.Context, req *profile.ConfigCreateOrUpdateReq) (resp *basic.Response, err error)
	ConfigUpdateInfo(ctx context.Context, req *profile.ConfigCreateOrUpdateReq) (resp *basic.Response, err error)
	ConfigGetByUnitID(ctx context.Context, req *profile.ConfigGetByUnitIdReq) (resp *profile.ConfigGetByUnitIdResp, err error)
}

type ConfigController struct {
	ConfigService *service.ConfigService
}

var ConfigControllerSet = wire.NewSet(
	wire.Struct(new(ConfigController), "*"),
	wire.Bind(new(IConfigController), new(*ConfigController)),
)

func (c *ConfigController) ConfigCreate(ctx context.Context, req *profile.ConfigCreateOrUpdateReq) (resp *basic.Response, err error) {
	logs.Info("ConfigCreate", req)
	return c.ConfigService.ConfigCreate(ctx, req)
}

func (c *ConfigController) ConfigUpdateInfo(ctx context.Context, req *profile.ConfigCreateOrUpdateReq) (resp *basic.Response, err error) {
	logs.Info("ConfigUpdateInfo", req)
	return c.ConfigService.ConfigUpdateInfo(ctx, req)
}

func (c *ConfigController) ConfigGetByUnitID(ctx context.Context, req *profile.ConfigGetByUnitIdReq) (resp *profile.ConfigGetByUnitIdResp, err error) {
	logs.Info("ConfigGetByUnitID", req)
	return c.ConfigService.ConfigGetByUnitID(ctx, req)
}
