package infra

import (
	"github.com/cloudwego/kitex/client"
	
	"github.com/imag-er/wendingcup/app/api/conf"
	boardservice "github.com/imag-er/wendingcup/rpc_gen/kitex_gen/board/board"
	submitservice "github.com/imag-er/wendingcup/rpc_gen/kitex_gen/submit/submit"
	userservice "github.com/imag-er/wendingcup/rpc_gen/kitex_gen/user/user"
	"github.com/imag-er/wendingcup/common"
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
	SubmitClient, err = submitservice.NewClient("submit", client.WithSuite(&common.CommonClientSuite{
		CurrentServiceName: conf.GetConf().Hertz.Service,
		RegistryAddr:       conf.GetConf().Registry.RegistryAddress[0],
	}))
	if err != nil {
		panic(err)
	}
}

func initUserClient() {
	UserClient, err = userservice.NewClient("user", client.WithSuite(&common.CommonClientSuite{
		CurrentServiceName: conf.GetConf().Hertz.Service,
		RegistryAddr:       conf.GetConf().Registry.RegistryAddress[0],
	}))
	if err != nil {
		panic(err)
	}
}


func initBoardClient() {
	BoardClient, err = boardservice.NewClient("board", client.WithSuite(&common.CommonClientSuite{
		CurrentServiceName: conf.GetConf().Hertz.Service,
		RegistryAddr:       conf.GetConf().Registry.RegistryAddress[0],
	}))
	if err != nil {
		panic(err)
	}
}

// common.CommonClientSuite 通用的client配置
