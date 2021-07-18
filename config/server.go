package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

type Server struct {
	ListenerEndpoint string `toml:"listenerEndpoint"`
	APIEndpoint      string `toml:"apiEndpoint"`
	LogFileLocation  string `toml:"logFileLocation"`
	DBLocation       string `toml:"dbLocation"`
}

func (cfg *Server) Parse(fname string) error {
	var err error = nil

	if _, err = toml.DecodeFile(fname, cfg); err != nil {
		log.Fatalf("Unable to parse config file [%s], error [%s]\n", fname, err.Error())
	}

	return err
}

func (cfg *Server) GenerateAlias() string {
	/* If auto generation of alias is configured, then generate an alias
	 * and overwrite the one read from config file (even if present); we'll
	 * emit a info in log for this though.
	 */
	return ""
}
