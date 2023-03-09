package cache

import "github.com/go-redis/redis/v8"

var redisClient *redis.Client

func Init() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	// initVideoCache()
	initMessageCache()
	initFavoriteCache()
}
