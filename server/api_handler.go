package server

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"pravaah/api"
	"pravaah/messaging"

	"github.com/google/uuid"
	"github.com/ledisdb/ledisdb/ledis"
)

/*
 * List all agents
 */
func ListAgents(res http.ResponseWriter, req *http.Request) {
	log.Printf("Received [%s] [%s] from [%s]\n", req.Method, req.URL, req.RemoteAddr)

	// Add usual headers
	res.Header().Add("Content-Type", "application/json")

	// Get DB
	store := me.db.GetAgentsDB()

	// Get all agent UIDs
	agentUIDs, _ := store.Scan(ledis.KV, nil, 0, false, "")

	// Agent info array
	var agents []api.AgentAdd

	// Fetch agent info
	for _, uid := range agentUIDs {
		var agent api.AgentAdd = api.AgentAdd{}

		// Get data
		bytes, _ := store.Get(uid)

		// Unmarshal
		json.Unmarshal(bytes, &agent)

		// Unmarshal and append
		agents = append(agents, agent)
	}

	// Marshal back
	bytes, _ := json.Marshal(agents)

	res.Write(bytes)
}

func ActivateAgent(res http.ResponseWriter, req *http.Request) {
	log.Printf("Received [%s] [%s] from [%s]\n", req.Method, req.URL, req.RemoteAddr)

}

func DeactivateAgent(res http.ResponseWriter, req *http.Request) {
	log.Printf("Received [%s] [%s] from [%s]\n", req.Method, req.URL, req.RemoteAddr)

}

/*
 *
 */
func RemoveAgent(res http.ResponseWriter, req *http.Request) {
	var agents []api.AgentDel

	log.Printf("Received [%s] [%s] from [%s]\n", req.Method, req.URL, req.RemoteAddr)

	// Add usual headers
	res.Header().Add("Content-Type", "application/json")

	// Read body
	bytes, _ := io.ReadAll(req.Body)

	// Unmarshal body
	json.Unmarshal(bytes, &agents)

	// Get DB
	store := me.db.GetAgentsDB()

	// Iterate and delete
	for _, agent := range agents {
		// Delete entry
		store.Del([]byte(agent.ID))

		log.Printf("Deleting agent [%s] alias [%s]\n", agent.ID, agent.Alias)
	}

	// Respond back
	rsp := api.AgentDelRsp{
		Count:  len(agents),
		Status: messaging.MsgStatusStr[messaging.MSG_STATUS_OK],
	}

	// Marshal response
	bytes, _ = json.Marshal(rsp)

	res.Write(bytes)
}

func ListSources(res http.ResponseWriter, req *http.Request) {
	log.Printf("Received [%s] [%s] from [%s]\n", req.Method, req.URL, req.RemoteAddr)

	str := "OKAY"

	res.Write([]byte(str))
}

func AddSource(res http.ResponseWriter, req *http.Request) {
	// Add usual headers
	res.Header().Add("Content-Type", "application/json")

	log.Printf("Received [%s] [%s] from [%s]\n", req.Method, req.URL, req.RemoteAddr)

	// Read body
	body, _ := io.ReadAll(req.Body)

	// Unmarshall body
	var addReq api.AddReq = api.AddReq{}
	if err := json.Unmarshal(body, &addReq); err != nil {
		log.Printf("Unable to parse [%s] body from [%s]", req.URL, req.RemoteAddr)
		return
	}

	// Send a message to agent

	rspData, _ := json.Marshal(api.AddRsp{
		UUID: uuid.New().String(),
	})

	res.Write(rspData)
}

func RemoveSource(res http.ResponseWriter, req *http.Request) {
	log.Printf("Received [%s] [%s] from [%s]\n", req.Method, req.URL, req.RemoteAddr)

}

func StartSource(res http.ResponseWriter, req *http.Request) {
	log.Printf("Received [%s] [%s] from [%s]\n", req.Method, req.URL, req.RemoteAddr)

}

func StopSource(res http.ResponseWriter, req *http.Request) {
	log.Printf("Received [%s] [%s] from [%s]\n", req.Method, req.URL, req.RemoteAddr)

}
