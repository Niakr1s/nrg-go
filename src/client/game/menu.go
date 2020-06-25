package game

import (
	"os"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/niakr1s/nrg-go/src/client/state"
	"github.com/niakr1s/nrg-go/src/config"
	"github.com/niakr1s/nrg-go/src/ecs/component"
	"github.com/niakr1s/nrg-go/src/widget"
)

type Menu struct {
	*state.NextState
	layout widget.Layout
}

func NewMenu() *Menu {
	menu := &Menu{layout: widget.NewLayout(), NextState: state.NewNext()}
	return menu
}

func NewMainMenu() *Menu {
	menu := NewMenu()
	menu.SetButtons(
		button{"Start", func() { menu.SetNext(NewGame()) }},
		button{"Exit", func() { os.Exit(0) }},
	)
	return menu
}

func NewPauseMenu(paused state.State) *Menu {
	menu := NewMenu()
	menu.SetButtons(
		button{"Continue", func() { menu.SetNext(paused) }},
		button{"Main menu", func() { menu.SetNext(NewMainMenu()) }},
		button{"Exit", func() { os.Exit(0) }},
	)
	return menu
}

type button struct {
	text    string
	onClick func()
}

func (m *Menu) SetButtons(btns ...button) {
	layout := widget.NewLayout()
	l := len(btns)
	const (
		buttonW = 300
		buttonH = 60
		spaceH  = 40
	)
	centeredX := (config.ScreenWidth - buttonW) / 2
	allButtonsH := l*buttonH + (l-1)*spaceH
	y := (config.ScreenHeight - allButtonsH) / 2
	for _, btn := range btns {
		layout.AddWidget(component.NewPos(float64(centeredX), float64(y)), widget.NewButton(buttonW, buttonH, btn.text, btn.onClick))
		y += buttonH + spaceH
	}
	m.layout = layout
}

func (m *Menu) Update(screen *ebiten.Image) error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		e := widget.NewMouseClickEvent(component.NewPos(float64(x), float64(y)))
		m.layout.NewEvent(e)
	}
	return nil
}

func (m *Menu) Draw(screen *ebiten.Image) {
	for _, w := range m.layout {
		op := ebiten.DrawImageOptions{}
		op.GeoM.Translate(w.Pos.X, w.Pos.Y)
		screen.DrawImage(w.Draw(), &op)
	}
}
