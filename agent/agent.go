package agent

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"pravaah/config"
	"pravaah/db"
	"pravaah/messaging"

	"github.com/gorilla/websocket"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logger lumberjack.Logger

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
	logger.Filename = config.Agent.LogFileLocation
	logger.MaxSize = 1
	log.SetPrefix("[PRAVAAH AGENT] ")
	log.SetOutput(&logger)

	// Initialize DB
	if err = db.Init(); err != nil {
		return
	}

	// Remote endpoint details
	server_url := url.URL{
		Scheme: "ws",
		Host:   config.Agent.ServerEndpoint,
		Path:   "/connect",
	}

	// Initiate connection
	conn, _, err := websocket.DefaultDialer.Dial(server_url.String(), nil)
	if err != nil {
		log.Fatalf("Unable to connect to server [%s], error [%s]", server_url.String(), err.Error())
		return
	}

	defer conn.Close()

	// Begin connection
	BeginConnect(conn)

	for {
		// Read messages from websocket
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Error receiving message from server [%s], error [%s]", conn.RemoteAddr().String(), err.Error())
			break
		}

		// Unmarshall
		var rspMsg messaging.RspMsg = messaging.RspMsg{}
		if err := json.Unmarshal(msg, &rspMsg); err != nil {
			log.Printf("Unable to unmarshall message to JSON from server [%s], error [%s]\n", conn.RemoteAddr().String(), err.Error())
			break
		}

		// Check message status
		log.Printf("Status is %s\n", rspMsg.Status)

		// Dispatch message to handler
		if rspMsg.Type == messaging.MSG_CONNECT_RSP {
			HandleConnectionRsp(&rspMsg)
		}
	}
}
