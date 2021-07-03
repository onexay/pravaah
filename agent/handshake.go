package agent

import (
	"log"
	"pravaah/messaging"

	"github.com/gorilla/websocket"
)

func BeginHandshake(conn *websocket.Conn) {
	var msg messaging.CapabilitiesMsg = messaging.CapabilitiesMsg{
		Version: "v0.2",
		Secret:  "ec52b2292c19bc4087f0a7dfb0092f85d7725a30344f88bb0dd98a1354fd287e",
	}

	if err := conn.WriteJSON(msg); err != nil {
		log.Fatalf("Unable to send handshake message, error [%s]\n", err.Error())
		return
	}
}

func ProcessHandshake() {

}
