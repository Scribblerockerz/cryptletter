package cryptletter

import (
	"fmt"
	"github.com/Scribblerockerz/cryptletter/pkg/database"
	"github.com/Scribblerockerz/cryptletter/pkg/logger"
	"github.com/Scribblerockerz/cryptletter/pkg/router"
	"github.com/go-redis/redis"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"net/http"
)

// NewCmd builds a new analyse url command
func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "serve",
		Short: "Start the server",
		Run:   runCmd(),
	}

	return cmd
}

func runCmd() func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {

		// Connect to redis
		database.ConnectRedisClient(&redis.Options{
			Addr:     viper.GetString("redis.address"),
			Password: viper.GetString("redis.password"),
			DB:       viper.GetInt("redis.database"),
		})

		port := viper.GetString("app.server.port")
		env := viper.GetString("app.env")

		logger.LogInfo(fmt.Sprintf("Serving cryptletter on http://localhost:%s in %s environment\n", port, env))
		logger.LogFatal(http.ListenAndServe(":"+port, router.NewRouter()))
	}
}
