package psych_profile

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	"github.com/xh-polaris/gopkg/kitex/middleware"
	"github.com/xh-polaris/psych-idl/kitex_gen/user/psychuserservice"
	"github.com/xh-polaris/psych-profile/biz/infra/config"
	"github.com/xh-polaris/psych-profile/pkg/logs"
	"github.com/xh-polaris/psych-profile/provider"
	"net"
)

func main() {
	klog.SetLogger(logs.NewKlogLogger())
	s, err := provider.NewProvider()
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", config.GetConfig().ListenOn)
	if err != nil {
		panic(err)
	}
	svr := psychuserservice.NewServer(
		s,
		server.WithServiceAddr(addr),
		server.WithSuite(tracing.NewServerSuite()),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: config.GetConfig().Name}),
		server.WithMiddleware(middleware.LogMiddleware(config.GetConfig().Name)),
	)

	err = svr.Run()

	if err != nil {
		logs.Error(err.Error())
	}
}
