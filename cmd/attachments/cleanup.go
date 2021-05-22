package attachments

import (
	"github.com/Scribblerockerz/cryptletter/pkg/attachment"
	"github.com/Scribblerockerz/cryptletter/pkg/database"
	"github.com/go-redis/redis"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// NewCleanupCmd builds a new cleanup command
func NewCleanupCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "attachments:cleanup",
		Short: "Trigger a cleanup of stored attachments",
		Run:   runCleanupCmd(),
	}

	return cmd
}

func runCleanupCmd() func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {

		//// Connect to redis
		database.ConnectRedisClient(&redis.Options{
			Addr:     viper.GetString("redis.address"),
			Password: viper.GetString("redis.password"),
			DB:       viper.GetInt("redis.database"),
		})

		attachmentHandler := attachment.NewAttachmentHandler(viper.GetString("app.attachments.driver"))
		err := attachmentHandler.Cleanup()
		if err != nil {
			panic(err)
		}
	}
}

