package mw

import (
	"github.com/hertz-contrib/cors"
	"time"
)

var CORSMiddleware = cors.New(cors.Config{
	AllowOrigins:     []string{"*"},
	AllowMethods:     []string{"*"},
	AllowHeaders:     []string{"*"},
	ExposeHeaders:    []string{"*"},
	AllowCredentials: true,

	MaxAge: 12 * time.Hour,
})
