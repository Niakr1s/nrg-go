package tag

type ID int

// constants
const (
	// PlayerID is abstract player
	PlayerID ID = iota
	// UserID is keyboard-controlled player
	UserID
	// GroundID is used to distinguish ground bodies from dynamic
	GroundID
)

type Tag interface {
	TagID() ID
}

type Player struct{}

func NewPlayer() *Player {
	return &Player{}
}

func (p *Player) ID() ID {
	return PlayerID
}

type User struct{}

func NewUser() *User {
	return &User{}
}

func (u *User) ID() ID {
	return UserID
}

type Ground struct{}

func (d Ground) ID() ID {
	return GroundID
}
