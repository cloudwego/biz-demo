package dal

import (
	"github.com/baiyutang/gomall/app/frontend/biz/dal/mysql"
)

func Init() {
	//redis.Init()
	mysql.Init()
}
