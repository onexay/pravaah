package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

type agent struct {
	ServerEndpoint  string `toml:"serverEndpoint"`
	ServerSecret    string `toml:"serverSecret"`
	LogFileLocation string `toml:"logFileLocation"`
	DBLocation      string `toml:"dbLocation"`
}

var Agent agent

func ParseAgentFile(fname string) error {
	var err error = nil

	if _, err = toml.DecodeFile(fname, &Agent); err != nil {
		log.Fatalf("Unable to parse config file [%s], error [%s]\n", fname, err.Error())
	}

	return err
}
