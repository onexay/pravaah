package logger

import "pravaah/config"

type Logger_T interface {
	// Init
	Init(config *config.Config_T) error
}
