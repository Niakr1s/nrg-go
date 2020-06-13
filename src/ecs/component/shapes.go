package component

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/niakr1s/nrg-go/src/img"
	log "github.com/sirupsen/logrus"
)

// Circle ..
type Circle struct {
	R float64

	image *ebiten.Image
}

func NewCircle(r float64, path string) *Circle {
	image, err := img.Load(path)
	if err != nil {
		log.Fatal(err)
	}
	return &Circle{R: r, image: image}
}

func (c *Circle) ID() ID {
	return ShapeID
}

func (c *Circle) Draw(board *ebiten.Image, pos Pos) {
	op := &ebiten.DrawImageOptions{}
	w, h := c.image.Size()
	op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
	scale := float64(c.R) / float64(w)
	op.GeoM.Scale(scale, scale)
	op.GeoM.Translate(pos.X, pos.Y)

	board.DrawImage(c.image, op)
}
