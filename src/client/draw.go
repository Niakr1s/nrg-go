package client

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten"
	"github.com/niakr1s/nrg-go/src/config"
	"github.com/niakr1s/nrg-go/src/ecs/component"
)

func (c *Client) produceBoard() *ebiten.Image {
	board, _ := ebiten.NewImage(config.BoardWidth, config.BoardHeight, ebiten.FilterDefault)
	board.Fill(color.Gray16{0xaaaf})
	c.Reg.RLock()
	for _, e := range c.Reg.Entities {
		e.RLock()
		if cs := e.GetComponents(component.ShapeID, component.PosID); cs != nil {
			shape := cs[0].(component.Shape)
			pos := cs[1].(component.Pos)
			shape.Draw(board, pos)
		}
		e.RUnlock()
	}
	c.Reg.RUnlock()
	return board
}

func (c *Client) drawBoard(screen *ebiten.Image) {
	board := c.produceBoard()
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
