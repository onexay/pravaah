package db

import (
	"pravaah/config"

	ledisDB "github.com/ledisdb/ledisdb/ledis"
)

const (
	// Prefixes
	PREFIX_AGENT  = "AGENT_"
	PREFIX_SOURCE = "SOURCE_"
)

const (
	// Database keys
	AGENT_ID = "AGENT_ID"
	SECRET   = "SECRET"
)

const (
	DB_STATE   = 0 // Runtime info
	DB_SOURCES = 1 // Sources
	DB_AGENTS  = 2 // Agents
	DB_ALIASES = 3 // Aliases
	DB_MAX     = 4 // MUST BE LAST
)

var DBName []string = []string{
	"STATE",
	"SOURCES",
	"AGENTS",
	"ALIASES",
	"MAX",
}

type DB_T interface {
	// Private
	initDB(num int) error
	getDB(num int) *ledisDB.DB

	// Init
	InitHandle(config *config.Config_T) error
	InitStateDB() error
	InitSourcesDB() error
	InitAgentsDB() error

	// Get DBs
	GetStateDB() *ledisDB.DB
	GetSourcesDB() *ledisDB.DB
	GetAgentsDB() *ledisDB.DB
}
