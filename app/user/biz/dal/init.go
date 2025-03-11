package dal

import (
	"github.com/imag-er/wendingcup/app/user/biz/dal/mysql"
	"github.com/imag-er/wendingcup/app/user/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
