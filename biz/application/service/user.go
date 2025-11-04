package service

import (
	"github.com/google/wire"
	"github.com/xh-polaris/psych-profile/biz/infra/mapper/user"
)

var _ IUserService = (*UserService)(nil)

type IUserService interface {
}

type UserService struct {
	UserMapper user.IMongoMapper
}

var UserServiceSet = wire.NewSet(
	wire.Struct(new(UserService), "*"),
	wire.Bind(new(IUserService), new(*UserService)),
)
