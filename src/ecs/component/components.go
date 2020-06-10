package component

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/niakr1s/nrg-go/src/geo"
)

type ID int

// IDs
const (
	DrawableID ID = iota
	IntersectableID
	MovableID
)

type Component interface{}

type Drawable interface {
	Draw(board *ebiten.Image)
}

type Intersectable interface {
	Intersects(rhs Intersectable) bool
}

type Movable interface {
	Move(vec geo.Vector, dist geo.Distance)
}
