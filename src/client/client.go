package client

import (
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/niakr1s/nrg-go/src/config"
	"github.com/niakr1s/nrg-go/src/ecs/registry"
	"github.com/niakr1s/nrg-go/src/ecs/system"
	log "github.com/sirupsen/logrus"
)

// Client ...
type Client struct {
	Reg *registry.Registry

	systems []system.System
}

// New ...
func New() *Client {
	return &Client{Reg: registry.NewRegistry(), systems: make([]system.System, 0)}
}

// Init ...
func (c *Client) Init() {
	ebiten.SetWindowSize(config.ScreenWidth, config.ScreenHeight)
	ebiten.SetWindowTitle("Hello, World!")
	ebiten.SetRunnableOnUnfocused(true)

	go func() {
		for {
			<-time.After(time.Second * 10)
			c.Reg.RLock()
			log.Tracef("%d entities", len(c.Reg.Entities))
			c.Reg.RUnlock()
		}
	}()

	c.systems = append(c.systems,
		system.NewKeyBoard(c.Reg),
		system.NewMove(c.Reg, config.BoardWidth, config.BoardHeight),
		system.NewContain(c.Reg, config.BoardWidth, config.BoardHeight),
		system.NewBounce(c.Reg),
		system.NewWeapon(c.Reg),
		system.NewDamage(c.Reg),
		system.NewDestroy(c.Reg, config.BoardWidth, config.BoardHeight),
		system.NewClean(c.Reg),
	)
}

// Update ...
func (c *Client) Update(screen *ebiten.Image) error {
	for _, s := range c.systems {
		s.Step()
	}

	return nil
}

// Draw ...
func (c *Client) Draw(screen *ebiten.Image) {
	c.drawBoard(screen)
}

// Layout ...
func (c *Client) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return config.ScreenWidth, config.ScreenHeight
}
