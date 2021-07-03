package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

var ConfigFile string // Config file location
var ServerMode bool   // Server persona

type server struct {
	ListenerEndpoint string `toml:"listenerEndpoint"`
	APIEndpoint      string `toml:"apiEndpoint"`
	LogFileLocation  string `toml:"logFileLocation"`
}

type client struct {
	ServerEndpoint  string `toml:"serverEndpoint"`
	LogFileLocation string `toml:"logFileLocation"`
}

var Server server
var Client client

func ParseServerFile(fname string) error {
	var err error = nil

	if _, err = toml.DecodeFile(fname, &Server); err != nil {
		log.Fatalf("Unable to parse config file [%s], error [%s]\n", fname, err.Error())
	}

	return err
}

func ParseClientFile(fname string) error {
	var err error = nil

	if _, err = toml.DecodeFile(fname, &Client); err != nil {
		log.Fatalf("Unable to parse config file [%s], error [%s]\n", fname, err.Error())
	}

	return err
}
