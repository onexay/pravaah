package main

import (
	"flag"
	"log"
	"pravaah/config"
	"pravaah/version"
)

func main() {
	// Show build info
	version.ShowBuildVersion()

	// Parse command line options
	flag.StringVar(&config.RemoteEndPoint, "remote", "", "Remote endpoint address")
	flag.Parse()

	log.Printf("Option provided %s\n", config.RemoteEndPoint)
}
