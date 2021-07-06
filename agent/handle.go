package agent

import (
	"encoding/json"
	"log"
	"pravaah/db"
	"pravaah/messaging"
	"pravaah/version"
)

func (me *Agent) ConnectReq() {
	// Fetch agent ID from db
	id, err := me.db.GetAgentID()
	if err != nil {
		log.Printf("Unable to fetch value of [%s] from DB\n", db.AGENT_ID)
		return
	}

	connectReqMsg, err := json.Marshal(messaging.ConnectReqMsg{
		Version: version.GITInfo,
		Secret:  me.config.ServerSecret,
		ID:      id,
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
	me.wsConn.WriteJSON(reqMsg)
}

func (me *Agent) ConnectRsp(rsp *messaging.RspMsg) {
	// Unmarshall data
	var connectRspMsg messaging.ConnectRspMsg = messaging.ConnectRspMsg{}
	if err := json.Unmarshal([]byte(rsp.Data), &connectRspMsg); err != nil {
		log.Printf("Unable to unmarshall connect rsp to JSON from server, error [%s]\n", err.Error())
		return
	}

	// Fetch agent ID from db
	id, err := me.db.GetAgentID()
	if err != nil {
		log.Printf("Unable to fetch value of [%s] from DB\n", db.AGENT_ID)
		return
	}

	if id == "" {
		log.Printf("No existing agent ID found, using server assigned ID [%s]\n", connectRspMsg.ID)

		// Save agent ID received from server
		me.db.SetAgentId(connectRspMsg.ID)
	} else {
		log.Printf("Existing agent ID found [%s]\n", id)

		if string(id) != connectRspMsg.ID {
			log.Printf("Local agent ID [%s] doesn't match server agent ID [%s]\n", id, connectRspMsg.ID)
		}
	}

	log.Printf("Agent version [%s], server version [%s]\n", version.GITInfo, connectRspMsg.Version)
}

func (me *Agent) AddSourceRsp() {

}

func (me *Agent) AddSourceReq(req *messaging.ReqMsg) {
	// Unmarshall data
	var addSourceReq messaging.AddSourceReqMsg = messaging.AddSourceReqMsg{}
	if err := json.Unmarshal([]byte(req.Data), &addSourceReq); err != nil {
		return
	}
}
