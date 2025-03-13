package infra

import (
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/transport"
	"github.com/imag-er/wendingcup/app/api/conf"
	boardservice "github.com/imag-er/wendingcup/rpc_gen/kitex_gen/board/Board"
	submitservice "github.com/imag-er/wendingcup/rpc_gen/kitex_gen/submit/Submit"
	userservice "github.com/imag-er/wendingcup/rpc_gen/kitex_gen/user/User"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	consul "github.com/kitex-contrib/registry-consul"
	"sync"
)

var (
	once         sync.Once
	UserClient   userservice.Client
	SubmitClient submitservice.Client
	BoardClient  boardservice.Client
	err          error
)

func Init() {
	once.Do(func() {
		initBoardClient()
		initSubmitClient()
		initUserClient()
	})
}

func initSubmitClient() {
	SubmitClient, err = submitservice.NewClient("submit", client.WithSuite(&CommonClientSuite{
		CurrentServiceName: conf.GetConf().Hertz.Service,
		RegistryAddr:       conf.GetConf().Registry.RegistryAddress[0],
	}))
	if err != nil {
		panic(err)
	}
}

func initUserClient() {
	UserClient, err = userservice.NewClient("user", client.WithSuite(&CommonClientSuite{
		CurrentServiceName: conf.GetConf().Hertz.Service,
		RegistryAddr:       conf.GetConf().Registry.RegistryAddress[0],
	}))
	if err != nil {
		panic(err)
	}
}


func initBoardClient() {
	BoardClient, err = boardservice.NewClient("board", client.WithSuite(&CommonClientSuite{
		CurrentServiceName: conf.GetConf().Hertz.Service,
		RegistryAddr:       conf.GetConf().Registry.RegistryAddress[0],
	}))
	if err != nil {
		panic(err)
	}
}

// CommonClientSuite 通用的client配置

type CommonClientSuite struct {
	CurrentServiceName string
	RegistryAddr       string
}

func (s CommonClientSuite) Options() []client.Option {
	r, err := consul.NewConsulResolver(s.RegistryAddr)
	if err != nil {
		panic(err)
	}
	return []client.Option{
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: s.CurrentServiceName,
		}),
		client.WithMetaHandler(transmeta.ClientHTTP2Handler),
		client.WithTransportProtocol(transport.GRPC),
		client.WithResolver(r),
		client.WithSuite(tracing.NewClientSuite()),
	}
}
