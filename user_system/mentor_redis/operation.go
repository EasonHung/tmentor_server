package mentor_redis

import (
	"time"
)

// expireTime -> seconds
func SetWithTTL(key string, value string, expireTime int) {
	expire := time.Second * time.Duration(expireTime)
	err := Client.Set(key, value, expire).Err()
	if err != nil {
		panic(err)
	}
}

func Get(key string) string {
	val, err := Client.Get(key).Result()
	if val == "" {
		return "no value"
	}
	if err != nil {
		panic(err)
	}

	return val
}
