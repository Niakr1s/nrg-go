package startmenu

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/niakr1s/nrg-go/src/client/state"
)

type StartMenu struct {
	next state.State
}

func New() *StartMenu {
	return &StartMenu{}
}

func (m *StartMenu) Update(screen *ebiten.Image) error {
	return nil
}

func (m *StartMenu) Draw(screen *ebiten.Image) {

}

func (m *StartMenu) Next() state.State {
	return m.next
}
