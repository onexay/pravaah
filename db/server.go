package db

import (
	"log"
	"pravaah/config"

	ledisDBCfg "github.com/ledisdb/ledisdb/config"
	ledisDB "github.com/ledisdb/ledisdb/ledis"
)

type Server struct {
	cfg    *ledisDBCfg.Config  // DB configuration
	handle *ledisDB.Ledis      // DB handle
	db     [DB_MAX]*ledisDB.DB // DB
}

func (server *Server) InitHandle(cfg *config.Server) error {
	var err error = nil

	// Initialize DB configuration
	server.cfg = ledisDBCfg.NewConfigDefault()

	// Set DB location
	server.cfg.DataDir = cfg.DBLocation

	// Open LEDIS
	server.handle, err = ledisDB.Open(server.cfg)
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
func (server *Server) initDB(num int) error {
	var err error = nil

	server.db[num], err = server.handle.Select(num)
	if err != nil {
		log.Fatalf("Unable to open db [%d] in [%s], error [%s]\n", server.db[num].Index(), server.cfg.DataDir, err.Error())
		return err
	}

	log.Printf("Opened db [%s] index [%d] successfully.\n", server.cfg.DataDir, server.db[num].Index())

	return err
}

/*
 *
 */
func (server *Server) InitStateDB() error {
	return server.initDB(DB_STATE)
}

/*
 *
 */
func (server *Server) InitAgentsDB() error {
	return server.initDB(DB_AGENTS)
}

/*
 *
 */
func (server *Server) InitSourcesDB() error {
	return server.initDB(DB_SOURCES)
}

/*
 *
 */
func (server *Server) InitAliasesDB() error {
	return server.initDB(DB_ALIASES)
}

/*
 *
 */
func (server *Server) getDB(num int) *ledisDB.DB {
	return server.db[num]
}

/*
 *
 */
func (server *Server) GetStateDB() *ledisDB.DB {
	return server.getDB(DB_STATE)
}

/*
 *
 */
func (server *Server) GetSourceDB() *ledisDB.DB {
	return server.getDB(DB_SOURCES)
}

/*
 *
 */
func (server *Server) GetAgentsDB() *ledisDB.DB {
	return server.getDB(DB_AGENTS)
}

/*
 *
 */
func (server *Server) GetAliasesDB() *ledisDB.DB {
	return server.getDB(DB_ALIASES)
}
