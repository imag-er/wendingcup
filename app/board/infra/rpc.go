package infra

import (
	"github.com/cloudwego/kitex/client"

	"sync"

	"github.com/imag-er/wendingcup/app/board/conf"
	"github.com/imag-er/wendingcup/common"
	userservice "github.com/imag-er/wendingcup/rpc_gen/kitex_gen/user/User"
)

var (
	once        sync.Once
	UserClient  userservice.Client
	err         error
)

func Init() {
	once.Do(func() {
		initUserClient()
	})
}

func initUserClient() {
	UserClient, err = userservice.NewClient("user", client.WithSuite(&common.CommonClientSuite{
		CurrentServiceName: conf.GetConf().Kitex.Service,
		RegistryAddr:       conf.GetConf().Registry.RegistryAddress[0],
	}))
	if err != nil {
		panic(err)
	}
}
