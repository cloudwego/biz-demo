package dal

import (
	"github.com/baiyutang/gomall/app/payment/biz/dal/mysql"
	"github.com/baiyutang/gomall/app/payment/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
