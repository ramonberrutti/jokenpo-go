package server

import "github.com/gofiber/fiber/v2"

type webhookConfig struct {
	Url     string            `json:"url"`
	Headers map[string]string `json:"headers"`
}

type createRoomBodyReq struct {
	// Id internal id of the room.
	Id string `json:"id"`
	// Id used to identify the room for the creator of the room.
	// You can't use this to join the room.
	ExternalId string `json:"external_id"`
	// Name of the room that will be vissible to the users.
	Name string `json:"name"`
	// Webhook to notify about events in the room.
	EventsWebhook webhookConfig `json:"events_webhook"`
	// Player 1 of the room.
	Player1 string `json:"player1"`
	// Player 2 of the room.
	Player2 string `json:"player2"`
	// Rounds to play in the room.
	Rounds int `json:"rounds"`
}

type createRoomBodyResp struct {
	// Id used to identify the room for the creator of the room
	Id string `json:"name"`
	// Player 1 join code.
	Player1JoinCode string `json:"player1_join_code"`
	// Player 2 join code.
	Player2JoinCode string `json:"player2_join_code"`
}

func (s *Server) createRoomHandler(c *fiber.Ctx) error {
	req := createRoomBodyReq{}
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(createRoomBodyResp{
		Id: "123",
	})
}

func (s *Server) getRoomHandler(c *fiber.Ctx) error {
	return nil
}

func (s *Server) deleteRoomHandler(c *fiber.Ctx) error {
	return nil
}
