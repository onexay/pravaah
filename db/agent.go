package db

import (
	"log"
	"pravaah/config"

	ledisDBCfg "github.com/ledisdb/ledisdb/config"
	ledisDB "github.com/ledisdb/ledisdb/ledis"
)

const (
	DB_STATE   = 0 // Runtime state keys
	DB_SOURCES = 1 // Sources
	DB_MAX     = 2 // MUST BE LAST
)

type Agent struct {
	cfg    *ledisDBCfg.Config  // DB configuration
	handle *ledisDB.Ledis      // DB handle
	db     [DB_MAX]*ledisDB.DB // DB
}

func (agent *Agent) InitHandle(cfg *config.Agent) error {
	var err error = nil

	// Initialize DB configuration
	agent.cfg = ledisDBCfg.NewConfigDefault()

	// Set DB location
	agent.cfg.DataDir = cfg.DBLocation

	// Open LEDIS
	agent.handle, err = ledisDB.Open(agent.cfg)
	if err != nil {
		log.Fatalf("Unable to open db [%s] handle, error [%s]", cfg.DBLocation, err.Error())
		return err
	}

	log.Printf("Opened db [%s] handle successfully.\n", cfg.DBLocation)

	return err
}

func (agent *Agent) InitState() error {
	var err error = nil

	agent.db[DB_STATE], err = agent.handle.Select(DB_STATE)
	if err != nil {
		log.Fatalf("Unable to open db [%d] in [%s], error [%s]\n", agent.db[DB_STATE].Index(), agent.cfg.DataDir, err.Error())
		return err
	}

	log.Printf("Opened db [%s] index [%d] successfully.\n", agent.cfg.DataDir, agent.db[DB_STATE].Index())

	return err
}

func (agent *Agent) GetAgentID() (string, error) {
	// Try to retrieve
	val, err := agent.db[DB_STATE].Get([]byte(AGENT_ID))

	return string(val), err
}

func (agent *Agent) SetAgentId(val string) error {
	// Try to set
	err := agent.db[DB_STATE].Set([]byte(AGENT_ID), []byte(val))

	return err
}
