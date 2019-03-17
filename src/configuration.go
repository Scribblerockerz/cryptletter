package main

import (
	"os"

	"github.com/BurntSushi/toml"

	"github.com/jessevdk/go-flags"
)

// Configuration is the type of the configuration structure
type Configuration struct {
	Server   server
	Database database
	App      app
	Debug    debug
}

type server struct {
	Port int16
}

type database struct {
	Address  string
	Password string
	Database int
}

type app struct {
	TemplatesDir             string
	AssetsDir                string
	DefaultTTLForNewMessages int64
}

type debug struct {
	LogLevel int
}

// CliOptions of the possible cli arguments
type CliOptions struct {
	ConfigFile string `short:"f" long:"file" description:"Configuration file eg. config.toml"`
	Verbose    []bool `short:"v" long:"verbose" description:"Show verbose debug information"`
}

// ParseArguments will pars cli arguments
func ParseArguments() CliOptions {
	var opts CliOptions

	_, err := flags.Parse(&opts)
	if err != nil {
		os.Exit(0)
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

	cfg.Database.Address = "localhost:6379"
	cfg.Database.Password = ""
	cfg.Database.Database = 0
	cfg.Server.Port = 8080
	cfg.App.TemplatesDir = "./theme/templates"
	cfg.App.AssetsDir = "public"
	cfg.App.DefaultTTLForNewMessages = 43830

	cfg.Debug.LogLevel = 0

	return cfg
}

// AssembleConfiguration will collect and merge configuration from different sources
func AssembleConfiguration() Configuration {
	Config = NewConfiguration()
	opts := ParseArguments()

	if opts.ConfigFile != "" {
		_, err := toml.DecodeFile(opts.ConfigFile, &Config)
		if err != nil {
			LogFatal(err)
		}
	}

	Config.Debug.LogLevel = len(opts.Verbose)

	return Config
}
