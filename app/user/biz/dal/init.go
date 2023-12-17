package dal

import (
	"github.com/baiyutang/gomall/app/user/biz/dal/mysql"
	"github.com/baiyutang/gomall/app/user/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
