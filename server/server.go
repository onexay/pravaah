/*
 *
 *
 */
package server

import (
	"fmt"
	"log"
	"net/http"
	"pravaah/api"
	"pravaah/config"
	"pravaah/stream"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Logger
var logger lumberjack.Logger

func Server_main(configFile string) {
	// Server mode of operation
	fmt.Printf("Starting server ...\n")

	// Check for config file
	if configFile == "" {
		fmt.Printf("No config file specified. Exiting.\n")
		return
	}

	// Logger setup
	logger.Filename = config.Server.LogFileLocation
	logger.MaxSize = 1
	log.SetOutput(&logger)

	// Create a websocket router
	ws_router := chi.NewRouter()
	ws_router.Use(middleware.Logger)

	// Create a HTTP router
	http_router := chi.NewRouter()
	http_router.Use(middleware.Logger)

	// Setup routes for websocket
	ws_router.HandleFunc("/stream", stream.Handle)

	// Setup routes for frontend
	http_router.Get(api.URL_LIST, api.List)
	http_router.Post(api.URL_ADD, api.Add)
	http_router.Delete(api.URL_REMOVE, api.Remove)

	// Wait for websocket connections from client
	go http.ListenAndServe(":10080", ws_router)

	// Wait for connections from frontend
	http.ListenAndServe(":8080", http_router)
}
