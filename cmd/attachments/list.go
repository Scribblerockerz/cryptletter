package attachments

import (
	"fmt"
	"github.com/Scribblerockerz/cryptletter/pkg/attachment"
	"github.com/Scribblerockerz/cryptletter/pkg/database"
	"github.com/go-redis/redis"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// NewListCmd builds a new init config command
func NewListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "attachments:list",
		Short: "List all known attachments",
		Long: "A list of all known attachment identifier. The list might diverge from the actual storage list.",
		Run:   runListCmd(),
	}

	return cmd
}

func runListCmd() func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {

		//// Connect to redis
		database.ConnectRedisClient(&redis.Options{
			Addr:     viper.GetString("redis.address"),
			Password: viper.GetString("redis.password"),
			DB:       viper.GetInt("redis.database"),
		})

		attachmentHandler := attachment.NewAttachmentHandler(viper.GetString("app.attachments.driver"))
		list, err := attachmentHandler.ListTimetable()
		if err != nil {
			panic(err)
		}

		for _, v := range list {
			fmt.Println(v)
		}
	}
}

