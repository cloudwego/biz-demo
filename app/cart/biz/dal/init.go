package dal

import (
	"github.com/baiyutang/gomall/app/cart/biz/dal/mysql"
	"github.com/baiyutang/gomall/app/cart/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
