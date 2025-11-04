package service

import (
	"github.com/google/wire"
	"github.com/xh-polaris/psych-profile/biz/infra/mapper/unit"
)

var _ IUnitService = (*UnitService)(nil)

type IUnitService interface {
}

type UnitService struct {
	UnitMapper unit.IMongoMapper
}

var UnitServiceSet = wire.NewSet(
	wire.Struct(new(UnitService), "*"),
	wire.Bind(new(IConfigService), new(*UnitService)),
)
