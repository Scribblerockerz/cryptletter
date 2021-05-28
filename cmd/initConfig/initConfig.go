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
