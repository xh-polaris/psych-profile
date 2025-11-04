package adaptor

import "github.com/xh-polaris/psych-profile/biz/adaptor/controller"

type Server struct {
	controller.IUserController
	controller.IUnitController
	controller.IConfigController
}
