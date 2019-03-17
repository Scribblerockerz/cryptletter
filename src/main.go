package main

import (
	"fmt"
	"net/http"

	"github.com/go-redis/redis"
)

func main() {
	AssembleConfiguration()
	RegisterPartials()
	ConnectRedisClient(&redis.Options{
		Addr:     Config.Database.Address,
		Password: Config.Database.Password,
		DB:       Config.Database.Database,
	})

	port := fmt.Sprintf("%d", Config.Server.Port)

	LogInfo(fmt.Sprintf("Serving cryptletter on http://localhost:%s\n", port))
	LogFatal(http.ListenAndServe(":"+port, NewRouter()))
}
