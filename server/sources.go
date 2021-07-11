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
	"log"
	"pravaah/api"

	"github.com/google/uuid"
	"github.com/ledisdb/ledisdb/ledis"
)

/*
 * Add a source to prepare for streaming
 */
func SourceAdd(source *api.Source) (string, bool, error) {
	// Get agents DB
	agent_store := me.db.GetAgentsDB()

	// Get aliases DB
	alias_store := me.db.GetAliasesDB()

	// Check agent
	i, _ := agent_store.Exists([]byte(source.AgentID))

	log.Printf("Checking source agent [%s], [%d]\n", source.AgentID, i)

	if i == 0 {
		return "", false, ledis.ErrScoreMiss
	}

	// Get sources DB
	source_store := me.db.GetSourceDB()

	// Marshal data
	bytes, _ := json.Marshal(source)

	// Check source alias
	if bytes, err := alias_store.Get([]byte(source.Alias)); err == nil {
		// Existing entry found
		return string(bytes), true, nil
	}

	// Generate a new ID for this source
	id := uuid.New().String()

	// Create source binding and persist
	source_store.Set([]byte(id), bytes)

	// Create alias binding and persist
	alias_store.Set([]byte(source.Alias), []byte(id))

	return id, false, nil
}
