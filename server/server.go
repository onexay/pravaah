/* MIT License
 *
 * Copyright (c) 2021 Akshay Ranjan
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */
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

	// Open ALIASES db
	if err = me.db.InitAliasesDB(); err != nil {
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
	http_router.Get(api.URL_SECRET_LIST, ListSecrets)
	http_router.Get(api.URL_SECRET_RENEW, RenewSecrets)

	http_router.Get(api.URL_AGENT_LIST, HandleAgentsList)
	http_router.Post(api.URL_AGENT_ACTIVE, HandleAgentActivate)
	http_router.Post(api.URL_AGENT_INACTIVE, HandleAgentDeactivate)
	http_router.Post(api.URL_AGENT_DELETE, HandleAgentDelete)

	http_router.Get(api.URL_AGENT_SOURCES_LIST, HandleSourcesList)
	http_router.Post(api.URL_AGENT_SOURCES_ADD, HandleSourceAdd)
	http_router.Post(api.URL_AGENT_SOURCES_DELETE, HandleSourceDelete)
	http_router.Post(api.URL_AGENT_SOURCES_START, HandleSourceStart)
	http_router.Post(api.URL_AGENT_SOURCES_STOP, HandleSourceStop)

	// Wait for websocket connections from client
	go http.ListenAndServe(me.config.ListenerEndpoint, ws_router)

	// Wait for connections from frontend
	http.ListenAndServe(me.config.APIEndpoint, http_router)
}
