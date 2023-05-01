package server

import (
	"encoding/json"
	"log"

	"github.com/gofiber/websocket/v2"
)

type wsMessage struct {
	Type string            `json:"type"`
	Data map[string]string `json:"data"`
}

func (s *Server) wsHandler(c *websocket.Conn) {
	defer func() {
		c.Close()
	}()

	for {
		_, msg, err := c.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Println("read error:", err)
			}

			return
		}

		wsMsg := wsMessage{}
		if err := json.Unmarshal(msg, &wsMsg); err != nil {
			log.Println("unmarshal error:", err)
			return
		}

		switch wsMsg.Type {
		case "join":
			s.hub.JoinRoom(wsMsg.Data["room_id"], wsMsg.Data["player_id"], wsMsg.Data["join_code"])
		}

	}

}
