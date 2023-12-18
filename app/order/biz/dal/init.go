package dal

import (
	"github.com/baiyutang/gomall/app/order/biz/dal/mysql"
	"github.com/baiyutang/gomall/app/order/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
