package main

import (
	"github.com/BurntSushi/toml"

	"github.com/jessevdk/go-flags"
)

// Configuration is the type of the configuration structure
type Configuration struct {
	Server   server
	Database database
	App      app
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

// Config will contain the assembled configuration from file and acli arguments
var Config Configuration

// DefaultConfiguration will contain just the defaults
var DefaultConfiguration = NewConfiguration()

// NewConfiguration will generate a fresh configuration with defaults
func NewConfiguration() Configuration {
	cfg := Configuration{}

	cfg.Server.Port = 8080
	cfg.App.TemplatesDir = "./src/templates"

	return cfg
}

// AssembleConfiguration will collect and merge configuration from different sources
func AssembleConfiguration() Configuration {
	Config = NewConfiguration()
	opts := ParseArguments()

	if opts.ConfigFile != "" {
		_, err := toml.DecodeFile(opts.ConfigFile, &Config)
		if err != nil {
			panic(err)
		}
	}

	return Config
}
