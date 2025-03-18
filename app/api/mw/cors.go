package mw

import (
	"github.com/hertz-contrib/cors"
	"time"
)

var CORSMiddleware = cors.New(cors.Config{
	AllowOrigins:     []string{"https://foo.com"},
	AllowMethods:     []string{"PUT", "PATCH"},
	AllowHeaders:     []string{"Origin"},
	ExposeHeaders:    []string{"Content-Length"},
	AllowCredentials: true,
	AllowOriginFunc: func(origin string) bool {
		return origin == "https://github.com"
	},
	MaxAge: 12 * time.Hour,
})
