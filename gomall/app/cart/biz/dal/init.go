package dal

import (
	"github.com/baiyutang/gomall/app/cart/biz/dal/mysql"
)

func Init() {
	//redis.Init()
	mysql.Init()
}
