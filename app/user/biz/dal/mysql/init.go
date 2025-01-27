package mysql

import (
	"fmt"
	"os"

	"github.com/bz-2021/gomall/app/user/biz/model"
	"github.com/bz-2021/gomall/app/user/conf"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	dsn := fmt.Sprintf(conf.GetConf().MySQL.DSN,
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_DATABASE"))
	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}
	fmt.Printf("MYSQL VERSION: %#v\n", DB.Debug().Exec("select version()"))
	
	DB.AutoMigrate(&model.User{})
}