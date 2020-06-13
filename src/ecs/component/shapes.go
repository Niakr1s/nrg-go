package component

import (
	"math"

	"github.com/hajimehoshi/ebiten"
	"github.com/niakr1s/nrg-go/src/img"
	log "github.com/sirupsen/logrus"
)

type Shape interface {
	Component
	Draw(board *ebiten.Image, pos Pos)
	Bound(center Pos) Bound
	Intersects(selfCenter, rhsCenter Pos, rhs Shape) bool
}

type Bound struct {
	TopLeft, BotRight Pos
}

func NewBound(center Pos, w, h float64) Bound {
	left := center.X - w/2
	right := center.X + w/2
	top := center.Y - h/2
	bot := center.Y + h/2
	return Bound{TopLeft: NewPos(left, top), BotRight: NewPos(right, bot)}
}

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

func (c *Circle) Bound(center Pos) Bound {
	return NewBound(center, c.R, c.R)
}

func (c *Circle) Intersects(selfCenter, rhsCenter Pos, rhs Shape) bool {
	switch rhs := rhs.(type) {
	case *Circle:
		dist := math.Sqrt(math.Pow(selfCenter.X-rhsCenter.X, 2) +
			math.Pow(selfCenter.Y-rhsCenter.Y, 2))
		return dist <= c.R+rhs.R

	default:
		log.Errorf("Circle.Intersects: got unknown rhs: %v", rhs)
		return false
	}
}
