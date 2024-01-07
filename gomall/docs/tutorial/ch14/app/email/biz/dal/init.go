package dal

import (
	"github.com/cloudwego/biz-demo/gomall/app/email/biz/dal/mysql"
	"github.com/cloudwego/biz-demo/gomall/app/email/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
