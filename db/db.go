package db

import (
	"log"
	"pravaah/config"

	ledisDBCfg "github.com/ledisdb/ledisdb/config"
	ledisDB "github.com/ledisdb/ledisdb/ledis"
)

var DB *ledisDB.DB

func Init() error {
	// Initialize LEDIS
	dbCfg := ledisDBCfg.NewConfigDefault()
	dbCfg.DataDir = config.Agent.DBLocation

	// Open LEDIS
	db, err := ledisDB.Open(dbCfg)
	if err != nil {
		log.Fatalf("Unable to open db instance [%s], error [%s]", config.Agent.DBLocation, err.Error())
		return err
	}

	log.Printf("Opened db instance [%s] successfully.\n", config.Agent.DBLocation)

	// For agent its always one DB
	DB, err = db.Select(0)
	if err != nil {
		log.Fatalf("Unable to select default db in instance [%s], error [%s]\n", config.Agent.DBLocation, err.Error())
		return err
	}

	log.Printf("Selected default db in instance [%s] successfully.\n", config.Agent.DBLocation)

	return nil
}
