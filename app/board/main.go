package main

import (
	"time"
	"github.com/imag-er/wendingcup/common"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	"github.com/imag-er/wendingcup/app/board/conf"
	"github.com/imag-er/wendingcup/rpc_gen/kitex_gen/board/board"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"context"
)

func main() {
	p := common.InitTracing(conf.GetConf().Kitex.Service)
	defer p.Shutdown(context.Background())	

	opts := kitexInit()

	svr := board.NewServer(new(BoardImpl), opts...)

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
	asyncWriter := &zapcore.BufferedWriteSyncer{
		WS: zapcore.AddSync(&lumberjack.Logger{
			Filename:   conf.GetConf().Kitex.LogFileName,
			MaxSize:    conf.GetConf().Kitex.LogMaxSize,
			MaxBackups: conf.GetConf().Kitex.LogMaxBackups,
			MaxAge:     conf.GetConf().Kitex.LogMaxAge,
		}),
		FlushInterval: time.Second,
	}
	klog.SetOutput(asyncWriter)
	server.RegisterShutdownHook(func() {
		asyncWriter.Sync()
	})
	return
}
