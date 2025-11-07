package provider

import (
	"github.com/google/wire"
	"github.com/xh-polaris/psych-profile/biz/adaptor/controller"
	"github.com/xh-polaris/psych-profile/biz/application/service"
	infraconfig "github.com/xh-polaris/psych-profile/biz/infra/config"
	"github.com/xh-polaris/psych-profile/biz/infra/mapper/config"
	"github.com/xh-polaris/psych-profile/biz/infra/mapper/unit"
	"github.com/xh-polaris/psych-profile/biz/infra/mapper/user"
)

var ControllerSet = wire.NewSet(
	controller.UserControllerSet,
	controller.UnitControllerSet,
	controller.ConfigControllerSet,
)

var ApplicationSet = wire.NewSet(
	service.UserServiceSet,
	service.UnitServiceSet,
	service.ConfigServiceSet,
)

var MapperSet = wire.NewSet(
	user.NewMongoMapper,
	unit.NewMongoMapper,
	config.NewMongoMapper,
)

var InfraSet = wire.NewSet(
	infraconfig.NewConfig,
	MapperSet,
)

var ServerProvider = wire.NewSet(
	ControllerSet,
	ApplicationSet,
	InfraSet,
)
