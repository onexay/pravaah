package logger

import (
	"log"
	"pravaah/config"

	"gopkg.in/natefinch/lumberjack.v2"
)

type Server struct {
	logger lumberjack.Logger // Log file rotation
}

func (server *Server) Init(cfg *config.Server) error {
	// Setup log rotation
	server.logger.Filename = cfg.LogFileLocation
	server.logger.MaxSize = 1 // MB
	log.SetPrefix("[SERVER] ")
	log.SetOutput(&server.logger)

	return nil
}
