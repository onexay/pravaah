package logger

import (
	"log"
	"pravaah/config"

	"gopkg.in/natefinch/lumberjack.v2"
)

type Agent struct {
	logger lumberjack.Logger // Log file rotation
}

func (agent *Agent) Init(cfg *config.Agent) error {
	// Setup log rotation
	agent.logger.Filename = cfg.LogFileLocation
	agent.logger.MaxSize = 1 // MB
	log.SetPrefix("[AGENT] ")
	log.SetOutput(&agent.logger)

	return nil
}
