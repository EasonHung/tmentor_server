package mentor_redis

import (
	"mentor_app/user_system/initialize"

	"github.com/go-redis/redis"
)

var url = initialize.GLOBAL_CONFIG.Redis.Url

var Client = NewClient()

func NewClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     url + ":6379",
		Password: "123456",
		DB:       0,
	})

	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}

	return client
}
