package tag

type ID int

// constants
const (
	// Player is abstract player
	Player ID = iota
	// User is keyboard-controlled player
	User
)
