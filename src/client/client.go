package client

import (
	"github.com/hajimehoshi/ebiten"
	startmenu "github.com/niakr1s/nrg-go/src/client/startMenu"
	"github.com/niakr1s/nrg-go/src/client/state"
	"github.com/niakr1s/nrg-go/src/config"
)

// Client ...
type Client struct {
	state state.State
}

// New ...
func New() *Client {

	mainMenu := startmenu.New()
	return &Client{state: mainMenu}
}

// Init ...
func (c *Client) Init() {
	ebiten.SetWindowSize(config.ScreenWidth, config.ScreenHeight)
	ebiten.SetWindowTitle("Hello, World!")
	ebiten.SetRunnableOnUnfocused(true)
}

// Update ...
func (c *Client) Update(screen *ebiten.Image) error {
	if next := c.state.Next(); next != nil {
		c.state = next
	}
	return c.state.Update(screen)
}

// Draw ...
func (c *Client) Draw(screen *ebiten.Image) {
	c.state.Draw(screen)
}

// Layout ...
func (c *Client) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return config.ScreenWidth, config.ScreenHeight
}
