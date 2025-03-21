package dal

import (
	"github.com/imag-er/wendingcup/app/submit/biz/dal/mysql"
	"github.com/imag-er/wendingcup/app/submit/biz/dal/redis"
	"github.com/imag-er/wendingcup/app/submit/biz/dal/file"
)

func Init() {
	redis.Init()
	mysql.Init()
	file.InitFileManager()
}
