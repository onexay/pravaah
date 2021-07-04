package agent

import (
	"encoding/json"
	"log"
	"pravaah/messaging"
	"pravaah/version"

	"github.com/gorilla/websocket"
)

func BeginConnect(conn *websocket.Conn) {
	connectReqMsg, err := json.Marshal(messaging.ConnectReqMsg{
		Version: version.GITInfo,
		Secret:  "d0830b03171ba29e47e0148ab8629d04b48cf74cb78c3b8ad763d4ada41f6cc2",
	})

	if err != nil {
		log.Printf("Unable to encode connect request as JSON, error [%s]\n", err.Error())
		return
	}

	var reqMsg messaging.ReqMsg = messaging.ReqMsg{
		Type: messaging.MSG_CONNECT_REQ,
		Data: string(connectReqMsg),
	}

	// Send message
	conn.WriteJSON(reqMsg)
}

func HandleConnectionRsp(rspMsg *messaging.RspMsg) {
	// Unmarshall data
	var connectRspMsg messaging.ConnectRspMsg = messaging.ConnectRspMsg{}
	if err := json.Unmarshal([]byte(rspMsg.Data), &connectRspMsg); err != nil {
		log.Printf("Unable to unmarshall connect rsp to JSON from server, error [%s]\n", err.Error())
		return
	}

	log.Printf("Agent ID is now [%s]\n", connectRspMsg.ID)
	log.Printf("Agent version [%s], server version [%s]\n", version.GITInfo, connectRspMsg.Version)
}
