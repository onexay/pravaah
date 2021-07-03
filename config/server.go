package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

type server struct {
	ListenerEndpoint string `toml:"listenerEndpoint"`
	APIEndpoint      string `toml:"apiEndpoint"`
	LogFileLocation  string `toml:"logFileLocation"`
	DBLocation       string `toml:"dbLocation"`
}

var Secret string
var Server server

func ParseServerFile(fname string) error {
	var err error = nil

	if _, err = toml.DecodeFile(fname, &Server); err != nil {
		log.Fatalf("Unable to parse config file [%s], error [%s]\n", fname, err.Error())
	}

	return err
}
