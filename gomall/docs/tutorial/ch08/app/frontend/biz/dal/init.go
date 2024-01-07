package dal

import (
	"github.com/cloudwego/biz-demo/gomall/app/frontend/biz/dal/mysql"
	"github.com/cloudwego/biz-demo/gomall/app/frontend/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
