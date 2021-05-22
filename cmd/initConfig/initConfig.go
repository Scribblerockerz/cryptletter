package initConfig

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// NewCmd builds a new init config command
func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config:init",
		Short: "Generate a fresh config in the current directory",
		Run:   runCmd(),
	}

	return cmd
}

func runCmd() func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {

		// Defaults
		viper.SetDefault("redis.address", "redis:6379")
		viper.SetDefault("redis.password", "")
		viper.SetDefault("redis.database", 0)
		viper.SetDefault("app.default_message_ttl", 43830) // in minutes
		viper.SetDefault("app.log_level", 4) // enum
		viper.SetDefault("app.env", "prod")
		viper.SetDefault("app.server.port", 8080)
		viper.SetDefault("app.creation_protection_password", "")
		viper.SetDefault("app.attachments.driver", "local")
		viper.SetDefault("app.attachments.storage_path", "cryptletter-uploads")
		viper.SetDefault("app.attachments.cleanup_schedule", "local")

		viper.SetDefault("s3.endpoint", "http://127.0.0.1:9000")
		viper.SetDefault("s3.access_id", "minioadmin")
		viper.SetDefault("s3.access_secret", "minioadmin")


		err := viper.SafeWriteConfig()

		if _, ok := err.(viper.ConfigFileAlreadyExistsError); ok  {
			err = viper.WriteConfig()
			fmt.Println("Updating configuration file")
		} else {
			fmt.Println("Creating a new configuration")
		}

		if err != nil {
			panic(err)
		}
	}
}
