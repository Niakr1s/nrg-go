package geo

import (
	"image"

	"github.com/hajimehoshi/ebiten"
	"github.com/niakr1s/nrg-go/src/img"
	log "github.com/sirupsen/logrus"
)

// Circle ..
type Circle struct {
	Pos
	Radius

	image image.Image
}

func NewCircle(x, y, r float64, path string) *Circle {
	image, err := img.Load(path)
	if err != nil {
		log.Fatal(err)
	}
	return &Circle{Pos: NewPos(x, y), Radius: NewRadius(r), image: image}
}

func (c *Circle) Draw(board *ebiten.Image) {
	eImage, err := ebiten.NewImageFromImage(c.image, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	op := &ebiten.DrawImageOptions{}
	w, h := eImage.Size()
	op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
	scale := float64(c.Radius) / float64(w)
	op.GeoM.Scale(scale, scale)
	op.GeoM.Translate(c.X, c.Y)

	board.DrawImage(eImage, op)
}
