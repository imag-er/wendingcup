package dal

import (
	"github.com/imag-er/wendingcup/app/submit/biz/dal/mysql"
	"github.com/imag-er/wendingcup/app/submit/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
