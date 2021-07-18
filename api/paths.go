package api

const (
	// Base URL
	URL_BASE    = "/v1"
	URL_SECRET  = "/secret"
	URL_AGENTS  = "/agents"
	URL_SOURCES = "/sources"

	// Secret
	URL_SECRET_LIST  = URL_BASE + URL_SECRET + "/list"  // Show current secret
	URL_SECRET_RENEW = URL_BASE + URL_SECRET + "/renew" // Regenerate new secret

	// Agents
	URL_AGENT_LIST     = URL_BASE + URL_AGENTS + "/list"
	URL_AGENT_ACTIVE   = URL_BASE + URL_AGENTS + "/active"
	URL_AGENT_INACTIVE = URL_BASE + URL_AGENTS + "/inactive"
	URL_AGENT_DELETE   = URL_BASE + URL_AGENTS + "/delete"

	// Sources
	URL_AGENT_SOURCES_LIST   = URL_BASE + URL_SOURCES + "/list"
	URL_AGENT_SOURCES_ADD    = URL_BASE + URL_SOURCES + "/add"
	URL_AGENT_SOURCES_DELETE = URL_BASE + URL_SOURCES + "/delete"
	URL_AGENT_SOURCES_START  = URL_BASE + URL_SOURCES + "/start"
	URL_AGENT_SOURCES_STOP   = URL_BASE + URL_SOURCES + "/stop"
)
