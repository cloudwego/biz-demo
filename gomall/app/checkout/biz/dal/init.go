package dal

import (
	"github.com/cloudwego/biz-demo/gomall/app/checkout/biz/dal/mysql"
	"github.com/cloudwego/biz-demo/gomall/app/checkout/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
