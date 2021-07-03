package client

import (
	"fmt"
	"log"
	"net/url"
	"pravaah/config"

	"github.com/gorilla/websocket"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Logger
var logger lumberjack.Logger

func Client_main(configFile string) {
	// Server mode of operation
	fmt.Printf("Starting client ...\n")

	// Check for config file
	if configFile == "" {
		fmt.Printf("No config file specified. Exiting.\n")
		return
	}

	// Logger setup
	logger.Filename = config.Client.LogFileLocation
	logger.MaxSize = 1
	log.SetOutput(&logger)

	// Server connection details
	server_url := url.URL{
		Scheme: "ws",
		Host:   "localhost:10080",
		Path:   "/stream",
	}

	// Initiate connection
	conn, _, err := websocket.DefaultDialer.Dial(server_url.String(), nil)
	if err != nil {
		log.Fatalf("Unable to connect to server [%s], error [%s]", server_url.String(), err.Error())
		return
	}

	defer conn.Close()

	go func() {
		for {
			// Read messages from websocket
			_, _, err := conn.ReadMessage()
			if err != nil {
				log.Fatalf("Error receiving message on websocket, error [%s]", err.Error())
			}
		}
	}()
}
