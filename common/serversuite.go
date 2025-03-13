package common

import (
	"net"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	consul "github.com/kitex-contrib/registry-consul"
)

type CommonServerSuite struct {
	CurrentServiceName string
	RegistryAddr       string
	ServiceAddress     string
}

func (s CommonServerSuite) Options() []server.Option {
	r, err := consul.NewConsulRegister(s.RegistryAddr)
	if err != nil {
		klog.Fatal(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", s.ServiceAddress)
	if err != nil {
		panic(err)
	}
	opts := []server.Option{
		server.WithServiceAddr(addr),
		server.WithMetaHandler(transmeta.ServerHTTP2Handler),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: s.CurrentServiceName,
		}),
		server.WithRegistry(r),
		server.WithSuite(tracing.NewServerSuite()),
	}
	klog.Infof("Suite: %s\n\tService Address: %s\n\tRegistry:%s", s.CurrentServiceName, s.ServiceAddress, s.RegistryAddr)

	return opts
}
