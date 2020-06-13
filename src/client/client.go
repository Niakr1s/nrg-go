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
		if vec, shape := e.GetComponent(component.VectorID), e.GetComponent(component.ShapeID); vec != nil && shape != nil {
			vec := vec.(component.Vector)
			shape := shape.(component.Shape)
			shape.Move(vec.Vector(), 1)
		}
		e.Unlock()
	}

	return nil
}

// Draw ...
func (c *Client) Draw(screen *ebiten.Image) {
	c.drawBoard(screen)
}

func (c *Client) StartProduceBoard() {
	go func() {
		for {
			board, _ := ebiten.NewImage(1000, 1000, ebiten.FilterDefault)
			board.Fill(color.Gray16{0xaaaf})
			c.Reg.RLock()
			for _, e := range c.Reg.Entities {
				e.RLock()
				if c := e.GetComponent(component.ShapeID); c != nil {
					c := c.(component.Shape)
					c.Draw(board)
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

func min(first int, other ...int) int {
	for _, i := range other {
		if i < first {
			first = i
		}
	}
	return first
}

func max(first int, other ...int) int {
	for _, i := range other {
		if i > first {
			first = i
		}
	}
	return first
}

// Layout ...
func (c *Client) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return config.ScreenWidth, config.ScreenHeight
}
