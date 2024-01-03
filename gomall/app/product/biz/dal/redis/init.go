package redis

import (
	"context"

	"github.com/cloudwego/biz-demo/gomall/app/product/infra/mtl"

	"github.com/cloudwego/biz-demo/gomall/app/product/conf"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/redis/go-redis/extra/redisotel/v9"
	"github.com/redis/go-redis/extra/redisprometheus/v9"
	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func Init() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     conf.GetConf().Redis.Address,
		Username: conf.GetConf().Redis.Username,
		Password: conf.GetConf().Redis.Password,
		DB:       conf.GetConf().Redis.DB,
	})
	if err := RedisClient.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}
	if err := redisotel.InstrumentTracing(RedisClient); err != nil {
		klog.Error("redis tracing collect error ", err)
	}
	if err := mtl.Registry.Register(redisprometheus.NewCollector("default", "product", RedisClient)); err != nil {
		klog.Error("redis metric collect error ", err)
	}
	redisotel.InstrumentTracing(RedisClient)
}
