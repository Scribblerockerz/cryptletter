package attachments

import (
	"github.com/Scribblerockerz/cryptletter/pkg/attachment"
	"github.com/Scribblerockerz/cryptletter/pkg/database"
	"github.com/go-redis/redis"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// NewDropCmd builds a new init config command
func NewDropCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "attachments:drop",
		Short: "Remove all stored attachments",
		Run:   runDropCmd(),
	}

	return cmd
}

func runDropCmd() func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {

		//// Connect to redis
		database.ConnectRedisClient(&redis.Options{
			Addr:     viper.GetString("redis.address"),
			Password: viper.GetString("redis.password"),
			DB:       viper.GetInt("redis.database"),
		})

		attachmentHandler := attachment.NewAttachmentHandler(viper.GetString("app.attachments.driver"))
		err := attachmentHandler.DropAll()
		if err != nil {
			panic(err)
		}
	}
}

