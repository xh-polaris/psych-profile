package controller

import (
	"context"
	"github.com/xh-polaris/gopkg/util"
	"github.com/xh-polaris/psych-profile/pkg/errorx"
	"go.opentelemetry.io/otel/trace"

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
	resp, err = c.ConfigService.ConfigCreate(ctx, req)
	logs.CtxInfof(ctx, "req=%s, resp=%s, err=%s, trace=%s", util.JSONF(req), util.JSONF(resp), errorx.ErrorWithoutStack(err), trace.SpanContextFromContext(ctx).TraceID().String())
	return resp, err
}

func (c *ConfigController) ConfigUpdateInfo(ctx context.Context, req *profile.ConfigCreateOrUpdateReq) (resp *basic.Response, err error) {
	resp, err = c.ConfigService.ConfigUpdateInfo(ctx, req)
	logs.CtxInfof(ctx, "req=%s, resp=%s, err=%s, trace=%s", util.JSONF(req), util.JSONF(resp), errorx.ErrorWithoutStack(err), trace.SpanContextFromContext(ctx).TraceID().String())
	return resp, err
}

func (c *ConfigController) ConfigGetByUnitID(ctx context.Context, req *profile.ConfigGetByUnitIdReq) (resp *profile.ConfigGetByUnitIdResp, err error) {
	resp, err = c.ConfigService.ConfigGetByUnitID(ctx, req)
	logs.CtxInfof(ctx, "req=%s, resp=%s, err=%s, trace=%s", util.JSONF(req), util.JSONF(resp), errorx.ErrorWithoutStack(err), trace.SpanContextFromContext(ctx).TraceID().String())
	return resp, err
}
