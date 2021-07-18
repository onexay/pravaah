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

	// Response data
	var result []api.AgentListResultElem

	// Fetch agent info
	for _, key := range keys {
		// Get data
		bytes, _ := store.Get(key)

		var agent api.AgentListResultElem

		// Unmarshal
		json.Unmarshal(bytes, &agent)

		// Unmarshal and append
		result = append(result, agent)
	}

	// Marshal response
	bytes, _ := json.Marshal(api.AgentListResponse{
		Status: messaging.MsgStatusStr[messaging.MSG_STATUS_OK],
		Data:   result,
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
	var agents []api.AgentDeleteElem

	log.Printf("Received [%s] [%s] from [%s]\n", req.Method, req.URL, req.RemoteAddr)

	// Add usual headers
	res.Header().Add("Content-Type", "application/json")

	// Read body
	bytes, _ := io.ReadAll(req.Body)

	// Unmarshal body
	json.Unmarshal(bytes, &agents)

	// Response data
	var result []api.AgentDeleteResultElem
	var sources []api.SourceDeleteResultElem

	// Iterate through sources
	for _, agent := range agents {
		// Delete agent
		id, err := AgentDelete(&agent)

		if err != nil {
			result = append(result, api.AgentDeleteResultElem{
				Status:      messaging.MsgStatusStr[messaging.MSG_STATUS_ERROR],
				Description: err.Error(),
				Alias:       agent.Alias,
				ID:          id,
			})
		} else {
			// Delete all sources
			sources = SourceDeleteAll(agent.ID)

			result = append(result, api.AgentDeleteResultElem{
				Status:  messaging.MsgStatusStr[messaging.MSG_STATUS_OK],
				Alias:   agent.Alias,
				ID:      id,
				Sources: sources,
			})
		}
	}

	// Marshal response
	bytes, _ = json.Marshal(api.AgentDeleteResponse{
		Status: messaging.MsgStatusStr[messaging.MSG_STATUS_OK],
		Data:   result,
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
	var result []api.SourceListResultElem

	// Iterate through keys
	for _, key := range keys {
		// Get source info
		bytes, _ := store.Get(key)

		// Source
		var source api.SourceAddElem

		// Unmarshall
		json.Unmarshal(bytes, &source)

		// Append to response
		result = append(result, api.SourceListResultElem{
			Alias:       source.Alias,
			Description: source.Description,
			Path:        source.Path,
			ID:          string(key),
			AgentID:     source.AgentID,
		})
	}

	// Marshal response
	bytes, _ := json.Marshal(api.SourceListResponse{
		Status: messaging.MsgStatusStr[messaging.MSG_STATUS_OK],
		Data:   result,
	})

	res.Write(bytes)
}

/*
 *
 */
func HandleSourceAdd(res http.ResponseWriter, req *http.Request) {
	var sources []api.SourceAddElem

	log.Printf("Received [%s] [%s] from [%s]\n", req.Method, req.URL, req.RemoteAddr)

	// Add usual headers
	res.Header().Add("Content-Type", "application/json")

	// Read body
	bytes, _ := io.ReadAll(req.Body)

	// Unmarshal body
	json.Unmarshal(bytes, &sources)

	// Response data
	var result []api.SourceAddResultElem

	// Iterate through sources
	for _, source := range sources {
		// Add source
		id, err := SourceAdd(&source)

		if err != nil {
			log.Printf("Adding source [%s] id fail, error [%s]\n", source.Alias, err.Error())

			result = append(result, api.SourceAddResultElem{
				Status:      messaging.MsgStatusStr[messaging.MSG_STATUS_ERROR],
				Description: err.Error(),
				Alias:       source.Alias,
			})
		} else {
			log.Printf("Adding source [%s] id [%s] success\n", source.Alias, id)

			result = append(result, api.SourceAddResultElem{
				Status: messaging.MsgStatusStr[messaging.MSG_STATUS_OK],
				Alias:  source.Alias,
				ID:     id,
			})
		}
	}

	// Marshal response
	bytes, _ = json.Marshal(api.SourceAddResponse{
		Status: messaging.MsgStatusStr[messaging.MSG_STATUS_OK],
		Data:   result,
	})

	res.Write(bytes)
}

func HandleSourceDelete(res http.ResponseWriter, req *http.Request) {
	var sources []api.SourceDeleteElem

	log.Printf("Received [%s] [%s] from [%s]\n", req.Method, req.URL, req.RemoteAddr)

	// Add usual headers
	res.Header().Add("Content-Type", "application/json")

	// Read body
	bytes, _ := io.ReadAll(req.Body)

	// Unmarshal body
	json.Unmarshal(bytes, &sources)

	// Response data
	var result []api.SourceDeleteResultElem

	// Iterate through sources
	for _, source := range sources {
		// Delete source
		id, err := SourceDelete(&source)

		if err != nil {
			result = append(result, api.SourceDeleteResultElem{
				Status:      messaging.MsgStatusStr[messaging.MSG_STATUS_ERROR],
				Description: err.Error(),
				Alias:       source.Alias,
				ID:          id,
			})
		} else {
			result = append(result, api.SourceDeleteResultElem{
				Status: messaging.MsgStatusStr[messaging.MSG_STATUS_OK],
				Alias:  source.Alias,
				ID:     id,
			})
		}
	}

	// Marshal response
	bytes, _ = json.Marshal(api.SourceDeleteResponse{
		Status: messaging.MsgStatusStr[messaging.MSG_STATUS_OK],
		Data:   result,
	})

	res.Write(bytes)
}

func HandleSourceStart(res http.ResponseWriter, req *http.Request) {
	log.Printf("Received [%s] [%s] from [%s]\n", req.Method, req.URL, req.RemoteAddr)

}

func HandleSourceStop(res http.ResponseWriter, req *http.Request) {
	log.Printf("Received [%s] [%s] from [%s]\n", req.Method, req.URL, req.RemoteAddr)

}
