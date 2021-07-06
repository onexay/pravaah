package server

import (
	"encoding/json"
	"log"
	"net/http"
	"pravaah/config"
	"pravaah/messaging"
	"pravaah/version"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

/* Handle connection request from an Agent
 *
 */
func HandleAgentConnectionReq(conn *websocket.Conn, reqMsg *messaging.ReqMsg) {
	// Unmarshall data
	var connectReqMsg messaging.ConnectReqMsg = messaging.ConnectReqMsg{}
	if err := json.Unmarshal([]byte(reqMsg.Data), &connectReqMsg); err != nil {
		log.Printf("Unable to unmarshall connect req to JSON from agent [%s], error [%s]\n", conn.RemoteAddr().String(), err.Error())
		return
	}

	// Check agent ID
	var id uuid.UUID
	if connectReqMsg.ID == "" {
		log.Printf("Agent [%s] doesn't have an ID, generating a new one\n", conn.RemoteAddr().String())
		id = uuid.New()
	} else {
		log.Printf("Agent [%s] has an ID [%s], reuseing same\n", conn.RemoteAddr().String(), connectReqMsg.ID)
		id, _ = uuid.Parse(connectReqMsg.ID)
	}

	// Check server secret
	if connectReqMsg.Secret != config.Secret {
		log.Printf("Agent [%s] secret [%s] doesn't match current server secret [%s]\n", id.String(), connectReqMsg.Secret, config.Secret)
	}

	// Create a response for this Agent
	connectRspMsg, err := json.Marshal(messaging.ConnectRspMsg{
		Version: version.GITInfo,
		ID:      id.String(),
	})

	if err != nil {
		log.Printf("Unable to encode connect request as JSON, error [%s]\n", err.Error())
		return
	}

	// Prepare message
	rspMsg := messaging.RspMsg{
		Type:   messaging.MSG_CONNECT_RSP,
		Data:   string(connectRspMsg),
		Status: "OK",
	}

	// Send message
	conn.WriteJSON(rspMsg)
}

/* Handle incoming requests from Agents
 *
 * This method handles incoming requests from Agents and saves the
 * context in a runtime database for easy indexing and persistence
 * throughout the server runtime.
 */
func HandleAgent(res http.ResponseWriter, req *http.Request) {
	var upgrader websocket.Upgrader = websocket.Upgrader{}

	log.Printf("Received websocket connect request from [%s]\n", req.RemoteAddr)

	// Upgrade request to websocket
	conn, err := upgrader.Upgrade(res, req, nil)
	if err != nil {
		log.Printf("Unable to upgrade request from agent [%s] to websocket\n", req.RemoteAddr)
		return
	}

	// Close websocket
	defer conn.Close()

	// Process messages from this Agent
	for {
		// Read messages from this agent
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Unable to receive message from agent [%s]\n", conn.RemoteAddr().String())
			break
		}

		// Unmarshall
		var reqMsg messaging.ReqMsg = messaging.ReqMsg{}
		if err := json.Unmarshal(msg, &reqMsg); err != nil {
			log.Printf("Unable to unmarshall message to JSON from agent [%s], error [%s]\n", conn.RemoteAddr().String(), err.Error())
			break
		}

		// Dispatch message to handler
		if reqMsg.Type == messaging.MSG_CONNECT_REQ {
			HandleAgentConnectionReq(conn, &reqMsg)
		} else if reqMsg.Type == messaging.MSG_CAPABILITY {

		}
	}
}
