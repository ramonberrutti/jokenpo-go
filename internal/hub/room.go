package hub

import "sync"

type Room struct {
	// Lock to prevent concurrent access to the room.
	sync.RWMutex

	// Id of the room.
	id string

	// Name of the room.
	name string

	// Player 1 of the room.
	player1 string

	// Player 2 of the room.
	player2 string

	// Player 1 join code.
	player1JoinCode string

	// Player 2 join code.
	player2JoinCode string

	// State of the room.
	state RoomState

	// Rounds to play in the room.
	rounds int

	// Current round of the room.
	currentRound int

	// Current round of the room.
	results []RoundResult

	subscriptions  map[int]func(e *Event)
	subscriptionId int
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

func (m Move) String() string {
	return [...]string{"NoMove", "Rock", "Paper", "Scissors"}[m]
}

type RoundResult struct {
	// Player 1 move.
	Player1Move Move

	// Player 2 move.
	Player2Move Move
}

func NewRoom(id string, name string, player1 string, player2 string, rounds int) *Room {
	return &Room{
		id:      id,
		name:    name,
		player1: player1,
		player2: player2,
		rounds:  rounds,
		state:   WaitingForPlayers,
		results: make([]RoundResult, rounds),
	}
}

func (r *Room) Start() {
	r.Lock()
	defer r.Unlock()
	if r.state != WaitingForPlayers {
		return
	}

	r.state = Running
}

func (r *Room) AddPlayer1Move(move Move) {
	r.Lock()
	defer r.Unlock()
	if r.state != Running {
		return
	}

	// Prevent player 1 from changing the move.
	if r.results[r.currentRound].Player1Move != NoMove {
		return
	}

	r.results[r.currentRound].Player1Move = move
	r.processRound()
}

func (r *Room) AddPlayer2Move(move Move) {
	r.Lock()
	defer r.Unlock()
	if r.state != Running {
		return
	}

	// Prevent player 2 from changing the move.
	if r.results[r.currentRound].Player2Move != NoMove {
		return
	}

	r.results[r.currentRound].Player2Move = move
	r.processRound()
}

func (r *Room) processRound() {
	player1Move := r.results[r.currentRound].Player1Move
	player2Move := r.results[r.currentRound].Player2Move

	// If any of the players didn't make a move.
	if player1Move == NoMove || player2Move == NoMove {
		return
	}

	// If both players made the same move.
	if player1Move == player2Move {
		r.results[r.currentRound].Player1Move = NoMove
		r.results[r.currentRound].Player2Move = NoMove
		return
	}

	r.results[r.currentRound] = RoundResult{
		Player1Move: player1Move,
		Player2Move: player2Move,
	}

	// TODO: Add logic to determine the winner if win more than the half of the rounds.

	r.currentRound++
	if r.currentRound == r.rounds {
		// TODO: Add overtime logic.
		r.state = Finished
	}
}

// GetResults return the ongoing results of the room.
func (r *Room) GetResults() (int, int) {
	r.Lock()
	defer r.Unlock()

	player1Wins := 0
	player2Wins := 0
	for _, result := range r.results {
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

func (r *Room) GetId() string {
	r.RLock()
	defer r.RUnlock()
	return r.id
}

func (r *Room) GetName() string {
	r.RLock()
	defer r.RUnlock()
	return r.name
}

func (r *Room) GetState() RoomState {
	r.RLock()
	defer r.RUnlock()
	return r.state
}

func (r *Room) GetPlayer1() string {
	r.RLock()
	defer r.RUnlock()
	return r.player1
}

func (r *Room) GetPlayer2() string {
	r.RLock()
	defer r.RUnlock()
	return r.player2
}

func (r *Room) GetPlayer1JoinCode() string {
	r.RLock()
	defer r.RUnlock()
	return r.player1JoinCode
}

func (r *Room) GetPlayer2JoinCode() string {
	r.RLock()
	defer r.RUnlock()
	return r.player2JoinCode
}

func (r *Room) GetRounds() int {
	r.RLock()
	defer r.RUnlock()
	return r.rounds
}

func (r *Room) GetCurrentRound() int {
	r.RLock()
	defer r.RUnlock()
	return r.currentRound
}

func (r *Room) GetResultsForRound(round int) (Move, Move) {
	r.RLock()
	defer r.RUnlock()

	if round < 0 || round >= r.rounds {
		return NoMove, NoMove
	}

	return r.results[round].Player1Move, r.results[round].Player2Move
}

type Event struct {
	State        RoomState
	CurrentRound int
	Player1Move  Move
	Player2Move  Move
	Player1Wins  int
	Player2Wins  int
}

func (r *Room) Subscribe(cb func(*Event)) func() {
	r.Lock()
	defer r.Unlock()
	r.subscriptionId++
	nextId := r.subscriptionId

	r.subscriptions[nextId] = cb

	return func() {
		r.Lock()
		defer r.Unlock()

		delete(r.subscriptions, nextId)
	}
}
