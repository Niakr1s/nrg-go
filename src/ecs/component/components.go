package component

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/niakr1s/nrg-go/src/geo"
)

type ID int

// IDs
const (
	ShapeID ID = iota
	VectorID

	PlayerTagID
	UserTagID
)

type Component interface{}

type Shape interface {
	Draw(board *ebiten.Image)
	Intersects(rhs Shape) bool
	Move(vec geo.Vector, dist geo.Distance)
	Center() geo.Pos
}

type Vector interface {
	Vector() geo.Vector
}

// PlayerTag tag for a player
type PlayerTag interface{}

// UserTag is tag for user-controlled unit
type UserTag interface{}
