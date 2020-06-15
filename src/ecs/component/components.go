package component

type ID int

// IDs
const (
	ShapeID ID = iota
	VectorID
	SpeedID
	PosID
	GroundID
)

type Component interface {
	ID() ID
}

type Components []Component
