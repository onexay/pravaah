package main

import (
	"flag"
	"pravaah/agent"
	"pravaah/config"
	"pravaah/server"
	"pravaah/version"
)

func main() {
	// Show build info
	version.ShowBuildVersion()

	// Parse command line options
	flag.BoolVar(&config.ServerMode, "server", false, "Server mode")
	flag.StringVar(&config.ConfigFile, "config", "", "Config file location")
	flag.Parse()

	if config.ServerMode {
		server.Server_main(config.ConfigFile)
	} else {
		agent.Agent_main(config.ConfigFile)
	}
}
