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

func (server *Server) InitStateDB() error {
	var err error = nil

	server.db[DB_STATE], err = server.handle.Select(DB_STATE)
	if err != nil {
		log.Fatalf("Unable to open db [%d] in [%s], error [%s]\n", server.db[DB_STATE].Index(), server.cfg.DataDir, err.Error())
		return err
	}

	log.Printf("Opened db [%s] index [%d] successfully.\n", server.cfg.DataDir, server.db[DB_STATE].Index())

	return err
}

func (server *Server) InitSourcesDB() error {
	return nil
}

func (server *Server) InitAgentsDB() error {
	var err error = nil

	server.db[DB_AGENTS], err = server.handle.Select(DB_AGENTS)
	if err != nil {
		log.Fatalf("Unable to open db [%d] in [%s], error [%s]\n", server.db[DB_AGENTS].Index(), server.cfg.DataDir, err.Error())
		return err
	}

	log.Printf("Opened db [%s] index [%d] successfully.\n", server.cfg.DataDir, server.db[DB_AGENTS].Index())

	return err
}

func (server *Server) GetStateDB() *ledisDB.DB {
	return server.db[DB_STATE]
}

func (server *Server) GetSourceDB() *ledisDB.DB {
	return server.db[DB_SOURCES]
}

func (server *Server) GetAgentsDB() *ledisDB.DB {
	return server.db[DB_AGENTS]
}
