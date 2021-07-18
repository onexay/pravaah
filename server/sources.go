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
	"errors"
	"pravaah/api"
	"pravaah/db"
	"pravaah/messaging"

	"github.com/google/uuid"
)

/*
 * Add a source to prepare for streaming
 */
func SourceAdd(source *api.SourceAddElem) (string, error) {
	// Get agents DB
	agent_store := me.db.GetAgentsDB()

	// Check agent
	e, err := agent_store.Exists([]byte(source.AgentID))

	if err != nil {
		return "", err
	}

	if e == 0 {
		return "", errors.New("agent ID not found")
	}

	// Get aliases DB
	alias_store := me.db.GetAliasesDB()

	// Check source alias
	bytes, err := alias_store.Get([]byte(db.PREFIX_SOURCE + source.Alias))

	if err != nil {
		return "", err
	}

	if bytes != nil {
		return "", errors.New("alias [%s] for source [%s] already exists")
	}

	// Get sources DB
	source_store := me.db.GetSourceDB()

	// Marshal data
	bytes, _ = json.Marshal(source)

	// Generate a new ID for this source
	id := uuid.New().String()

	// Create source binding and persist
	source_store.Set([]byte(id), bytes)

	// Create alias binding and persist
	alias_store.Set([]byte(db.PREFIX_SOURCE+source.Alias), []byte(id))

	// Add to agent
	agent_store.SAdd([]byte(source.AgentID), []byte(id))

	return id, nil
}

/*
 * Delete a source
 */
func SourceDelete(source *api.SourceDeleteElem) (string, error) {
	// Get agents DB
	agents_store := me.db.GetAgentsDB()

	// Check agent id
	e, err := agents_store.Exists([]byte(source.AgentID))

	if err != nil {
		return source.AgentID, err
	}

	if e == 0 {
		return source.AgentID, errors.New("agent ID not found")
	}

	// Remove from agent
	agents_store.SRem([]byte(source.AgentID), []byte(source.ID))

	// Get sources DB
	sources_store := me.db.GetSourceDB()

	// Check source id
	e, err = sources_store.Del([]byte(source.ID))

	if err != nil {
		return source.ID, err
	}

	if e == 0 {
		return source.ID, errors.New("source ID not found")
	}

	// Get aliases DB
	aliases_store := me.db.GetAliasesDB()

	aliases_store.Del([]byte(db.PREFIX_SOURCE + source.Alias))

	return source.ID, nil
}

func SourceDeleteAll(agentId string) []api.SourceDeleteResultElem {
	// Get agents DB
	agents_store := me.db.GetAgentsDB()

	// Get all sources
	ids, _ := agents_store.SMembers([]byte(agentId))

	// Response data
	var rspData []api.SourceDeleteResultElem

	// Iterate
	for _, id := range ids {
		// Delete source
		source, err := SourceDeleteByID(string(id))

		if err == nil {
			rspData = append(rspData, *source)
		}
	}

	return rspData
}

func SourceDeleteByID(id string) (*api.SourceDeleteResultElem, error) {
	var source api.SourceAddElem

	// Get sources DB
	sources_store := me.db.GetSourceDB()

	// Lookup source
	sourceBytes, err := sources_store.Get([]byte(id))

	if err != nil {
		return nil, err
	}

	if sourceBytes == nil {
		return nil, errors.New("source not found")
	}

	json.Unmarshal(sourceBytes, &source)

	// Get agent DB
	agents_store := me.db.GetAgentsDB()

	// Remove from agent
	agents_store.SRem([]byte(source.AgentID), []byte(id))

	// Get aliases DB
	aliases_store := me.db.GetAliasesDB()

	aliases_store.Del([]byte(db.PREFIX_SOURCE + source.Alias))

	sources_store.Del([]byte(id))

	return &api.SourceDeleteResultElem{
		Status: messaging.MsgStatusStr[messaging.MSG_STATUS_OK], // Status
		Alias:  source.Alias,                                    // Alias
		ID:     id,                                              // Source UID
	}, nil
}

func SourceDeleteByAlias(alias string) (*api.SourceDeleteResultElem, error) {
	var source api.SourceAddElem

	// Get aliases DB
	aliases_store := me.db.GetAgentsDB()

	// Lookup alias
	idBytes, err := aliases_store.Get([]byte(db.PREFIX_SOURCE + alias))

	if err != nil {
		return nil, err
	}

	if idBytes == nil {
		return nil, errors.New("source alias not found")
	}

	// Get sources DB
	sources_store := me.db.GetSourceDB()

	// Lookup source
	sourceBytes, _ := sources_store.Get(idBytes)

	json.Unmarshal(sourceBytes, &source)

	// Get agent DB
	agents_store := me.db.GetAgentsDB()

	// Remove from agent
	agents_store.SRem([]byte(source.AgentID), idBytes)

	aliases_store.Del([]byte(db.PREFIX_SOURCE + alias))

	sources_store.Del(idBytes)

	return &api.SourceDeleteResultElem{
		Status: messaging.MsgStatusStr[messaging.MSG_STATUS_OK], // Status
		Alias:  source.Alias,                                    // Alias
		ID:     string(idBytes),                                 // Source UID
	}, nil
}
