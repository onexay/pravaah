package db

import "pravaah/config"

const (
	// Database keys
	AGENT_ID = "AGENT_ID"
)

type DB_T interface {
	// Init
	InitHandle(config *config.Config_T) error
	InitState() error
	InitSource() error

	// AGENT_ID ops
	GetAgentID() (string, error)
	SetAgentID(string) error
}
