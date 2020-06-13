package component

import (
	"github.com/hajimehoshi/ebiten"
)

type ID int

// IDs
const (
	ShapeID ID = iota
	VectorID
	SpeedID
	PosID
)

type Component interface {
	ID() ID
}

type Shape interface {
	Component
	Draw(board *ebiten.Image, pos Pos)
}
