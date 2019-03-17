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

	router := NewRouter()
	port := fmt.Sprintf("%d", Config.Server.Port)

	fmt.Printf("Serving taskronaut api on http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
