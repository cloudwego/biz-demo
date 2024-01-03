package dal

import (
	"github.com/baiyutang/gomall/app/checkout/biz/dal/mysql"
	"github.com/baiyutang/gomall/app/checkout/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
