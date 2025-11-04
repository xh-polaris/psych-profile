package controller

import (
	"github.com/google/wire"
	"github.com/xh-polaris/psych-profile/biz/application/service"
)

var _ IConfigController = (*ConfigController)(nil)

type IConfigController interface {
}

type ConfigController struct {
	ConfigService *service.ConfigService
}

var ConfigControllerSet = wire.NewSet(
	wire.Struct(new(ConfigController), "*"),
	wire.Bind(new(IConfigController), new(*ConfigController)),
)
