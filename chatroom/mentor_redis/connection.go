package mentor_redis

import (
	"mentor_app/chatroom/config"

	"github.com/go-redis/redis"
)

var url = config.GLOBAL_CONFIG.Redis.Url

var Client = NewClient()

func NewClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     url + ":6379",
		Password: config.GLOBAL_CONFIG.Redis.Password,
		DB:       0,
	})

	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}

	return client
}
