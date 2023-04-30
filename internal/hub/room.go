package hub

import "sync"

type Room struct {
	// Lock to prevent concurrent access to the room.
	sync.Mutex

	// Id of the room.
	Id string

	// Name of the room.
	Name string

	// Player 1 of the room.
	Player1 string

	// Player 2 of the room.
	Player2 string

	// Player 1 join code.
	Player1JoinCode string

	// Player 2 join code.
	Player2JoinCode string

	// State of the room.
	State RoomState

	// Rounds to play in the room.
	Rounds int

	// Current round of the room.
	CurrentRound int

	// Current round of the room.
	Results []RoundResult
}

type RoomState int

const (
	// Room is waiting for players to join.
	WaitingForPlayers RoomState = iota

	// Room is running.
	Running

	// Room is finished.
	Finished
)

type Move int

const (
	// No move.
	NoMove Move = iota

	// Rock move.
	Rock

	// Paper move.
	Paper

	// Scissors move.
	Scissors
)

type RoundResult struct {
	// Player 1 move.
	Player1Move Move

	// Player 2 move.
	Player2Move Move
}

func NewRoom(id string, name string, player1 string, player2 string, rounds int) *Room {
	return &Room{
		Id:      id,
		Name:    name,
		Player1: player1,
		Player2: player2,
		Rounds:  rounds,
		State:   WaitingForPlayers,
		Results: make([]RoundResult, rounds),
	}
}

func (r *Room) Start() {
	r.Lock()
	defer r.Unlock()
	if r.State != WaitingForPlayers {
		return
	}

	r.State = Running
}

func (r *Room) AddPlayer1Move(move Move) {
	r.Lock()
	defer r.Unlock()
	if r.State != Running {
		return
	}

	// Prevent player 1 from changing the move.
	if r.Results[r.CurrentRound].Player1Move != NoMove {
		return
	}

	r.Results[r.CurrentRound].Player1Move = move
	r.processRound()
}

func (r *Room) AddPlayer2Move(move Move) {
	r.Lock()
	defer r.Unlock()
	if r.State != Running {
		return
	}

	// Prevent player 2 from changing the move.
	if r.Results[r.CurrentRound].Player2Move != NoMove {
		return
	}

	r.Results[r.CurrentRound].Player2Move = move
	r.processRound()
}

func (r *Room) processRound() {
	player1Move := r.Results[r.CurrentRound].Player1Move
	player2Move := r.Results[r.CurrentRound].Player2Move

	// If any of the players didn't make a move.
	if player1Move == NoMove || player2Move == NoMove {
		return
	}

	// If both players made the same move.
	if player1Move == player2Move {
		r.Results[r.CurrentRound].Player1Move = NoMove
		r.Results[r.CurrentRound].Player2Move = NoMove
		return
	}

	r.Results[r.CurrentRound] = RoundResult{
		Player1Move: player1Move,
		Player2Move: player2Move,
	}

	// TODO: Add logic to determine the winner if win more than the half of the rounds.

	r.CurrentRound++
	if r.CurrentRound == r.Rounds {
		// TODO: Add overtime logic.
		r.State = Finished
	}
}

func (r *Room) GetResults() (int, int) {
	r.Lock()
	defer r.Unlock()

	player1Wins := 0
	player2Wins := 0
	for _, result := range r.Results {
		switch {
		case result.Player1Move == NoMove || result.Player2Move == NoMove:
			continue
		case result.Player1Move == result.Player2Move:
			continue
		case result.Player1Move == Rock && result.Player2Move == Scissors:
			player1Wins++
		case result.Player1Move == Paper && result.Player2Move == Rock:
			player1Wins++
		case result.Player1Move == Scissors && result.Player2Move == Paper:
			player1Wins++
		default:
			player2Wins++
		}
	}

	return player1Wins, player2Wins
}

func (r *Room) GetState() RoomState {
	r.Lock()
	defer r.Unlock()
	return r.State
}
