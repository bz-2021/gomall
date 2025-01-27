package dal

import (
	"github.com/bz-2021/gomall/app/user/biz/dal/mysql"
	"github.com/bz-2021/gomall/app/user/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
