package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

// RedisClient is a redis instance
var RedisClient *redis.Client

// ConnectRedisClient with a redis db
func ConnectRedisClient(o *redis.Options) {
	RedisClient = redis.NewClient(o)
	pong, err := RedisClient.Ping().Result()
	if err != nil {
		panic(err)
	} else if pong == "PONG" {
		fmt.Println("Successfuly established connection to redis")
	} else {
		fmt.Println("Connection to redis failed. Pong not received.")
	}
}
