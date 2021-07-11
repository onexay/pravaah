package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

type Agent struct {
	Alias           string `toml:"alias"`
	Description     string `toml:"description"`
	ServerEndpoint  string `toml:"serverEndpoint"`
	ServerSecret    string `toml:"serverSecret"`
	LogFileLocation string `toml:"logFileLocation"`
	DBLocation      string `toml:"dbLocation"`
}

func (cfg *Agent) Parse(fname string) error {
	var err error = nil

	if _, err = toml.DecodeFile(fname, cfg); err != nil {
		log.Fatalf("Unable to parse config file [%s], error [%s]\n", fname, err.Error())
	}

	return err
}
