package handler

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

var roomConnections = make(map[string][]*websocket.Conn)

func WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	roomID := r.URL.Query().Get("room_code")
	if roomID == "" {
		roomID = "room-123" // default room
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}

	roomConnections[roomID] = append(roomConnections[roomID], conn)
	log.Println("New connection in", roomID)

	defer func() {
		conn.Close()
		log.Println("Connection closed in", roomID)
	}()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}

		// Broadcast to others in the same room
		for _, peer := range roomConnections[roomID] {
			if peer != conn {
				peer.WriteMessage(websocket.TextMessage, msg)
			}
		}
	}
}
