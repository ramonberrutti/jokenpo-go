package hub

// The hub contain the information of all the connected clients.
// Each client is identified by a unique ID.
// They are going to be distributed in the available rooms to play the game.

func NewHub() *Hub {
	return &Hub{}
}

type Hub struct {
}
