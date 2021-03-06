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

type Tags map[ID]struct{}

func (e Tags) SetTags(ids ...ID) {
	for _, id := range ids {
		e[id] = struct{}{}
	}
}

// HasTags returns false if any is missing
func (e Tags) HasTags(ids ...ID) bool {
	for _, id := range ids {
		if _, ok := e[id]; !ok {
			return false
		}
	}
	return true
}

func (e Tags) RemoveTags(ids ...ID) {
	for _, id := range ids {
		delete(e, id)
	}
}
