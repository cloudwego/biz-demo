package dal

import (
	"github.com/cloudwego/biz-demo/gomall/app/product/biz/dal/mysql"
	"github.com/cloudwego/biz-demo/gomall/app/product/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
