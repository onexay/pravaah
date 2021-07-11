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
	"pravaah/api"
	"pravaah/messaging"
)

/*
 *
 */
func RegisterAgent(uid string, req *messaging.ConnectReqMsg) {
	store := me.db.GetAgentsDB()

	// Encode agent info
	encoded, _ := json.Marshal(api.AgentAdd{
		Alias:       req.Alias,
		Description: req.Description,
		Version:     req.Version,
		Secret:      req.Secret,
		ID:          uid,
	})

	// Persist info
	store.Set([]byte(uid), encoded)
}

/*
 *
 */
func RetrieveAgent(uid string) (api.AgentAdd, error) {
	var agent api.AgentAdd = api.AgentAdd{}

	// Get DB
	store := me.db.GetAgentsDB()

	// Lookup uid
	data, err := store.Get([]byte(uid))
	if err != nil {
		return agent, err
	}

	// Unmarshal data
	err = json.Unmarshal(data, &agent)

	return agent, err
}

/*
 *
 */
func DeregisterAgent(uid string) error {
	store := me.db.GetAgentsDB()

	// Delete agent info
	if _, err := store.Del([]byte(uid)); err != nil {
		return err
	}

	return nil
}
