package config

import (
	"fmt"
	"github.com/Scribblerockerz/cryptletter/pkg/logger"
	"github.com/spf13/viper"
	"strings"
)

//InitConfig reads in config file and ENV variables if set.
func InitConfig(file string) {
	initDefaults()

	if file != "" {
		viper.SetConfigFile(file)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigName("cryptletter")
		viper.SetConfigType("yaml")
	}

	viper.AutomaticEnv()                                    // read in environment variables that match
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "__")) // replaces APP__ENV to app.env

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		logger.LogInfo(fmt.Sprintf("Using config file: %s", viper.ConfigFileUsed()))
		viper.Set("viper.config_file", viper.ConfigFileUsed())
	}
}

//initDefaults will setup all configuration fallbacks. It is used for generating/initializing new config as well
func initDefaults() {
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
	viper.SetDefault("app.attachments.cleanup_schedule", "* * * * *")

	viper.SetDefault("s3.endpoint", "http://127.0.0.1:9000")
	viper.SetDefault("s3.secure", true)
	viper.SetDefault("s3.access_id", "minioadmin")
	viper.SetDefault("s3.access_secret", "minioadmin")
	viper.SetDefault("s3.bucket_name", "cryptletter-attachments")
	viper.SetDefault("s3.bucket_region", "eu-central-1")
}
