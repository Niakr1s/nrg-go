package menu

import (
	"os"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/niakr1s/nrg-go/src/client/game"
	"github.com/niakr1s/nrg-go/src/client/state"
	"github.com/niakr1s/nrg-go/src/config"
	"github.com/niakr1s/nrg-go/src/ecs/component"
	"github.com/niakr1s/nrg-go/src/widget"
)

type MainMenu struct {
	next   state.State
	layout widget.Layout
}

func NewMainMenu() *MainMenu {
	menu := &MainMenu{layout: widget.NewLayout()}
	w := 200
	x := (config.ScreenWidth - w) / 2
	menu.layout.AddWidget(component.NewPos(float64(x), 200), widget.NewButton(w, 60, "Start", func() { menu.next = game.NewGame() }))
	menu.layout.AddWidget(component.NewPos(float64(x), 300), widget.NewButton(w, 60, "Exit", func() { os.Exit(0) }))
	return menu
}

func (m *MainMenu) Update(screen *ebiten.Image) error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		e := widget.NewMouseClickEvent(component.NewPos(float64(x), float64(y)))
		m.layout.NewEvent(e)
	}
	return nil
}

func (m *MainMenu) Draw(screen *ebiten.Image) {
	for _, w := range m.layout {
		op := ebiten.DrawImageOptions{}
		op.GeoM.Translate(w.Pos.X, w.Pos.Y)
		screen.DrawImage(w.Draw(), &op)
	}
}

func (m *MainMenu) Next() state.State {
	return m.next
}
