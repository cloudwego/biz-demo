package dal

import (
	"github.com/baiyutang/gomall/app/product/biz/dal/mysql"
	"github.com/baiyutang/gomall/app/product/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
