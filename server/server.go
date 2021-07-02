package server

import (
	"net/http"
	"pravaah/stream"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func Server_main() {
	// Server mode of operation

	// Create a websocket router
	ws_router := chi.NewRouter()
	ws_router.Use(middleware.Logger)

	// Create a HTTP router
	http_router := chi.NewRouter()
	http_router.Use(middleware.Logger)

	// Setup routes
	ws_router.HandleFunc("/stream", stream.Handle)

	// Wait for websocket connections from client
	go http.ListenAndServe(":10080", ws_router)
}
