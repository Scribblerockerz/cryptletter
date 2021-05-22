package utils

import (
	"github.com/spf13/viper"
	"os"
	"path"
)

//GetConfigRootDir will return the dir from the provided/loaded configuration
func GetConfigRootDir() string {
	configFile := viper.GetString("viper.config_file")
	if configFile == "" {
		cwd, _ := os.Getwd()
		return cwd
	}

	return path.Dir(configFile)
}
