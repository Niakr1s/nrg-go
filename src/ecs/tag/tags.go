package tag

type ID int

// constants
const (
	// Enemy is abstract player
	Enemy ID = iota
	// User is keyboard-controlled player
	User

	Destroyed
	Bullet
)
