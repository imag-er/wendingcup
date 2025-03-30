package main

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/imag-er/wendingcup/app/api/conf"
	"github.com/imag-er/wendingcup/common"

	"github.com/imag-er/wendingcup/app/api/handler"
	"github.com/imag-er/wendingcup/app/api/infra"
	"github.com/imag-er/wendingcup/app/api/mw"
)

func main() {
	h := server.New(
		server.WithHostPorts(conf.GetConf().Hertz.Address),
		server.WithMaxRequestBodySize(10*1024*1024), // 100 MB
	)
	infra.Init()
	mw.InitJWT()
	h.Use(mw.CORSMiddleware)
	h.Use(mw.RequestLog)

	p := common.InitTracing(conf.GetConf().Hertz.Service)
	defer p.Shutdown(context.Background())

	h.GET("/board", handler.BoardHandler) // 获取排行榜

	h.GET("/ping", handler.PingHandler)                // ping
	h.POST("/login", mw.AuthMiddleware.LoginHandler)   // 登录
	h.POST("/register", handler.RegisterHandler)       // 注册
	h.POST("/logout", mw.AuthMiddleware.LogoutHandler) // 登出

	auth := h.Group("auth")
	auth.Use(mw.AuthMiddleware.MiddlewareFunc())
	{
		auth.GET("/team/:team_id", handler.TeamInfoHandler) // 获取队伍信息

		auth.GET("/submit/:team_id", handler.GetSubmitHandler) // 获取提交列表
		auth.POST("/submit", handler.PostSubmitHandler)        // 提交文件

		auth.POST("/refresh", mw.AuthMiddleware.RefreshHandler) // 刷新token
	}

	// 启动服务器，监听8080端口
	h.Spin()
}
