package menu

import (
	"os"

	"github.com/niakr1s/nrg-go/src/client/game"
)

func NewMainMenu() *Menu {
	menu := NewMenu()
	menu.SetButtons(
		button{"Start", func() { menu.next = game.NewGame() }},
		button{"Exit", func() { os.Exit(0) }},
	)
	return menu
}
