package main

import (
	"github.com/BurntSushi/toml"

	"github.com/jessevdk/go-flags"
)

// Configuration which is retrieved from a toml file
type Configuration struct {
	Server   server
	Database database
}

type server struct {
	Port int16
}

type database struct {
	Host     string
	User     string
	Password string
	Database string
}

type app struct {
	TemplatesDir string
	AssetsDir    string
}

// CliOptions of the possible cli arguments
type CliOptions struct {
	ConfigFile string `short:"f" long:"file" description:"Configuration file"`
}

// ParseArguments will pars cli arguments
func ParseArguments() CliOptions {
	var opts CliOptions

	_, err := flags.Parse(&opts)
	if err != nil {
		panic(err)
	}

	return opts
}

// AssembleConfiguration will collect and merge configuration from different sources
func AssembleConfiguration() Configuration {
	opts := ParseArguments()
	var config Configuration

	if opts.ConfigFile != "" {
		_, err := toml.DecodeFile(opts.ConfigFile, &config)
		if err != nil {
			panic(err)
		}
	}

	return config
}
