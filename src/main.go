package main

import (
	"fmt"
	"log"
	"net/http"

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
