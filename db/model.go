package db

import (
	"pravaah/config"

	ledisDB "github.com/ledisdb/ledisdb/ledis"
)

const (
	// Database keys
	AGENT_ID = "AGENT_ID"
)

const (
	DB_STATE   = 0 // Runtime info
	DB_SOURCES = 1 // Sources
	DB_AGENTS  = 2 // Agents
	DB_MAX     = 3 // MUST BE LAST
)

var DBName []string = []string{
	"STATE",
	"SOURCES",
	"AGENTS",
	"MAX",
}

type DB_T interface {
	// Init
	InitHandle(config *config.Config_T) error
	InitStateDB() error
	InitSourcesDB() error
	InitAgentsDB() error

	// Get DBs
	GetStateDB() *ledisDB.DB
	GetSourcesDB() *ledisDB.DB
	GetAgentsDB() *ledisDB.DB

	// AGENT_ID ops
	GetAgentID() (string, error)
	SetAgentID(string) error
}
