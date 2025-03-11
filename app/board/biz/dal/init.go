package dal

import (
	"github.com/imag-er/wendingcup/app/board/biz/dal/mysql"
	"github.com/imag-er/wendingcup/app/board/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
