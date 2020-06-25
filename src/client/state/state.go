package state

import "github.com/hajimehoshi/ebiten"

type State interface {
	Update(screen *ebiten.Image) error
	Draw(screen *ebiten.Image)
	Next() State
}
