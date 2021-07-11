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
