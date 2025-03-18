// 以下是一个简单的Hertz框架的示例代码，用于创建一个HTTP服务器并处理简单的路由。
package main

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/imag-er/wendingcup/app/api/conf"
	"github.com/imag-er/wendingcup/common"

	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/imag-er/wendingcup/app/api/handler"
	"github.com/imag-er/wendingcup/app/api/infra"
	"github.com/imag-er/wendingcup/app/api/mw"
)

func main() {
	h := server.New(
		server.WithHostPorts(conf.GetConf().Hertz.Address),
	)
	infra.Init()
	mw.InitJWT()
	h.Use(mw.RequestLog)
	h.Use(mw.CORSMiddleware)

	p := common.InitTracing(conf.GetConf().Hertz.Service)
	defer p.Shutdown(context.Background())

	h.GET("/", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(consts.StatusOK, utils.H{"message": "Hello, World!"})
	})

	h.GET("/board", handler.BoardHandler)

	h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(consts.StatusOK, utils.H{"message": "pong"})
	})
	h.POST("/login", mw.AuthMiddleware.LoginHandler)
	h.POST("/register", handler.RegisterHandler)
	h.POST("/logout", mw.AuthMiddleware.LogoutHandler)

	auth := h.Group("auth")
	auth.Use(mw.AuthMiddleware.MiddlewareFunc())
	{
		auth.GET("/teaminfo", handler.TeamInfoHandler)
		auth.POST("/submit", mw.AuthMiddleware.MiddlewareFunc(), handler.SubmitHandler)
		auth.POST("/refresh", mw.AuthMiddleware.RefreshHandler)
	}

	// 启动服务器，监听8080端口
	h.Spin()
}
