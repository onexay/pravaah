package config

var ConfigFile string // Config file location
var ServerMode bool   // Server persona
var Secret string     // Secret string

type Config_T interface {
	// Parse configuration
	Parse(fname *string) error
}
