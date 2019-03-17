package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-redis/redis"
)

func main() {
	AssembleConfiguration()
	RegisterPartials()

	ConnectRedisClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	// TODO: intergrate redis server
	// - connect to server
	// - set keys with default ttl of 30 days
	// - get keys and match them against users ip
	// 	- if no ip is set
	//		- update ttl to requested value

	router := NewRouter()
	port := fmt.Sprintf("%d", Config.Server.Port)

	fmt.Printf("Serving taskronaut api on http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

// Message asd
type Message struct {
	Content      string
	Token        string
	CreatedAt    time.Time
	Lifetime     int64
	AccessableIP string
}

func testRedisConnection() {
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
}

/*

OLD SQL

CREATE TABLE IF NOT EXISTS `messages` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `text` longtext COLLATE utf8_unicode_ci NOT NULL,
  `token` longtext COLLATE utf8_unicode_ci NOT NULL,
  `created_at` datetime NOT NULL,
  `active_until` datetime DEFAULT NULL,
  `mode` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `mode_value` int(11) NOT NULL,
  `accessable_ip` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=88 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

*/
