package client

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten"
	"github.com/niakr1s/nrg-go/src/config"
	"github.com/niakr1s/nrg-go/src/ecs/component"
	"github.com/niakr1s/nrg-go/src/ecs/registry"
	"github.com/niakr1s/nrg-go/src/ecs/system"
)

// Client ...
type Client struct {
	Reg *registry.Registry

	board chan *ebiten.Image

	systems []system.System
}

// New ...
func New() *Client {
	return &Client{Reg: registry.NewRegistry(), board: make(chan *ebiten.Image), systems: make([]system.System, 0)}
}

// Init ...
func (c *Client) Init() {
	ebiten.SetWindowSize(config.ScreenWidth, config.ScreenHeight)
	ebiten.SetWindowTitle("Hello, World!")
	ebiten.SetRunnableOnUnfocused(true)

	c.startProduceBoard()

	c.systems = append(c.systems, system.NewKeyBoard(c.Reg))
}

// Update ...
func (c *Client) Update(screen *ebiten.Image) error {
	for _, s := range c.systems {
		s.Step()
	}

	c.Reg.RLock()
	defer c.Reg.RUnlock()
	for _, e := range c.Reg.Entities {
		e.Lock()
		if vec, pos, speed := e.GetComponent(component.VectorID), e.GetComponent(component.PosID), e.GetComponent(component.SpeedID); vec != nil && pos != nil && speed != nil {
			vec := vec.(*component.Vector)
			pos := pos.(*component.Pos)
			speed := speed.(*component.Speed)
			pos.Move(*vec, *speed)
		}
		e.Unlock()
	}

	return nil
}

// Draw ...
func (c *Client) Draw(screen *ebiten.Image) {
	c.drawBoard(screen)
}

func (c *Client) startProduceBoard() {
	go func() {
		for {
			board, _ := ebiten.NewImage(1000, 1000, ebiten.FilterDefault)
			board.Fill(color.Gray16{0xaaaf})
			c.Reg.RLock()
			for _, e := range c.Reg.Entities {
				e.RLock()
				if c, pos := e.GetComponent(component.ShapeID), e.GetComponent(component.PosID); c != nil && pos != nil {
					shape := c.(component.Shape)
					pos := pos.(*component.Pos)
					shape.Draw(board, *pos)
				}
				e.RUnlock()
			}
			c.Reg.RUnlock()
			c.board <- board
		}
	}()
}

func (c *Client) drawBoard(screen *ebiten.Image) {
	board := <-c.board
	op := &ebiten.DrawImageOptions{}

	swI, shI := screen.Size()
	bwI, bhI := board.Size()

	// sizes of screen and board
	sw, sh := float64(swI), float64(shI)
	bw, bh := float64(bwI), float64(bhI)

	scale := 0.9 * math.Min(sw, sh) / math.Max(bw, bh)

	// scaled board size
	bw, bh = bw*scale, bh*scale

	op.GeoM.Scale(scale, scale)
	op.GeoM.Translate((sw-bw)/2, (sh-bh)/2)

	screen.DrawImage(board, op)
}

// Layout ...
func (c *Client) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return config.ScreenWidth, config.ScreenHeight
}
