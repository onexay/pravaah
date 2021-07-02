package session

import "github.com/gorilla/websocket"

type session struct {
	ws websocket.Conn // Per session websocket
}
