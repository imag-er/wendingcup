package dal

import (
	"github.com/imag-er/wendingcup/app/gateway/biz/dal/redis"
)

func Init() {
	redis.Init()
}
