/* MIT License
 *
 * Copyright (c) 2021 Akshay Ranjan
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */
package server

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"pravaah/api"
	"pravaah/messaging"

	"github.com/ledisdb/ledisdb/ledis"
)

/*
 * List current server secret(s)
 */
func ListSecrets(res http.ResponseWriter, req *http.Request) {
	log.Printf("Received [%s] [%s] from [%s]\n", req.Method, req.URL, req.RemoteAddr)
}

/*
 * Renew server secret(s)
 */
func RenewSecrets(res http.ResponseWriter, req *http.Request) {
	log.Printf("Received [%s] [%s] from [%s]\n", req.Method, req.URL, req.RemoteAddr)
}

/*
 * List all agents
 */
func HandleAgentsList(res http.ResponseWriter, req *http.Request) {
	log.Printf("Received [%s] [%s] from [%s]\n", req.Method, req.URL, req.RemoteAddr)

	// Add usual headers
	res.Header().Add("Content-Type", "application/json")

	// Get DB
	store := me.db.GetAgentsDB()

	// Get all agent UIDs
	keys, _ := store.Scan(ledis.KV, nil, 0, false, "")

	// Agent info array
	var agents []api.Agent

	// Fetch agent info
	for _, key := range keys {
		// Get data
		bytes, _ := store.Get(key)

		var agent api.Agent

		// Unmarshal
		json.Unmarshal(bytes, &agent)

		// Unmarshal and append
		agents = append(agents, agent)
	}

	// Marshal response
	bytes, _ := json.Marshal(api.AgentResponseList{
		Status: messaging.MsgStatusStr[messaging.MSG_STATUS_OK],
		Data:   agents,
	})

	res.Write(bytes)
}

func HandleAgentActivate(res http.ResponseWriter, req *http.Request) {
	log.Printf("Received [%s] [%s] from [%s]\n", req.Method, req.URL, req.RemoteAddr)

}

func HandleAgentDeactivate(res http.ResponseWriter, req *http.Request) {
	log.Printf("Received [%s] [%s] from [%s]\n", req.Method, req.URL, req.RemoteAddr)

}

/*
 *
 */
func HandleAgentDelete(res http.ResponseWriter, req *http.Request) {
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

	// Marshal response
	bytes, _ = json.Marshal(api.AgentDelRsp{
		Count:  len(agents),
		Status: messaging.MsgStatusStr[messaging.MSG_STATUS_OK],
	})

	res.Write(bytes)
}

/*
 *
 */
func HandleSourcesList(res http.ResponseWriter, req *http.Request) {
	log.Printf("Received [%s] [%s] from [%s]\n", req.Method, req.URL, req.RemoteAddr)

	// Add usual headers
	res.Header().Add("Content-Type", "application/json")

	// Get DB
	store := me.db.GetSourceDB()

	// Get all source keys
	keys, _ := store.Scan(ledis.KV, nil, 0, false, "")

	// Response data
	var sources []api.SourceExt

	// Iterate through keys
	for _, key := range keys {
		// Get source info
		bytes, _ := store.Get(key)

		// Source
		var source api.Source

		// Unmarshall
		json.Unmarshal(bytes, &source)

		// Append to response
		sources = append(sources, api.SourceExt{
			Source: source,
			ID:     string(key),
		})
	}

	// Marshal response
	bytes, _ := json.Marshal(api.SourceResponseList{
		Status: messaging.MsgStatusStr[messaging.MSG_STATUS_OK],
		Data:   sources,
	})

	res.Write(bytes)
}

/*
 *
 */
func HandleSourceAdd(res http.ResponseWriter, req *http.Request) {
	var sources []api.Source

	log.Printf("Received [%s] [%s] from [%s]\n", req.Method, req.URL, req.RemoteAddr)

	// Add usual headers
	res.Header().Add("Content-Type", "application/json")

	// Read body
	bytes, _ := io.ReadAll(req.Body)

	// Unmarshal body
	json.Unmarshal(bytes, &sources)

	// Source IDs
	var sourceIDs []api.SourceID

	// Iterate through sources
	for _, source := range sources {
		// Add source
		id, exists, err := SourceAdd(&source)
		if err != nil {
			sourceIDs = append(sourceIDs, api.SourceID{
				Alias:  source.Alias,
				Status: messaging.MsgStatusStr[messaging.MSG_STATUS_ERROR],
			})
		} else {
			sourceIDs = append(sourceIDs, api.SourceID{
				Alias:  source.Alias,
				ID:     id,
				New:    !exists,
				Status: messaging.MsgStatusStr[messaging.MSG_STATUS_OK],
			})
		}
	}

	// Marshal response
	bytes, _ = json.Marshal(api.SourceResponseAdd{
		Status: messaging.MsgStatusStr[messaging.MSG_STATUS_OK],
		Data:   sourceIDs,
	})

	res.Write(bytes)
}

func HandleSourceDelete(res http.ResponseWriter, req *http.Request) {
	log.Printf("Received [%s] [%s] from [%s]\n", req.Method, req.URL, req.RemoteAddr)

}

func HandleSourceStart(res http.ResponseWriter, req *http.Request) {
	log.Printf("Received [%s] [%s] from [%s]\n", req.Method, req.URL, req.RemoteAddr)

}

func HandleSourceStop(res http.ResponseWriter, req *http.Request) {
	log.Printf("Received [%s] [%s] from [%s]\n", req.Method, req.URL, req.RemoteAddr)

}
