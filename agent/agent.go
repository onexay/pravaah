package agent

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"pravaah/config"
	"pravaah/db"
	"pravaah/logger"
	"pravaah/messaging"

	"github.com/gorilla/websocket"
)

/*
 * Agent
 */
type Agent struct {
	config config.Agent    // Configuration
	db     db.Agent        // Database
	logger logger.Agent    // Log rotation
	wsConn *websocket.Conn // Websocket connection
}

// Agent instance
var me *Agent

func Agent_main(configFile string) {
	var err error = nil

	// Create agent instance
	me = new(Agent)

	// Agent mode of operation
	fmt.Printf("Starting agent ...\n")

	// Parse config file
	if err = me.config.Parse(configFile); err != nil {
		return
	}

	// Initialize logger
	if err = me.logger.Init(&me.config); err != nil {
		return
	}

	// Initialize DB
	if err = me.db.InitHandle(&me.config); err != nil {
		return
	}

	// Initialize STATE DB
	if err = me.db.InitState(); err != nil {
		return
	}

	// Remote endpoint details
	server_url := url.URL{
		Scheme: "ws",
		Host:   me.config.ServerEndpoint,
		Path:   "/connect",
	}

	// Initiate connection
	me.wsConn, _, err = websocket.DefaultDialer.Dial(server_url.String(), nil)
	if err != nil {
		log.Fatalf("Unable to connect to server [%s], error [%s]", server_url.String(), err.Error())
		return
	}

	defer me.wsConn.Close()

	// Send connection request to server
	me.ConnectReq()

	for {
		// Read messages from websocket
		_, msg, err := me.wsConn.ReadMessage()
		if err != nil {
			log.Printf("Error receiving message from server [%s], error [%s]", me.wsConn.RemoteAddr().String(), err.Error())
			break
		}

		// Unmarshall
		var rspMsg messaging.RspMsg = messaging.RspMsg{}
		if err := json.Unmarshal(msg, &rspMsg); err != nil {
			log.Printf("Unable to unmarshall message to JSON from server [%s], error [%s]\n", me.wsConn.RemoteAddr().String(), err.Error())
			break
		}

		// Check message status
		log.Printf("Status is %s\n", rspMsg.Status)

		// Dispatch message to handler
		if rspMsg.Type == messaging.MSG_CONNECT_RSP {
			me.ConnectRsp(&rspMsg)
		}
	}
}
