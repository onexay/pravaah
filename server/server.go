package server

import (
	"fmt"
	"net/http"
	"pravaah/api"
	"pravaah/config"
	"pravaah/db"
	"pravaah/logger"
	"pravaah/stream"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/gorilla/websocket"
)

/*
 * Server
 */
type Server struct {
	config config.Server   // Configuration
	db     db.Server       // Database
	logger logger.Server   // Log rotation
	wsConn *websocket.Conn // Websocket connection
}

// Agent instance
var me *Server

func Server_exit(server *Server) {
	fmt.Printf("Exiting server ...\n")
}

func Server_main(configFile string) {
	var err error = nil

	// Create server instance
	me = new(Server)

	// Server mode of operation
	fmt.Printf("Starting server ...\n")

	// Cleanup
	defer Server_exit(me)

	// Parse config file
	if err = me.config.Parse(configFile); err != nil {
		return
	}

	// Initialize logger
	if err = me.logger.Init(&me.config); err != nil {
		return
	}

	// Initialize DBs
	if err = me.db.InitHandle(&me.config); err != nil {
		return
	}

	// Open STATE db
	if err = me.db.InitStateDB(); err != nil {
		return
	}

	// Open SOURCES db
	if err = me.db.InitSourcesDB(); err != nil {
		return
	}

	// Open AGENTS db
	if err = me.db.InitAgentsDB(); err != nil {
		return
	}

	// Initialize a secret
	if err := InitSecret(); err != nil {
		fmt.Printf("Unable to generate secret, error [%s]\n", err.Error())
		return
	}

	// Dump secret on console for admin
	fmt.Printf("Server secret is [%s]\n", config.Secret)

	// Create a websocket router
	ws_router := chi.NewRouter()
	ws_router.Use(middleware.Logger)

	// Create a HTTP router
	http_router := chi.NewRouter()
	http_router.Use(middleware.Logger)

	// Setup routes for websocket
	ws_router.HandleFunc("/connect", HandleAgent)
	ws_router.HandleFunc("/stream", stream.Handle)

	// Setup routes for frontend
	http_router.Get(api.URL_AGENT_LIST, ListAgents)
	http_router.Post(api.URL_AGENT_ACTIVE, ActivateAgent)
	http_router.Post(api.URL_AGENT_INACTIVE, DeactivateAgent)
	http_router.Post(api.URL_AGENT_REMOVE, RemoveAgent)
	http_router.Get(api.URL_AGENT_SOURCES_LIST, ListSources)
	http_router.Post(api.URL_AGENT_SOURCES_ADD, AddSource)
	http_router.Delete(api.URL_AGENT_SOURCES_REMOVE, RemoveSource)
	http_router.Post(api.URL_AGENT_SOURCES_START, StartSource)
	http_router.Post(api.URL_AGENT_SOURCES_STOP, StopSource)

	// Wait for websocket connections from client
	go http.ListenAndServe(me.config.ListenerEndpoint, ws_router)

	// Wait for connections from frontend
	http.ListenAndServe(me.config.APIEndpoint, http_router)
}
