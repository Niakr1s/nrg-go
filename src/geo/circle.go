package geo

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/niakr1s/nrg-go/src/img"
	log "github.com/sirupsen/logrus"
)

// Circle ..
type Circle struct {
	Pos
	Radius
}

func (c *Circle) Draw(board *ebiten.Image) {
	image, err := img.Load("data/white_circle.png")
	if err != nil {
		log.Fatal(err)
	}
	eImage, err := ebiten.NewImageFromImage(image, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
	board.DrawImage(eImage, &ebiten.DrawImageOptions{})
}
