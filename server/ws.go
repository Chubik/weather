package server

import (
	"fmt"
	"net/http"
	"strings"
	"weather/config"
	"weather/data"

	"github.com/gorilla/websocket"
)

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func wshandler(w http.ResponseWriter, r *http.Request) {
	conn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Failed to set websocket upgrade: %+v", err)
		return
	}

	for {
		t, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}
		s := string(msg[:])

		pc := strings.ToLower(s)
		if len(pc) < config.C.Cap {
			se := string("Short name of City. Should be more > " + string(config.C.Cap))
			conn.WriteMessage(t, []byte(se))
			break
		}
		resp := data.Search(pc)
		if len(resp) > 0 {
			conn.WriteJSON(resp)
		} else {
			conn.WriteMessage(t, []byte("Cities not found"))
		}
	}
}
