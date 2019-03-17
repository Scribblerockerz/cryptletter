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

/*

client := redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       0,
})
pong, err := client.Ping().Result()
fmt.Println(pong, err)

msg := Message{
	Content:   "Hello World",
	Token:     "123456",
	CreatedAt: time.Now(),
}

bytes, err := json.Marshal(&msg)
if err != nil {
	panic(err)
}

statement := client.Set("test", string(bytes), 0)
fmt.Println(statement)

val, err := client.Get("test").Result()
if err != nil {
	panic(err)
}

loadedMessage := &Message{}
err = json.Unmarshal([]byte(val), loadedMessage)
if err != nil {
	panic(err)
}

fmt.Println("test", loadedMessage.Content)

*/
