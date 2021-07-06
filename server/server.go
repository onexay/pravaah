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

func Server_main(configFile string) {
	var err error = nil

	// Create server instance
	me = new(Server)

	// Server mode of operation
	fmt.Printf("Starting server ...\n")

	// Parse config file
	if err = me.config.Parse(configFile); err != nil {
		return
	}

	// Initialize logger
	if err = me.logger.Init(&me.config); err != nil {
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
	http_router.Get(api.URL_LIST, api.List)
	http_router.Post(api.URL_ADD, api.Add)
	http_router.Delete(api.URL_REMOVE, api.Remove)
	http_router.Post(api.URL_START, api.Start)
	http_router.Post(api.URL_STOP, api.Stop)

	// Wait for websocket connections from client
	go http.ListenAndServe(":10080", ws_router)

	// Wait for connections from frontend
	http.ListenAndServe(":8080", http_router)
}
