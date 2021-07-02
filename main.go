package main

import (
	"flag"
	"log"
	"pravaah/client"
	"pravaah/config"
	"pravaah/server"
	"pravaah/version"

	"gopkg.in/natefinch/lumberjack.v2"
)

// Logger
var logger lumberjack.Logger

func main() {
	// Setup logger
	logger.Filename = "./logs/pravaah.log"
	logger.MaxSize = 1
	log.SetOutput(&logger)

	// Show build info
	version.ShowBuildVersion()

	// Parse command line options
	flag.StringVar(&config.RemoteEndPoint, "remote", "", "Remote endpoint address")
	flag.BoolVar(&config.Server, "server", false, "Server mode")
	flag.StringVar(&config.LogFile, "logfile", "/var/log/pravaah", "Log file location")
	flag.Parse()

	if config.Server {
		server.Server_main()
	} else {
		client.Client_main()
	}
}
