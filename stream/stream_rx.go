package stream

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func Handle(res http.ResponseWriter, req *http.Request) {
	conn, err := upgrader.Upgrade(res, req, nil)
	if err != nil {
		log.Printf("Unable to upgrade to websocket from [%s], error [%s]", req.RequestURI, err.Error())
		return
	}

	defer conn.Close()

	for {
		mt, msg, rx_err := conn.ReadMessage()
		if rx_err != nil {
			log.Printf("Unable to receive message, error [%s]", rx_err.Error())
			return
		}

		log.Printf("Type: %d, Msg: %s\n", mt, msg)
	}
}
