package widget

import "github.com/hajimehoshi/ebiten"

type Widget interface {
	Width() int
	Height() int
	OnEvent(Event)
	Draw() *ebiten.Image
}
