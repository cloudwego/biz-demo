package dal

import (
	"github.com/baiyutang/gomall/app/product/infra/mysql"
	"github.com/baiyutang/gomall/app/product/infra/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
