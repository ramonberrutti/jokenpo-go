package hub

import (
	"errors"
	"sync"
)

// The hub contain the information of all the connected clients.
// Each client is identified by a unique ID.
// They are going to be distributed in the available rooms to play the game.

func NewHub() *Hub {
	return &Hub{
		rooms: make(map[string]*Room),
	}
}

type Hub struct {
	sync.RWMutex
	// Rooms of the hub.
	rooms map[string]*Room
}

func (h *Hub) RegisterRoom(room *Room) {
	h.Lock()
	defer h.Unlock()

	h.rooms[room.id] = room
}

func (h *Hub) GetRoom(id string) *Room {
	h.RLock()
	defer h.RUnlock()

	return h.rooms[id]
}

func (h *Hub) UnregisterRoom(id string) {
	h.Lock()
	defer h.Unlock()

	delete(h.rooms, id)
}

var ErrNoRoom = errors.New("no room")
var ErrInvalidJoinCode = errors.New("invalid join code")

func (h *Hub) JoinRoom(roomId, playerId, joinCode string) (*Room, error) {
	h.RLock()
	defer h.RUnlock()

	room, ok := h.rooms[roomId]
	if !ok {
		return nil, ErrNoRoom
	}

	if room.GetPlayer1() == playerId && room.GetPlayer1JoinCode() == joinCode {
		return room, nil
	} else if room.GetPlayer2() == playerId && room.GetPlayer2JoinCode() == joinCode {
		return room, nil
	}

	return nil, ErrInvalidJoinCode
}
