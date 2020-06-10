package shape

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/niakr1s/nrg-go/src/ecs/component"
	"github.com/niakr1s/nrg-go/src/geo"
	"github.com/niakr1s/nrg-go/src/img"
	log "github.com/sirupsen/logrus"
)

// Circle ..
type Circle struct {
	geo.Pos
	geo.Radius

	image *ebiten.Image
}

func NewCircle(x, y, r float64, path string) *Circle {
	image, err := img.Load(path)
	if err != nil {
		log.Fatal(err)
	}
	return &Circle{Pos: geo.NewPos(x, y), Radius: geo.NewRadius(r), image: image}
}

func (c *Circle) Draw(board *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	w, h := c.image.Size()
	op.GeoM.Translate(-float64(w)/2, -float64(h)/2)
	scale := float64(c.Radius) / float64(w)
	op.GeoM.Scale(scale, scale)
	op.GeoM.Translate(c.X, c.Y)

	board.DrawImage(c.image, op)
}

func (c *Circle) Center() geo.Pos {
	return c.Pos
}

func (c *Circle) Intersects(rhs component.Shape) bool {
	return false
}
