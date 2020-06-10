package component

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/niakr1s/nrg-go/src/geo"
)

type Component interface{}

type Drawable interface {
	Draw(*ebiten.Image)
}

type Intersectable interface {
	Intersects(Intersectable) bool
}

type Movable interface {
	Move(geo.Vector, geo.Distance)
}
