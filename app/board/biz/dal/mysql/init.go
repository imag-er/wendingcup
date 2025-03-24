package mysql

import (
	"fmt"

	"github.com/imag-er/wendingcup/app/board/biz/dal/model"
	"github.com/imag-er/wendingcup/app/board/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	DB, err = gorm.Open(mysql.Open(conf.GetConf().MySQL.DSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}


	// 自动迁移
	err = DB.AutoMigrate(&model.Result{})
	if err != nil {
		panic(err)
	}

	fmt.Println("mysql init success")

}
