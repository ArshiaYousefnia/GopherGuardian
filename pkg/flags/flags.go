package flags

import (
	"flag"
	"fmt"
)

type CLIConfig struct {
	ConfigPath string
	Verbose    bool
	Port       int
}

func ParseFlags() (*CLIConfig, error) {
	configPath := flag.String("config", "config.json", "path to config json file")
	verbose := flag.Bool("verbose", false, "verbose output")
	port := flag.Int("port", 8080, "port to listen on")

	flag.Parse()

	if *configPath == "" {
		return nil, fmt.Errorf("no config file specified")
	}

	return &CLIConfig{
		ConfigPath: *configPath,
		Verbose:    *verbose,
		Port:       *port,
	}, nil
}
