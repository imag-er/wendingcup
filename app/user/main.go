package main

import (
	"context"
	"os"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	"github.com/imag-er/wendingcup/app/user/biz/dal"
	"github.com/imag-er/wendingcup/app/user/conf"
	"github.com/imag-er/wendingcup/common"
	"github.com/imag-er/wendingcup/rpc_gen/kitex_gen/user/user"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
)

func main() {
	p := common.InitTracing(conf.GetConf().Kitex.Service)
	defer p.Shutdown(context.Background())
	opts := kitexInit()
	dal.Init()
	svr := user.NewServer(new(UserImpl), opts...)

	err := svr.Run()
	if err != nil {
		klog.Error(err.Error())
	}
}

func kitexInit() (opts []server.Option) {
	// address
	opts = append(opts, server.WithSuite(
		common.CommonServerSuite{
			CurrentServiceName: conf.GetConf().Kitex.Service,
			RegistryAddr:       conf.GetConf().Registry.RegistryAddress[0],
		},
	))

	// klog
	logger := kitexlogrus.NewLogger()
	klog.SetLogger(logger)
	klog.SetLevel(conf.LogLevel())
	klog.SetOutput(os.Stdout)

	return
}
