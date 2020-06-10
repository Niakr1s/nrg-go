package client

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/niakr1s/nrg-go/src/client/key"
	"github.com/niakr1s/nrg-go/src/config"
	log "github.com/sirupsen/logrus"
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
	ebiten.SetWindowSize(config.ScreenWidth, config.ScreenHeight)
	ebiten.SetWindowTitle("Hello, World!")

	c.keyCh = key.NewListener().StartPollKeys()
	go func() {
		for {
			log.Tracef("%s", <-c.keyCh)
		}
	}()
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
	return config.ScreenWidth, config.ScreenHeight
}
