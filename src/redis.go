package main

import (
	"time"

	"github.com/go-redis/redis"
)

// RedisClient is a redis instance
var RedisClient *redis.Client

// ConnectRedisClient with a redis db
func ConnectRedisClient(o *redis.Options) {
	connectWithRedis(o, 3*time.Second, 10)
}

//time.Sleep(2 * time.Second)

func connectWithRedis(o *redis.Options, waitDuration time.Duration, maxRetries int64) {
	RedisClient = redis.NewClient(o)
	pong, err := RedisClient.Ping().Result()

	if err != nil && maxRetries > 0 {
		LogWarning("Connection to redis failed. Retry...")
		time.Sleep(waitDuration)
		connectWithRedis(o, waitDuration, maxRetries-1)
	} else if err != nil {
		panic(err)
	} else if pong == "PONG" {
		LogInfo("Successfuly established connection to redis")
	} else {
		LogWarning("Connection to redis failed. Pong not received.")
	}
}
