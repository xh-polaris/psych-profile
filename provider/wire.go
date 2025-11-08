//go:build wireinject
// +build wireinject

package provider

import (
	"github.com/google/wire"
	"github.com/xh-polaris/psych-profile/biz/adaptor"
)

func NewProvider() (*adaptor.Server, error) {
	wire.Build(
		wire.Struct(new(adaptor.Server), "*"),
		ServerProvider,
	)
	return nil, nil
}
