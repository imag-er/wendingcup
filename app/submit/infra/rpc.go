package infra

import (
	"github.com/cloudwego/kitex/client"

	"sync"

	"github.com/imag-er/wendingcup/app/submit/conf"
	"github.com/imag-er/wendingcup/common"
	boardservice "github.com/imag-er/wendingcup/rpc_gen/kitex_gen/board/Board"
	userservice "github.com/imag-er/wendingcup/rpc_gen/kitex_gen/user/User"
)

var (
	once        sync.Once
	UserClient  userservice.Client
	BoardClient boardservice.Client
	err         error
)

func Init() {
	once.Do(func() {
		initBoardClient()
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

func initBoardClient() {
	BoardClient, err = boardservice.NewClient("board", client.WithSuite(&common.CommonClientSuite{
		CurrentServiceName: conf.GetConf().Kitex.Service,
		RegistryAddr:       conf.GetConf().Registry.RegistryAddress[0],
	}))
	if err != nil {
		panic(err)
	}
}
