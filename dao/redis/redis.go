package redis

import (
	"autoplay-hub/settings"
	"fmt"

	"github.com/go-redis/redis"
)

// client redis操作实例，小写保证只在redis层操作redis
var client *redis.Client

func Init(rc *settings.RedisConfig) (err error) {
	client = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			rc.Host,
			rc.Port,
		),
		Password: rc.Password,
		DB:       rc.DB,
		PoolSize: rc.PoolSize,
	})
	_, err = client.Ping().Result()
	return
}

func Close() {
	_ = client.Close()
}
