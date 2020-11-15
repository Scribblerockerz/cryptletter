package cmd

import (
	"fmt"
	"github.com/Scribblerockerz/cryptletter/cmd/cryptletter"
	"github.com/Scribblerockerz/cryptletter/cmd/initConfig"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cryptletter",
	Short: "Encrypted self-destructing messages",
	Long: `Cryptletter is a tiny service to exchange information securely.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initializeConfiguration)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is cryptletter.yaml)")

	// Add commands
	rootCmd.AddCommand(cryptletter.NewCmd())
	//rootCmd.AddCommand(analyzeUrl.NewCmd())
	rootCmd.AddCommand(initConfig.NewCmd())
}

// initializeConfiguration reads in config file and ENV variables if set.
func initializeConfiguration() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigName("cryptletter")
		viper.SetConfigType("yaml")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
