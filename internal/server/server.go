package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

type Server struct {
}

func (s *Server) Run() error {
	app := fiber.New()

	// api endpoints to manage rooms
	app.Route("/api", func(r fiber.Router) {
		r.Post("/rooms", s.createRoomHandler)
		r.Get("/rooms/:id", s.getRoomHandler)
		r.Delete("/rooms/:id", s.deleteRoomHandler)
	})

	app.Get("/ws", websocket.New(s.wsHandler))

	// TODO: Add options to configure the port
	return app.Listen(":8081")
}

func (s *Server) wsHandler(c *websocket.Conn) {

}
