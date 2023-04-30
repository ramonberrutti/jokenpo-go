package hub_test

import (
	"testing"

	"github.com/ramonberrutti/jokenpo-go/internal/hub"
	"github.com/stretchr/testify/assert"
)

// TODO: Create a table test for this.
func TestRoomCompleteGame(t *testing.T) {
	room := hub.NewRoom("id", "name", "player1", "player2", 3)

	room.Start()

	p1, p2 := room.GetResults()
	assert.Equal(t, 0, p1)
	assert.Equal(t, 0, p2)

	room.AddPlayer1Move(hub.Rock)

	p1, p2 = room.GetResults()
	assert.Equal(t, 0, p1)
	assert.Equal(t, 0, p2)

	room.AddPlayer2Move(hub.Paper)

	p1, p2 = room.GetResults()
	assert.Equal(t, 0, p1)
	assert.Equal(t, 1, p2)

	room.AddPlayer1Move(hub.Paper)

	p1, p2 = room.GetResults()
	assert.Equal(t, 0, p1)
	assert.Equal(t, 1, p2)

	room.AddPlayer2Move(hub.Rock)

	p1, p2 = room.GetResults()
	assert.Equal(t, 1, p1)
	assert.Equal(t, 1, p2)

	room.AddPlayer1Move(hub.Paper)

	p1, p2 = room.GetResults()
	assert.Equal(t, 1, p1)
	assert.Equal(t, 1, p2)

	room.AddPlayer2Move(hub.Paper)

	p1, p2 = room.GetResults()
	assert.Equal(t, 1, p1)
	assert.Equal(t, 1, p2)

	room.AddPlayer1Move(hub.Scissors)

	p1, p2 = room.GetResults()
	assert.Equal(t, 1, p1)
	assert.Equal(t, 1, p2)

	room.AddPlayer2Move(hub.Rock)

	p1, p2 = room.GetResults()
	assert.Equal(t, 1, p1)
	assert.Equal(t, 2, p2)

	assert.Equal(t, hub.Finished, room.GetState())
}
