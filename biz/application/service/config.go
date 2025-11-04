package service

import (
	"github.com/google/wire"
	"github.com/xh-polaris/psych-profile/biz/infra/mapper/config"
)

var _ IConfigService = (*ConfigService)(nil)

type IConfigService interface {
}

type ConfigService struct {
	ConfigMapper config.IMongoMapper
}

var ConfigServiceSet = wire.NewSet(
	wire.Struct(new(ConfigService), "*"),
	wire.Bind(new(IConfigService), new(*ConfigService)),
)
