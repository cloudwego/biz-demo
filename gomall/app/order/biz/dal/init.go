package dal

import (
	"github.com/cloudwego/biz-demo/gomall/app/order/biz/dal/mysql"
	"github.com/cloudwego/biz-demo/gomall/app/order/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
