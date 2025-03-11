package dal

import (
	"github.com/imag-er/wendingcup/app/judge/biz/dal/mysql"
	"github.com/imag-er/wendingcup/app/judge/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
