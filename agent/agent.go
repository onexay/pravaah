package agent

import (
	"fmt"
	"log"
	"net/url"
	"pravaah/config"
	"pravaah/db"
	"sync"

	"github.com/gorilla/websocket"
	"gopkg.in/natefinch/lumberjack.v2"
)

type StateIntf struct {
	// Private
	logger lumberjack.Logger // Log station

	// Public
	wsConn *websocket.Conn // Websocket connection
}

var State StateIntf

func Agent_main(configFile string) {
	var err error = nil

	// Server mode of operation
	fmt.Printf("Starting agent ...\n")

	// Check for config file
	if configFile == "" {
		fmt.Printf("No config file specified. Exiting.\n")
		return
	}

	// Parse config file
	_ = config.ParseAgentFile(configFile)

	// Logger setup
	State.logger.Filename = config.Agent.LogFileLocation
	State.logger.MaxSize = 1
	log.SetPrefix("[PRAVAAH AGENT] ")
	log.SetOutput(&State.logger)

	// Initialize DB
	if err = db.Init(); err != nil {
		return
	}

	// Remote endpoint details
	server_url := url.URL{
		Scheme: "ws",
		Host:   config.Agent.ServerEndpoint,
		Path:   "/stream",
	}

	var wg sync.WaitGroup

	// Initiate connection
	State.wsConn, _, err = websocket.DefaultDialer.Dial(server_url.String(), nil)
	if err != nil {
		log.Fatalf("Unable to connect to server [%s], error [%s]", server_url.String(), err.Error())
		return
	}

	// Begin handshake
	BeginHandshake(State.wsConn)

	wg.Add(1)

	go func() {
		for {
			// Read messages from websocket
			_, _, err := State.wsConn.ReadMessage()
			if err != nil {
				log.Fatalf("Error receiving message on websocket, error [%s]", err.Error())
			}
		}
	}()

	wg.Wait()
}
