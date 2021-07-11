package server

import (
	"encoding/json"
	"log"
	"net/http"
	"pravaah/messaging"
	"pravaah/version"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

/*
 * Handle connection request from an Agent
 */
func HandleAgentConnectReq(conn *websocket.Conn, reqMsg *messaging.ReqMsg) {
	// Unmarshall data
	var connectReqMsg messaging.ConnectReqMsg = messaging.ConnectReqMsg{}
	if err := json.Unmarshal([]byte(reqMsg.Data), &connectReqMsg); err != nil {
		log.Printf("Unable to unmarshall connect req to JSON from agent [%s], error [%s]\n",
			conn.RemoteAddr().String(),
			err.Error())
		return
	}

	/* This could either be a fresh agent or a returning agent. In case of a
	 * fresh agent, server will try to assign a UID and will register and
	 * persist agent info. Optionally for a returning agent, it will validate
	 * agent particulars.
	 */

	var uid string = connectReqMsg.ID

	//--- New Agent ---//

	if len(uid) == 0 {
		// This is a new agent, generate a new uid
		uid = uuid.NewString()

		log.Printf("Agent [%s] doesn't have an ID, generated id [%s]\n", conn.RemoteAddr().String(), uid)

		// Register agent
		RegisterAgent(uid, &connectReqMsg)

		// Prepare and marshal response data
		rspData, _ := json.Marshal(messaging.ConnectRspMsg{
			Version: version.GITInfo,
			ID:      uid,
		})

		// Send response and leave connection open
		conn.WriteJSON(messaging.RspMsg{
			Type:   messaging.MSG_CONNECT_RSP,
			Data:   string(rspData),
			Status: messaging.MSG_STATUS_OK,
		})

		return
	}

	//--- Returning Agent ---//

	log.Printf("Agent [%s] has an ID [%s]\n", conn.RemoteAddr().String(), connectReqMsg.ID)

	// Retrieve agent info
	_, err := RetrieveAgent(connectReqMsg.ID)
	if err != nil {
		log.Printf("Unable to find info for agent [%s], error [%s]", connectReqMsg.ID, err)

		// Send response and close connection
		conn.WriteJSON(messaging.RspMsg{
			Type:   messaging.MSG_CONNECT_RSP,
			Status: messaging.MSG_STATUS_ERROR,
		})
		conn.Close()

		return
	}

	// Prepare and marshal response data
	rspData, _ := json.Marshal(messaging.ConnectRspMsg{
		Version: version.GITInfo,
		ID:      uid,
	})

	// Send respond and leave connection open
	conn.WriteJSON(messaging.RspMsg{
		Type:   messaging.MSG_CONNECT_RSP,
		Data:   string(rspData),
		Status: messaging.MSG_STATUS_OK,
	})
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
			HandleAgentConnectReq(conn, &reqMsg)
		} else if reqMsg.Type == messaging.MSG_CAPABILITY {

		}
	}
}
