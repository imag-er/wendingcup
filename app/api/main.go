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
)

func main() {
	h := server.New(
		server.WithHostPorts(conf.GetConf().Hertz.Address),
	)
	infra.Init()
	infra.InitJWT()

	p := common.InitTracing(conf.GetConf().Hertz.Service)
	defer p.Shutdown(context.Background())

	h.GET("/", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(consts.StatusOK, utils.H{
			"message": "Hello, World!",
		})
	})
	auth := h.Group("auth")

	auth.POST("/login", infra.AuthMiddleware.LoginHandler)
	auth.POST("/register", handler.RegisterHandler)

	auth.Use(infra.AuthMiddleware.MiddlewareFunc())
	{
		auth.POST("/logout", infra.AuthMiddleware.LogoutHandler)
		auth.POST("/refresh", infra.AuthMiddleware.RefreshHandler)
	}
	h.POST("/submit", infra.AuthMiddleware.MiddlewareFunc(), handler.SubmitHandler)
	h.GET("/ping", infra.AuthMiddleware.MiddlewareFunc(), func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(consts.StatusOK, utils.H{
			"message": "pong",
		})
	})
	// 启动服务器，监听8080端口
	h.Spin()
}
