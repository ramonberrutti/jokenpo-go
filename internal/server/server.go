package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/ramonberrutti/jokenpo-go/internal/hub"
)

type Server struct {
	hub *hub.Hub
}

func (s *Server) Run() error {
	app := fiber.New()
	s.hub = hub.NewHub()

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
