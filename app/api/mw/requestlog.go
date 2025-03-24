package mw

import (
	"context"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"
)

func RequestLog(ctx context.Context, c *app.RequestContext) {
	// 打印请求的方法和 URL
	fmt.Printf("🛜  Received request: %s %s\n", c.Request.Method(), c.Request.URI())
	// 继续处理请求
	c.Next(ctx)
}
