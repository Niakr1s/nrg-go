package client

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/niakr1s/nrg-go/src/client/key"
)

// Client ...
type Client struct {
	keyCh <-chan key.Event
}

// New ...
func New() *Client {
	return &Client{}
}

// Init ...
func (c *Client) Init() {
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("Hello, World!")

	c.keyCh = key.NewListener().StartPollKeys()
}

// Update ...
func (c *Client) Update(screen *ebiten.Image) error {
	return nil
}

// Draw ...
func (c *Client) Draw(screen *ebiten.Image) {
}

// Layout ...
func (c *Client) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 800, 600
}
