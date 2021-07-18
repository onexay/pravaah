package db

import (
	"log"
	"pravaah/config"

	ledisDBCfg "github.com/ledisdb/ledisdb/config"
	ledisDB "github.com/ledisdb/ledisdb/ledis"
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

/*
 *
 */
func (agent *Agent) initDB(num int) error {
	var err error = nil

	agent.db[num], err = agent.handle.Select(num)
	if err != nil {
		log.Fatalf("Unable to open db [%d] in [%s], error [%s]\n", agent.db[num].Index(), agent.cfg.DataDir, err.Error())
		return err
	}

	log.Printf("Opened db [%s] index [%d] successfully.\n", agent.cfg.DataDir, agent.db[num].Index())

	return err
}

/*
 *
 */
func (agent *Agent) InitStateDB() error {
	return agent.initDB(DB_STATE)
}

/*
 *
 */
func (Agent *Agent) InitSourcesDB() error {
	return Agent.initDB(DB_SOURCES)
}

/*
 *
 */
func (agent *Agent) getDB(num int) *ledisDB.DB {
	return agent.db[num]
}

/*
 *
 */
func (agent *Agent) GetStateDB() *ledisDB.DB {
	return agent.getDB(DB_STATE)
}

/*
 *
 */
func (agent *Agent) GetSourceDB() *ledisDB.DB {
	return agent.getDB(DB_SOURCES)
}
