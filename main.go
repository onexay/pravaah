package main

import (
	"flag"
	"pravaah/client"
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
		client.Client_main(config.ConfigFile)
	}
}
