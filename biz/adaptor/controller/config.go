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
	ConfigUpdate(ctx context.Context, req *profile.ConfigCreateOrUpdateReq) (resp *basic.Response, err error)
	ConfigFindByUnitID(ctx context.Context, req *profile.ConfigGetByUnitIdReq) (resp *profile.ConfigGetByUnitIdResp, err error)
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

func (c *ConfigController) ConfigUpdate(ctx context.Context, req *profile.ConfigCreateOrUpdateReq) (resp *basic.Response, err error) {
	logs.Info("ConfigUpdate", req)
	return c.ConfigService.ConfigUpdate(ctx, req)
}

func (c *ConfigController) ConfigFindByUnitID(ctx context.Context, req *profile.ConfigGetByUnitIdReq) (resp *profile.ConfigGetByUnitIdResp, err error) {
	logs.Info("ConfigGetByUnitID", req)
	return c.ConfigService.ConfigGetByUnitID(ctx, req)
}
